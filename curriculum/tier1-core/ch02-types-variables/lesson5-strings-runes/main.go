// Lesson 5: Strings, Runes, and UTF-8 Encoding
// Demonstrates: byte vs rune length, for-range over runes, and UTF-8 encoding
//
// Run: go run .
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== Strings, Runes & UTF-8 ===")
	fmt.Println()

	// A string with multi-byte characters
	greeting := "Hello, 世界! 🌍"

	fmt.Printf("String:     %s\n", greeting)
	fmt.Printf("Byte len:   %d\n", len(greeting))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(greeting))

	fmt.Println()
	fmt.Println("--- Iterating over runes ---")
	for i, r := range greeting {
		fmt.Printf("  byte %2d: rune %c (U+%04X)\n", i, r, r)
	}

	fmt.Println()
	fmt.Println("--- Byte indexing vs rune iteration ---")
	word := "café"
	fmt.Printf("String: %q\n", word)
	fmt.Printf("len():  %d bytes\n", len(word))
	fmt.Printf("Runes:  %d characters\n", utf8.RuneCountInString(word))

	fmt.Println()
	fmt.Println("Byte-by-byte:")
	for i := 0; i < len(word); i++ {
		fmt.Printf("  byte[%d] = 0x%02X\n", i, word[i])
	}

	fmt.Println()
	fmt.Println("Rune-by-rune:")
	for _, r := range word {
		fmt.Printf("  rune: %c (U+%04X)\n", r, r)
	}

	fmt.Println()
	fmt.Println("Key insight: len() counts bytes, not characters.")
	fmt.Println("Use for-range to iterate over runes safely.")
}
