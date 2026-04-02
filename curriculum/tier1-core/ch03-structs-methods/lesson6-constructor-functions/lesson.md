# Lesson 6: Constructor Functions (NewXxx Pattern)

## Concept

Go doesn't have constructors like Java or Python's `__init__`. Instead, the convention is to write a plain function named `NewXxx` that returns an initialized instance of your type. This function validates inputs, sets defaults, and returns either a value or a pointer depending on the use case.

The `NewXxx` pattern is used when the zero value of a struct isn't sufficient — when you need to enforce invariants, validate inputs, or perform setup that can't happen automatically. For example, `NewServer` might require a port number and return an error if the port is invalid.

Constructor functions typically return a pointer (`*MyType`) for two reasons: it signals that the caller should share the instance rather than copy it, and it allows the struct's methods to use pointer receivers consistently. When construction can fail, the function returns `(*MyType, error)` — the standard Go pattern for fallible operations.

## Analogy

Think of a constructor function like a car factory's assembly line. You don't hand someone a pile of parts (zero-value struct) and hope they assemble it correctly. Instead, the factory (`NewCar`) takes an order (parameters), validates it (error checking), assembles everything properly (field initialization), and hands back a finished car (pointer to struct). The factory ensures every car leaves in a valid state.

## What the Code Demonstrates

The `main.go` file defines a `Server` struct with a `NewServer` constructor that validates inputs and sets defaults, and a `Temperature` struct with a `NewTemperature` constructor that enforces physical constraints.

## Key Takeaways

- Go uses `NewXxx` functions instead of built-in constructors
- Use constructors when the zero value isn't sufficient or inputs need validation
- Return `(*Type, error)` when construction can fail
- Return `*Type` (pointer) to signal shared ownership and enable pointer receiver methods
- If the zero value is useful, you may not need a constructor at all
