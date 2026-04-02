# Exercise 3: Use Struct Embedding for Composition

## Task

Create a program that models employees using struct embedding. Define a base `ContactInfo` struct, embed it in an `Employee` struct, and embed `Employee` in a `Manager` struct. Show how fields and methods are promoted through multiple levels.

## Requirements

- Define a `ContactInfo` struct with `Email` (string) and `Phone` (string) fields
- Define an `Employee` struct that embeds `ContactInfo` and adds `Name` (string) and `ID` (int) fields
- Define a `Manager` struct that embeds `Employee` and adds `Department` (string)
- Add a `ContactCard` method on `ContactInfo` that returns a formatted string
- Demonstrate accessing promoted fields and methods from the outermost struct

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Struct Embedding Demo
=====================
Manager: Bob (ID: 101)
Department: Engineering
Email: bob@example.com (promoted from ContactInfo)
Contact: bob@example.com / 555-0101 (promoted method)
```

## Hints

<details>
<summary>Hint 1</summary>

Embed without a field name: `type Employee struct { Name string; ID int; ContactInfo }`. This promotes `ContactInfo`'s fields to `Employee`.

</details>

<details>
<summary>Hint 2</summary>

When `Manager` embeds `Employee`, and `Employee` embeds `ContactInfo`, the `Manager` can access `ContactInfo` fields directly: `mgr.Email` works because promotion chains through both levels.

</details>
