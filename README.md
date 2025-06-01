# find_whats_hidden

A fast and efficient CLI tool for discovering hidden files and directories in your file system.

## Features

- 🔍 Recursive directory scanning
- 🚀 Fast file system traversal
- 💻 Cross-platform support (Windows, macOS, Linux)
- 🎯 Simple and intuitive CLI interface
- 🎨 Colored output for better readability
- 📊 File size and modification time display
- 🚫 Custom ignore patterns support

## Installation

```bash
go get github.com/pcristin/find_whats_hidden
```

Or build from source:

```bash
git clone https://github.com/pcristin/find_whats_hidden.git
cd find_whats_hidden
go build -o find_whats_hidden
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

# Disable colored output
./find_whats_hidden -no-color

# Using go run
go run main.go -dir=/path/to/search -v -ignore=".git"
```

## What are hidden files?

Hidden files are files that begin with a dot (.) in Unix-like systems. These files are typically:
- Configuration files
- System files
- Application data

## Common Hidden Files

- `.gitignore` - Git ignore rules
- `.env` - Environment variables
- `.bashrc` - Bash configuration
- `.ssh/` - SSH keys and config
- `.config/` - Application configurations

## Contributing

Pull requests are welcome! For major changes, please open an issue first.

## License

MIT