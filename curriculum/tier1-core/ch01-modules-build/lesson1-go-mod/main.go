// Lesson 1: go.mod and Module Declaration
// Demonstrates: reading and displaying the go.mod file to understand module structure
//
// Run: go run .
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Go Module Declaration ===")
	fmt.Println()

	// Read our own go.mod file
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println("Error reading go.mod:", err)
		os.Exit(1)
	}

	fmt.Println("Contents of go.mod:")
	fmt.Println("---")
	fmt.Print(string(data))
	fmt.Println("---")
	fmt.Println()

	fmt.Println("Key points:")
	fmt.Println("  - 'module' declares this project's unique path")
	fmt.Println("  - 'go 1.24' sets the minimum Go version")
	fmt.Println("  - Dependencies would appear in a 'require' block")
	fmt.Println()
	fmt.Println("Try: go mod init github.com/you/myproject")
}
