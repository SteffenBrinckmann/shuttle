//go:build go1.11
// +build go1.11

package main

import (
	"os"
)

func ReadFile(p string) ([]byte, error) {
	return os.ReadFile(args.crt)
}

func ReadDir(p string) ([]os.DirEntry, error) {
	return os.ReadDir(args.crt)
}
