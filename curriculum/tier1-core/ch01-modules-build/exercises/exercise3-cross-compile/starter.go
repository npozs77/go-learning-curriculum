// Exercise 3: Cross-Compile a Go Binary
// Task: Generate cross-compilation commands for multiple platforms
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
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

	// TODO: Define a slice of Target with at least 4 platforms
	// Hint: Include linux/amd64, linux/arm64, darwin/arm64, windows/amd64
	targets := []Target{
		// TODO: Add targets here
	}

	fmt.Println("Cross-compilation commands:")
	for _, t := range targets {
		_ = t // TODO: Replace with your formatting logic
		// TODO: Build output filename as "myapp-GOOS-GOARCH"
		// TODO: Mark the current platform with "← current"
		// TODO: Print the GOOS=... GOARCH=... go build command
	}

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
