# Chapter 01 Cheat Sheet: Modules, Build & Project Structure

| Pattern | Code |
|---------|------|
| Initialize a new module | `go mod init github.com/user/myapp` |
| Sync dependencies | `go mod tidy` |
| Verify dependency hashes | `go mod verify` |
| Compiler-enforced privacy | Place packages in `internal/` directory |
| Build for current platform | `go build -o myapp .` |
| Cross-compile for Linux | `GOOS=linux GOARCH=amd64 go build -o myapp .` |
| Run all tests recursively | `go test ./...` |
| Run tests with verbose output | `go test -v ./...` |
| Makefile build target | `build:\n\tgo build -o myapp ./cmd/myapp` |
| Godoc comment for a function | `// FuncName does something specific.` |
| View package documentation | `go doc ./mypackage` |
