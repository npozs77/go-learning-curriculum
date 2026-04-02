# Exercise 2: Add Methods with Value and Pointer Receivers

## Task

Create a `BankAccount` struct with methods that demonstrate the difference between value receivers (for reading) and pointer receivers (for modifying). Implement deposit, withdraw, and balance display methods.

## Requirements

- Define a `BankAccount` struct with `Owner` (string) and `Balance` (float64) fields
- Implement a `Summary` method with a value receiver that returns a formatted string
- Implement a `Deposit` method with a pointer receiver that adds to the balance
- Implement a `Withdraw` method with a pointer receiver that subtracts from the balance (only if sufficient funds)
- Show that `Deposit` and `Withdraw` modify the original struct

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Bank Account Demo
=================
Alice's account: $100.00
After depositing $50.00: $150.00
After withdrawing $30.00: $120.00
Withdraw $200.00: insufficient funds
Final: Alice's account: $120.00
```

## Hints

<details>
<summary>Hint 1</summary>

Use a pointer receiver for `Deposit` and `Withdraw` so they modify the original balance: `func (a *BankAccount) Deposit(amount float64)`.

</details>

<details>
<summary>Hint 2</summary>

For `Withdraw`, check if the balance is sufficient before subtracting. Return a `bool` to indicate success or failure, or print a message directly.

</details>
