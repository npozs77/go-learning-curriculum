# Chapter 02 Cheat Sheet: Types, Variables & Zero Values

| Pattern | Code |
|---------|------|
| Basic types | `int`, `float64`, `string`, `bool`, `byte`, `rune` |
| Print a variable's type | `fmt.Printf("%T", x)` |
| Explicit var declaration | `var count int = 42` |
| Short declaration (functions only) | `count := 42` |
| Constant declaration | `const Pi = 3.14159` |
| Zero value for numeric types | `var n int` → `0` |
| Zero value for string | `var s string` → `""` |
| Zero value for bool | `var b bool` → `false` |
| Type conversion (int → float64) | `f := float64(myInt)` |
| Type assertion (safe form) | `s, ok := val.(string)` |
| Count runes in a string | `utf8.RuneCountInString(s)` |
| Iterate over runes | `for i, r := range s { ... }` |
| Classic for loop | `for i := 0; i < 10; i++ { ... }` |
| While-style for loop | `for condition { ... }` |
| If with initializer | `if v := compute(); v > 0 { ... }` |
| Conditionless switch | `switch { case x > 0: ... }` |
