# Chapter 02: Types, Variables & Zero Values

**Learning Goal:** After this chapter, you will understand Go's type system, know how to declare variables and constants, rely on zero-value semantics, convert between types safely, work with strings and runes, and use Go's control flow constructs.

**Prerequisites:** [Chapter 01: Modules, Build & Project Structure](../ch01-modules-build/)

## Lessons

1. [Lesson 1: Go's Type System and Basic Types](lesson1-basic-types/) — Learn Go's basic types: int, float64, string, bool, byte, and rune
2. [Lesson 2: Variable Declarations](lesson2-declarations/) — Master var, short declaration :=, and const
3. [Lesson 3: Zero-Value Semantics](lesson3-zero-values/) — Understand Go's guarantee that every variable starts in a known state
4. [Lesson 4: Type Conversions and Type Assertions](lesson4-type-conversions/) — Convert between types explicitly and assert interface types safely
5. [Lesson 5: Strings, Runes, and UTF-8 Encoding](lesson5-strings-runes/) — Work with multi-byte strings and Unicode code points
6. [Lesson 6: Control Flow](lesson6-control-flow/) — Use if, for, and switch — Go's single loop construct

## Exercises

1. [Exercise 1: Explore Zero Values](exercises/exercise1-zero-values/) — Declare variables without values and inspect their zero values
2. [Exercise 2: Practice Type Conversions](exercises/exercise2-type-conversions/) — Convert temperatures between Celsius and Fahrenheit
3. [Exercise 3: Work with Strings and Runes](exercises/exercise3-string-processing/) — Analyze a Unicode string's bytes and runes
4. [Exercise 4: Implement Control Flow Patterns](exercises/exercise4-control-flow/) — Write FizzBuzz using a conditionless switch

## Concepts Reinforced from Earlier Chapters

- **Module structure (Chapter 01):** Every lesson and exercise in this chapter is a self-contained Go module with its own `go.mod` — the same pattern you learned in Chapter 01. Each directory works with `go run .` independently.
