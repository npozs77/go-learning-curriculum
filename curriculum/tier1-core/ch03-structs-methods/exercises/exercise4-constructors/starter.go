// Exercise 4: Implement Constructor Functions
// Task: Create a Player struct with a NewPlayer constructor that validates inputs
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import (
	"fmt"
)

// Player represents a game character.
type Player struct {
	Name   string
	Health int
	Level  int
}

// TODO: Implement NewPlayer(name string, health int) (*Player, error)
// - Return error if name is empty
// - Return error if health is negative
// - Default health to 100 if 0 is passed
// - Set Level to 1 for all new players
// Hint: import "errors" and use errors.New("message")

// TODO: Implement a Status method on *Player that returns a formatted string
// like "Player: Hero (Health: 100, Level: 1)"
// Hint: func (p *Player) Status() string { ... }

func main() {
	fmt.Println("Constructor Functions Demo")
	fmt.Println("==========================")

	// TODO: Create a player with name "Hero" and health 0 (should default to 100)
	// TODO: Create a player with name "Tank" and health 200
	// TODO: Try creating a player with empty name — print the error
	// TODO: Try creating a player with negative health — print the error

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
