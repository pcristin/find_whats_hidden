package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findHiddenFiles(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if file/directory is hidden
		if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			fmt.Printf("Found hidden: %s\n", path)
		}

		return nil
	})
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "Directory to search")
	flag.Parse()

	fmt.Printf("Searching for hidden files in: %s\n", dir)

	if err := findHiddenFiles(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
