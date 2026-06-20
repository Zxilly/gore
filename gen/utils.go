// Copyright (C) 2019-2026 GoRE Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// diffCode returns false if a and b have different other than the date.
func diffCode(a, b string) bool {
	if a == b {
		return false
	}

	aLines := strings.Split(a, "\n")
	bLines := strings.Split(b, "\n")

	// ignore the license and the date
	aLines = aLines[21:]
	bLines = bLines[21:]

	if len(aLines) != len(bLines) {
		return true
	}

	for i := 0; i < len(aLines); i++ {
		if aLines[i] != bLines[i] {
			return true
		}
	}

	return false
}

func writeOnDemand(new []byte, target string) {
	old, err := os.ReadFile(target)
	if err != nil {
		fmt.Println("Error when reading the old file:", target, err)
		return
	}

	old, _ = format.Source(old)
	new, _ = format.Source(new)

	// Compare the old and the new.
	if !diffCode(string(old), string(new)) {
		fmt.Println(target + " no changes.")
		return
	}

	fmt.Println(target + " changes detected.")

	// Write the new file.
	err = os.WriteFile(target, new, 0664)
	if err != nil {
		fmt.Println("Error when writing the new file:", err)
		return
	}
}

func getSourceDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return filepath.Join(filepath.Dir(filename), "..")
}

func getCsvStoredGoversions(f *os.File) (map[string]*goversion, error) {
	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	vers := make(map[string]*goversion)
	c, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range c {
		vers[line[0]] = &goversion{Name: line[0], Sha: line[1], Date: line[2]}
	}

	return vers, err
}

func getFileHash(f *os.File) (string, error) {
	h := sha256.New()
	_, err := io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
