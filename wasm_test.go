// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"bytes"
	"encoding/binary"
	goversion "go/version"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/eliben/watgo/wasmir"
	"github.com/stretchr/testify/require"
)

func TestWasmOpen(t *testing.T) {
	for _, goos := range []string{"js", "wasip1"} {
		t.Run(goos, func(t *testing.T) {
			testWasmOpen(t, goos)
		})
	}
}

func testWasmOpen(t *testing.T, goos string) {
	if testing.Short() {
		t.Skip("building a WebAssembly fixture")
	}

	wasmPath := filepath.Join(t.TempDir(), "gore-test.wasm")
	cmd := exec.Command("go", "build", "-o", wasmPath, "./testdata/wasm")
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH=wasm", "CGO_ENABLED=0")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))

	file, err := Open(wasmPath)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, file.Close()) })

	require.Equal(t, ArchWASM, file.FileInfo.Arch)
	require.Equal(t, goos, file.FileInfo.OS)
	require.Equal(t, intSize64, file.FileInfo.WordSize)
	require.NotEmpty(t, file.BuildID)

	version, err := file.GetCompilerVersion()
	require.NoError(t, err)
	require.NotNil(t, version)
	require.True(t, goversion.IsValid(version.Name), "invalid Go version %q", version.Name)

	require.NotNil(t, file.BuildInfo)
	require.NotNil(t, file.BuildInfo.ModInfo)
	require.Equal(t, version.Name, file.BuildInfo.ModInfo.GoVersion)

	packages, err := file.GetPackages()
	require.NoError(t, err)
	require.NotEmpty(t, packages)

	standardLibrary, err := file.GetSTDLib()
	require.NoError(t, err)
	require.NotEmpty(t, standardLibrary)

	types, err := file.GetTypes()
	require.NoError(t, err)
	require.NotEmpty(t, types)

	parsed, ok := file.GetParsedFile().(WasmInfo)
	require.True(t, ok)
	require.NotNil(t, parsed.Module)
	require.NotEmpty(t, parsed.Memory)

	_, err = file.fh.getDwarf()
	require.ErrorContains(t, err, "DWARF is not supported for WebAssembly binaries")

	raw, err := os.ReadFile(wasmPath)
	require.NoError(t, err)
	readerFile, err := OpenReader(bytes.NewReader(raw))
	require.NoError(t, err)
	require.NoError(t, readerFile.Close())
}

func TestBuildWasmMemory(t *testing.T) {
	module := &wasmir.Module{
		Memories: []wasmir.Memory{{Min: 1}},
		Data: []wasmir.DataSegment{
			{
				Mode:       wasmir.DataSegmentModeActive,
				OffsetExpr: []wasmir.Instruction{{Kind: wasmir.InstrI32Const, I32Const: 4}, {Kind: wasmir.InstrEnd}},
				Init:       []byte("gore"),
			},
			{
				Mode:      wasmir.DataSegmentModeActive,
				OffsetI64: 16,
				Init:      []byte("wasm"),
			},
			{
				Mode: wasmir.DataSegmentModePassive,
				Init: []byte("ignored"),
			},
		},
	}

	memory, err := buildWasmMemory(module)
	require.NoError(t, err)
	require.Len(t, memory, wasmPageSize)
	require.Equal(t, []byte("gore"), memory[4:8])
	require.Equal(t, []byte("wasm"), memory[16:20])
}

func TestBuildWasmMemoryRejectsOutOfBoundsSegment(t *testing.T) {
	module := &wasmir.Module{
		Memories: []wasmir.Memory{{Min: 1}},
		Data: []wasmir.DataSegment{{
			Mode:       wasmir.DataSegmentModeActive,
			OffsetExpr: []wasmir.Instruction{{Kind: wasmir.InstrI32Const, I32Const: wasmPageSize - 1}, {Kind: wasmir.InstrEnd}},
			Init:       []byte("too large"),
		}},
	}

	_, err := buildWasmMemory(module)
	require.ErrorContains(t, err, "outside linear memory")
}

func TestBuildWasmMemoryRejectsExcessiveMinimum(t *testing.T) {
	module := &wasmir.Module{
		Memories: []wasmir.Memory{{Min: maxWasmMemorySize/wasmPageSize + 1}},
	}

	_, err := buildWasmMemory(module)
	require.ErrorContains(t, err, "reconstruction limit")
}

func TestGoVersionFromProducers(t *testing.T) {
	var data []byte
	data = binary.AppendUvarint(data, 2)
	data = appendWasmName(data, "processed-by")
	data = binary.AppendUvarint(data, 1)
	data = appendWasmName(data, "some tool")
	data = appendWasmName(data, "1.0")
	data = appendWasmName(data, "language")
	data = binary.AppendUvarint(data, 2)
	data = appendWasmName(data, "Rust")
	data = appendWasmName(data, "1.90")
	data = appendWasmName(data, "Go")
	data = appendWasmName(data, "go1.26.1")

	version, err := goVersionFromProducers(data)
	require.NoError(t, err)
	require.Equal(t, "go1.26.1", version)
}

func appendWasmName(dst []byte, value string) []byte {
	dst = binary.AppendUvarint(dst, uint64(len(value)))
	return append(dst, value...)
}
