package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	filepath.WalkDir(".", fileWalker)
	fmt.Println("Hello World")
}

func fileWalker(path string, d fs.DirEntry, err error) error {
	return nil
}
