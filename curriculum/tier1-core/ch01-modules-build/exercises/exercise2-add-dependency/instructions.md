# Exercise 2: Add a Dependency

## Task

Create a Go program that simulates dependency management by tracking packages and their versions. This exercise reinforces how `go.mod` and `go.sum` manage dependencies without requiring actual network access.

## Requirements

- Define a `Dependency` struct with `Name` and `Version` fields
- Create a function that adds a dependency to a slice (no duplicates)
- Create a function that lists all dependencies in a formatted table
- Demonstrate adding, listing, and attempting to add a duplicate

## Expected Behavior

When you run `go run .`, you should see:

```
Adding: github.com/spf13/cobra v1.10.2
Adding: pgregory.net/rapid v1.2.0
Adding: github.com/spf13/cobra v1.10.2 (already exists, skipping)

Dependencies:
  github.com/spf13/cobra   v1.10.2
  pgregory.net/rapid        v1.2.0

Total: 2 dependencies
```

## Hints

<details>
<summary>Hint 1</summary>

Define a struct: `type Dependency struct { Name, Version string }` and use a `[]Dependency` slice to store them.

</details>

<details>
<summary>Hint 2</summary>

To check for duplicates, loop through the existing slice and compare the `Name` field before appending.

</details>
