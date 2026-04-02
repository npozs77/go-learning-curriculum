# Lesson 7: Godoc Comment Conventions

## Concept

Go generates documentation directly from comments in your source code. There's no special syntax like Javadoc's `@param` or Python's docstring format — you just write a regular comment above any exported (capitalized) name, starting with that name.

The rules are simple: start the comment with the symbol name (`// LoadConfig reads...`), use complete sentences ending with periods, and describe what the function does rather than how it works. Package-level comments start with `// Package <name>` and describe the package's purpose.

You can view documentation from the terminal with `go doc`. For example, `go doc fmt.Println` shows the docs for `fmt.Println`, and `go doc ./mypackage` shows the package overview. The `go doc` tool reads your comments and formats them as documentation.

## Analogy

Think of Godoc comments like name tags at a conference. Each name tag (comment) starts with the person's name (symbol name) and a brief description of what they do. You don't need a special badge printer (documentation tool) — just a marker and a blank tag. Anyone can read the tag and know who this person is and what they do.

## What the Code Demonstrates

The `main.go` file defines several well-documented functions and types following Godoc conventions, then prints the documentation rules and shows how to view docs with `go doc`.

## Key Takeaways

- Start comments with the symbol name: `// FuncName does...`
- Use complete sentences ending with periods
- View docs with `go doc <package>` or `go doc <package.Symbol>`
