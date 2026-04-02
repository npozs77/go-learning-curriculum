# Exercise 3: Cross-Compile a Go Binary

## Task

Create a Go program that detects the current platform and generates cross-compilation commands for all other supported platforms. This exercise reinforces how `GOOS` and `GOARCH` work for cross-compilation.

## Requirements

- Detect and display the current `GOOS` and `GOARCH` using the `runtime` package
- Define a list of at least 4 target platforms (OS/architecture pairs)
- Generate the `go build` command for each target, marking the current platform
- Print the output filename for each target (e.g., `myapp-linux-amd64`)

## Expected Behavior

When you run `go run .`, you should see (output varies by platform):

```
Current platform: linux/amd64

Cross-compilation commands:
  GOOS=linux   GOARCH=amd64 go build -o myapp-linux-amd64     . ← current
  GOOS=linux   GOARCH=arm64 go build -o myapp-linux-arm64     .
  GOOS=darwin  GOARCH=arm64 go build -o myapp-darwin-arm64    .
  GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64   .
```

## Hints

<details>
<summary>Hint 1</summary>

Use `runtime.GOOS` and `runtime.GOARCH` to get the current platform values.

</details>

<details>
<summary>Hint 2</summary>

Use `fmt.Sprintf` to build the output filename: `fmt.Sprintf("myapp-%s-%s", goos, goarch)`.

</details>
