# Chapter 03: Structs, Methods & Receivers

**Learning Goal:** After this chapter, you will be able to define custom data types with structs, attach behavior using methods with value and pointer receivers, compose types through embedding, control serialization with struct tags, and create validated instances using constructor functions.

**Prerequisites:** [Chapter 02: Types, Variables & Zero Values](../ch02-types-variables/)

## Lessons

1. [Lesson 1: Struct Definitions and Field Access](lesson1-struct-definitions/) — Define structs, create instances, and understand copy-on-assignment behavior
2. [Lesson 2: Methods and Receivers](lesson2-methods-receivers/) — Attach methods to types using value and pointer receivers
3. [Lesson 3: Struct Embedding](lesson3-struct-embedding/) — Compose types through embedding with field and method promotion
4. [Lesson 4: Struct Tags](lesson4-struct-tags/) — Control JSON serialization with struct tags, omitempty, and field exclusion
5. [Lesson 5: Zero-Value Structs](lesson5-zero-value-structs/) — Understand how uninitialized struct fields behave and design useful zero values
6. [Lesson 6: Constructor Functions](lesson6-constructor-functions/) — Use the NewXxx pattern to create validated struct instances

## Exercises

1. [Exercise 1: Define and Use Structs](exercises/exercise1-define-structs/) — Create a Movie struct and demonstrate copy behavior
2. [Exercise 2: Methods with Receivers](exercises/exercise2-methods/) — Build a BankAccount with value and pointer receiver methods
3. [Exercise 3: Struct Embedding](exercises/exercise3-embedding/) — Model employees using multi-level struct embedding
4. [Exercise 4: Constructor Functions](exercises/exercise4-constructors/) — Implement a validated NewPlayer constructor

## Concepts Reinforced from Earlier Chapters

- **Types and zero values (Chapter 02):** Struct fields follow the same zero-value rules you learned in Chapter 02 — an `int` field starts at `0`, a `string` at `""`, a `bool` at `false`. Lesson 5 builds directly on this foundation, showing how zero-value semantics apply to composite types, not just basic ones.
