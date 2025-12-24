# godev 🚀 [![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org/dl/)

A modern Go development kit that streamlines project setup, environment diagnostics, and development workflow.

## Features ✨

- **Project Initialization**: Quickly scaffold new Go projects with best-practice templates
- **Environment Diagnostics**: Comprehensive health checks for your Go development environment
- **Modern Tooling**: Integrated with popular Go development tools
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Beautiful CLI**: Enhanced terminal user interface with colors and emojis

## Installation 📦

### Using Go Install
```bash
go install github.com/thought2code/godev@latest
```

### Building from Source
```bash
git clone https://github.com/thought2code/godev.git
cd godev
go build -o godev main.go
```

### Prerequisites
- Go 1.25 or later

## Usage 🚀

### Initialize a New Project
```bash
# Initialize in current directory
godev init

# Initialize with project name
godev init myproject

# Initialize in specific directory
godev init /path/to/myproject
```

The `init` command creates a new Go project with:
- Pre-configured VS Code settings
- `.gitignore` file
- `golangci-lint-v2` configuration
- Latest Go module setup
- Professional project structure

### Environment Health Check
```bash
godev doctor
```

The `doctor` command diagnoses your development environment:
- Validates Go module file (`go.mod`)
- Checks Go version compatibility
- Verifies essential Go tools installation
- Provides actionable remediation advice

### Testing Framework
```bash
# Run unit tests
godev test unit

# Run integration tests  
godev test integ
```

## Project Structure 📁

When you initialize a new project, godev creates:

```
myproject/
├── .vscode/
│   ├── extensions.json    # Recommended VS Code extensions
│   ├── launch.json        # Debug configuration
│   └── settings.json      # VS Code settings
├── .gitignore             # Git ignore rules
├── .golangci.yml         # Linting configuration
├── go.mod                 # Go module file
└── README.md              # Project documentation
```

## Commands Reference 📖

| Command | Description | Example |
|---------|-------------|---------|
| `godev` | Show help information | `godev` |
| `godev init [project]` | Initialize new Go project | `godev init myapp` |
| `godev doctor` | Diagnose development environment | `godev doctor` |
| `godev test unit` | Run unit tests | `godev test unit` |
| `godev test integ` | Run integration tests | `godev test integ` |

## Development Tools Integration 🔧

godev automatically sets up and recommends these essential Go development tools:

- **[gofumpt](https://github.com/mvdan/gofumpt)**: Enhanced Go formatter
- **[goimports](https://golang.org/x/tools/cmd/goimports)**: Automatic import management
- **[golangci-lint-v2](https://golangci-lint.run/)**: Fast Go linters runner

## Configuration ⚙️

### VS Code Integration
The generated `.vscode/settings.json` includes:
- Go extension configuration
- Format on save with gofumpt
- Auto-import organization
- Linting integration

### Linting Configuration
The `.golangci.yml` file provides:
- Comprehensive linter rules
- Performance optimizations
- Custom rule configurations

## Architecture 🏗️

The project follows a clean architecture:

```
godev/
├── cmd/                    # CLI commands
│   ├── root.go          # Root command setup
│   ├── init.go          # Project initialization
│   ├── doctor.go        # Environment diagnostics
│   └── test.go          # Testing commands
├── internal/             # Internal packages
│   ├── osutil/          # OS utilities
│   ├── strconst/        # String constants
│   └── tui/             # Terminal UI utilities
├── template/             # Project templates
└── main.go              # Application entry point
```

## Contributing 🤝

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Setup
1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/godev.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## Testing 🧪

Run the test suite:
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...

# Run tests verbosely
go test -v ./...
```

## Dependencies 📚

- [Cobra](https://github.com/spf13/cobra): CLI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss): Terminal styling
- [golang.org/x/mod](https://golang.org/x/mod): Go module utilities

## License 📄

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## Support 💬

- 🐛 Issues: [GitHub Issues](https://github.com/thought2code/godev/issues)
- 🔧 Pull Requests: [Pull Requests](https://github.com/thought2code/godev/pulls)

## Changelog 📝

See [CHANGELOG.md](CHANGELOG.md) for a list of changes.

---

**Made with ❤️ by [Thought2Code](https://thought2code.com)**
