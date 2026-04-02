# Exercise 1: Initialize a Go Module

## Task

Create a Go program that reads its own `go.mod` file and prints the module name and Go version. This exercise reinforces how `go.mod` works and how to read files in Go.

## Requirements

- Read the `go.mod` file from the current directory
- Parse and print the module name (the value after `module`)
- Parse and print the Go version (the value after `go`)
- Handle the case where `go.mod` doesn't exist with a clear error message

## Expected Behavior

When you run `go run .`, you should see:

```
Module name: exercise1-init-module
Go version:  1.24
```

## Hints

<details>
<summary>Hint 1</summary>

Use `os.ReadFile("go.mod")` to read the file contents into a byte slice.

</details>

<details>
<summary>Hint 2</summary>

Use `strings.Split` to break the file into lines, then check if each line starts with `"module "` or `"go "` using `strings.HasPrefix`.

</details>
