# Exercise 3: Work with Strings and Runes

## Task

Write a program that analyzes a string containing multi-byte Unicode characters. Count the bytes, count the runes, and print each rune with its byte position and Unicode code point.

## Requirements

- Define a string that contains both ASCII and multi-byte characters (use `"Go is fun! 🚀🌍"`)
- Print the byte length and rune count
- Iterate over the string using `for range` and print each rune's byte position, character, and Unicode code point
- Count how many runes are multi-byte (take more than 1 byte)

## Expected Behavior

When you complete the exercise and run `go run .`, you should see:

```
String Analyzer
===============
Text:       "Go is fun! 🚀🌍"
Byte count: 19
Rune count: 13

Rune Details:
  byte  0: 'G' (U+0047)
  byte  1: 'o' (U+006F)
  byte  2: ' ' (U+0020)
  byte  3: 'i' (U+0069)
  byte  4: 's' (U+0073)
  byte  5: ' ' (U+0020)
  byte  6: 'f' (U+0066)
  byte  7: 'u' (U+0075)
  byte  8: 'n' (U+006E)
  byte  9: '!' (U+0021)
  byte 10: ' ' (U+0020)
  byte 11: '🚀' (U+1F680)
  byte 15: '🌍' (U+1F30D)

Multi-byte runes: 2 out of 13
```

## Hints

<details>
<summary>Hint 1</summary>

Use `utf8.RuneCountInString(s)` from the `unicode/utf8` package to count runes. Use `len(s)` for byte count.

</details>

<details>
<summary>Hint 2</summary>

In a `for i, r := range s` loop, `i` is the byte position and `r` is the rune. Use `utf8.RuneLen(r)` to check if a rune takes more than 1 byte.

</details>
