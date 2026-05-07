// Copyright (C) 2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !gore_agpl

package gore

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractGoRoot(t *testing.T) {
	goldFiles, err := getGoldenResources()
	if err != nil || len(goldFiles) == 0 {
		// Golden folder does not exist
		t.Skip("No golden files")
	}

	const expectGoRoot = "/usr/local/go"

	for _, test := range goldFiles {
		t.Run("get goroot form "+test, func(t *testing.T) {
			r := require.New(t)

			// TODO: Remove this check when arm support has been added.
			if strings.Contains(test, "arm64") {
				t.Skip("ARM currently not supported")
			}

			fp, err := getTestResourcePath("gold/" + test)
			r.NoError(err, "Failed to get path to resource")
			if _, err = os.Stat(fp); os.IsNotExist(err) {
				// Skip this file because it doesn't exist
				// t.Skip will cause the parent test to be skipped.
				fmt.Printf("[SKIPPING TEST] golden fille %s does not exist\n", test)
				return
			}
			r.NoError(err)
			f, err := Open(fp)
			r.NoError(err)
			defer f.Close()
			goroot, err := f.GetGoRoot()
			r.Error(err)
			r.Equal(err, ErrNotAGPLBuild)
			r.Equal(goroot, "")
		})
	}
}
