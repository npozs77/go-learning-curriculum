# Implementation Plan: Go Learning Curriculum

## Overview

Generate a complete, interactive Go learning curriculum with 19 chapters across 3 tiers, ~200+ files total. Each chapter contains self-contained lesson folders (lesson.md + main.go + go.mod), exercise folders (instructions.md + starter.go + solution.go + go.mod), and chapter-level files (README.md, quiz.md, cheat-sheet.md, real-project.md). Chapters are generated in tier order to ensure concept reinforcement works correctly. All Go code targets Go 1.24+ with only Cobra and rapid as third-party dependencies.

## Tasks

### Phase 1: Tier 1 — Core Go Programming (Chapters 01–08)

- [ ] 1. Generate Chapter 01: Modules, Build & Project Structure (7 lessons, 3 exercises)
  - [ ] 1.1 Create chapter folder structure at `curriculum/tier1-core/ch01-modules-build/` with lesson dirs (lesson1–lesson7), exercises/ dir (exercise1–exercise3)
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 1.2 Generate lesson subfolders: lesson1-go-mod, lesson2-go-sum, lesson3-internal-pattern, lesson4-cross-compilation, lesson5-go-test, lesson6-makefile, lesson7-godoc
    - Each subfolder gets lesson.md (concept + one analogy), main.go (compiles, runs, produces output), go.mod (module name matches folder, go 1.24)
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 4.1, 4.4, 4.5, 15.1–15.7_
  - [ ] 1.3 Generate exercise subfolders: exercise1-init-module, exercise2-add-dependency, exercise3-cross-compile
    - Each subfolder gets instructions.md (title, task, expected behavior, collapsible hints), starter.go (compiles, has TODOs), solution.go (compiles, runs, correct output), go.mod
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 4.2, 4.3_
  - [ ] 1.4 Generate quiz.md with 5 questions (4 options a–d each, collapsible answers with explanations)
    - _Requirements: 7.1, 7.2, 7.3, 7.4_
  - [ ] 1.5 Generate cheat-sheet.md as two-column table (Pattern | Code), ≥7 rows (one per lesson)
    - _Requirements: 8.1, 8.2, 8.3_
  - [ ] 1.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings and concrete code references
    - _Requirements: 9.1, 9.2, 9.3_
  - [ ] 1.7 Generate README.md with title, learning goal, prerequisites (None), lesson list with links, exercise list with links, reinforced concepts section ("first chapter — no prior concepts")
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 1.8 Validate chapter: all .go files compile (`go build ./...`), all lesson main.go and solution.go produce output, all starter.go compile and contain TODOs
    - _Requirements: 4.1, 4.2, 4.3, 12.5, 12.6_

