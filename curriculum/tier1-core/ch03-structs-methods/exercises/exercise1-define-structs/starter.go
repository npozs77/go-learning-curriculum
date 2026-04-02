// Exercise 1: Define and Use Structs
// Task: Define a Movie struct, create instances, and demonstrate copy behavior
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

// TODO: Define a Movie struct with fields:
//   Title    string
//   Director string
//   Year     int
//   Rating   float64

func main() {
	fmt.Println("Movie Collection")
	fmt.Println("================")

	// TODO: Create movie1 using named field initialization
	// Hint: movie1 := Movie{Title: "The Matrix", Director: "The Wachowskis", Year: 1999, Rating: 8.7}

	// TODO: Create movie2 using named field initialization
	// Hint: movie2 := Movie{Title: "Inception", Director: "Christopher Nolan", Year: 2010, Rating: 8.8}

	// TODO: Print each movie's details using fmt.Printf
	// Hint: fmt.Printf("%q (%d) directed by %s — Rating: %.1f\n", movie1.Title, movie1.Year, movie1.Director, movie1.Rating)

	fmt.Println()

	// TODO: Copy movie1 to a new variable, change the copy's rating to 9.0,
	// then print both ratings to show the original is unchanged
	// Hint: movieCopy := movie1

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
