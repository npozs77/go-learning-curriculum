# Lesson 5: Running Tests with go test

## Concept

Go has testing built into the toolchain — no separate test runner to install, no configuration files to write. Test files live right next to the code they test: if you have `math.go`, your tests go in `math_test.go` in the same directory.

Test functions follow a simple convention: they must start with `Test`, take a single parameter of type `*testing.T`, and live in files ending with `_test.go`. The `go test` command finds and runs them automatically.

The `./...` glob pattern is one of Go's most useful shortcuts. It means "this directory and all subdirectories, recursively." So `go test ./...` runs every test in your entire project. You can also target a specific package: `go test ./internal/config/` runs only the tests in that package.

## Analogy

Think of `go test` like a building inspector who knows exactly where to look. Test files (`_test.go`) are like inspection checklists placed in each room. The inspector walks through every room (`./...`), finds the checklists, and runs each check. You don't need to tell the inspector where to go — the naming convention handles it.

## What the Code Demonstrates

The `main.go` file contains a simple `Add` function and prints its results. The `main_test.go` file contains a `TestAdd` function that verifies the `Add` function works correctly. Run `go test -v .` from this directory to see the test output.

## Key Takeaways

- Test files end with `_test.go` and live next to the code they test
- Test functions start with `Test` and take `*testing.T`
- `go test ./...` runs all tests recursively; `go test -v` shows verbose output
