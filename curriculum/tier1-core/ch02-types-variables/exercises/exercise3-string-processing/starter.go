// Exercise 3: Work with Strings and Runes
// Task: Analyze a Unicode string — count bytes, runes, and identify multi-byte characters
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "Go is fun! 🚀🌍"

	fmt.Println("String Analyzer")
	fmt.Println("===============")
	fmt.Printf("Text:       %q\n", text)

	// TODO: Print the byte count using len()
	// Hint: fmt.Printf("Byte count: %d\n", len(text))

	// TODO: Print the rune count using utf8.RuneCountInString()
	// Hint: fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))

	fmt.Println()
	fmt.Println("Rune Details:")

	// TODO: Use for range to iterate over runes and print each one
	// Hint: for i, r := range text { ... }
	// Print format: fmt.Printf("  byte %2d: %q (U+%04X)\n", i, r, r)
	_ = utf8.RuneLen('a') // remove this line when you implement the loop

	// TODO: Count multi-byte runes (where utf8.RuneLen(r) > 1)
	// Hint: keep a counter and increment it inside the loop

	// TODO: Print the multi-byte rune count
	// Hint: fmt.Printf("\nMulti-byte runes: %d out of %d\n", multiByte, total)

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
