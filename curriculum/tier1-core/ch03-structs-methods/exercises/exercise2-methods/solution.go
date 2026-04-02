//go:build ignore

// Exercise 2: Add Methods with Value and Pointer Receivers — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

// BankAccount represents a simple bank account.
type BankAccount struct {
	Owner   string
	Balance float64
}

// Summary returns a formatted account summary.
// Value receiver — only reads data.
func (a BankAccount) Summary() string {
	return fmt.Sprintf("%s's account: $%.2f", a.Owner, a.Balance)
}

// Deposit adds the given amount to the balance.
// Pointer receiver — modifies the original.
func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

// Withdraw subtracts the given amount if sufficient funds exist.
// Pointer receiver — modifies the original. Returns true if successful.
func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}

func main() {
	fmt.Println("Bank Account Demo")
	fmt.Println("=================")

	acct := BankAccount{Owner: "Alice", Balance: 100.00}
	fmt.Println(acct.Summary())

	acct.Deposit(50.00)
	fmt.Printf("After depositing $50.00: $%.2f\n", acct.Balance)

	acct.Withdraw(30.00)
	fmt.Printf("After withdrawing $30.00: $%.2f\n", acct.Balance)

	if !acct.Withdraw(200.00) {
		fmt.Printf("Withdraw $200.00: insufficient funds\n")
	}

	fmt.Printf("Final: %s\n", acct.Summary())
}
