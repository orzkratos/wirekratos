[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/wirekratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/wirekratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/wirekratos)](https://pkg.go.dev/github.com/orzkratos/wirekratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/wirekratos/main.svg)](https://coveralls.io/github/orzkratos/wirekratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/wirekratos.svg)](https://github.com/orzkratos/wirekratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/wirekratos)](https://goreportcard.com/report/github.com/orzkratos/wirekratos)

# wirekratos

Wire workspace mode fix package with Kratos integration in Go projects.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Problem

When using Google Wire with Go workspace mode (go.work), `go generate ./...` fails:

```
go generate ./...
go: -mod may only be set to readonly or vendor when in workspace mode, but it is set to "mod"
        Remove the -mod flag to use the default readonly value,
        or set GOWORK=off to disable workspace mode.
xxx/xxx/wire_gen.go:3: running "go": exit status 1
```

Wire generates code with this directive that conflicts with workspace mode:
```go
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject
```

## Solution

Auto removes the `-mod=mod` flag from Wire-generated files, making them workspace-compatible:
```go
//go:generate go run github.com/google/wire/cmd/wire
```

## Wire Status (2025)

- Latest version: v0.7.0 (August 22, 2025)
- Repo: Archived, read-access-mode (August 25, 2025)
- Workspace issue: NOT fixed (PR #410 remains open)
- This package remains needed as Wire won't receive more updates

## Installation

```bash
go install github.com/orzkratos/wirekratos/cmd/wirekratos@latest
```

## Usage

### Auto mode (Kratos projects)
```bash
wirekratos -framework=kratos
```
Auto detects `cmd/PROJECT_NAME/wire_gen.go` in Kratos projects.

### Relative path mode
```bash
wirekratos -name=cmd/myproject/wire_gen.go
```

### Absolute path mode
```bash
wirekratos -path=/absolute/path/to/wire_gen.go
```

### Debug mode
```bash
wirekratos -framework=kratos -debug
```

## Demo Projects

Complete demo projects showing integration in production Kratos applications:

**[wirekratos-demos](https://github.com/orzkratos/wirekratos-demos)** - Full-featured Kratos applications with Wire integration

* [demo1kratos](https://github.com/orzkratos/wirekratos-demos/tree/main/demo1kratos) - Basic Kratos project setup
* [demo2kratos](https://github.com/orzkratos/wirekratos-demos/tree/main/demo2kratos) - Advanced Kratos project with workspace

## How It Works

1. Locates `wire_gen.go` file
2. Checks Wire-generated code signature
3. Finds the `//go:generate` directive (line 3)
4. Removes `-mod=mod` flag if present
5. Keeps everything else unchanged

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-22 10:00:00.000000 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize via reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/wirekratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/wirekratos)
