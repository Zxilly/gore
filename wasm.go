// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"bytes"
	"debug/dwarf"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/eliben/watgo"
	"github.com/eliben/watgo/wasmir"
)

const (
	wasmPageSize = 64 * 1024
	// Eager reconstruction must be bounded independently of the input file size.
	maxWasmMemorySize = 256 * 1024 * 1024
)

// WasmModule is watgo's semantic representation of a WebAssembly module.
type WasmModule = wasmir.Module

// WasmInfo is returned by GoFile.GetParsedFile for WebAssembly binaries.
type WasmInfo struct {
	Module *wasmir.Module
	Memory []byte
}

func openWasm(r io.ReaderAt) (*wasmFile, error) {
	data, err := io.ReadAll(io.NewSectionReader(r, 0, int64(^uint64(0)>>1)))
	if err != nil {
		return nil, fmt.Errorf("error when reading the WebAssembly file: %w", err)
	}

	module, err := watgo.DecodeWASM(data)
	if err != nil {
		return nil, fmt.Errorf("error when parsing the WebAssembly file: %w", err)
	}

	memory, err := buildWasmMemory(module)
	if err != nil {
		return nil, err
	}

	return &wasmFile{module: module, memory: memory, reader: r}, nil
}

func buildWasmMemory(module *wasmir.Module) ([]byte, error) {
	if module == nil || len(module.Memories) == 0 {
		return nil, errors.New("WebAssembly module has no linear memory")
	}

	minimumPages := module.Memories[0].Min
	if minimumPages > maxWasmMemorySize/wasmPageSize {
		return nil, fmt.Errorf("WebAssembly minimum memory size exceeds the %d-byte reconstruction limit", maxWasmMemorySize)
	}
	maxInt := int(^uint(0) >> 1)
	if minimumPages > uint64(maxInt)/wasmPageSize {
		return nil, errors.New("WebAssembly minimum memory size is too large")
	}
	memory := make([]byte, int(minimumPages*wasmPageSize))

	for i, segment := range module.Data {
		if segment.Mode != wasmir.DataSegmentModeActive {
			continue
		}
		if segment.MemoryIndex != 0 {
			return nil, fmt.Errorf("WebAssembly data segment %d targets unsupported memory %d", i, segment.MemoryIndex)
		}

		offset, err := wasmConstExpression(module, segment.OffsetExpr, segment.OffsetI64, nil)
		if err != nil {
			return nil, fmt.Errorf("WebAssembly data segment %d has invalid offset: %w", i, err)
		}
		if offset < 0 || uint64(offset) > uint64(len(memory)) || uint64(len(segment.Init)) > uint64(len(memory))-uint64(offset) {
			return nil, fmt.Errorf("WebAssembly data segment %d is outside linear memory", i)
		}
		copy(memory[int(offset):], segment.Init)
	}

	return memory, nil
}

func wasmConstExpression(module *wasmir.Module, expression []wasmir.Instruction, fallback int64, visiting map[uint32]bool) (int64, error) {
	if len(expression) == 0 {
		return fallback, nil
	}

	var value int64
	valueSet := false
	for _, instruction := range expression {
		switch instruction.Kind {
		case wasmir.InstrI32Const:
			if valueSet {
				return 0, errors.New("constant expression produces multiple values")
			}
			value = int64(instruction.I32Const)
			valueSet = true
		case wasmir.InstrI64Const:
			if valueSet {
				return 0, errors.New("constant expression produces multiple values")
			}
			value = instruction.I64Const
			valueSet = true
		case wasmir.InstrGlobalGet:
			if valueSet {
				return 0, errors.New("constant expression produces multiple values")
			}
			index := instruction.GlobalIndex
			if int(index) >= len(module.Globals) {
				return 0, fmt.Errorf("global index %d is out of bounds", index)
			}
			global := module.Globals[index]
			if global.Mutable || len(global.Init) == 0 {
				return 0, fmt.Errorf("global %d cannot be evaluated statically", index)
			}
			if visiting == nil {
				visiting = make(map[uint32]bool)
			}
			if visiting[index] {
				return 0, fmt.Errorf("global %d has a cyclic initializer", index)
			}
			visiting[index] = true
			globalValue, err := wasmConstExpression(module, global.Init, 0, visiting)
			delete(visiting, index)
			if err != nil {
				return 0, err
			}
			value = globalValue
			valueSet = true
		case wasmir.InstrEnd:
			// The binary constant expression terminator carries no value.
		default:
			return 0, fmt.Errorf("unsupported instruction kind %d", instruction.Kind)
		}
	}

	if !valueSet {
		return 0, errors.New("constant expression produces no value")
	}
	return value, nil
}

