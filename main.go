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
	ColorCyan   = "\033[36m"
)

// formatSize converts bytes to human readable format
func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// printFileInfo prints formatted file information
func printFileInfo(path string, info os.FileInfo, noColor bool, verbose bool) {
	if info.IsDir() {
		if noColor {
			fmt.Printf("[DIR]  %s", path)
		} else {
			fmt.Printf("%s[DIR]%s  %s%s%s", ColorBlue, ColorReset, ColorYellow, path, ColorReset)
		}
	} else {
		if noColor {
			fmt.Printf("[FILE] %s", path)
		} else {
			fmt.Printf("%s[FILE]%s %s%s%s", ColorGreen, ColorReset, ColorPurple, path, ColorReset)
		}
	}

	if verbose {
		modTime := info.ModTime().Format("2006-01-02 15:04")
		size := formatSize(info.Size())
		if noColor {
			fmt.Printf(" (Size: %s, Modified: %s)", size, modTime)
		} else {
			fmt.Printf(" %s(Size: %s, Modified: %s)%s", ColorCyan, size, modTime, ColorReset)
		}
	}

	fmt.Println()
}

// findHiddenFiles walks the directory tree and finds hidden files
func findHiddenFiles(root string, noColor bool, verbose bool) error {
	count := 0
	totalSize := int64(0)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if file/directory is hidden
		if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			count++
			totalSize += info.Size()
			// https://imgur.com/a/JNEhhwT
			printFileInfo(path, info, noColor, verbose)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Print summary
	if noColor {
		fmt.Printf("\nTotal hidden items found: %d\n", count)
		if verbose {
			fmt.Printf("Total size: %s\n", formatSize(totalSize))
		}
	} else {
		fmt.Printf("\n%sTotal hidden items found: %s%d%s\n", ColorGreen, ColorYellow, count, ColorReset)
		if verbose {
			fmt.Printf("%sTotal size: %s%s%s\n", ColorGreen, ColorYellow, formatSize(totalSize), ColorReset)
		}
	}

	return nil
}

func main() {
	var dir string
	var noColor bool
	var verbose bool

	flag.StringVar(&dir, "dir", ".", "Directory to search")
	flag.BoolVar(&noColor, "no-color", false, "Disable colored output")
	flag.BoolVar(&verbose, "v", false, "Verbose output (show size and modification time)")
	flag.Parse()

	fmt.Printf("Searching for hidden files in: %s\n\n", dir)

	if err := findHiddenFiles(dir, noColor, verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
