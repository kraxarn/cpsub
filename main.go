package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func copyFile(from, to string) error {
	inFile, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(to, inFile, 0644)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: cpsub <source> <target> <extension>")
		return
	}
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Full path of destination file
		dest := strings.ReplaceAll(path, os.Args[1], "")
		fullDest := filepath.Join(os.Args[2], dest)
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
	if err != nil {
		fmt.Println(err)
	}
}