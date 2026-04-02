//go:build ignore

// Exercise 1: Define and Use Structs — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

// Movie represents a film with basic metadata.
type Movie struct {
	Title    string
	Director string
	Year     int
	Rating   float64
}

func main() {
	fmt.Println("Movie Collection")
	fmt.Println("================")

	movie1 := Movie{
		Title:    "The Matrix",
		Director: "The Wachowskis",
		Year:     1999,
		Rating:   8.7,
	}

	movie2 := Movie{
		Title:    "Inception",
		Director: "Christopher Nolan",
		Year:     2010,
		Rating:   8.8,
	}

	fmt.Printf("%q (%d) directed by %s — Rating: %.1f\n",
		movie1.Title, movie1.Year, movie1.Director, movie1.Rating)
	fmt.Printf("%q (%d) directed by %s — Rating: %.1f\n",
		movie2.Title, movie2.Year, movie2.Director, movie2.Rating)

	fmt.Println()

	// Demonstrate copy behavior
	movieCopy := movie1
	movieCopy.Rating = 9.0

	fmt.Println("Copy behavior:")
	fmt.Printf("  Original rating: %.1f\n", movie1.Rating)
	fmt.Printf("  Copy rating:     %.1f\n", movieCopy.Rating)
	fmt.Println("  (Original unchanged — structs are value types)")
}
