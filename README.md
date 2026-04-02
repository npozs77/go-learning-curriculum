# Go Learning Curriculum

A complete, interactive Go learning course with 19 chapters across 3 tiers. Every lesson includes runnable `.go` files, exercises with starter/solution code, quizzes, cheat sheets, and real-project connections.

## Curriculum

### Tier 1: Core Go Programming

Start here. These chapters cover the fundamentals every Go developer needs.

| # | Chapter | What You'll Learn |
|---|---------|-------------------|
| 01 | [Modules, Build & Project Structure](curriculum/tier1-core/ch01-modules-build/) | Set up, build, and organize Go projects |
| 02 | [Types, Variables & Zero Values](curriculum/tier1-core/ch02-types-variables/) | Go's type system, declarations, and zero-value semantics |
| 03 | [Structs, Methods & Receivers](curriculum/tier1-core/ch03-structs-methods/) | Define custom types with behavior |
| 04 | [Slices, Arrays & Pointers](curriculum/tier1-core/ch04-slices-arrays-pointers/) | Collections, memory layout, and reference mechanics |
| 05 | [Error Handling](curriculum/tier1-core/ch05-error-handling/) | Handle failures explicitly with wrapping and inspection |
| 06 | [Interfaces & Abstraction](curriculum/tier1-core/ch06-interfaces-abstraction/) | Write flexible, testable code with implicit interfaces |
| 07 | [Testing: Table-Driven & Property-Based](curriculum/tier1-core/ch07-testing/) | Build robust test suites with subtests and rapid |
| 08 | [Concurrency: Goroutines, Channels & Patterns](curriculum/tier1-core/ch08-concurrency/) | Write concurrent programs with goroutines and channels |

### Tier 2: Practical Go Backend Development

Real-world skills for building production Go applications.

| # | Chapter | What You'll Learn |
|---|---------|-------------------|
| 09 | [Context, Defer & Resource Management](curriculum/tier2-practical/ch09-context-defer/) | Manage timeouts, cancellation, and cleanup |
| 10 | [Encoding: JSON, YAML & IO Patterns](curriculum/tier2-practical/ch10-encoding-io/) | Serialize data and work with files and streams |
| 11 | [database/sql & Data Persistence](curriculum/tier2-practical/ch11-database-sql/) | Build data layers with queries and transactions |
| 12 | [Test Doubles & Service Architecture](curriculum/tier2-practical/ch12-test-doubles-service/) | Architect testable apps with fakes, stubs, and service layers |
| 13 | [CLI Tools: Cobra, slog & Build Injection](curriculum/tier2-practical/ch13-cli-tools/) | Build production CLI tools with logging and versioning |
| 14 | [HTTP Servers, Middleware & Templates](curriculum/tier2-practical/ch14-http-servers/) | Serve web apps with routing, templates, and SSE |
| 15 | [Generics: Type Parameters & Constraints](curriculum/tier2-practical/ch15-generics/) | Write reusable type-safe code |

### Tier 3: Expert Topics (Optional)

Advanced material for when you want to go deeper. Each chapter is independent.

| # | Chapter | What You'll Learn |
|---|---------|-------------------|
| 16 | [Performance: Benchmarks, Fuzzing & Profiling](curriculum/tier3-expert/ch16-performance/) | Measure and optimize Go code |
| 17 | [Memory Model & Escape Analysis](curriculum/tier3-expert/ch17-memory-model/) | Understand how Go manages memory under the hood |
| 18 | [Advanced Build, Release & Workspaces](curriculum/tier3-expert/ch18-build-release/) | Master build tags, goreleaser, and multi-module workspaces |
| 19 | [Networking, Crypto & Protocol Design](curriculum/tier3-expert/ch19-networking-crypto/) | Build TCP/UDP servers and use TLS |

## Prerequisites

- Go 1.24+ installed
- Basic programming experience in any language
- A terminal and text editor

## Quick Start

```bash
# Clone the repo
git clone https://github.com/npozs77/Go-Learning-Curriculum.git
cd Go-Learning-Curriculum

# Start with Chapter 01, Lesson 1
cd curriculum/tier1-core/ch01-modules-build/lesson1-go-mod/
cat lesson.md        # Read the concept
go run .             # Run the demo
```

## How to Use

1. Open a chapter folder (start with `tier1-core/ch01-modules-build/`)
2. Read the chapter `README.md` for an overview and learning goal
3. Work through lessons in order — read `lesson.md`, then run `main.go`
4. Try the exercises — edit `starter.go`, then check against `solution.go`
5. Take the quiz in `quiz.md` to test yourself
6. Use `cheat-sheet.md` as a quick reference anytime

## Project Layout

```
curriculum/
├── index.md                    # Course navigation
├── docs/                       # Master guide, concept map
├── tier1-core/                 # Chapters 01–08
├── tier2-practical/            # Chapters 09–15
└── tier3-expert/               # Chapters 16–19

concept-learning-modules/       # Source reference material
.kiro/specs/                    # SDD specification
.kiro/steering/                 # Project rules
```

## Spec-Driven Development

This curriculum was designed using SDD methodology. See `.kiro/specs/go-learning-curriculum/` for requirements, design, and tasks.

## License

MIT License — see [LICENSE](LICENSE) for details.
