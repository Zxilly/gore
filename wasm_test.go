// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"encoding/binary"
	"testing"

	"github.com/eliben/watgo/wasmir"
	"github.com/stretchr/testify/require"
)

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
	require.ErrorIs(t, err, ErrWasmMemoryTooLarge)
}

func TestBuildWasmMemoryRequiresLinearMemory(t *testing.T) {
	_, err := buildWasmMemory(&wasmir.Module{})
	require.ErrorIs(t, err, ErrWasmNoLinearMemory)
}

func TestWasmConstExpressionErrors(t *testing.T) {
	module := &wasmir.Module{}
	_, err := wasmConstExpression(module, []wasmir.Instruction{
		{Kind: wasmir.InstrI32Const},
		{Kind: wasmir.InstrI64Const},
	}, 0, nil)
	require.ErrorIs(t, err, ErrWasmConstExpressionMultipleValues)

	_, err = wasmConstExpression(module, []wasmir.Instruction{{Kind: wasmir.InstrEnd}}, 0, nil)
	require.ErrorIs(t, err, ErrWasmConstExpressionNoValue)
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

func TestWasmVersionErrors(t *testing.T) {
	_, err := (&wasmFile{module: &wasmir.Module{}}).getVersion()
	require.ErrorIs(t, err, ErrNoGoVersionFound)

	data := binary.AppendUvarint(nil, 0)
	_, err = goVersionFromProducers(data)
	require.ErrorIs(t, err, ErrNoGoVersionFound)
}

func TestWasmDwarfUnsupported(t *testing.T) {
	_, err := (&wasmFile{}).getDwarf()
	require.ErrorIs(t, err, ErrUnsupportedDwarf)
}

func appendWasmName(dst []byte, value string) []byte {
	dst = binary.AppendUvarint(dst, uint64(len(value)))
	return append(dst, value...)
}
