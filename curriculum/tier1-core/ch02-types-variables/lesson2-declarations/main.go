// Lesson 2: Variable Declarations
// Demonstrates: var, short declaration :=, and const
//
// Run: go run .
package main

import "fmt"

// Package-level variables must use var (cannot use :=)
var appName = "TypeDemo"

// Constants — set at compile time, never change
const Pi = 3.14159
const MaxRetries = 3

func main() {
	fmt.Println("=== Variable Declarations ===")
	fmt.Println()

	// var with explicit type and value
	var count int = 42
	fmt.Printf("var count int = 42       → %d (type: %T)\n", count, count)

	// var with type inference
	var greeting = "Hello, Go!"
	fmt.Printf("var greeting = \"Hello\"   → %s (type: %T)\n", greeting, greeting)

	// var with zero value (no initial value)
	var score int
	fmt.Printf("var score int            → %d (type: %T) [zero value]\n", score, score)

	// Short declaration — most common inside functions
	language := "Go"
	version := 1.24
	fmt.Printf("language := \"Go\"         → %s (type: %T)\n", language, language)
	fmt.Printf("version := 1.24          → %.2f (type: %T)\n", version, version)

	// Multiple declarations in one line
	x, y := 10, 20
	fmt.Printf("x, y := 10, 20          → x=%d, y=%d\n", x, y)

	fmt.Println()
	fmt.Println("--- Constants ---")
	fmt.Printf("appName     = %s\n", appName)
	fmt.Printf("Pi          = %f\n", Pi)
	fmt.Printf("MaxRetries  = %d\n", MaxRetries)

	fmt.Println()
	fmt.Println("Key insight: use := inside functions, var at package level.")
}
