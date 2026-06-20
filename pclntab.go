// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"bytes"
	"encoding/binary"
)

// keep sync with debug/gosym/pclntab.go
const (
	gopclntab12magic  uint32 = 0xfffffffb
	gopclntab116magic uint32 = 0xfffffffa
	gopclntab118magic uint32 = 0xfffffff0
	gopclntab120magic uint32 = 0xfffffff1
)

// searchSectionForTab looks for the PCLN table within the section.
func searchSectionForTab(secData []byte, order binary.ByteOrder) ([]byte, error) {
	// First check for the current magic used. If this fails, it could be
	// an older version. So check for the old header.
MagicLoop:
	for _, magic := range []uint32{gopclntab120magic, gopclntab118magic, gopclntab116magic, gopclntab12magic} {
		bMagic := make([]byte, 6) // 4 bytes for the magic, 2 bytes for padding.
		order.PutUint32(bMagic, magic)

		off := bytes.LastIndex(secData, bMagic)
		if off == -1 {
			continue // Try other magic.
		}
		for off != -1 {
			if off != 0 {
				buf := secData[off:]
				if len(buf) < 16 || buf[4] != 0 || buf[5] != 0 ||
					(buf[6] != 1 && buf[6] != 2 && buf[6] != 4) || // pc quantum
					(buf[7] != 4 && buf[7] != 8) { // pointer size
					// Header doesn't match.
					if off-1 <= 0 {
						continue MagicLoop
					}
					off = bytes.LastIndex(secData[:off-1], bMagic)
					continue
				}
				// Header match
				return secData[off:], nil
			}
			break
		}
	}
	return nil, ErrNoPCLNTab
}
