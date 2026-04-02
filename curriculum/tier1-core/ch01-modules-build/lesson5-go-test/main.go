// Lesson 5: Running Tests with go test
// Demonstrates: test file conventions and the go test command
//
// Run: go run .
// Test: go test -v .
package main

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("=== Running Tests with go test ===")
	fmt.Println()

	fmt.Printf("Add(2, 3) = %d\n", Add(2, 3))
	fmt.Printf("Add(-1, 1) = %d\n", Add(-1, 1))
	fmt.Printf("Multiply(4, 5) = %d\n", Multiply(4, 5))
	fmt.Println()

	fmt.Println("Test conventions:")
	fmt.Println("  - Test files:     *_test.go (e.g., main_test.go)")
	fmt.Println("  - Test functions: TestXxx(t *testing.T)")
	fmt.Println("  - Same package:   tests live next to the code")
	fmt.Println()
	fmt.Println("Try running the tests:")
	fmt.Println("  go test -v .       ← verbose output")
	fmt.Println("  go test ./...      ← all tests recursively")
	fmt.Println("  go test -run Add . ← only tests matching 'Add'")
}