var _ fileHandler = (*wasmFile)(nil)

type wasmFile struct {
	module *wasmir.Module
	memory []byte
	reader io.ReaderAt
}

func (w *wasmFile) Close() error {
	return tryClose(w.reader)
}

func (w *wasmFile) getSymbol(string) (Symbol, error) {
	return Symbol{}, ErrSymbolNotFound
}

func (w *wasmFile) getRData() ([]byte, error) {
	return w.memory, nil
}

func (w *wasmFile) getCodeSection() (uint64, []byte, error) {
	return 0, nil, errors.New("WebAssembly code is not stored in linear memory")
}

func (w *wasmFile) getSectionDataFromAddress(address uint64) (uint64, []byte, error) {
	if address >= uint64(len(w.memory)) {
		return 0, nil, fmt.Errorf("address %#x is outside WebAssembly linear memory", address)
	}
	return 0, w.memory, nil
}

func (w *wasmFile) getSectionData(name string) (uint64, []byte, error) {
	if name != "memory" {
		return 0, nil, ErrSectionDoesNotExist
	}
	return 0, w.memory, nil
}

func (w *wasmFile) getVersion() (*GoVersion, error) {
	section := w.customSection("producers")
	if section == nil {
		return nil, errors.New("WebAssembly producers section does not exist")
	}

	version, err := goVersionFromProducers(section)
	if err != nil {
		return nil, err
	}
	if resolved := ResolveGoVersion(version); resolved != nil {
		return resolved, nil
	}
	return &GoVersion{Name: version}, nil
}

func goVersionFromProducers(data []byte) (string, error) {
	r := bytes.NewReader(data)
	fieldCount, err := binary.ReadUvarint(r)
	if err != nil {
		return "", fmt.Errorf("invalid WebAssembly producers field count: %w", err)
	}

	for i := uint64(0); i < fieldCount; i++ {
		fieldName, err := readWasmName(r)
		if err != nil {
			return "", fmt.Errorf("invalid WebAssembly producers field %d: %w", i, err)
		}
		valueCount, err := binary.ReadUvarint(r)
		if err != nil {
			return "", fmt.Errorf("invalid WebAssembly producers value count: %w", err)
		}
		for j := uint64(0); j < valueCount; j++ {
			name, err := readWasmName(r)
			if err != nil {
				return "", fmt.Errorf("invalid WebAssembly producer name: %w", err)
			}
			version, err := readWasmName(r)
			if err != nil {
				return "", fmt.Errorf("invalid WebAssembly producer version: %w", err)
			}
			if fieldName == "language" && name == "Go" {
				return version, nil
			}
		}
	}

	return "", errors.New("WebAssembly producers section has no Go language entry")
}

func readWasmName(r *bytes.Reader) (string, error) {
	length, err := binary.ReadUvarint(r)
	if err != nil {
		return "", err
	}
	if length > uint64(r.Len()) {
		return "", io.ErrUnexpectedEOF
	}
	data := make([]byte, int(length))
	if _, err := io.ReadFull(r, data); err != nil {
		return "", err
	}
	return string(data), nil
}

func (w *wasmFile) getFileInfo() *FileInfo {
	version, _ := w.getVersion()
	return &FileInfo{
		Arch:      ArchWASM,
		OS:        wasmOS(w.module),
		ByteOrder: binary.LittleEndian,
		WordSize:  intSize64,
		goversion: version,
	}
}

func wasmOS(module *wasmir.Module) string {
	for _, imported := range module.Imports {
		switch imported.Module {
		case "go", "gojs":
			return "js"
		case "wasi_snapshot_preview1":
			return "wasip1"
		}
	}
	return "wasm"
}

func (w *wasmFile) getPCLNTABData() (uint64, []byte, error) {
	data, err := searchSectionForTab(w.memory, binary.LittleEndian)
	if err != nil {
		return 0, nil, err
	}
	return uint64(len(w.memory) - len(data)), data, nil
}

func (w *wasmFile) moduledataSection() string {
	return "memory"
}

func (w *wasmFile) getBuildID() (string, error) {
	section := w.customSection("go:buildid")
	if section == nil {
		return "", nil
	}
	return string(section), nil
}

func (w *wasmFile) getReader() io.ReaderAt {
	return w.reader
}

func (w *wasmFile) getParsedFile() any {
	return WasmInfo{Module: w.module, Memory: w.memory}
}

func (w *wasmFile) getDwarf() (*dwarf.Data, error) {
	return nil, errors.New("DWARF is not supported for WebAssembly binaries")
}

func (w *wasmFile) customSection(name string) []byte {
	for _, section := range w.module.CustomSections {
		if section.Name == name {
			return section.Payload
		}
	}
	return nil
}
