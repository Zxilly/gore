// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gore_bsd && slow_test

package gore

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDwarfString(t *testing.T) {
	noStrip := false
	getMatrix(t, nil, &noStrip, "dwarfString", func(t *testing.T, exe string) {
		r := require.New(t)

		f, err := Open(exe)
		r.NoError(err)
		r.NotNil(f)
		defer f.Close()

		gover, ok := getBuildVersionFromDwarf(f.fh)
		r.True(ok)
		r.Equal(gover, runtime.Version())

		goroot, ok := getGoRootFromDwarf(f.fh)
		r.True(ok)
		r.Equal(goroot, runtime.GOROOT())
	})
}
