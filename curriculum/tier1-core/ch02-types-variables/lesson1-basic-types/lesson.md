# Lesson 1: Go's Type System and Basic Types

## Concept

Go is a statically typed language — every variable has a type known at compile time, and that type never changes. This catches bugs early: if you try to add a string to an integer, the compiler tells you before your program ever runs.

Go's basic types cover the essentials. `int` holds whole numbers (sized to your platform — 64-bit on modern machines). `float64` handles decimal numbers with double precision. `string` holds text as an immutable sequence of bytes. `bool` is either `true` or `false`. Two less obvious types round out the set: `byte` (an alias for `uint8`, used for raw data) and `rune` (an alias for `int32`, used for Unicode code points).

Unlike dynamically typed languages where a variable can hold a string one moment and a number the next, Go locks in the type at declaration. This rigidity is a feature — it makes your code predictable and lets the compiler optimize aggressively.

## Analogy

Think of Go's type system like labeled containers in a kitchen. You have a flour jar, a sugar jar, and a salt jar. Each container only holds one kind of ingredient. You can't accidentally put salt in the sugar jar — the label (the type) prevents mix-ups. The compiler is the kitchen inspector who checks every jar before you start cooking.

## What the Code Demonstrates

The `main.go` file declares variables of each basic type, prints their values and types using `fmt.Printf` with the `%T` verb, and shows how the compiler enforces type safety.

## Key Takeaways

- Go has six essential basic types: `int`, `float64`, `string`, `bool`, `byte`, and `rune`
- Every variable has a fixed type determined at compile time
- `byte` is an alias for `uint8`; `rune` is an alias for `int32`
- The `%T` format verb in `fmt.Printf` prints a variable's type
