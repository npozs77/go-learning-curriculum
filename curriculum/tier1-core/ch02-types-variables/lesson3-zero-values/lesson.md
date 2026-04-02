# Lesson 3: Zero-Value Semantics

## Concept

In Go, every variable has a well-defined initial value even if you don't assign one. This is called the **zero value**. There are no uninitialized variables in Go — the language guarantees a safe default for every type.

The zero values are predictable: `0` for numeric types (`int`, `float64`, `byte`, `rune`), `false` for `bool`, `""` (empty string) for `string`, and `nil` for pointers, slices, maps, channels, interfaces, and functions. This means you can declare a variable and immediately use it without worrying about garbage data.

Zero values are a deliberate design choice. They eliminate an entire class of bugs — the "uninitialized variable" problems common in C and C++. In Go, a freshly declared `int` is always `0`, a `string` is always empty, and a `bool` is always `false`. Your code can rely on this.

## Analogy

Imagine moving into a new apartment. In some cities, you get the apartment "as-is" — leftover furniture, mystery stains, who knows what's in the closets. Go is like an apartment that's always professionally cleaned before you move in. Every room starts in a known, clean state. The bedroom has a bed frame (zero value), the kitchen has empty shelves (empty string), and the lights are off (false). You know exactly what you're working with from day one.

## What the Code Demonstrates

The `main.go` file declares variables of every basic type without assigning values, then prints each one to show Go's zero-value guarantees.

## Key Takeaways

- Every Go type has a defined zero value — no uninitialized variables
- Numeric types zero to `0`, bools to `false`, strings to `""`
- Pointers, slices, maps, channels, and interfaces zero to `nil`
- Zero values eliminate "uninitialized variable" bugs by design
