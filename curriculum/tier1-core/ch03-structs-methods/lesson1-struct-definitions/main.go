// Lesson 1: Struct Definitions and Field Access
// Demonstrates: defining structs, creating instances, accessing fields, and copy behavior
//
// Run: go run .
package main

import "fmt"

// Book represents a book with basic metadata.
type Book struct {
	Title  string
	Author string
	Pages  int
	Read   bool
}

func main() {
	fmt.Println("=== Struct Definitions and Field Access ===")
	fmt.Println()

	// Create a struct with named fields (preferred)
	book1 := Book{
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
		Pages:  380,
		Read:   true,
	}
	fmt.Printf("Book 1: %s by %s (%d pages, read: %t)\n",
		book1.Title, book1.Author, book1.Pages, book1.Read)

	// Create a struct with zero values, then set fields
	var book2 Book
	book2.Title = "Learning Go"
	book2.Author = "Jon Bodner"
	book2.Pages = 375
	fmt.Printf("Book 2: %s by %s (%d pages, read: %t)\n",
		book2.Title, book2.Author, book2.Pages, book2.Read)

	fmt.Println()

	// Structs are value types — assignment copies the data
	book3 := book1
	book3.Title = "Concurrency in Go"
	fmt.Println("After copying book1 to book3 and changing book3's title:")
	fmt.Printf("  book1.Title = %q (unchanged)\n", book1.Title)
	fmt.Printf("  book3.Title = %q (independent copy)\n", book3.Title)

	fmt.Println()
	fmt.Println("Key insight: structs are value types — copies are independent.")
}
