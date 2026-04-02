// Lesson 6: Control Flow — if, for, switch
// Demonstrates: if with initializer, three for-loop forms, and switch statements
//
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("=== Control Flow in Go ===")
	fmt.Println()

	// --- if with initializer ---
	fmt.Println("--- if with initializer ---")
	if length := len("Gopher"); length > 3 {
		fmt.Printf("\"Gopher\" has %d characters (more than 3)\n", length)
	}

	// --- Classic for loop ---
	fmt.Println()
	fmt.Println("--- Classic for loop ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}

	// --- While-style for loop ---
	fmt.Println()
	fmt.Println("--- While-style for loop ---")
	n := 1
	for n < 32 {
		fmt.Printf("  n = %d\n", n)
		n *= 2
	}

	// --- for range over a string ---
	fmt.Println()
	fmt.Println("--- for range over a string ---")
	for i, ch := range "Go!" {
		fmt.Printf("  index %d: %c\n", i, ch)
	}

	// --- switch statement ---
	fmt.Println()
	fmt.Println("--- switch statement ---")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek — keep going!")
	case "Friday":
		fmt.Println("Almost weekend!")
	default:
		fmt.Println("Weekend!")
	}

	// --- Conditionless switch (like if/else if) ---
	fmt.Println()
	fmt.Println("--- Conditionless switch ---")
	score := 85
	switch {
	case score >= 90:
		fmt.Printf("Score %d: A\n", score)
	case score >= 80:
		fmt.Printf("Score %d: B\n", score)
	case score >= 70:
		fmt.Printf("Score %d: C\n", score)
	default:
		fmt.Printf("Score %d: needs improvement\n", score)
	}

	fmt.Println()
	fmt.Println("Key insight: Go has one loop (for) and switch cases don't fall through.")
}
