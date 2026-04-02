// Lesson 3: Struct Embedding for Composition
// Demonstrates: embedding structs, field promotion, and method promotion
//
// Run: go run .
package main

import "fmt"

// Address holds a street address.
type Address struct {
	Street string
	City   string
	State  string
}

// FullAddress returns a formatted address string.
func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.State)
}

// Person embeds Address — all Address fields and methods are promoted.
type Person struct {
	Name string
	Age  int
	Address // embedded (no field name)
}

func main() {
	fmt.Println("=== Struct Embedding for Composition ===")
	fmt.Println()

	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "Portland",
			State:  "OR",
		},
	}

	// Access promoted fields directly
	fmt.Printf("Name:   %s\n", p.Name)
	fmt.Printf("Age:    %d\n", p.Age)
	fmt.Printf("City:   %s (promoted from Address)\n", p.City)
	fmt.Printf("Street: %s (promoted from Address)\n", p.Street)

	fmt.Println()

	// Access via the explicit embedded field name
	fmt.Printf("Explicit: p.Address.City = %s\n", p.Address.City)

	fmt.Println()

	// Promoted method from Address
	fmt.Printf("Full address: %s\n", p.FullAddress())

	fmt.Println()
	fmt.Println("Key insight: embedding promotes fields and methods — composition, not inheritance.")
}
