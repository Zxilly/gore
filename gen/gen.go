// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run ./gen [stdpkgs|goversion|moduledata]")
		return
	}

	switch os.Args[1] {
	case "stdpkgs":
		generateStdPkgs()
	case "goversion":
		generateGoVersions()
	case "moduledata":
		generateModuleData()
	default:
		fmt.Println("go run ./gen [stdpkgs|goversion|moduledata]")
	}
}
