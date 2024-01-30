// This file is part of GoRE.
//
// Copyright (C) 2019-2021 GoRE Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package gore

import (
	"fmt"
	"path"
	"runtime/debug"
	"sort"
	"strings"
)

//go:generate go run gen.go

var (
	knownRepos = []string{"golang.org", "github.com", "gitlab.com"}
)

// Package is a representation of a Go package.
type Package struct {
	// Name is the name of the package.
	Name string `json:"name"`
	// Filepath is the extracted file path for the package.
	Filepath string `json:"path"`
	// Functions is a list of functions that are part of the package.
	Functions []*Function `json:"functions"`
	// Methods a list of methods that are part of the package.
	Methods []*Method `json:"methods"`
}

// GetSourceFiles returns a slice of source files within the package.
// The source files are a representations of the source code files in the package.
func (f *GoFile) GetSourceFiles(p *Package) []*SourceFile {
	tmp := make(map[string]*SourceFile)
	getSourceFile := func(fileName string) *SourceFile {
		sf, ok := tmp[fileName]
		if !ok {
			return &SourceFile{Name: path.Base(fileName)}
		}
		return sf
	}

	// Sort functions and methods by source file.
	for _, fn := range p.Functions {
		fileName, _, _ := f.pclntab.PCToLine(fn.Offset)
		start, end := findSourceLines(fn.Offset, fn.End, f.pclntab)

		e := FileEntry{Name: fn.Name, Start: start, End: end}

		sf := getSourceFile(fileName)
		sf.entries = append(sf.entries, e)
		tmp[fileName] = sf
	}
	for _, m := range p.Methods {
		fileName, _, _ := f.pclntab.PCToLine(m.Offset)
		start, end := findSourceLines(m.Offset, m.End, f.pclntab)

		e := FileEntry{Name: fmt.Sprintf("%s%s", m.Receiver, m.Name), Start: start, End: end}

		sf := getSourceFile(fileName)
		sf.entries = append(sf.entries, e)
		tmp[fileName] = sf
	}

	// Create final slice and populate it.
	files := make([]*SourceFile, len(tmp))
	i := 0
	for _, sf := range tmp {
		files[i] = sf
		i++
	}

	// Sort the file list.
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
	return files
}

// PackageClass is a type used to indicate the package kind.
type PackageClass uint8

const (
	// ClassUnknown is used for packages that could not be classified.
	ClassUnknown PackageClass = iota
	// ClassSTD is used for packages that are part of the standard library.
	ClassSTD
	// ClassMain is used for the main package and its subpackages.
	ClassMain
	// ClassVendor is used for vendor packages.
	ClassVendor
	// ClassGenerated are used for packages generated by the compiler.
	ClassGenerated
)

// PackageClassifier classifies a package to the correct class type.
type PackageClassifier interface {
	// Classify performs the classification.
	Classify(pkg *Package) PackageClass
}

// NewPathPackageClassifier constructs a new classifier based on the main package's filepath.
func NewPathPackageClassifier(mainPkgFilepath string) *PathPackageClassifier {
	return &PathPackageClassifier{
		mainFilepath: mainPkgFilepath, mainFolders: []string{
			path.Dir(mainPkgFilepath),
			path.Clean(mainPkgFilepath),
		},
	}
}

// PathPackageClassifier can classify the class of a go package.
type PathPackageClassifier struct {
	mainFilepath string
	mainFolders  []string
}

