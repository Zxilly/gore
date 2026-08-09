// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

type record struct {
	Name  string
	Value int
}

var sample = record{Name: "gore-wasm", Value: 42}

func main() {
	println(sample.Name, sample.Value)
}
