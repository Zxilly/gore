// Copyright (C) 2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gore_bsd

package gore

// GetGoRoot returns the Go Root path used to compile the binary.
func (f *GoFile) GetGoRoot() (string, error) {
	return "", ErrNotAGPLBuild
}
