// Lesson 2: go.sum and Dependency Verification
// Demonstrates: how cryptographic hashes verify dependency integrity
//
// Run: go run .
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("=== go.sum: Dependency Verification ===")
	fmt.Println()

	// Simulate how go.sum works: hash a dependency's content
	original := "func Hello() string { return \"hello\" }"
	tampered := "func Hello() string { return \"hacked\" }"

	originalHash := sha256.Sum256([]byte(original))
	tamperedHash := sha256.Sum256([]byte(tampered))

	fmt.Println("Original code:", original)
	fmt.Printf("Original hash: %x\n", originalHash[:8])
	fmt.Println()

	fmt.Println("Tampered code:", tampered)
	fmt.Printf("Tampered hash: %x\n", tamperedHash[:8])
	fmt.Println()

	if originalHash == tamperedHash {
		fmt.Println("Result: Hashes match — safe to use")
	} else {
		fmt.Println("Result: Hashes DON'T match — Go would reject this!")
	}

	fmt.Println()
	fmt.Println("This is exactly what go.sum does:")
	fmt.Println("  - Records hashes when you first download a dependency")
	fmt.Println("  - Verifies hashes on every subsequent build")
	fmt.Println("  - Rejects builds if any hash doesn't match")
	fmt.Println()
	fmt.Println("Key commands:")
	fmt.Println("  go mod tidy      — sync dependencies")
	fmt.Println("  go mod download  — fetch all deps to cache")
	fmt.Println("  go mod verify    — check all deps against go.sum")
}
