# Lesson 6: Control Flow — if, for, switch

## Concept

Go has three control flow statements: `if`, `for`, and `switch`. Notably, Go has only one loop construct — `for` — which replaces `while`, `do-while`, and traditional `for` loops from other languages.

The `if` statement doesn't require parentheses around the condition, but braces are mandatory. Go also supports an initializer in the `if` statement: `if err := doSomething(); err != nil { ... }`. The variable declared in the initializer is scoped to the `if`/`else` block.

The `for` loop comes in three forms. The classic form: `for i := 0; i < 10; i++`. The while-style form: `for condition { ... }`. The infinite loop: `for { ... }` (use `break` to exit). The `for range` form iterates over slices, maps, strings, and channels.

Go's `switch` is cleaner than in most languages — cases don't fall through by default (no `break` needed). You can switch on any comparable type, and cases can be expressions. A `switch` without a condition acts like a chain of `if/else if` statements.

## Analogy

Think of Go's `for` loop like a Swiss Army knife for repetition. Other languages give you a separate tool for each job — a `while` wrench, a `for` screwdriver, a `do-while` bottle opener. Go hands you one tool that transforms to fit every situation. Fewer tools to learn, same jobs done.

## What the Code Demonstrates

The `main.go` file shows `if` with an initializer, all three `for` loop forms, and a `switch` statement with multiple cases and a conditionless switch.

## Key Takeaways

- `if` doesn't need parentheses; braces are required
- `if` supports an initializer: `if v := compute(); v > 0 { ... }`
- `for` is Go's only loop — it covers classic, while-style, and infinite loops
- `switch` cases don't fall through by default — no `break` needed
- A conditionless `switch` replaces long `if/else if` chains