// Classify returns the package class for the package.
func (c *PathPackageClassifier) Classify(pkg *Package) PackageClass {
	if pkg.Name == "type" || strings.HasPrefix(pkg.Name, "type..") {
		return ClassGenerated
	}

	if IsStandardLibrary(pkg.Name) {
		return ClassSTD
	}

	if isGeneratedPackage(pkg) {
		return ClassGenerated
	}

	// Detect internal/golang.org/x/net/http2/hpack type/
	tmp := strings.Split(pkg.Name, "/golang.org")[0]
	if len(tmp) < len(pkg.Name) && IsStandardLibrary(tmp) {
		return ClassSTD
	}

	// cgo packages.
	if strings.HasPrefix(pkg.Name, "_cgo_") || strings.HasPrefix(pkg.Name, "x_cgo_") {
		return ClassSTD
	}

	// If the file path contains "@v", it's a 3rd party package.
	if strings.Contains(pkg.Filepath, "@v") {
		return ClassVendor
	}

	parentFolder := path.Dir(pkg.Filepath)

	if strings.HasPrefix(pkg.Filepath, c.mainFilepath+"/vendor/") ||
		strings.HasPrefix(pkg.Filepath, path.Dir(c.mainFilepath)+"/vendor/") ||
		strings.HasPrefix(pkg.Filepath, path.Dir(path.Dir(c.mainFilepath))+"/vendor/") {
		return ClassVendor
	}

	for _, folder := range c.mainFolders {
		if parentFolder == folder {
			return ClassMain
		}
	}

	// If the package name starts with "vendor/" assume it's a vendor package.
	if strings.HasPrefix(pkg.Name, "vendor/") {
		return ClassVendor
	}

	// Start with repo url.and has it in the path.
	for _, url := range knownRepos {
		if strings.HasPrefix(pkg.Name, url) && strings.Contains(pkg.Filepath, url) {
			return ClassVendor
		}
	}

	// If the path does not contain the "vendor" in a path but has the main package folder name, assume part of main.
	if !strings.Contains(pkg.Filepath, "vendor/") &&
		(path.Base(path.Dir(pkg.Filepath)) == path.Base(c.mainFilepath)) {
		return ClassMain
	}
	// Special case for entry point package.
	if pkg.Name == "" && path.Base(pkg.Filepath) == "runtime" {
		return ClassSTD
	}

	// At this point, if it's a subpackage of the main assume main.
	if strings.HasPrefix(pkg.Filepath, c.mainFilepath) {
		return ClassMain
	}

	// Check if it's the main parent package.
	if pkg.Name != "" && (!strings.Contains(pkg.Name, "/") && strings.Contains(c.mainFilepath, pkg.Name)) {
		return ClassMain
	}

	// At this point, if the main package has a file path of "command-line-arguments" and we haven't figured out
	// what class it is. We assume it is part of the main package.
	if c.mainFilepath == "command-line-arguments" {
		return ClassMain
	}

	return ClassUnknown
}

// IsStandardLibrary returns true if the package is from the standard library.
// Otherwise, false is retuned.
func IsStandardLibrary(pkg string) bool {
	_, ok := stdPkgs[pkg]
	if ok {
		return true
	}

	// Detect regexp.(*onePassInst).regexp/syntax type packages
	tmp := strings.Split(pkg, ".")[0]
	if len(tmp) < len(pkg) && IsStandardLibrary(tmp) {
		return true
	}

	return false
}

func isGeneratedPackage(pkg *Package) bool {
	if pkg.Filepath == "<autogenerated>" {
		return true
	}

	// Special case for no package name and path of "".
	if pkg.Name == "" && pkg.Filepath == "" {
		return true
	}

	// Some internal stuff, classify it as Generated
	if pkg.Filepath == "" && (pkg.Name == "__x86" || pkg.Name == "__i686") {
		return true
	}

	return false
}

// NewModPackageClassifier creates a new mod based package classifier.
func NewModPackageClassifier(buildInfo *debug.BuildInfo) *ModPackageClassifier {
	return &ModPackageClassifier{modInfo: buildInfo}
}

// ModPackageClassifier uses the mod info extracted from the binary to classify packages.
type ModPackageClassifier struct {
	modInfo *debug.BuildInfo
}

// Classify performs the classification.
func (c *ModPackageClassifier) Classify(pkg *Package) PackageClass {
	if IsStandardLibrary(pkg.Name) {
		return ClassSTD
	}

	// Main package.
	if pkg.Name == "main" {
		return ClassMain
	}

	// If the build info path is not an empty string and the package has the path as a substring, it is part of the main module.
	if c.modInfo.Path != "" && (strings.HasPrefix(pkg.Filepath, c.modInfo.Path) || strings.HasPrefix(pkg.Name, c.modInfo.Path)) {
		return ClassMain
	}

	// If the main module path is not an empty string and the package has the path as a substring, it is part of the main module.
	if c.modInfo.Main.Path != "" && (strings.HasPrefix(pkg.Filepath, c.modInfo.Main.Path) || strings.HasPrefix(pkg.Name, c.modInfo.Main.Path)) {
		return ClassMain
	}

	// Check if the package is a direct dependency.
	for _, dep := range c.modInfo.Deps {
		if strings.HasPrefix(pkg.Filepath, dep.Path) || strings.HasPrefix(pkg.Name, dep.Path) {
			// If the vendor it matched on has the version of "(devel)", it is treated as part of
			// the main module.
			if dep.Version == "(devel)" {
				return ClassMain
			}
			return ClassVendor
		}
	}

	if isGeneratedPackage(pkg) {
		return ClassGenerated
	}

	// cgo packages.
	if strings.HasPrefix(pkg.Name, "_cgo_") || strings.HasPrefix(pkg.Name, "x_cgo_") {
		return ClassSTD
	}

	// Only indirect dependencies should be left.
	return ClassVendor
}
