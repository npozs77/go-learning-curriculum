# Lesson 2: Variable Declarations

## Concept

Go gives you three ways to introduce variables: `var`, the short declaration `:=`, and `const`. Each serves a different purpose, and knowing when to use which is a core Go skill.

The `var` keyword is the most explicit form. You can specify the type, provide an initial value, or both: `var count int = 10`. If you provide a value, Go can infer the type: `var count = 10`. If you only specify the type, the variable gets its zero value: `var count int` gives you `0`.

The short declaration `:=` is the most common form inside functions. It declares and initializes in one step: `count := 10`. Go infers the type from the value on the right. You can't use `:=` outside a function — that's where `var` is required.

Constants declared with `const` are values that never change. They must be set at compile time — you can't assign the result of a function call to a constant. Go constants are untyped by default, which means they adapt to the context where they're used.

## Analogy

Think of `var` as filling out a formal application form — you write the name, the type, and the value in separate fields. The short declaration `:=` is like a sticky note — quick, informal, and you just write the value (the type is obvious from context). A `const` is like a plaque on the wall — once it's engraved, nobody can change it.

## What the Code Demonstrates

The `main.go` file shows all three declaration styles, demonstrates type inference, and illustrates how constants differ from variables.

## Key Takeaways

- `var x int = 10` — explicit type and value
- `var x = 10` — type inferred from value
- `x := 10` — short declaration, only inside functions
- `const Pi = 3.14159` — immutable, set at compile time
- Prefer `:=` inside functions for brevity; use `var` for package-level or zero-value declarations
