// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/filemode"
	"github.com/go-git/go-git/v5/plumbing/object"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Generate with version hash: {{ .Hash }}
var pkgHashMatcher = regexp.MustCompile(`// Generate with version hash: (\b[A-Fa-f0-9]{64}\b)`)

func getCurrentStdPkgHash() (string, error) {
	f, err := os.Open(stdpkgOutputFile)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}

	matches := pkgHashMatcher.FindStringSubmatch(buf.String())
	if len(matches) != 2 {
		return "", nil
	}

	return matches[1], nil
}

func generateStdPkgs() {
	fmt.Println("Generating " + stdpkgOutputFile)

	collect := func(tag string) ([]string, error) {
		reference, err := goRepo.Tag(tag)
		if err != nil {
			return nil, err
		}

		commit, err := goRepo.CommitObject(reference.Hash())
		if err != nil {
			return nil, err
		}

		tree, err := commit.Tree()
		if err != nil {
			return nil, err
		}

		var stdPkgs []string

		addPkg := func(s string) {
			if strings.HasSuffix(s, "_asm") ||
				strings.Contains(s, "testdata") {
				return
			}
			stdPkgs = append(stdPkgs, s)
		}

		var dive func(prefix string, tree *object.Tree) error
		dive = func(prefix string, tree *object.Tree) error {
			for _, entry := range tree.Entries {
				if entry.Mode == filemode.Dir {
					subTree, err := tree.Tree(entry.Name)
					if err != nil {
						return fmt.Errorf("error when getting tree for %s: %w", entry.Name, err)
					}
					p := path.Join(prefix, entry.Name)
					addPkg(p)
					err = dive(p, subTree)
					if err != nil {
						return fmt.Errorf("error when diving into %s: %w", p, err)
					}
				}
			}
			return nil
		}

		srcTree, err := tree.Tree("src")
		if err != nil {
			return nil, err
		}

		for _, entry := range srcTree.Entries {
			if entry.Name == "cmd" {
				continue
			}

			if entry.Mode == filemode.Dir {
				subTree, err := srcTree.Tree(entry.Name)
				if err != nil {
					log.Println("Error when getting tree for", entry.Name, ":", err)
					return nil, err
				}
				addPkg(entry.Name)
				err = dive(entry.Name, subTree)
				if err != nil {
					log.Println("Error when diving into", entry.Name, ":", err)
					return nil, err
				}
			}
		}

		return stdPkgs, nil
	}

	f, err := os.OpenFile(goversionCsv, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("Error when opening goversions.csv:", err)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	hash, err := getFileHash(f)
	if err != nil {
		fmt.Println("Error when getting file hash:", err)
		return
	}
	currentHash, err := getCurrentStdPkgHash()
	if err != nil {
		fmt.Println("Error when getting current hash:", err)
		return
	}

	if hash == currentHash {
		fmt.Println("No need to update " + stdpkgOutputFile)
		return
	}

	knownVersions, err := getCsvStoredGoversions(f)

	stdpkgsSet := map[string]struct{}{}

	for tag := range knownVersions {
		ps, err := collect(tag)
		if err != nil {
			fmt.Println("Error when collecting std pkgs for tag "+tag+":", err)
			return
		}
		for _, p := range ps {
			stdpkgsSet[p] = struct{}{}
		}
	}

	pkgs := make([]string, 0, len(stdpkgsSet))
	for pkg := range stdpkgsSet {
		pkgs = append(pkgs, pkg)
	}
	sort.Slice(pkgs, func(i, j int) bool {
		return pkgs[i] < pkgs[j]
	})

	// Generate the code.
	buf := bytes.NewBuffer(nil)

	err = packageTemplate.Execute(buf, struct {
		Timestamp time.Time
		StdPkg    []string
		Hash      string
	}{
		Timestamp: time.Now().UTC(),
		StdPkg:    pkgs,
		Hash:      hash,
	})
	if err != nil {
		fmt.Println("Error when generating the code:", err)
		return
	}

	writeOnDemand(buf.Bytes(), stdpkgOutputFile)

	fmt.Println("Generated " + stdpkgOutputFile)
}
