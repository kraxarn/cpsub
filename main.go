package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Copies file from source to target
func copyFile(from, to string) error {
	// Read source
	inFile, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}
	// Write target
	return ioutil.WriteFile(to, inFile, 0644)
}

func main() {
	// Expect (at least) 4 arguments
	if len(os.Args) < 4 {
		fmt.Println("usage: cpsub <source> <target> <extension>")
		return
	}
	// Recursively walk source directory
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		// Something went wrong
		if err != nil {
			return err
		}
		// Full path of destination file
		dest := strings.ReplaceAll(path, os.Args[1], "")
		fullDest := filepath.Join(os.Args[2], dest)
		// Create directory in target directory
		if info.IsDir() {
			if err = os.MkdirAll(fullDest, 0755);  err != nil {
				return err
			}
			return nil
		}
		if !strings.HasSuffix(info.Name(), os.Args[3]) {
			return nil
		}
		// Copy file to new directory
		fmt.Println(dest)
		if err := copyFile(path, fullDest); err != nil {
			return err
		}
		return nil
	})
	// Print if something went wrong
	if err != nil {
		fmt.Println(err)
	}
}