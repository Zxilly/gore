// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringer(t *testing.T) {
	assert := assert.New(t)

	// Setup
	tests := []struct {
		name     string
		obj      fmt.Stringer
		expected string
	}{
		{"main", &Function{Name: "main", PackageName: "main"}, "main"},
		{"setup", &Function{Name: "setup", PackageName: "main"}, "setup"},
		{"method", &Method{Receiver: "(T)", Function: &Function{Name: "MyMethod"}}, "(T)MyMethod"},
	}

	for _, test := range tests {
		t.Run("stringer_"+test.name, func(t *testing.T) {
			assert.Equal(test.expected, test.obj.String())
		})
	}

}
