# Clipboard Monitor

![CI](https://github.com/Qyroxen/Clipboard-Monitor/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Clipboard-Monitor/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Clipboard-Monitor?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Clipboard-Monitor)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Clipboard-Monitor)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Clipboard-Monitor?style=social)](https://github.com/Qyroxen/Clipboard-Monitor/stargazers)

## What is it?

Clipboard Monitor is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Clipboard-Monitor.git
cd Clipboard-Monitor
go build -o clipboardmonitor .

# Run
./clipboardmonitor --help
```

## CLI Usage

```bash
# Basic usage
./clipboardmonitor

# With flags
./clipboardmonitor --verbose --output json

# Get help
./clipboardmonitor --help
```

## Examples

```bash
# Example 1
./clipboardmonitor example1

# Example 2
./clipboardmonitor example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o clipboardmonitor .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Clipboard-Monitor/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Clipboard-Monitor?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Clipboard-Monitor/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Clipboard-Monitor?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Clipboard-Monitor/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Clipboard-Monitor" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Clipboard-Monitor/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Clipboard-Monitor" alt="Pull Requests">
  </a>
</p>
