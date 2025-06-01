package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
)

func findHiddenFiles(root string, noColor bool) error {
	count := 0

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if file/directory is hidden
		if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			count++

			if info.IsDir() {
				if noColor {
					fmt.Printf("[DIR]  %s\n", path)
				} else {
					fmt.Printf("%s[DIR]%s  %s%s%s\n", ColorBlue, ColorReset, ColorYellow, path, ColorReset)
				}
			} else {
				if noColor {
					fmt.Printf("[FILE] %s\n", path)
				} else {
					fmt.Printf("%s[FILE]%s %s%s%s\n", ColorGreen, ColorReset, ColorPurple, path, ColorReset)
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if noColor {
		fmt.Printf("\nTotal hidden items found: %d\n", count)
	} else {
		fmt.Printf("\n%sTotal hidden items found: %s%d%s\n", ColorGreen, ColorYellow, count, ColorReset)
	}

	return nil
}

func main() {
	var dir string
	var noColor bool

	flag.StringVar(&dir, "dir", ".", "Directory to search")
	flag.BoolVar(&noColor, "no-color", false, "Disable colored output")
	flag.Parse()

	fmt.Printf("Searching for hidden files in: %s\n\n", dir)

	if err := findHiddenFiles(dir, noColor); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
