// Lesson 4: Building Go Binaries and Cross-Compilation
// Demonstrates: runtime platform detection and cross-compilation targets
//
// Run: go run .
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Go Build and Cross-Compilation ===")
	fmt.Println()

	// Show current platform
	fmt.Printf("Current platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Go version:       %s\n", runtime.Version())
	fmt.Println()

	// Show common build commands
	fmt.Println("Build for current platform:")
	fmt.Println("  go build -o myapp .")
	fmt.Println()

	// Show cross-compilation examples
	targets := []struct {
		goos, goarch, desc string
	}{
		{"linux", "amd64", "Linux (x86_64)"},
		{"linux", "arm64", "Linux (ARM64, e.g., AWS Graviton)"},
		{"darwin", "arm64", "macOS (Apple Silicon)"},
		{"darwin", "amd64", "macOS (Intel)"},
		{"windows", "amd64", "Windows (x86_64)"},
	}

	fmt.Println("Cross-compilation commands:")
	for _, t := range targets {
		marker := " "
		if t.goos == runtime.GOOS && t.goarch == runtime.GOARCH {
			marker = "←"
		}
		fmt.Printf("  GOOS=%s GOARCH=%s go build -o myapp . %s %s\n",
			t.goos, t.goarch, marker, t.desc)
	}

	fmt.Println()
	fmt.Println("No Docker, no VMs — just two environment variables!")
}
