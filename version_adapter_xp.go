//go:build !go1.11
// +build !go1.11

package main

import "io/ioutil"
import "io/fs"

func ReadFile(p string) ([]byte, error) {
	return ioutil.ReadFile(args.crt)
}

func ReadDir(p string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(args.crt)
}
