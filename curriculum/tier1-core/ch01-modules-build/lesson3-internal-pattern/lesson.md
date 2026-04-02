# Lesson 3: The internal/ Directory Pattern

## Concept

Go has a unique way of enforcing code privacy at the directory level. If you place packages inside a directory named `internal/`, the Go compiler will prevent any code outside your module from importing those packages. This is not a convention or a linter rule — it's enforced by the compiler itself.

Inside your own module, you can import `internal/` packages freely. The restriction only applies to external consumers. This makes `internal/` perfect for code that supports your public API but shouldn't be used directly by others — things like helper utilities, configuration loaders, or implementation details.

The pattern works at any level of your directory tree. A package at `myapp/internal/config/` can be imported by `myapp/cmd/server/` but not by `github.com/someone/other-project`. The compiler draws the boundary at the parent of the `internal/` directory.

## Analogy

Think of `internal/` like the kitchen in a restaurant. Customers (external code) can order from the menu (your public API), but they can't walk into the kitchen (internal packages). The kitchen staff (your own module's code) moves freely between the dining room and the kitchen. The door isn't just labeled "Staff Only" — it's actually locked.

## What the Code Demonstrates

The `main.go` file imports a helper function from `internal/helper/helper.go`, showing that code within the same module can access internal packages. The helper function returns a greeting that main.go prints.

## Key Takeaways

- Place packages in `internal/` to make them compiler-enforced private
- Your own module can import `internal/` packages freely
- External modules cannot import your `internal/` packages — the compiler rejects it
