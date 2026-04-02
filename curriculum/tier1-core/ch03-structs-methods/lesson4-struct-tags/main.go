// Lesson 4: JSON and YAML Struct Tags for Serialization Control
// Demonstrates: struct tags with encoding/json — key naming, omitempty, and field exclusion
//
// Run: go run .
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents a user profile with JSON tags.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"` // never include in JSON
	Bio      string `json:"bio,omitempty"`
}

func main() {
	fmt.Println("=== Struct Tags for Serialization ===")
	fmt.Println()

	// Full user — all fields populated
	user1 := User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Age:      30,
		Password: "s3cret",
		Bio:      "Go developer",
	}

	data1, err := json.MarshalIndent(user1, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Full user (Password excluded by json:\"-\"):")
	fmt.Println(string(data1))

	fmt.Println()

	// User with zero-value fields — omitempty skips them
	user2 := User{
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "hidden",
		// Age is 0 (zero value) — omitempty will skip it
		// Bio is "" (zero value) — omitempty will skip it
	}

	data2, err := json.MarshalIndent(user2, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User with zero-value fields (omitempty skips Age and Bio):")
	fmt.Println(string(data2))

	fmt.Println()
	fmt.Println("Key insight: struct tags control serialization — not Go behavior.")
}
