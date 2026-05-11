// Copyright (C) 2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gore_bsd

package gore

func getGoRootFromDwarf(fh fileHandler) (string, bool)       { return "", false }
func getBuildVersionFromDwarf(fh fileHandler) (string, bool) { return "", false }
