# Chapter 02 Real-Project Connection: Types, Variables & Zero Values

## How These Concepts Apply: Vocabulary Generator

In a vocabulary generator CLI application, the concepts from this chapter appear in these ways:

### Basic Types and Variable Declarations

The vocabulary generator uses Go's basic types throughout its codebase. In `internal/config/config.go`, configuration values are declared with explicit types: `var maxWords int = 50` for the word limit, `var outputFormat string = "json"` for the default format, and `var verbose bool = false` for debug output. Short declarations (`:=`) appear inside functions where the type is obvious from context, like `words := generateWords(topic, count)`.

### Zero Values and Safe Defaults

Zero values play a critical role in the vocabulary generator's configuration loading. When `internal/config/Load()` reads a YAML file, any missing fields automatically get their zero values — `0` for numeric limits, `""` for unset paths, `false` for disabled features. The code checks `if cfg.MaxWords == 0 { cfg.MaxWords = 50 }` to apply sensible defaults, relying on Go's zero-value guarantee to detect unset fields without needing nullable types or sentinel values.

### Type Conversions

The vocabulary generator converts between types when processing user input and formatting output. In `cmd/vocabgen/main.go`, command-line arguments arrive as strings and are converted to integers with `strconv.Atoi()`. In `internal/service/generator.go`, word frequency scores stored as `float64` are converted to `int` for display ranking: `rank := int(score * 100)`. The truncation behavior is intentional — fractional rankings aren't meaningful to the user.

### Strings, Runes, and UTF-8

Since the vocabulary generator supports multiple languages, string handling is rune-aware. In `internal/service/tokenizer.go`, the tokenizer iterates over input text with `for _, r := range text` to correctly split words at Unicode boundaries. The function `internal/service/truncate.go` uses `utf8.RuneCountInString(word)` to enforce maximum word length in runes (not bytes), ensuring that a 20-character limit means 20 characters regardless of whether they're ASCII or CJK.
