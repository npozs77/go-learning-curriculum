# Lesson 4: Type Conversions and Type Assertions

## Concept

Go does not perform implicit type conversions. If you have an `int` and need a `float64`, you must explicitly convert it: `float64(myInt)`. This strictness prevents subtle bugs where a language silently converts types and produces unexpected results.

The syntax for type conversion is `TargetType(value)`. You can convert between numeric types (`int` to `float64`, `float64` to `int`), between `string` and `[]byte`, and between `string` and `[]rune`. Be aware that converting `float64` to `int` truncates the decimal — `int(3.9)` gives you `3`, not `4`.

Type assertions are different from conversions. They apply to interface values and extract the concrete type stored inside. The syntax is `value.(TargetType)`. A type assertion can panic if the type doesn't match, so Go provides a safe two-value form: `v, ok := value.(TargetType)` where `ok` is `false` if the assertion fails.

## Analogy

Type conversion is like exchanging currency at an airport. You have euros and need dollars — you go to the exchange counter (the conversion function) and explicitly ask for the swap. The exchange might lose some precision (like cents being rounded), but you always know it's happening. Go never secretly converts your euros to dollars behind your back.

## What the Code Demonstrates

The `main.go` file shows numeric conversions between `int` and `float64`, string-to-byte-slice conversions, and a safe type assertion on an interface value.

## Key Takeaways

- Go requires explicit type conversions — no implicit coercion
- Syntax: `TargetType(value)` for conversions (e.g., `float64(42)`)
- Converting `float64` to `int` truncates the decimal part
- Type assertions extract concrete types from interfaces: `v, ok := i.(string)`
- Always use the two-value form of type assertions to avoid panics
