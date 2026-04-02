# Lesson 3: Struct Embedding for Composition

## Concept

Go doesn't have inheritance. Instead, it uses composition through struct embedding. When you embed one struct inside another without giving it a field name, the outer struct "promotes" all the inner struct's fields and methods — you can access them directly as if they belonged to the outer struct.

Embedding is Go's answer to code reuse. Rather than building deep class hierarchies, you compose small, focused types together. A `Manager` embeds an `Employee` and gains all its fields and methods. You can still access the embedded struct explicitly (`m.Employee.Name`) or use the promoted shorthand (`m.Name`).

This is not inheritance — there's no "is-a" relationship. A `Manager` is not an `Employee` in the type system. It's a "has-a" relationship: a `Manager` has an `Employee` inside it. The distinction matters when you work with interfaces later.

## Analogy

Think of embedding like a smartphone case with a built-in wallet. The case (outer struct) doesn't inherit from the wallet — it contains one. But you can reach into the card slots directly without saying "case, then wallet, then card." The wallet's features are promoted to the case level for convenience, but the wallet is still a separate thing inside.

## What the Code Demonstrates

The `main.go` file defines an `Address` struct and a `Person` struct that embeds `Address`. It shows how embedded fields are promoted and accessible directly, and how methods on the embedded type are also promoted.

## Key Takeaways

- Embedding is composition, not inheritance — it's a "has-a" relationship
- Embedded struct fields and methods are promoted to the outer struct
- Access promoted fields directly (`p.City`) or explicitly (`p.Address.City`)
- You can embed multiple structs in one type
- If two embedded structs have a field with the same name, you must use the explicit form to disambiguate
