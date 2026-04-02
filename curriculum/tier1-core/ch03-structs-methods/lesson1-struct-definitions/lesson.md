# Lesson 1: Struct Definitions and Field Access

## Concept

A struct is Go's way of grouping related data together under one name. Unlike basic types that hold a single value, a struct bundles multiple fields — each with its own name and type — into a single unit. You define a struct with the `type` keyword, then create instances and access fields with dot notation.

Structs are value types in Go. When you assign a struct to a new variable or pass it to a function, Go copies the entire thing. This is different from languages where objects are passed by reference. Understanding this copy behavior is essential — it affects how you design your programs.

Fields in a struct can be any type: basic types, other structs, slices, maps, or even function types. Exported fields (capitalized names) are visible outside the package; unexported fields (lowercase) are private to the package.

## Analogy

Think of a struct like a paper form. A "Patient Intake Form" has labeled fields: Name (string), Age (int), Insured (bool). Each form is a copy — filling in one form doesn't change another. The form template is the struct definition; each filled-out form is an instance.

## What the Code Demonstrates

The `main.go` file defines a `Book` struct with several fields, creates instances using both named and positional initialization, accesses and modifies fields with dot notation, and shows that structs are copied on assignment.

## Key Takeaways

- Structs group related fields of different types into a single unit
- Define structs with `type Name struct { ... }`
- Access fields with dot notation: `book.Title`
- Structs are value types — assignment copies the entire struct
- Prefer named field initialization (`Book{Title: "...", Author: "..."}`) over positional
