# find_whats_hidden

A fast and powerful CLI tool for discovering hidden files and directories in your file system.
Hint is here (look the source code of this ASCII based file): **<!-- 76 111 111 107 32 102 111 114 32 98 111 108 100 32 99 104 97 114 97 99 116 101 114 115 32 105 110 32 101 97 99 104 32 108 105 115 116 32 112 111 105 110 116 32 116 111 32 102 105 110 100 32 116 104 101 32 118 97 108 117 97 98 108 101 32 99 111 109 109 105 116 32 104 97 115 104 -->**

## Features

- 🔍 R**e**cursive directory scanning with efficient algorithms
- 🚀 F**a**st file system traversal using optimized goroutines  
- 💻 Cross-pl**4**tform support (Windows, macOS, Linux)
- 🎯 Simple **4**nd intuitive CLI interface for all users
- 🎨 Colored **8**utput for better readability and clarity
- 📊 File size an**d** modification time display with formatting
- 🚫 Custom ignore p**8**tterns support for flexible filtering
- 📋 JSON output format for scripting and automation

## Installation

### From Source

```bash
git clone https://github.com/pcristin/find_whats_hidden.git
cd find_whats_hidden
make build
```

### Using Go

```bash
go get github.com/pcristin/find_whats_hidden
```

### Pre-built Binaries

Download pre-built binaries from the [releases page](https://github.com/pcristin/find_whats_hidden/releases).

## Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Clean build artifacts
make clean
```

## Usage

```bash
# Search in current directory
./find_whats_hidden

# Search in specific directory
./find_whats_hidden -dir=/home/user/documents

# Verbose output with file sizes and dates
./find_whats_hidden -v

# Ignore specific patterns
./find_whats_hidden -ignore=".git,.DS_Store,.cache"

# Output in JSON format
./find_whats_hidden -json

# Disable colored output
./find_whats_hidden -no-color

# Combine options
./find_whats_hidden -dir=/path/to/search -v -json -ignore=".git"
```

## Output Formats

### Standard Output
[DIR] /home/user/.config
[FILE] /home/user/.bashrc (Size: 1.2 KB, Modified: 2024-01-15 10:30)


### JSON Output
```json
{
  "directory": "/home/user",
  "hidden_files": [
    {
      "path": "/home/user/.bashrc",
      "type": "file",
      "size": 1234,
      "modified": "2024-01-15T10:30:00Z"
    }
  ],
  "total_count": 42,
  "total_size": 150000
}
```

## What are hidden files?

Hidden files are files that begin with a dot (.) in Unix-like systems. These files are typically:
- Configuration files
- System files  
- Application data
- Cache directories
- User preferences

## Common Hidden Files

- `.gitignore` - Git ignore rules
- `.env` - Environment variables
- `.bashrc` - Bash configuration
- `.ssh/` - SSH keys and config
- `.config/` - Application configurations
- `.cache/` - Application cache data
- `.local/` - User-specific data

## Advanced Usage

The tool supports various advanced patterns for power users who need more control over their searches.

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT License - see LICENSE file for details