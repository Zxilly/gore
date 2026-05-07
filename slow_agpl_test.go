// This file is part of GoRE.
//
// Copyright (C) 2019-2021 GoRE Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

//go:build slow_test && gore_agpl

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
