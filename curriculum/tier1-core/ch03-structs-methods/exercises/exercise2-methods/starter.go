// Exercise 2: Add Methods with Value and Pointer Receivers
// Task: Create a BankAccount with deposit, withdraw, and summary methods
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

// BankAccount represents a simple bank account.
type BankAccount struct {
	Owner   string
	Balance float64
}

// TODO: Implement a Summary method with a value receiver
// that returns a formatted string like "Alice's account: $100.00"
// Hint: func (a BankAccount) Summary() string { ... }

// TODO: Implement a Deposit method with a pointer receiver
// that adds the given amount to the balance
// Hint: func (a *BankAccount) Deposit(amount float64) { ... }

// TODO: Implement a Withdraw method with a pointer receiver
// that subtracts the amount if sufficient funds exist
// Return true if successful, false if insufficient funds
// Hint: func (a *BankAccount) Withdraw(amount float64) bool { ... }

func main() {
	fmt.Println("Bank Account Demo")
	fmt.Println("=================")

	// TODO: Create an account for "Alice" with balance 100.00
	// TODO: Print the account summary
	// TODO: Deposit 50.00 and print the new balance
	// TODO: Withdraw 30.00 and print the new balance
	// TODO: Try to withdraw 200.00 and print "insufficient funds" if it fails
	// TODO: Print the final summary

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
