# NiFi Parser

A lightweight CLI tool for parsing and analyzing Apache NiFi flow configuration files.

## Overview

NiFi Parser is a command-line utility written in Go that helps you quickly read and analyze NiFi flow configurations. It provides a clean, structured view of your NiFi flows, processors, connections, and statistics.

## Features

- Parse JSON-formatted NiFi flow configuration files
- Display comprehensive flow summaries
- View detailed processor information including properties and status
- Analyze flow connections and relationships
- Generate flow statistics (running/stopped processors, connection counts)
- Fast and lightweight with minimal dependencies
- Well-tested with comprehensive unit test coverage

## Project Structure

```
nifiparser/
├── cmd/
│   └── nifiparser/       # Main CLI application entry point
│       └── main.go
├── internal/             # Private application code
│   ├── config/           # Data models for NiFi flows
│   │   └── models.go
│   ├── parser/           # File parsing logic
│   │   ├── parser.go
│   │   └── parser_test.go
│   └── display/          # Output formatting and display
│       ├── display.go
│       └── display_test.go
├── testdata/             # Test fixtures
│   └── sample_nifi_config.json
├── .vscode/              # VS Code configuration
│   └── settings.json
├── .golangci.yml         # Linter configuration
├── gopls.toml            # LSP server configuration
├── go.mod                # Go module definition
├── plan.md               # Project planning document
└── README.md             # This file
```

## Installation

### Prerequisites

- Go 1.21 or higher
- (Optional) golangci-lint for development

### From Source

```bash
# Clone the repository
git clone https://github.com/ghawsshafadonia/nifiparser.git
cd nifiparser

# Build the binary
go build -o nifiparser ./cmd/nifiparser

# Run the tool
./nifiparser -file testdata/sample_nifi_config.json
```

### Using Go Install

```bash
go install github.com/ghawsshafadonia/nifiparser/cmd/nifiparser@latest
```

## Usage

### Basic Usage

Display complete flow information:

```bash
nifiparser -file /path/to/nifi_config.json
```

### Command-line Options

- `-file` (required): Path to NiFi configuration JSON file
- `-processors`: Show detailed processor information only
- `-connections`: Show connection information only
- `-stats`: Show flow statistics only
- `-version`: Show version information

### Examples

**Show only processors:**
```bash
nifiparser -file config.json -processors
```

**Show only statistics:**
```bash
nifiparser -file config.json -stats
```

**Show connections:**
```bash
nifiparser -file config.json -connections
```

**Check version:**
```bash
nifiparser -version
```

### Sample Output

```
============================================================
NiFi Flow Summary: Sample Data Processing Flow
============================================================
Flow ID:       a1b2c3d4-e5f6-7890-abcd-ef1234567890
Version:       1.0.0
Processors:    3
Connections:   2

Processors:
------------------------------------------------------------
1. GetFile (ID: proc-001)
   Type:   org.apache.nifi.processors.standard.GetFile
   Status: running
   Properties:
     - Input Directory: /data/input
     - Keep Source File: false

...
```

## Development

### Setting Up Development Environment

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ghawsshafadonia/nifiparser.git
   cd nifiparser
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Install development tools:**
   ```bash
   # Install golangci-lint
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   
   # Verify installation
   golangci-lint --version
   ```

### Running Tests

Run all tests with coverage:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Linting

This project uses `golangci-lint` for comprehensive code quality checks.

```bash
# Run all configured linters
golangci-lint run

# Run with verbose output
golangci-lint run -v

# Auto-fix issues where possible
golangci-lint run --fix

# Run specific linters only
golangci-lint run --disable-all -E errcheck -E govet
```

**Configured Linters:**
- errcheck: Check for unchecked errors
- gosimple: Simplify code suggestions
- govet: Official Go static analysis
- staticcheck: Advanced static analysis
- gofmt, goimports: Code formatting
- gosec: Security checks
- gocritic: Additional code quality checks
- And many more (see `.golangci.yml`)

### LSP (Language Server Protocol) Setup

This project is configured to work with `gopls`, the official Go language server.

#### VS Code

Settings are pre-configured in `.vscode/settings.json`. Just open the project in VS Code with the Go extension installed.

Features enabled:
- Auto-formatting on save
- Auto-import organization
- Semantic highlighting
- Inline documentation
- Code completion with placeholders

#### Neovim

Add to your Neovim configuration:

```lua
require('lspconfig').gopls.setup{
  settings = {
    gopls = {
      analyses = {
        unusedparams = true,
        shadow = true,
        nilness = true,
      },
      staticcheck = true,
    }
  }
}
```

#### Other Editors

The `gopls.toml` file provides editor-agnostic configuration for gopls. Most modern editors with Go support will automatically detect and use this configuration.

### Code Quality Standards

Before submitting code, ensure:

1. **All tests pass:**
   ```bash
   go test ./...
   ```

2. **Linter checks pass:**
   ```bash
   golangci-lint run
   ```

3. **Code is formatted:**
   ```bash
   gofmt -w .
   goimports -w .
   ```

## Project Architecture

### Package Overview

- **cmd/nifiparser**: Main application entry point, handles CLI flags and orchestration
- **internal/config**: Data models representing NiFi flow structures
- **internal/parser**: File I/O and JSON parsing logic
- **internal/display**: Output formatting and terminal display

### Design Principles

- **Separation of Concerns**: Business logic separated from CLI and display logic
- **Testability**: All core functionality has unit test coverage
- **Idiomatic Go**: Follows Go best practices and community standards
- **Minimal Dependencies**: Uses only standard library where possible

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for new functionality
4. Ensure all tests pass and linter checks succeed
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to your branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Coding Standards

- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Write meaningful commit messages
- Add tests for new features
- Update documentation as needed
- Keep the codebase clean and maintainable

## Roadmap

- [ ] Support for XML-formatted NiFi flows
- [ ] Interactive mode for flow exploration
- [ ] Export flow diagrams (ASCII art or graphical)
- [ ] Performance metrics analysis
- [ ] Configuration validation against NiFi schemas
- [ ] Support for comparing multiple flows

## FAQ

### Q: What NiFi configuration formats are supported?
**A:** Currently, only JSON format is supported. XML support is planned for future releases.

### Q: Can I use this with NiFi REST API responses?
**A:** Yes, as long as the API returns flow data in the expected JSON structure.

### Q: Does this tool modify my NiFi configurations?
**A:** No, this is a read-only analysis tool. It never modifies input files.

### Q: What Go version is required?
**A:** Go 1.21 or higher is required for building and running the tool.

## Troubleshooting

### "Error parsing file" message
- Ensure the file is valid JSON
- Check that the file follows the expected NiFi flow structure
- Verify file permissions allow reading

### LSP not working in editor
- Ensure `gopls` is installed: `go install golang.org/x/tools/gopls@latest`
- Check that your editor's Go extension is up to date
- Verify the `gopls.toml` configuration is in the project root

### Linter errors
- Run `golangci-lint cache clean` to clear the cache
- Ensure you're using a compatible golangci-lint version
- Check `.golangci.yml` for configuration issues

## License

MIT License - see LICENSE file for details

## Acknowledgments

- Apache NiFi project for the excellent data flow platform
- Go community for amazing tooling and best practices
- golangci-lint maintainers for comprehensive linting tools

## Contact

For questions, issues, or contributions, please open an issue on GitHub.

---

**Built with ❤️ using Go**
