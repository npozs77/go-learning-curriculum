// Lesson 7: Godoc Comment Conventions
// Demonstrates: writing documentation comments that follow Go conventions
//
// Run: go run .
// Docs: go doc -all .
package main

import "fmt"

// Greeting represents a message with a sender and text.
type Greeting struct {
	Sender string
	Text   string
}

// NewGreeting creates a Greeting with the given sender and text.
func NewGreeting(sender, text string) Greeting {
	return Greeting{Sender: sender, Text: text}
}

// String returns a formatted representation of the Greeting.
func (g Greeting) String() string {
	return fmt.Sprintf("[%s]: %s", g.Sender, g.Text)
}

// FormatWelcome returns a welcome message for the given name.
// It prepends "Welcome, " to the name and appends an exclamation mark.
func FormatWelcome(name string) string {
	return fmt.Sprintf("Welcome, %s!", name)
}

func main() {
	fmt.Println("=== Godoc Comment Conventions ===")
	fmt.Println()

	// Show the documented types in action
	g := NewGreeting("Go", "Documentation is just comments!")
	fmt.Println(g)
	fmt.Println(FormatWelcome("Gopher"))
	fmt.Println()

	// Print the rules
	fmt.Println("Godoc rules:")
	fmt.Println("  1. Start with the symbol name: // FuncName does...")
	fmt.Println("  2. Use complete sentences ending with periods.")
	fmt.Println("  3. Package comments: // Package name provides...")
	fmt.Println("  4. No @param or @return tags — plain English.")
	fmt.Println()

	fmt.Println("Good examples:")
	fmt.Println(`  // NewGreeting creates a Greeting with the given sender and text.`)
	fmt.Println(`  // FormatWelcome returns a welcome message for the given name.`)
	fmt.Println()

	fmt.Println("Bad examples:")
	fmt.Println(`  // This function creates a greeting  ← doesn't start with name`)
	fmt.Println(`  // @param sender string              ← not Go style`)
	fmt.Println()

	fmt.Println("View docs from the terminal:")
	fmt.Println("  go doc .                ← package overview")
	fmt.Println("  go doc -all .           ← all exported symbols")
	fmt.Println("  go doc fmt.Println      ← specific function")
}