- [ ] 2. Generate Chapter 02: Types, Variables & Zero Values (6 lessons, 4 exercises)
  - [ ] 2.1 Create chapter folder structure at `curriculum/tier1-core/ch02-types-variables/` with lesson dirs (lesson1–lesson6), exercises/ dir (exercise1–exercise4)
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 2.2 Generate lesson subfolders: lesson1-basic-types, lesson2-declarations, lesson3-zero-values, lesson4-type-conversions, lesson5-strings-runes, lesson6-control-flow
    - Each subfolder gets lesson.md (concept + one analogy), main.go (compiles, runs, produces output), go.mod
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 16.1–16.6_
  - [ ] 2.3 Generate exercise subfolders (4 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 2.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 2.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 2.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 2.7 Generate README.md with prerequisites (Ch01), reinforced concepts: Ch01 module structure
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 2.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 3. Generate Chapter 03: Structs, Methods & Receivers (6 lessons, 4 exercises)
  - [ ] 3.1 Create chapter folder structure at `curriculum/tier1-core/ch03-structs-methods/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 3.2 Generate lesson subfolders: lesson1-struct-definitions, lesson2-methods-receivers, lesson3-struct-embedding, lesson4-struct-tags, lesson5-zero-value-structs, lesson6-constructor-functions
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 17.1–17.6_
  - [ ] 3.3 Generate exercise subfolders (4 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 3.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 3.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 3.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 3.7 Generate README.md with prerequisites (Ch02), reinforced concepts: Ch02 types and zero values
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 3.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 4. Generate Chapter 04: Slices, Arrays & Pointers (6 lessons, 4 exercises)
  - [ ] 4.1 Create chapter folder structure at `curriculum/tier1-core/ch04-slices-arrays-pointers/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 4.2 Generate lesson subfolders: lesson1-arrays-vs-slices, lesson2-slice-internals, lesson3-append-behavior, lesson4-pointer-mechanics, lesson5-pass-by-value, lesson6-maps
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 18.1–18.6_
  - [ ] 4.3 Generate exercise subfolders (4 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 4.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 4.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 4.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 4.7 Generate README.md with prerequisites (Ch02), reinforced concepts: Ch02 value types
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 4.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 5. Generate Chapter 05: Error Handling (6 lessons, 4 exercises)
  - [ ] 5.1 Create chapter folder structure at `curriculum/tier1-core/ch05-error-handling/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 5.2 Generate lesson subfolders: lesson1-errors-as-values, lesson2-error-wrapping, lesson3-errors-is, lesson4-errors-as, lesson5-custom-error-types, lesson6-sentinel-errors
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 19.1–19.6_
  - [ ] 5.3 Generate exercise subfolders (4 exercises) — exercises MUST use structs from Ch03 for concept reinforcement
    - Each exercise uses struct types (e.g., `type ValidationError struct { Field string; Message string }`) to reinforce Ch03 concepts
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.1_
  - [ ] 5.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 5.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 5.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 5.7 Generate README.md with prerequisites (Ch03, Ch04), reinforced concepts: Ch03 structs (custom error types are structs)
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 5.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs, exercises use struct types from Ch03
    - _Requirements: 4.1–4.3, 6.1, 12.5, 12.6_

- [ ] 6. Generate Chapter 06: Interfaces & Abstraction (6 lessons, 3 exercises)
  - [ ] 6.1 Create chapter folder structure at `curriculum/tier1-core/ch06-interfaces-abstraction/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 6.2 Generate lesson subfolders: lesson1-implicit-satisfaction, lesson2-accept-interfaces-return-structs, lesson3-compile-time-checks, lesson4-interface-testability, lesson5-empty-interface, lesson6-stdlib-interfaces
    - Lessons MUST reference error handling from Ch05 to motivate the error interface
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 6.2, 20.1–20.6_
  - [ ] 6.3 Generate exercise subfolders (3 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 6.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 6.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 6.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 6.7 Generate README.md with prerequisites (Ch05), reinforced concepts: Ch05 error handling (error is an interface)
    - _Requirements: 2.2, 2.8, 6.2, 6.8_
  - [ ] 6.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 7. Generate Chapter 07: Testing: Table-Driven & Property-Based (7 lessons, 4 exercises)
  - [ ] 7.1 Create chapter folder structure at `curriculum/tier1-core/ch07-testing/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 7.2 Generate lesson subfolders: lesson1-table-driven-tests, lesson2-t-helper, lesson3-t-cleanup-tempdir, lesson4-rapid-concepts, lesson5-rapid-generators, lesson6-property-patterns, lesson7-table-vs-property
    - Lesson .go files include _test.go files where appropriate; rapid lessons require go.mod with `pgregory.net/rapid` dependency
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 4.6, 21.1–21.7_
  - [ ] 7.3 Generate exercise subfolders (4 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 7.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 7.5 Generate cheat-sheet.md (≥7 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 7.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 7.7 Generate README.md with prerequisites (Ch06), reinforced concepts: Ch02–Ch06 (all prior concepts appear in test examples)
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 7.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 8. Generate Chapter 08: Concurrency: Goroutines, Channels & Patterns (7 lessons, 4 exercises)
  - [ ] 8.1 Create chapter folder structure at `curriculum/tier1-core/ch08-concurrency/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 8.2 Generate lesson subfolders: lesson1-goroutines, lesson2-channels, lesson3-select, lesson4-worker-pools, lesson5-pipelines, lesson6-waitgroup, lesson7-concurrency-pitfalls
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 22.1–22.7_
  - [ ] 8.3 Generate exercise subfolders (4 exercises) — exercises MUST use slices from Ch04 for concept reinforcement
    - Exercises process slices concurrently (e.g., parallel map over a slice, fan-out workers processing slice elements)
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.3_
  - [ ] 8.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 8.5 Generate cheat-sheet.md (≥7 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 8.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 8.7 Generate README.md with prerequisites (Ch04, Ch07), reinforced concepts: Ch04 slices (concurrent processing of collections)
    - _Requirements: 2.2, 2.8, 6.3, 6.8_
  - [ ] 8.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs, exercises use slices from Ch04
    - _Requirements: 4.1–4.3, 6.3, 12.5, 12.6_

- [ ] 9. Checkpoint — Tier 1 Complete
  - Ensure all 8 Tier 1 chapters compile and validate. Verify concept reinforcement links (Ch05→Ch03 structs, Ch06→Ch05 errors, Ch08→Ch04 slices). Ask the user if questions arise.

### Phase 2: Tier 2 — Practical Go Backend Development (Chapters 09–15)

- [ ] 10. Generate Chapter 09: Context, Defer & Resource Management (6 lessons, 3 exercises)
  - [ ] 10.1 Create chapter folder structure at `curriculum/tier2-practical/ch09-context-defer/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 10.2 Generate lesson subfolders: lesson1-context-basics, lesson2-with-timeout-cancel, lesson3-context-propagation, lesson4-nested-timeouts, lesson5-defer-cleanup, lesson6-defer-lifo
    - Lessons MUST reference concurrency patterns from Ch08 for cancellation propagation
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 6.4, 23.1–23.6_
  - [ ] 10.3 Generate exercise subfolders (3 exercises) — exercises use channels with context for timeout patterns (Ch08 reinforcement)
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.4_
  - [ ] 10.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 10.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 10.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 10.7 Generate README.md with prerequisites (Ch08), reinforced concepts: Ch08 concurrency (context cancellation for goroutines)
    - _Requirements: 2.2, 2.8, 6.4, 6.8_
  - [ ] 10.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 11. Generate Chapter 10: Encoding: JSON, YAML & IO Patterns (7 lessons, 4 exercises)
  - [ ] 11.1 Create chapter folder structure at `curriculum/tier2-practical/ch10-encoding-io/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 11.2 Generate lesson subfolders: lesson1-json-marshal, lesson2-yaml-marshal, lesson3-custom-marshal, lesson4-io-reader-writer, lesson5-bufio, lesson6-file-operations, lesson7-strings-replacer
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 24.1–24.7_
  - [ ] 11.3 Generate exercise subfolders (4 exercises) — exercises use struct tags from Ch03 for JSON/YAML marshaling
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 11.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 11.5 Generate cheat-sheet.md (≥7 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 11.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 11.7 Generate README.md with prerequisites (Ch03), reinforced concepts: Ch03 structs + tags (JSON/YAML marshaling uses struct tags)
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 11.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 12. Generate Chapter 11: database/sql & Data Persistence (6 lessons, 3 exercises)
  - [ ] 12.1 Create chapter folder structure at `curriculum/tier2-practical/ch11-database-sql/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 12.2 Generate lesson subfolders: lesson1-sql-open, lesson2-parameterized-queries, lesson3-scan-results, lesson4-transactions, lesson5-nullable-columns, lesson6-exec-context
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 25.1–25.6_
  - [ ] 12.3 Generate exercise subfolders (3 exercises) — exercises reference error handling (Ch05) and context (Ch09)
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 12.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 12.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 12.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 12.7 Generate README.md with prerequisites (Ch09), reinforced concepts: Ch05 error handling (database operations return errors), Ch09 context (query context)
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 12.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 13. Generate Chapter 12: Test Doubles & Service Architecture (6 lessons, 3 exercises)
  - [ ] 13.1 Create chapter folder structure at `curriculum/tier2-practical/ch12-test-doubles-service/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 13.2 Generate lesson subfolders: lesson1-test-doubles-no-framework, lesson2-stubs-spies-fakes, lesson3-failure-doubles, lesson4-service-layer, lesson5-thin-handlers, lesson6-di-via-interfaces
    - Lessons MUST reference interface concepts from Ch06 as the enabling mechanism for test doubles
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 6.5, 26.1–26.6_
  - [ ] 13.3 Generate exercise subfolders (3 exercises) — exercises create doubles by implementing interfaces from Ch06
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.5_
  - [ ] 13.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 13.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 13.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 13.7 Generate README.md with prerequisites (Ch06, Ch07), reinforced concepts: Ch06 interfaces (test doubles implement interfaces)
    - _Requirements: 2.2, 2.8, 6.5, 6.8_
  - [ ] 13.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 14. Generate Chapter 13: CLI Tools: Cobra, slog & Build Injection (7 lessons, 4 exercises)
  - [ ] 14.1 Create chapter folder structure at `curriculum/tier2-practical/ch13-cli-tools/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 14.2 Generate lesson subfolders: lesson1-cobra-commands, lesson2-cobra-flags, lesson3-persistent-pre-run, lesson4-slog-basics, lesson5-slog-handlers, lesson6-ldflags, lesson7-signal-handling
    - Lessons MUST reference error handling patterns from Ch05 for Cobra RunE error returns
    - Cobra lessons require go.mod with `github.com/spf13/cobra` dependency
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 4.6, 6.6, 27.1–27.7_
  - [ ] 14.3 Generate exercise subfolders (4 exercises) — exercises use `errors.Is` for error checking in CLI commands (Ch05 reinforcement)
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.6_
  - [ ] 14.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 14.5 Generate cheat-sheet.md (≥7 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 14.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 14.7 Generate README.md with prerequisites (Ch05), reinforced concepts: Ch05 error handling (Cobra RunE returns errors)
    - _Requirements: 2.2, 2.8, 6.6, 6.8_
  - [ ] 14.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 15. Generate Chapter 14: HTTP Servers, Middleware & Templates (8 lessons, 4 exercises)
  - [ ] 15.1 Create chapter folder structure at `curriculum/tier2-practical/ch14-http-servers/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 15.2 Generate lesson subfolders: lesson1-go-embed, lesson2-embed-fs, lesson3-html-template, lesson4-template-composition, lesson5-servemux-routing, lesson6-middleware, lesson7-httptest, lesson8-sse
    - Lessons MUST reference the service layer pattern from Ch12 for handler design
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 6.7, 28.1–28.8_
  - [ ] 15.3 Generate exercise subfolders (4 exercises) — exercises test handlers using service interface doubles (Ch12 reinforcement)
    - _Requirements: 5.1–5.6, 4.2, 4.3, 6.7_
  - [ ] 15.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 15.5 Generate cheat-sheet.md (≥8 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 15.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 15.7 Generate README.md with prerequisites (Ch12), reinforced concepts: Ch12 service layer (handlers call services), Ch06 interfaces (handler dependencies)
    - _Requirements: 2.2, 2.8, 6.7, 6.8_
  - [ ] 15.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 16. Generate Chapter 15: Generics: Type Parameters & Constraints (6 lessons, 3 exercises)
  - [ ] 16.1 Create chapter folder structure at `curriculum/tier2-practical/ch15-generics/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 16.2 Generate lesson subfolders: lesson1-type-parameters, lesson2-constraints, lesson3-generic-data-structures, lesson4-generics-vs-interfaces, lesson5-type-inference, lesson6-comparable-any
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 29.1–29.6_
  - [ ] 16.3 Generate exercise subfolders (3 exercises) — exercises use generic slice utilities (Ch04 reinforcement) and compare constraints vs interfaces (Ch06)
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 16.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 16.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 16.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 16.7 Generate README.md with prerequisites (Ch06), reinforced concepts: Ch04 slices (generic slice utilities), Ch06 interfaces (constraints vs interfaces)
    - _Requirements: 2.2, 2.8, 6.8_
  - [ ] 16.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 17. Checkpoint — Tier 2 Complete
  - Ensure all 7 Tier 2 chapters compile and validate. Verify concept reinforcement links (Ch09→Ch08 concurrency, Ch12→Ch06 interfaces, Ch13→Ch05 errors, Ch14→Ch12 service layer). Ask the user if questions arise.

### Phase 3: Tier 3 — Expert Topics (Chapters 16–19)

- [ ] 18. Generate Chapter 16: Performance: Benchmarks, Fuzzing & Profiling (6 lessons, 3 exercises)
  - [ ] 18.1 Create chapter folder structure at `curriculum/tier3-expert/ch16-performance/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 18.2 Generate lesson subfolders: lesson1-benchmarks, lesson2-fuzz-testing, lesson3-pprof, lesson4-runtime-tracing, lesson5-benchmark-interpretation, lesson6-race-detector
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 30.1–30.6_
  - [ ] 18.3 Generate exercise subfolders (3 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 18.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 18.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 18.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 18.7 Generate README.md with prerequisites (Ch07, Ch08), reinforced concepts: Ch07 testing (benchmarks extend testing.B), Ch08 concurrency (race detector). Note: this chapter can be consumed independently of other Tier 3 chapters.
    - _Requirements: 2.2, 2.8, 6.8, 1.8, 10.8_
  - [ ] 18.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 19. Generate Chapter 17: Memory Model & Escape Analysis (6 lessons, 3 exercises)
  - [ ] 19.1 Create chapter folder structure at `curriculum/tier3-expert/ch17-memory-model/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 19.2 Generate lesson subfolders: lesson1-stack-vs-heap, lesson2-escape-analysis, lesson3-memory-model, lesson4-gc-tuning, lesson5-sync-pool, lesson6-allocation-patterns
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 31.1–31.6_
  - [ ] 19.3 Generate exercise subfolders (3 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 19.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 19.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 19.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 19.7 Generate README.md with prerequisites (Ch04), reinforced concepts: Ch04 pointers (stack vs heap for pointer escapes). Note: this chapter can be consumed independently of other Tier 3 chapters.
    - _Requirements: 2.2, 2.8, 6.8, 1.8, 10.8_
  - [ ] 19.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 20. Generate Chapter 18: Advanced Build, Release & Workspaces (6 lessons, 3 exercises)
  - [ ] 20.1 Create chapter folder structure at `curriculum/tier3-expert/ch18-build-release/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 20.2 Generate lesson subfolders: lesson1-goreleaser, lesson2-build-tags, lesson3-docker-multistage, lesson4-go-workspaces, lesson5-build-caching, lesson6-ci-cd-patterns
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 32.1–32.6_
  - [ ] 20.3 Generate exercise subfolders (3 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 20.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 20.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 20.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 20.7 Generate README.md with prerequisites (Ch01), reinforced concepts: Ch01 modules (workspaces extend module system). Note: this chapter can be consumed independently of other Tier 3 chapters.
    - _Requirements: 2.2, 2.8, 6.8, 1.8, 10.8_
  - [ ] 20.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 21. Generate Chapter 19: Networking, Crypto & Protocol Design (6 lessons, 3 exercises)
  - [ ] 21.1 Create chapter folder structure at `curriculum/tier3-expert/ch19-networking-crypto/`
    - _Requirements: 1.1, 1.2, 2.1_
  - [ ] 21.2 Generate lesson subfolders: lesson1-tcp-patterns, lesson2-udp-networking, lesson3-protocol-design, lesson4-crypto-hashing, lesson5-tls-config, lesson6-net-listener
    - _Requirements: 3.1–3.7, 4.1, 4.4, 4.5, 33.1–33.6_
  - [ ] 21.3 Generate exercise subfolders (3 exercises) with instructions.md, starter.go, solution.go, go.mod each
    - _Requirements: 5.1–5.6, 4.2, 4.3_
  - [ ] 21.4 Generate quiz.md (5 questions, collapsible answers)
    - _Requirements: 7.1–7.4_
  - [ ] 21.5 Generate cheat-sheet.md (≥6 rows)
    - _Requirements: 8.1–8.3_
  - [ ] 21.6 Generate real-project.md referencing "vocabulary generator" with ≥3 concept mappings
    - _Requirements: 9.1–9.3_
  - [ ] 21.7 Generate README.md with prerequisites (Ch08, Ch09), reinforced concepts: Ch08 concurrency (goroutines for server connections), Ch09 context (connection timeouts). Note: this chapter can be consumed independently of other Tier 3 chapters.
    - _Requirements: 2.2, 2.8, 6.8, 1.8, 10.8_
  - [ ] 21.8 Validate chapter: all .go files compile, solutions run, starters compile with TODOs
    - _Requirements: 4.1–4.3, 12.5, 12.6_

- [ ] 22. Checkpoint — Tier 3 Complete
  - Ensure all 4 Tier 3 chapters compile and validate. Verify each Tier 3 chapter notes independent consumability and has no inter-Tier-3 prerequisites. Ask the user if questions arise.

### Phase 4: Documentation

- [ ] 23. Generate top-level documentation
  - [ ] 23.1 Generate `curriculum/index.md` listing all 19 chapters grouped by tier
    - Include prerequisites note (basic programming experience in at least one language)
    - Include tier overview section describing purpose and scope of each tier
    - Each chapter entry has: relative link to chapter folder, one-sentence summary, prerequisite chapter numbers
    - Tier 3 section notes chapters can be consumed independently
    - _Requirements: 11.1, 11.4, 11.5, 1.8, 10.8_
  - [ ] 23.2 Generate `curriculum/docs/master-guide.md`
    - Complete curriculum overview, learning path recommendations, time estimates per tier and per chapter, tier descriptions
    - _Requirements: 11.2_
  - [ ] 23.3 Generate `curriculum/docs/concept-map.md`
    - Document how concepts link across chapters with a detailed connections table
    - Show concept flow: Ch03→Ch05 (structs→errors), Ch04→Ch08 (slices→concurrency), Ch05→Ch06 (errors→interfaces), Ch06→Ch12 (interfaces→test doubles), Ch08→Ch09 (concurrency→context), Ch05→Ch13 (errors→CLI), Ch12→Ch14 (service layer→HTTP)
    - _Requirements: 11.3, 6.1–6.8_
  - [ ] 23.4 Validate documentation: all 19 chapter links in index.md resolve to existing folders, master-guide.md and concept-map.md exist and are non-empty
    - _Requirements: 14.2, 14.4_

- [ ] 24. Checkpoint — Documentation Complete
  - Ensure index.md lists all 19 chapters with correct relative links. Verify master-guide.md and concept-map.md exist. Ask the user if questions arise.

### Phase 5: Validation

- [ ] 25. Final structural and compilation validation
  - [ ] 25.1 Verify tier directory structure: exactly 3 tier directories, 8 chapters in tier1-core, 7 in tier2-practical, 4 in tier3-expert
    - _Requirements: 1.1, 1.2, 14.1_
    - _Validates: Property 1_
  - [ ] 25.2 Verify chapter folder completeness: every chapter has README.md, quiz.md, cheat-sheet.md, real-project.md, 6–10 lesson subfolders, exercises/ with 3–5 exercise subfolders
    - _Requirements: 2.1, 2.3, 2.4, 2.5, 2.6, 2.7_
    - _Validates: Property 4_
  - [ ] 25.3 Verify lesson folder completeness: every lesson has lesson.md (with Analogy section), ≥1 .go file, go.mod (go 1.24)
    - _Requirements: 3.1, 3.2, 3.3, 4.4, 4.5_
    - _Validates: Property 6_
  - [ ] 25.4 Verify exercise folder completeness: every exercise has instructions.md (title, task, expected behavior), starter.go (with TODOs), solution.go, go.mod
    - _Requirements: 5.1, 5.2, 5.3_
    - _Validates: Property 7_
  - [ ] 25.5 Compile all lesson .go files (`go build ./...` from each lesson directory) and verify all produce output (`go run .`)
    - _Requirements: 3.4, 3.5, 4.1_
    - _Validates: Properties 8_
  - [ ] 25.6 Compile all starter.go files and verify they contain TODO comments; compile and run all solution.go files and verify they produce output
    - _Requirements: 4.2, 4.3, 5.3, 5.4_
    - _Validates: Properties 9, 10_
  - [ ] 25.7 Verify third-party library constraint: only Cobra and rapid appear in go.mod require directives across the entire curriculum
    - _Requirements: 4.6, 13.4_
    - _Validates: Property 15_
  - [ ] 25.8 Verify total file count ≥ 200 and all internal links are relative paths within curriculum/
    - _Requirements: 14.3, 14.4_
    - _Validates: Properties 16, 17_

- [ ] 26. Final checkpoint — All validation complete
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Chapters must be generated in tier order (Tier 1 → Tier 2 → Tier 3) because later chapters reference earlier concepts in code
- Each chapter task follows the 8-step pipeline from the design: folder structure → lessons → exercises → quiz → cheat sheet → real-project → README → validate
- Concept reinforcement cross-references are called out explicitly in the relevant chapter tasks (Ch05→Ch03, Ch06→Ch05, Ch08→Ch04, Ch09→Ch08, Ch12→Ch06, Ch13→Ch05, Ch14→Ch12)
- All Go code targets Go 1.24+ with only Cobra and rapid as third-party dependencies
- The "vocabulary generator" is the consistent real-project reference across all 19 chapters
- Tier 3 chapters (16–19) must each note they can be consumed independently of other Tier 3 chapters
- Checkpoints at task 9 (Tier 1), 17 (Tier 2), 22 (Tier 3), 24 (Docs), and 26 (Final) ensure incremental validation
