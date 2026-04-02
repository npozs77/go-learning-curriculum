//go:build ignore

// Exercise 3: Use Struct Embedding for Composition — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

// ContactInfo holds contact details.
type ContactInfo struct {
	Email string
	Phone string
}

// ContactCard returns a formatted contact string.
func (c ContactInfo) ContactCard() string {
	return fmt.Sprintf("%s / %s", c.Email, c.Phone)
}

// Employee embeds ContactInfo and adds employee-specific fields.
type Employee struct {
	Name string
	ID   int
	ContactInfo
}

// Manager embeds Employee and adds a department.
type Manager struct {
	Department string
	Employee
}

func main() {
	fmt.Println("Struct Embedding Demo")
	fmt.Println("=====================")

	mgr := Manager{
		Department: "Engineering",
		Employee: Employee{
			Name: "Bob",
			ID:   101,
			ContactInfo: ContactInfo{
				Email: "bob@example.com",
				Phone: "555-0101",
			},
		},
	}

	// Access promoted fields from multiple embedding levels
	fmt.Printf("Manager: %s (ID: %d)\n", mgr.Name, mgr.ID)
	fmt.Printf("Department: %s\n", mgr.Department)
	fmt.Printf("Email: %s (promoted from ContactInfo)\n", mgr.Email)
	fmt.Printf("Contact: %s (promoted method)\n", mgr.ContactCard())
}
