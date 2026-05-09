// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import "errors"

var (
	// ErrNotEnoughBytesRead is returned if read call returned less bytes than what is needed.
	ErrNotEnoughBytesRead = errors.New("not enough bytes read")
	// ErrUnsupportedFile is returned if the file process is unsupported.
	ErrUnsupportedFile = errors.New("unsupported file")
	// ErrSectionDoesNotExist is returned when accessing a section that does not exist.
	ErrSectionDoesNotExist = errors.New("section does not exist")
	// ErrNoGoVersionFound is returned if no goversion was found in the binary.
	ErrNoGoVersionFound = errors.New("no goversion found")
	// ErrNoPCLNTab is returned if no PCLN table can be located.
	ErrNoPCLNTab = errors.New("no pclntab located")
	// ErrInvalidGoVersion is returned if the go version set for the file is either invalid
	// or does not match a known version by the library.
	ErrInvalidGoVersion = errors.New("invalid go version")
	// ErrNoGoRootFound is returned if no goroot was found in the binary.
	ErrNoGoRootFound = errors.New("no goroot found")
	// ErrNotAGPLBuild is returned for a functionality that is not supported in the BSD licensed
	// version of this library. To use this functionality, compile using the build tag "gore_agpl"
	// to enabled the AGPL licensed code.
	ErrNotAGPLBuild = errors.New("not a AGPL build")
)
