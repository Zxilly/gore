// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"debug/buildinfo"
	"errors"
	"fmt"
	"runtime/debug"
)

var (
	// ErrNoBuildInfo is returned if the file has no build information available.
	ErrNoBuildInfo = errors.New("no build info available")
)

// BuildInfo that was extracted from the file.
type BuildInfo struct {
	// Compiler version. Can be nil.
	Compiler *GoVersion
	// ModInfo holds information about the Go modules in this file.
	// Can be nil.
	ModInfo *debug.BuildInfo
}

func (f *GoFile) extractBuildInfo() (*BuildInfo, error) {
	info, err := buildinfo.Read(f.fh.getReader())
	if err != nil {
		return nil, fmt.Errorf("error when extracting build information: %w", err)
	}

	result := &BuildInfo{
		Compiler: ResolveGoVersion(info.GoVersion),
		ModInfo:  info,
	}

	return result, nil
}
