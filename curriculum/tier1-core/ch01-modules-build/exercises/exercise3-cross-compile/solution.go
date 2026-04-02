//go:build ignore

// Exercise 3: Cross-Compile a Go Binary — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import (
	"fmt"
	"runtime"
)

// Target represents a cross-compilation target.
type Target struct {
	GOOS   string
	GOARCH string
}

func main() {
	fmt.Printf("Current platform: %s/%s\n\n", runtime.GOOS, runtime.GOARCH)

	targets := []Target{
		{"linux", "amd64"},
		{"linux", "arm64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
	}

	fmt.Println("Cross-compilation commands:")
	for _, t := range targets {
		outName := fmt.Sprintf("myapp-%s-%s", t.GOOS, t.GOARCH)
		marker := ""
		if t.GOOS == runtime.GOOS && t.GOARCH == runtime.GOARCH {
			marker = " ← current"
		}
		fmt.Printf("  GOOS=%-8s GOARCH=%-6s go build -o %-24s .%s\n",
			t.GOOS, t.GOARCH, outName, marker)
	}
}
