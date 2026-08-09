// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"bytes"
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
	if wasm, ok := f.fh.(*wasmFile); ok {
		return extractBuildInfoFromWasm(wasm)
	}

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

// These sentinels are used by cmd/go to delimit the module information string.
var (
	modInfoStart = []byte{0x30, 0x77, 0xaf, 0x0c, 0x92, 0x74, 0x08, 0x02, 0x41, 0xe1, 0xc1, 0x07, 0xe6, 0xd6, 0x18, 0xe6}
	modInfoEnd   = []byte{0xf9, 0x32, 0x43, 0x31, 0x86, 0x18, 0x20, 0x72, 0x00, 0x82, 0x42, 0x10, 0x41, 0x16, 0xd8, 0xf2}
)

func extractBuildInfoFromWasm(wasm *wasmFile) (*BuildInfo, error) {
	start := bytes.Index(wasm.memory, modInfoStart)
	if start < 0 {
		return nil, ErrNoBuildInfo
	}
	start += len(modInfoStart)

	end := bytes.Index(wasm.memory[start:], modInfoEnd)
	if end < 0 {
		return nil, ErrNoBuildInfo
	}
	if end == 0 {
		return nil, ErrNoBuildInfo
	}

	info, err := debug.ParseBuildInfo(string(wasm.memory[start : start+end]))
	if err != nil {
		return nil, fmt.Errorf("error parsing WebAssembly build information: %w", err)
	}
	version, _ := wasm.getVersion()
	if version != nil {
		info.GoVersion = version.Name
	}

	return &BuildInfo{Compiler: version, ModInfo: info}, nil
}
