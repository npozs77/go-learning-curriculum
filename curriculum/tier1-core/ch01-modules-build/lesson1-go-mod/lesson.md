# Lesson 1: go.mod and Module Declaration

## Concept

Every Go project starts with a `go.mod` file. This file declares your project as a **module** — a collection of related Go packages that are versioned together. The `go.mod` file lives at the root of your project and tells Go two essential things: the module's unique name (its path) and the minimum Go version required.

When you run `go mod init <module-path>`, Go creates this file for you. The module path is typically a URL-like identifier (e.g., `github.com/user/myapp`), but for local projects that won't be published, any name works. Once you have a `go.mod`, you can use `go run .`, `go build .`, and `go test ./...` from anywhere inside the module.

The `go.mod` file also tracks your dependencies. When you import a package from another module, Go adds it to the `require` section automatically. You can clean up unused dependencies with `go mod tidy`.

## Analogy

Think of `go.mod` like the label on a shipping container. The label says "this container belongs to Company X" (the module path) and "packed according to Standard v1.24" (the Go version). Everything inside the container — your packages, your code — is organized and shipped together as one unit.

## What the Code Demonstrates

The `main.go` file reads and prints the contents of its own `go.mod` file, showing you exactly what a module declaration looks like from inside a running Go program.

## Key Takeaways

- Every Go project needs a `go.mod` file — create one with `go mod init <path>`
- The module path is your project's unique identifier
- `go.mod` tracks the Go version and all dependencies
