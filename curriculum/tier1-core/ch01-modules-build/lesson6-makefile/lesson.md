# Lesson 6: Makefile Usage for Common Project Tasks

## Concept

A Makefile gives you one-word shortcuts for common project tasks. Instead of remembering `go build -o myapp ./cmd/myapp`, you type `make build`. Instead of `go test -race ./...`, you type `make test`. It's a command dashboard for your project.

In Go projects, a Makefile typically defines targets for building, testing, vetting, linting, and cleaning. Each target is a recipe: a name followed by the shell commands to run. Targets can depend on other targets — for example, `make ci` might run `vet`, `test`, and `build` in sequence.

While Makefiles aren't Go-specific, they're widely used in the Go ecosystem because Go's toolchain is command-line driven. The combination of `go build`, `go test`, `go vet`, and `go fmt` maps naturally to Makefile targets.

## Analogy

Think of a Makefile like a TV remote control. Each button (target) triggers a specific action. You don't need to walk to the TV and press buttons manually (type long commands) — you just press one button on the remote (`make test`). And some buttons are combos — pressing "Movie Mode" dims the lights, turns on surround sound, and switches the input all at once.

## What the Code Demonstrates

The `main.go` file prints a typical Go project Makefile, showing the targets and commands that a real Makefile would contain. It demonstrates the pattern without requiring `make` to be installed.

## Key Takeaways

- A Makefile provides one-word shortcuts for common Go commands
- Standard targets: `build`, `test`, `vet`, `fmt`, `clean`, `lint`
- Targets can depend on each other to create multi-step workflows
