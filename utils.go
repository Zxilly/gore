// Copyright (C) 2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gore

import "io"

func tryClose(r io.ReaderAt) error {
	if c, ok := r.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
