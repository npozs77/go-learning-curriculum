# Lesson 4: JSON and YAML Struct Tags for Serialization Control

## Concept

Struct tags are metadata strings attached to struct fields. They don't affect how Go compiles or runs your code — they're read at runtime by libraries (via reflection) to control behavior like serialization. The most common use is controlling how struct fields map to JSON or YAML keys.

A struct tag is a raw string literal placed after the field type: `` `json:"field_name"` ``. The `encoding/json` package reads these tags to decide what key name to use when marshaling (Go struct → JSON) or unmarshaling (JSON → Go struct). Without tags, the JSON key matches the Go field name exactly.

Tags support options separated by commas. The `omitempty` option tells the encoder to skip the field if it has its zero value. The `-` value tells the encoder to ignore the field entirely — useful for sensitive data like passwords or API keys that should never appear in output.

## Analogy

Think of struct tags like shipping labels on boxes. The box contents (field values) stay the same, but the label tells the delivery service (JSON encoder) what name to put on the package, whether to skip empty boxes (`omitempty`), and which boxes should never be shipped at all (`-`). The labels are instructions for the handler, not for the contents.

## What the Code Demonstrates

The `main.go` file defines a `User` struct with JSON tags, marshals it to JSON to show how tags control key names, demonstrates `omitempty` behavior, and shows how the `-` tag excludes fields from output.

## Key Takeaways

- Struct tags are metadata strings that libraries read via reflection
- `json:"name"` controls the JSON key name for a field
- `json:",omitempty"` skips the field when it has its zero value
- `json:"-"` excludes the field from JSON output entirely
- Tags don't affect Go code behavior — only library behavior (like `encoding/json`)
