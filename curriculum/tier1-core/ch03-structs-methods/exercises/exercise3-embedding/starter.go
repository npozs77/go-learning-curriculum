// Exercise 3: Use Struct Embedding for Composition
// Task: Model employees using struct embedding with promoted fields and methods
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

// TODO: Define a ContactInfo struct with Email (string) and Phone (string) fields

// TODO: Add a ContactCard method on ContactInfo that returns "email / phone"
// Hint: func (c ContactInfo) ContactCard() string { ... }

// TODO: Define an Employee struct that embeds ContactInfo
// and adds Name (string) and ID (int) fields

// TODO: Define a Manager struct that embeds Employee
// and adds Department (string)

func main() {
	fmt.Println("Struct Embedding Demo")
	fmt.Println("=====================")

	// TODO: Create a Manager with nested initialization
	// Hint: mgr := Manager{Department: "Engineering", Employee: Employee{Name: "Bob", ...}}

	// TODO: Print the manager's name and ID (promoted from Employee)
	// TODO: Print the department
	// TODO: Print the email (promoted from ContactInfo through Employee)
	// TODO: Print the contact card (promoted method from ContactInfo)

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
