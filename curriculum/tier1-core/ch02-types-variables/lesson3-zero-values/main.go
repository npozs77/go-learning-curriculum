// Lesson 3: Zero-Value Semantics
// Demonstrates: Go's zero-value guarantees for every basic type
//
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("=== Zero Values in Go ===")
	fmt.Println()

	// Declare variables without assigning values
	var i int
	var f float64
	var b bool
	var s string
	var by byte
	var r rune

	fmt.Println("Type        | Zero Value")
	fmt.Println("------------|----------")
	fmt.Printf("int         | %d\n", i)
	fmt.Printf("float64     | %g\n", f)
	fmt.Printf("bool        | %t\n", b)
	fmt.Printf("string      | %q\n", s)
	fmt.Printf("byte        | %d\n", by)
	fmt.Printf("rune        | %d\n", r)

	fmt.Println()

	// Zero values are usable immediately
	fmt.Println("--- Zero values are safe to use ---")
	total := i + 100
	fmt.Printf("int zero + 100 = %d\n", total)

	greeting := s + "Hello!"
	fmt.Printf("string zero + \"Hello!\" = %q\n", greeting)

	flag := !b
	fmt.Printf("!bool zero = %t\n", flag)

	fmt.Println()
	fmt.Println("Key insight: Go guarantees every variable starts in a known state.")
	fmt.Println("There are no uninitialized variables in Go.")
}
