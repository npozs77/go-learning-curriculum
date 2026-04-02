# Chapter 01 Real-Project Connection: Modules, Build & Project Structure

## How These Concepts Apply: Vocabulary Generator

In a vocabulary generator CLI application, the concepts from this chapter appear in these ways:

### go.mod and Module Declaration

The vocabulary generator's `go.mod` file at the project root declares the module as `github.com/user/vocabgen` and pins the Go version to 1.24. It lists two dependencies in the `require` block: `github.com/spf13/cobra` for CLI command parsing and `pgregory.net/rapid` for property-based testing. Running `go mod tidy` keeps this file in sync whenever imports change.

### The internal/ Directory Pattern

All business logic lives inside `internal/` to prevent external consumers from depending on implementation details. The project uses `internal/service/` for the core vocabulary generation logic, `internal/config/` for YAML configuration loading, `internal/db/` for SQLite caching, and `internal/llm/` for the LLM provider interface. Only `cmd/vocabgen/main.go` — the CLI entry point — wires these internal packages together.

### Cross-Compilation and Building

The vocabulary generator uses `go build -o vocabgen ./cmd/vocabgen` to produce a single self-contained binary. For distribution, cross-compilation builds binaries for macOS (`GOOS=darwin GOARCH=arm64`), Linux (`GOOS=linux GOARCH=amd64`), and Windows (`GOOS=windows GOARCH=amd64`). A Makefile automates these builds with targets like `make build`, `make test`, and `make release`.

### Godoc and Testing

Every exported function in the project follows Godoc conventions — for example, `// GenerateVocab produces a vocabulary list for the given topic.` in `internal/service/generator.go`. Tests live alongside the code: `internal/service/generator_test.go` tests the generation logic, and `go test ./...` from the project root runs the entire test suite.
