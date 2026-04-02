# Lesson 5: Strings, Runes, and UTF-8 Encoding

## Concept

A Go string is a read-only slice of bytes. It's not a slice of characters — this distinction matters when you work with non-ASCII text. Go strings are UTF-8 encoded by default, which means a single character like `é` or `🎉` can occupy multiple bytes.

The `len()` function on a string returns the number of bytes, not the number of characters. For the string `"café"`, `len()` returns 5 (the `é` takes 2 bytes), not 4. To count characters, you need to work with **runes** — Go's name for Unicode code points. The `utf8.RuneCountInString()` function gives you the character count.

When you range over a string with a `for range` loop, Go automatically decodes UTF-8 and gives you one rune at a time, along with its byte position. This is the idiomatic way to process strings character by character. If you index a string directly with `s[i]`, you get a byte, not a rune.

## Analogy

Think of a Go string like a bookshelf where each slot holds one byte. Most English letters fit in one slot, but characters from other languages — Chinese, Arabic, emoji — need two, three, or even four slots. If someone asks "how many slots?" (`len()`), you get the shelf count. If they ask "how many books?" (rune count), you need to walk the shelf and count the multi-slot books as single items.

## What the Code Demonstrates

The `main.go` file compares byte length vs rune count for a multi-language string, iterates over runes with `for range`, and shows the difference between byte indexing and rune iteration.

## Key Takeaways

- Go strings are byte slices, not character arrays
- `len(s)` returns byte count; `utf8.RuneCountInString(s)` returns character count
- `for i, r := range s` iterates over runes (Unicode code points)
- `s[i]` gives you a byte, not a rune
- A `rune` is an alias for `int32` and represents one Unicode code point
