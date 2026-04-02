// Lesson 3: The internal/ Directory Pattern
// Demonstrates: compiler-enforced privacy using the internal/ directory
//
// Run: go run .
package main

import (
	"fmt"

	"lesson3-internal-pattern/internal/helper"
)

func main() {
	fmt.Println("=== The internal/ Directory Pattern ===")
	fmt.Println()

	// We can import internal/helper because we're in the same module
	greeting := helper.Greet("Gopher")
	fmt.Println(greeting)
	fmt.Println()

	info := helper.ModuleInfo()
	fmt.Println(info)
	fmt.Println()

	fmt.Println("Directory structure:")
	fmt.Println("  lesson3-internal-pattern/")
	fmt.Println("  ├── main.go              ← can import internal/helper ✓")
	fmt.Println("  ├── go.mod")
	fmt.Println("  └── internal/")
	fmt.Println("      └── helper/")
	fmt.Println("          └── helper.go    ← locked to this module")
	fmt.Println()
	fmt.Println("If an external module tried to import internal/helper,")
	fmt.Println("the Go compiler would reject it with an error.")
}
