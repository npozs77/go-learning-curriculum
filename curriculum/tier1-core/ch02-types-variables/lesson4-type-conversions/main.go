// Lesson 4: Type Conversions and Type Assertions
// Demonstrates: explicit type conversions and safe type assertions
//
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("=== Type Conversions ===")
	fmt.Println()

	// int to float64
	count := 42
	precise := float64(count)
	fmt.Printf("int → float64:  %d → %f\n", count, precise)

	// float64 to int (truncates!)
	pi := 3.14159
	rounded := int(pi)
	fmt.Printf("float64 → int:  %f → %d (truncated, not rounded)\n", pi, rounded)

	// int to string — converts to Unicode code point, NOT digit string
	code := 65
	letter := string(rune(code))
	fmt.Printf("int 65 → string(rune): %q (ASCII 'A')\n", letter)

	// Use fmt.Sprintf for number-to-string
	numStr := fmt.Sprintf("%d", count)
	fmt.Printf("fmt.Sprintf(\"%%d\", 42): %q\n", numStr)

	// string to []byte and back
	msg := "Hello"
	bytes := []byte(msg)
	back := string(bytes)
	fmt.Printf("string → []byte → string: %q → %v → %q\n", msg, bytes, back)

	fmt.Println()
	fmt.Println("=== Type Assertions ===")
	fmt.Println()

	// Type assertion on an interface value
	var anything interface{} = "I am a string"

	// Safe form — two-value return
	str, ok := anything.(string)
	fmt.Printf("assert string: value=%q, ok=%t\n", str, ok)

	// Failed assertion — ok is false, value is zero
	num, ok := anything.(int)
	fmt.Printf("assert int:    value=%d, ok=%t\n", num, ok)

	fmt.Println()
	fmt.Println("Key insight: Go never converts types silently.")
	fmt.Println("Always use the two-value form for type assertions.")
}
