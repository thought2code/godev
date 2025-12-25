# godev 🚀

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org/dl/) [![Go Report Card](https://goreportcard.com/badge/github.com/thought2code/godev)](https://goreportcard.com/report/github.com/thought2code/godev) [![PkgGoDev](https://pkg.go.dev/badge/github.com/thought2code/godev)](https://pkg.go.dev/github.com/thought2code/godev)

**Less boilerplate, more building.**
The all-in-one command center and project scaffolder for modern Go development.

## 💡 Why godev?

> As a Go novice, I was once overwhelmed by the 'preliminary chores'. From choosing between `go fmt` and `gofumpt` to configuring `golangci-lint` and better project structures, various CLI commands like `gofumpt -w .`, `golangci-lint run ./...`, `go mod tidy`, etc. Every new project felt like a tedious loop of 'copy-paste' from the last one. I thought I need a tool that just worked and I don't like to use `Makefile`, so I built **godev**. It's born to free you from these chores, so you can focus on what matters: **writing your code.**

It provides a unified CLI to:
1. **Manage**: A single entry point for project init, testing, linting, building, and releasing.
2. **Diagnose**: A "Doctor" that doesn't just find issues, but understands Go development environments.
3. **Scaffold**: Initialize modern Go project with preset template and configuration.

## ✨ Key Features

- 📦 **One-Stop Scaffolding**: Create Go projects following best practices (`.golangci.yml`, `go.mod`, etc.).
- 🩺 **Interactive Doctor**: Comprehensive health checks for your Go version, modules, and toolchain.
- 🛠️ **Unified Workflow**: Integrated commands for unit and integration tests, stop remembering complex flags.
- ⚙️ **Standardized Tooling**: Automatically set up VS Code settings, `gofumpt`, and `golangci-lint`.

## 📦 Installation

### Using Go Install (Recommended)
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
- Go 1.25 or later (latest stable recommended)

## 🚀 Usage

### 1. Initialize a New Project

Don't waste time on folder structures.

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

### 2. Check Environment Health

Is your `GOPATH` messed up? Are you missing tools?

```bash
godev doctor
```

The `doctor` command diagnoses your development environment:
- Validates Go module file (`go.mod`)
- Checks Go version compatibility
- Verifies essential Go tools installation
- Provides actionable remediation advice

### 3. Smart Testing

No more long, messy `go test ./...` flags.

```bash
godev test unit           # Run unit tests
godev test unit -v        # Run unit tests with verbose output
godev test unit -c        # Run unit tests with coverage and save the cover profile
godev test unit --html    # Run unit tests with coverage and open report in your browser
godev test integ          # Run integration tests
```

## 📁 Project Structure

When you initialize a new project, godev creates:

```
myproject/
├── .vscode/
│   ├── extensions.json    # Recommended VS Code extensions
│   ├── launch.json        # Debug configuration
│   └── settings.json      # Recommended VS Code settings
├── .gitignore             # Git ignore rules
├── .golangci.yml          # Linting configuration
├── go.mod                 # Go module file
└── README.md              # Project documentation
```

## 📚 Commands Reference

| Command                | Description                      | Example            |
|------------------------|----------------------------------|--------------------|
| `godev`                | Show help information            | `godev`            |
| `godev init [project]` | Initialize new Go project        | `godev init myapp` |
| `godev doctor`         | Diagnose development environment | `godev doctor`     |
| `godev test unit`      | Run unit tests                   | `godev test unit`  |
| `godev test integ`     | Run integration tests            | `godev test integ` |

## 🔧 Development Tools Integration

`godev` automatically set up and recommend these essential Go development tools:

- **[gofumpt](https://github.com/mvdan/gofumpt)**: Enhanced Go formatter
- **[goimports](https://golang.org/x/tools/cmd/goimports)**: Automatic import management
- **[golangci-lint-v2](https://golangci-lint.run/)**: Fast Go linters runner

## ⚙️ Configuration

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

## 🏗️ Project Architecture

The project follows a clean architecture:

```
godev/
├── cmd/                 # CLI commands
│   ├── root.go          # Root command setup
│   ├── init.go          # Project initialization
│   ├── doctor.go        # Environment diagnostics
│   ├── tools.go         # Go tools management
│   └── test.go          # Testing commands
├── internal/            # Internal packages
│   ├── osutil/          # OS utilities (filesystem, exec, etc.)
│   ├── strconst/        # String constants
│   └── tui/             # Terminal UI utilities (colorized output, etc.)
├── template/            # Preset project templates
└── main.go              # Application entry point
```

## Contributing 🤝

We welcome and appreciate contributions ❤️

### Development Setup
1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/godev.git`
3. Create a feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes and add tests
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## 💬 Support

- 🐛 Issues: [GitHub Issues](https://github.com/thought2code/godev/issues)
- 🔧 Pull Requests: [Pull Requests](https://github.com/thought2code/godev/pulls)

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](https://github.com/thought2code/godev/blob/main/LICENSE) file for details.
