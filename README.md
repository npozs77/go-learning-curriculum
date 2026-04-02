# Go Learning Curriculum

A complete, interactive Go learning course with 19 chapters across 3 tiers. Every lesson includes runnable `.go` files, exercises with starter/solution code, quizzes, cheat sheets, and real-project connections.

## Structure

| Tier | Chapters | Focus | Priority |
|------|----------|-------|----------|
| Tier 1: Core | 01–08 | Go fundamentals | HIGH |
| Tier 2: Practical | 09–15 | Backend development | MEDIUM |
| Tier 3: Expert | 16–19 | Advanced/optional | LOW |

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
