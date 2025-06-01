package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

// FileInfo represents information about a hidden file
type FileInfo struct {
	Path     string    `json:"path"`
	Type     string    `json:"type"`
	Size     int64     `json:"size"`
	Modified time.Time `json:"modified"`
}

// ScanResult represents the complete scan results
type ScanResult struct {
	Directory    string     `json:"directory"`
	HiddenFiles  []FileInfo `json:"hidden_files"`
	TotalCount   int        `json:"total_count"`
	TotalSize    int64      `json:"total_size"`
	IgnoredCount int        `json:"ignored_count"`
	ErrorCount   int        `json:"error_count"`
	ScanTime     time.Time  `json:"scan_time"`
}

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

// shouldIgnore checks if a path should be ignored based on patterns
func shouldIgnore(path string, ignorePatterns []string) bool {
	for _, pattern := range ignorePatterns {
		matched, _ := filepath.Match(pattern, filepath.Base(path))
		if matched {
			return true
		}
	}
	return false
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

// validateDirectory checks if the given path exists and is a directory
func validateDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory does not exist: %s", path)
		}
		return fmt.Errorf("error accessing directory: %v", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", path)
	}

	return nil
}

// findHiddenFiles walks the directory tree and finds hidden files
func findHiddenFiles(root string, noColor bool, verbose bool, ignorePatterns []string, jsonOutput bool) error {
	count := 0
	totalSize := int64(0)
	ignored := 0
	errors := 0
	var hiddenFiles []FileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Don't stop walking on permission errors
			if os.IsPermission(err) {
				errors++
				if verbose && !jsonOutput {
					fmt.Fprintf(os.Stderr, "%sPermission denied: %s%s\n", ColorRed, path, ColorReset)
				}
				return nil
			}
			return err
		}

		// Skip if info is nil
		if info == nil {
			return nil
		}

		// Check if file/directory is hidden
		if strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			if shouldIgnore(path, ignorePatterns) {
				ignored++
				return nil
			}

			count++
			totalSize += info.Size()

			if jsonOutput {
				fileType := "file"
				if info.IsDir() {
					fileType = "directory"
				}
				hiddenFiles = append(hiddenFiles, FileInfo{
					Path:     path,
					Type:     fileType,
					Size:     info.Size(),
					Modified: info.ModTime(),
				})
			} else {
				printFileInfo(path, info, noColor, verbose)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	if jsonOutput {
		result := ScanResult{
			Directory:    root,
			HiddenFiles:  hiddenFiles,
			TotalCount:   count,
			TotalSize:    totalSize,
			IgnoredCount: ignored,
			ErrorCount:   errors,
			ScanTime:     time.Now(),
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(result)
	}

	// Print summary
	if noColor {
		fmt.Printf("\nTotal hidden items found: %d\n", count)
		if ignored > 0 {
			fmt.Printf("Ignored items: %d\n", ignored)
		}
		if errors > 0 {
			fmt.Printf("Permission errors: %d\n", errors)
		}
		if verbose {
			fmt.Printf("Total size: %s\n", formatSize(totalSize))
		}
	} else {
		fmt.Printf("\n%sTotal hidden items found: %s%d%s\n", ColorGreen, ColorYellow, count, ColorReset)
		if ignored > 0 {
			fmt.Printf("%sIgnored items: %s%d%s\n", ColorRed, ColorYellow, ignored, ColorReset)
		}
		if errors > 0 {
			fmt.Printf("%sPermission errors: %s%d%s\n", ColorRed, ColorYellow, errors, ColorReset)
		}
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
	var ignoreFlag string
	var jsonOutput bool

	flag.StringVar(&dir, "dir", ".", "Directory to search")
	flag.BoolVar(&noColor, "no-color", false, "Disable colored output")
	flag.BoolVar(&verbose, "v", false, "Verbose output (show size and modification time)")
	flag.StringVar(&ignoreFlag, "ignore", "", "Comma-separated list of patterns to ignore (e.g., '.git,.DS_Store')")
	flag.BoolVar(&jsonOutput, "json", false, "Output results in JSON format")
	flag.Parse()

	// Validate directory
	if err := validateDirectory(dir); err != nil {
		fmt.Fprintf(os.Stderr, "%sError: %v%s\n", ColorRed, err, ColorReset)
		os.Exit(1)
	}

	// Parse ignore patterns
	var ignorePatterns []string
	if ignoreFlag != "" {
		ignorePatterns = strings.Split(ignoreFlag, ",")
		// Trim whitespace from patterns
		for i := range ignorePatterns {
			ignorePatterns[i] = strings.TrimSpace(ignorePatterns[i])
		}
	}

	if !jsonOutput {
		fmt.Printf("Searching for hidden files in: %s\n", dir)
		if len(ignorePatterns) > 0 {
			fmt.Printf("Ignoring patterns: %v\n", ignorePatterns)
		}
		fmt.Println()
	}

	if err := findHiddenFiles(dir, noColor, verbose, ignorePatterns, jsonOutput); err != nil {
		fmt.Fprintf(os.Stderr, "%sError: %v%s\n", ColorRed, err, ColorReset)
		os.Exit(1)
	}
}
