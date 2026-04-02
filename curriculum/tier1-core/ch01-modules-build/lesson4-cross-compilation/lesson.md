# Lesson 4: Building Go Binaries and Cross-Compilation

## Concept

Go compiles your entire project into a single, self-contained binary. There's no runtime to install, no interpreter to configure, no virtual environment to activate. You run `go build` and get one executable file that includes everything it needs.

What makes Go especially powerful is built-in cross-compilation. By setting two environment variables — `GOOS` (target operating system) and `GOARCH` (target CPU architecture) — you can build a binary for any supported platform from your current machine. Want to build a Linux binary on your Mac? Just set `GOOS=linux GOARCH=amd64` before `go build`.

The `go build` command has useful flags: `-o` sets the output filename, and `./cmd/myapp` specifies which package to build. For simple projects with a single `main.go`, just `go build .` works.

## Analogy

Think of Go compilation like baking a cake. All your ingredients (source code, dependencies) go into the oven (compiler), and one finished cake (binary) comes out. You can even bake a cake shaped for a different pan (cross-compile for a different OS) without changing the recipe — just tell the oven which pan to use.

## What the Code Demonstrates

The `main.go` file uses `runtime.GOOS` and `runtime.GOARCH` to display the current platform, then prints the cross-compilation commands you would use to build for other platforms.

## Key Takeaways

- `go build` produces a single self-contained binary — no runtime needed
- Cross-compile with `GOOS` and `GOARCH` environment variables
- Common targets: `linux/amd64`, `darwin/arm64` (Apple Silicon), `windows/amd64`
