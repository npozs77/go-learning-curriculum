# Chapter 02 Quiz: Types, Variables & Zero Values

Test your understanding of Go's type system, variable declarations, and control flow.

---

**Question 1:** What is the zero value of a `string` variable in Go?

- a) `nil`
- b) `"null"`
- c) `""` (empty string)
- d) `" "` (single space)

---

**Question 2:** Which declaration is only valid inside a function?

- a) `var x int = 10`
- b) `const x = 10`
- c) `x := 10`
- d) `var x = 10`

---

**Question 3:** What does `len("café")` return in Go?

- a) 4 (the number of characters)
- b) 5 (the number of bytes, since `é` is 2 bytes in UTF-8)
- c) 8 (two bytes per character)
- d) An error, because the string contains non-ASCII characters

---

**Question 4:** What happens when you convert a `float64` value of `3.9` to `int`?

- a) It rounds up to `4`
- b) It rounds to the nearest integer: `4`
- c) It truncates to `3`
- d) It causes a compile error

---

**Question 5:** How many loop constructs does Go have?

- a) Three: `for`, `while`, and `do-while`
- b) Two: `for` and `while`
- c) One: `for` (which covers all loop patterns)
- d) Four: `for`, `while`, `loop`, and `foreach`

---

<details>
<summary>Answers</summary>

1. **c)** — The zero value of a `string` is the empty string `""`, not `nil` (which applies to pointers, slices, maps, and interfaces).
2. **c)** — The short declaration `:=` can only be used inside functions; `var` and `const` work at both package and function level.
3. **b)** — `len()` returns the byte count of a string; the `é` character takes 2 bytes in UTF-8, so `"café"` is 5 bytes total.
4. **c)** — Converting `float64` to `int` truncates the decimal part; `int(3.9)` gives `3`, not `4`.
5. **c)** — Go has a single loop construct `for` that serves as a classic for-loop, a while-loop, and an infinite loop.

</details>
