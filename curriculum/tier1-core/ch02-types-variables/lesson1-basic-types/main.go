// Lesson 1: Go's Type System and Basic Types
// Demonstrates: declaring and inspecting Go's basic types (int, float64, string, bool, byte, rune)
//
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("=== Go's Basic Types ===")
	fmt.Println()

	// Integer — whole numbers
	var age int = 30
	fmt.Printf("age     = %d    (type: %T)\n", age, age)

	// Float — decimal numbers
	var price float64 = 19.99
	fmt.Printf("price   = %.2f  (type: %T)\n", price, price)

	// String — text
	var name string = "Gopher"
	fmt.Printf("name    = %s  (type: %T)\n", name, name)

	// Bool — true or false
	var active bool = true
	fmt.Printf("active  = %t    (type: %T)\n", active, active)

	// Byte — alias for uint8, used for raw data
	var initial byte = 'G'
	fmt.Printf("initial = %c      (type: %T, value: %d)\n", initial, initial, initial)

	// Rune — alias for int32, used for Unicode code points
	var emoji rune = '🎉'
	fmt.Printf("emoji   = %c    (type: %T, value: %d)\n", emoji, emoji, emoji)

	fmt.Println()
	fmt.Println("Key insight: every variable has a fixed type.")
	fmt.Println("The compiler catches type mismatches before your code runs.")
}
