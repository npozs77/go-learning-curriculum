# Lesson 2: go.sum and Dependency Verification

## Concept

When your project depends on external packages, Go creates a second file alongside `go.mod` called `go.sum`. This file stores cryptographic hashes (checksums) of every dependency your project uses. Each time Go downloads a dependency, it computes a hash and compares it against what's recorded in `go.sum`.

You never edit `go.sum` by hand — Go manages it automatically. The three commands you'll use most often are: `go mod tidy` (add missing deps, remove unused ones), `go mod download` (fetch all deps to your local cache), and `go get pkg@version` (add or update a specific dependency).

If someone tampers with a dependency — changing the code but keeping the same version number — the hash won't match and Go will refuse to build. This protects your project from supply-chain attacks without any extra tooling.

## Analogy

Think of `go.sum` as a receipt with fingerprints. When you buy ingredients (dependencies), the receipt records exactly what you got and a fingerprint for each item. Next time you cook (build), you check the fingerprints — if an ingredient has been swapped out, you'll know immediately.

## What the Code Demonstrates

The `main.go` file simulates how Go verifies dependencies by computing SHA-256 hashes of strings, showing the principle behind `go.sum` verification without requiring actual external dependencies.

## Key Takeaways

- `go.sum` stores cryptographic hashes of all dependencies for tamper detection
- Never edit `go.sum` manually — use `go mod tidy` to manage it
- The checksum database at `sum.golang.org` provides an additional layer of verification
