# Lesson 5: Zero Values for Structs

## Concept

In Chapter 02, you learned that every Go type has a zero value. Structs follow the same rule: the zero value of a struct is a struct where every field is set to its own zero value. An `int` field starts at `0`, a `string` field starts at `""`, a `bool` field starts at `false`, and a pointer field starts at `nil`.

This means you can declare a struct variable without initializing it and immediately use it — no null pointer exceptions, no "undefined" surprises. The struct is always in a valid, predictable state. This is a powerful design tool: if you design your structs so that zero values are meaningful defaults, callers can use them without any setup.

The Go standard library uses this pattern extensively. `sync.Mutex` works with its zero value — you don't need to call a constructor. `bytes.Buffer` is ready to use immediately. When you design your own structs, ask: "Does the zero value make sense as a default?" If yes, you've created a more ergonomic API.

## Analogy

Think of a zero-value struct like a brand-new notebook. Every page is blank (zero value), but the notebook is fully functional — you can start writing immediately. You don't need to "initialize" the notebook or call a setup function. The blank state is the starting state, and it's perfectly usable.

## What the Code Demonstrates

The `main.go` file declares structs without initialization, inspects their zero-value fields, and shows a practical example of a struct designed to be useful at its zero value.

## Key Takeaways

- A struct's zero value has every field set to its type's zero value
- You can use a zero-value struct immediately — no initialization required
- Design structs so their zero value is a useful default when possible
- The standard library uses this pattern: `sync.Mutex{}`, `bytes.Buffer{}` work out of the box
- Zero-value design reduces the need for constructor functions
