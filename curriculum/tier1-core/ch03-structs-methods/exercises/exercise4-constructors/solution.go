//go:build ignore

// Exercise 4: Implement Constructor Functions — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import (
	"errors"
	"fmt"
)

// Player represents a game character.
type Player struct {
	Name   string
	Health int
	Level  int
}

// NewPlayer creates a Player with validation and defaults.
func NewPlayer(name string, health int) (*Player, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if health < 0 {
		return nil, errors.New("health cannot be negative")
	}
	if health == 0 {
		health = 100
	}
	return &Player{
		Name:   name,
		Health: health,
		Level:  1,
	}, nil
}

// Status returns a formatted player status string.
func (p *Player) Status() string {
	return fmt.Sprintf("Player: %s (Health: %d, Level: %d)", p.Name, p.Health, p.Level)
}

func main() {
	fmt.Println("Constructor Functions Demo")
	fmt.Println("==========================")

	// Default health (0 → 100)
	p1, err := NewPlayer("Hero", 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(p1.Status())
	}

	// Explicit health
	p2, err := NewPlayer("Tank", 200)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(p2.Status())
	}

	// Empty name — error
	_, err = NewPlayer("", 100)
	if err != nil {
		fmt.Printf("NewPlayer(\"\", 100): %v\n", err)
	}

	// Negative health — error
	_, err = NewPlayer("Bad", -5)
	if err != nil {
		fmt.Printf("NewPlayer(\"Bad\", -5): %v\n", err)
	}
}
