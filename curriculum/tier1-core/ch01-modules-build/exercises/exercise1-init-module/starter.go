// Exercise 1: Initialize a Go Module
// Task: Read go.mod and print the module name and Go version
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Read the go.mod file using os.ReadFile
	// Hint: os.ReadFile("go.mod") returns ([]byte, error)
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	// TODO: Loop through lines and find the module name and Go version
	// Hint: Check if a line starts with "module " or "go "
	// Hint: Use strings.TrimPrefix to extract the value
	for _, line := range lines {
		_ = line // TODO: Replace this with your parsing logic
	}

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
