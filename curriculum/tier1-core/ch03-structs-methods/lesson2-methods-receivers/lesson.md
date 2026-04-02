# Lesson 2: Methods and Value Receivers vs Pointer Receivers

## Concept

In Go, you attach behavior to a type by defining methods. A method is just a function with a special parameter called a receiver that appears before the function name. The receiver binds the method to a type, so you call it with dot notation: `myRect.Area()`.

There are two kinds of receivers. A value receiver (`func (r Rectangle) Area() float64`) gets a copy of the struct — any changes inside the method don't affect the original. A pointer receiver (`func (r *Rectangle) Scale(factor float64)`) gets a pointer to the original struct — changes inside the method modify the caller's data.

The rule of thumb: use a pointer receiver when the method needs to modify the struct, or when the struct is large and copying would be expensive. Use a value receiver when the method only reads data and the struct is small. In practice, most methods use pointer receivers for consistency — if any method on a type uses a pointer receiver, convention says all methods on that type should too.

## Analogy

Think of a value receiver like photocopying a document before marking it up — your marks don't change the original. A pointer receiver is like editing the original document directly — every change is permanent. If you need to update a patient's chart, you work on the original (pointer receiver). If you just need to read the chart, a photocopy is fine (value receiver).

## What the Code Demonstrates

The `main.go` file defines a `Rectangle` struct with both a value receiver method (`Area`) and a pointer receiver method (`Scale`). It shows that `Scale` modifies the original struct while `Area` works on a copy.

## Key Takeaways

- Methods are functions with a receiver parameter before the function name
- Value receivers get a copy — they can't modify the original struct
- Pointer receivers get a reference — they can modify the original struct
- Use pointer receivers when the method mutates state or the struct is large
- Go automatically takes the address or dereferences when calling methods
