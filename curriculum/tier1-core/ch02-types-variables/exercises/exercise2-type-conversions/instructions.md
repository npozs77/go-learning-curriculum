# Exercise 2: Practice Type Conversions

## Task

Write a program that converts between numeric types and demonstrates the precision differences. You'll convert temperatures between Celsius and Fahrenheit using both `int` and `float64` to see how truncation affects results.

## Requirements

- Define a `float64` temperature in Celsius (e.g., 36.6)
- Convert Celsius to Fahrenheit using the formula: `F = C × 9/5 + 32`
- Print the result as `float64` (precise) and as `int` (truncated)
- Convert a Fahrenheit `int` value back to Celsius as `float64`
- Print all results with clear labels

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
Temperature Converter
=====================
Celsius (float64):    36.6
Fahrenheit (float64): 97.88
Fahrenheit (int):     97 (truncated)

Converting back: 97°F → 36.11°C (float64)
Precision lost:  36.6°C → 97°F → 36.11°C
```

## Hints

<details>
<summary>Hint 1</summary>

Use `float64()` and `int()` for type conversions. Remember that `int()` truncates — it doesn't round.

</details>

<details>
<summary>Hint 2</summary>

For the Fahrenheit formula, make sure you're doing float division: `celsius * 9.0 / 5.0 + 32.0`. If you use integer literals, Go may do integer division.

</details>
