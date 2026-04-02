# Exercise 1: Explore Zero Values

## Task

Write a program that declares variables of different types without assigning values, then prints each variable's zero value alongside its type name. This reinforces Go's guarantee that every variable starts in a known state.

## Requirements

- Declare variables of types: `int`, `float64`, `bool`, `string`, `byte`, and `rune` — without assigning values
- Print each variable's type name and zero value in a formatted table
- Include a summary line stating how many types were checked

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Zero Value Explorer
===================
Type       | Zero Value
-----------|----------
int        | 0
float64    | 0
bool       | false
string     | ""
byte       | 0
rune       | 0
-----------|----------
Checked 6 types — all have defined zero values!
```

## Hints

<details>
<summary>Hint 1</summary>

Use `var x int` (no `=`) to declare a variable with its zero value. Do this for each type.

</details>

<details>
<summary>Hint 2</summary>

Use `fmt.Printf` with format verbs like `%d` for integers, `%g` for floats, `%t` for bools, and `%q` for strings (quoted).

</details>
