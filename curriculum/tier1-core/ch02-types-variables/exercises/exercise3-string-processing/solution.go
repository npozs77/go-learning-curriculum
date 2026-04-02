//go:build ignore

// Exercise 3: Work with Strings and Runes — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
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
	fmt.Printf("Byte count: %d\n", len(text))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))

	fmt.Println()
	fmt.Println("Rune Details:")

	multiByte := 0
	total := 0
	for i, r := range text {
		fmt.Printf("  byte %2d: %q (U+%04X)\n", i, r, r)
		if utf8.RuneLen(r) > 1 {
			multiByte++
		}
		total++
	}

	fmt.Printf("\nMulti-byte runes: %d out of %d\n", multiByte, total)
}
