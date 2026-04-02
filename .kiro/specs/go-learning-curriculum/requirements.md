# Requirements Document

## Introduction

This specification defines an interactive, runnable Go learning curriculum organized into 3 tiers and 19 chapters. The curriculum is no longer a collection of flat Markdown files — it is a hands-on course where every lesson includes runnable `.go` files, every exercise has `starter.go` and `solution.go` files, and later lessons deliberately reuse and build on earlier concepts. Each lesson folder is self-contained: `go run .` works from that directory. The curriculum targets Go 1.24+ and limits third-party library references to Cobra (CLI) and rapid (property-based testing). The six Markdown files in `concept-learning-modules/` serve as reference material for content quality and style, while `concept-learning-modules/go-topic-ranking.md` defines HIGH/MEDIUM/LOW priority classifications.

## Glossary

- **Curriculum**: The complete collection of 19 chapters organized into 3 tiers, plus top-level documentation (index.md, master-guide.md, concept-map.md), all under the `curriculum/` directory.
- **Tier**: A grouping of chapters by difficulty and priority. Tier 1 is Core (HIGH priority, chapters 01–08), Tier 2 is Practical (MEDIUM priority, chapters 09–15), Tier 3 is Expert (LOW priority, chapters 16–19).
- **Chapter**: A folder (e.g., `ch01-modules-build/`) containing lesson subfolders, an exercises subfolder, README.md, quiz.md, cheat-sheet.md, and real-project.md.
- **Lesson**: A subfolder within a chapter (e.g., `lesson1-go-mod/`) containing lesson.md (explanation + analogy), one or more runnable .go files, and a go.mod file. Each lesson is self-contained and executable via `go run .` from its directory.
- **Exercise**: A subfolder within a chapter's `exercises/` directory containing instructions.md, starter.go (incomplete code with TODO comments), and solution.go (complete reference answer).
- **Starter_Code**: A .go file that compiles but is intentionally incomplete — contains TODO comments and missing implementations for the learner to fill in.
- **Solution_Code**: A .go file that compiles and runs correctly, serving as the reference answer for an exercise.
- **Quiz**: A quiz.md file in the chapter folder with exactly 5 multiple-choice questions, collapsible answers, and explanations.
- **Cheat_Sheet**: A cheat-sheet.md file in the chapter folder with a two-column Markdown table (Pattern | Code) summarizing key patterns.
- **Real_Project_Section**: A real-project.md file in the chapter folder connecting the chapter's concepts to a concrete Go project.
- **Curriculum_Index**: The top-level `curriculum/index.md` file listing all 19 chapters grouped by tier, with summaries, prerequisites, and links.
- **Master_Guide**: The `curriculum/docs/master-guide.md` file providing a complete curriculum overview.
- **Concept_Map**: The `curriculum/docs/concept-map.md` file documenting how concepts link across chapters.
- **Curriculum_Generator**: The system that organizes, validates, and produces the final curriculum output files.
- **Source_MD_Files**: The six Markdown files in `concept-learning-modules/` used as reference material for content quality, style, and topic coverage.
- **Topic_Ranking**: The priority classification (HIGH/MEDIUM/LOW) defined in `go-topic-ranking.md` that determines tier assignment.

## Requirements

### Requirement 1: Three-Tier Curriculum Structure

**User Story:** As a learner, I want the curriculum organized into three progressive tiers, so that I can focus on foundational skills first and advance to specialized topics at my own pace.

#### Acceptance Criteria

1. THE Curriculum SHALL organize all 19 chapters into exactly 3 tiers: Tier 1 (Core Go Programming, chapters 01–08), Tier 2 (Practical Go Backend Development, chapters 09–15), and Tier 3 (Expert Topics, chapters 16–19).
2. THE Curriculum SHALL place Tier 1 chapters in a `tier1-core/` directory, Tier 2 chapters in a `tier2-practical/` directory, and Tier 3 chapters in a `tier3-expert/` directory.
3. THE Curriculum SHALL assign Tier 1 chapters to HIGH-priority topics as defined by the Topic_Ranking.
4. THE Curriculum SHALL assign Tier 2 chapters to MEDIUM-priority topics as defined by the Topic_Ranking.
5. THE Curriculum SHALL assign Tier 3 chapters to LOW-priority topics as defined by the Topic_Ranking.
6. THE Curriculum SHALL designate Tier 1 chapters as mandatory and foundational for all learners.
7. THE Curriculum SHALL designate Tier 2 chapters as practical chapters that build on Tier 1 knowledge.
8. THE Curriculum SHALL designate Tier 3 chapters as optional expert-level chapters that can be consumed independently of each other.

#### Properties

- P1.1: Curriculum contains exactly 3 tier directories with correct chapter ranges (01–08, 09–15, 16–19) (Req 1, AC 1–2)
- P1.2: All Tier 1 chapters map to HIGH-priority topics (Req 1, AC 3)
- P1.3: All Tier 2 chapters map to MEDIUM-priority topics (Req 1, AC 4)
- P1.4: All Tier 3 chapters map to LOW-priority topics (Req 1, AC 5)
- P1.5: Tier 3 chapters have no inter-dependencies (Req 1, AC 8)

### Requirement 2: Chapter Folder Structure

**User Story:** As a learner, I want every chapter to follow the same predictable folder structure, so that I can navigate the curriculum efficiently and know what to expect in each chapter.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each Chapter as a folder containing: README.md, lesson subfolders, an exercises/ subfolder, quiz.md, cheat-sheet.md, and real-project.md.
2. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create a README.md with the chapter title, learning goal, prerequisites, a list of lessons with links, and a list of prior concepts reinforced.
3. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create between 6 and 10 lesson subfolders.
4. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create between 3 and 5 exercise subfolders inside the exercises/ directory.
5. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create exactly one quiz.md file with 5 quiz questions.
6. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create exactly one cheat-sheet.md file.
7. WHEN a Chapter is generated, THE Curriculum_Generator SHALL create exactly one real-project.md file.
8. WHEN a Chapter is generated, THE Curriculum_Generator SHALL list prerequisite chapters (if any) in the README.md that the learner should complete first.

#### Properties

- P2.1: Every chapter folder contains README.md, lesson subfolders, exercises/, quiz.md, cheat-sheet.md, and real-project.md (Req 2, AC 1)
- P2.2: Every chapter README.md has title, learning goal, prerequisites, lesson list, and reinforced concepts (Req 2, AC 2)
- P2.3: Every chapter has between 6 and 10 lesson subfolders (Req 2, AC 3)
- P2.4: Every chapter has between 3 and 5 exercise subfolders (Req 2, AC 4)
- P2.5: Every chapter has exactly 1 quiz.md with 5 questions (Req 2, AC 5)
- P2.6: Every chapter has exactly 1 cheat-sheet.md (Req 2, AC 6)
- P2.7: Every chapter has exactly 1 real-project.md (Req 2, AC 7)

### Requirement 3: Lesson Structure — Self-Contained Runnable Lessons

**User Story:** As a learner, I want each lesson to be a self-contained folder with an explanation and runnable Go code, so that I can read the concept and immediately run working code that demonstrates it.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each Lesson as a subfolder containing lesson.md and one or more .go files.
2. WHEN a Lesson is generated, THE Curriculum_Generator SHALL include a go.mod file in the lesson folder so that `go run .` works from that directory.
3. WHEN a Lesson is generated, THE Curriculum_Generator SHALL ensure the lesson.md contains a concise explanation of the concept and exactly one real-world analogy.
4. WHEN a Lesson is generated, THE Curriculum_Generator SHALL ensure the main .go file compiles and runs successfully with `go run .` from the lesson directory.
5. WHEN a Lesson is generated, THE Curriculum_Generator SHALL ensure the .go file produces visible output (via fmt.Println, log, or other observable behavior) that demonstrates the concept being taught.
6. WHEN a Lesson is generated, THE Curriculum_Generator SHALL ensure the .go file demonstrates exactly the concept described in lesson.md — no unrelated code.
7. THE Curriculum_Generator SHALL name lesson folders with a sequential number and descriptive slug (e.g., `lesson1-go-mod/`, `lesson2-go-sum/`).

#### Properties

- P3.1: Every lesson folder contains lesson.md and at least one .go file (Req 3, AC 1)
- P3.2: Every lesson folder contains a go.mod file (Req 3, AC 2)
- P3.3: Every lesson.md contains exactly one analogy (Req 3, AC 3)
- P3.4: Every lesson .go file compiles with `go build ./...` from its directory (Req 3, AC 4)
- P3.5: Every lesson .go file produces visible output when run (Req 3, AC 5)
- P3.6: Every lesson .go file demonstrates only the concept from its lesson.md (Req 3, AC 6)

### Requirement 4: Runnable Code Requirements

**User Story:** As a learner, I want every code file in the curriculum to compile and run, so that I can trust the examples and focus on learning rather than debugging broken code.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL ensure every .go file in a lesson folder compiles with `go build ./...` from its directory.
2. THE Curriculum_Generator SHALL ensure every solution.go file in an exercise folder compiles and runs correctly.
3. THE Curriculum_Generator SHALL ensure every starter.go file in an exercise folder compiles (even though it is incomplete).
4. WHEN a .go file uses external packages, THE Curriculum_Generator SHALL include a go.mod (and go.sum if needed) in the same directory.
5. THE Curriculum_Generator SHALL target Go 1.24+ for all .go files and use standard library packages as the primary reference.
6. WHERE third-party libraries are used, THE Curriculum_Generator SHALL limit references to Cobra for CLI construction and rapid for property-based testing.

#### Properties

- P4.1: All lesson .go files compile with `go build ./...` (Req 4, AC 1)
- P4.2: All solution.go files compile and run correctly (Req 4, AC 2)
- P4.3: All starter.go files compile (Req 4, AC 3)
- P4.4: All .go files using external packages have accompanying go.mod (Req 4, AC 4)
- P4.5: All .go files target Go 1.24+ standard library (Req 4, AC 5)
- P4.6: Only Cobra and rapid appear as third-party dependencies (Req 4, AC 6)

### Requirement 5: Exercise Structure — Starter and Solution Files

**User Story:** As a learner, I want exercises with incomplete starter code and a reference solution, so that I can practice filling in the gaps and then check my work.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each Exercise as a subfolder inside the chapter's `exercises/` directory, containing instructions.md, starter.go, and solution.go.
2. WHEN an Exercise is generated, THE Curriculum_Generator SHALL ensure instructions.md contains a descriptive title, clear task statement, and expected behavior.
3. WHEN an Exercise is generated, THE Curriculum_Generator SHALL ensure starter.go compiles but is intentionally incomplete, with TODO comments marking where the learner should write code.
4. WHEN an Exercise is generated, THE Curriculum_Generator SHALL ensure solution.go compiles and runs correctly, producing the expected output described in instructions.md.
5. THE Curriculum_Generator SHALL ensure each Exercise references concepts from the same Chapter's lessons.
6. THE Curriculum_Generator SHALL name exercise folders with a sequential number and descriptive slug (e.g., `exercise1-init-module/`, `exercise2-add-dependency/`).

#### Properties

- P5.1: Every exercise folder contains instructions.md, starter.go, and solution.go (Req 5, AC 1)
- P5.2: Every instructions.md has a title, task statement, and expected behavior (Req 5, AC 2)
- P5.3: Every starter.go compiles but contains TODO comments (Req 5, AC 3)
- P5.4: Every solution.go compiles and runs correctly (Req 5, AC 4)
- P5.5: Every exercise references concepts from its chapter's lessons (Req 5, AC 5)

### Requirement 6: Concept Reinforcement Across Chapters

**User Story:** As a learner, I want later chapters to deliberately reuse concepts from earlier chapters in their code, so that my understanding deepens progressively through practice.

#### Acceptance Criteria

1. WHEN Chapter 05 (Error Handling) exercises are generated, THE Curriculum_Generator SHALL use structs from Chapter 03 in the exercise code.
2. WHEN Chapter 06 (Interfaces) is generated, THE Curriculum_Generator SHALL reference error handling concepts from Chapter 05 to motivate the error interface.
3. WHEN Chapter 08 (Concurrency) exercises are generated, THE Curriculum_Generator SHALL use slices from Chapter 04 in the exercise code.
4. WHEN Chapter 09 (Context, Defer) is generated, THE Curriculum_Generator SHALL reference concurrency patterns from Chapter 08 for cancellation propagation.
5. WHEN Chapter 12 (Test Doubles) is generated, THE Curriculum_Generator SHALL reference interface concepts from Chapter 06 as the enabling mechanism for test doubles.
6. WHEN Chapter 13 (CLI Tools) is generated, THE Curriculum_Generator SHALL reference error handling patterns from Chapter 05 for Cobra RunE error returns.
7. WHEN Chapter 14 (HTTP Servers) is generated, THE Curriculum_Generator SHALL reference the service layer pattern from Chapter 12 for handler design.
8. WHEN a Chapter README.md is generated, THE Curriculum_Generator SHALL list which prior concepts from earlier chapters are reinforced in this chapter.

#### Properties

- P6.1: Chapter 05 exercises use structs (Ch03) (Req 6, AC 1)
- P6.2: Chapter 06 references error handling (Ch05) (Req 6, AC 2)
- P6.3: Chapter 08 exercises use slices (Ch04) (Req 6, AC 3)
- P6.4: Chapter 09 references concurrency (Ch08) (Req 6, AC 4)
- P6.5: Chapter 12 references interfaces (Ch06) (Req 6, AC 5)
- P6.6: Chapter 13 references error handling (Ch05) (Req 6, AC 6)
- P6.7: Chapter 14 references service layer (Ch12) (Req 6, AC 7)
- P6.8: Every chapter README lists reinforced prior concepts (Req 6, AC 8)

### Requirement 7: Quiz Format and Correctness

**User Story:** As a learner, I want quizzes with clear multiple-choice questions and verifiable answers, so that I can self-assess my understanding after each chapter.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each quiz.md with exactly 5 questions, each having exactly 4 answer options labeled a through d.
2. THE Curriculum_Generator SHALL provide answers for all quiz questions in a collapsible details/summary HTML section.
3. WHEN a quiz answer is provided, THE Curriculum_Generator SHALL include a one-sentence explanation of why the answer is correct.
4. THE Curriculum_Generator SHALL ensure each quiz question references concepts covered in the same Chapter's lessons.

#### Properties

- P7.1: Every quiz.md has exactly 5 questions with 4 options labeled a–d (Req 7, AC 1)
- P7.2: Quiz answers are in a collapsible details/summary section (Req 7, AC 2)
- P7.3: Every quiz answer includes a one-sentence explanation (Req 7, AC 3)
- P7.4: Every quiz question references concepts from its chapter's lessons (Req 7, AC 4)

### Requirement 8: Cheat Sheet Completeness

**User Story:** As a learner, I want a quick-reference cheat sheet per chapter, so that I can look up patterns without re-reading the full lesson.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each cheat-sheet.md as a two-column Markdown table with columns "Pattern" and "Code".
2. THE Curriculum_Generator SHALL include at least one cheat-sheet entry for each Lesson in the same Chapter.
3. THE Curriculum_Generator SHALL ensure every code snippet in the cheat-sheet is syntactically valid Go or a valid shell command.

#### Properties

- P8.1: Every cheat-sheet.md is a two-column table (Pattern | Code) (Req 8, AC 1)
- P8.2: Cheat sheet has at least as many entries as the chapter has lessons (Req 8, AC 2)
- P8.3: All cheat-sheet code snippets are syntactically valid (Req 8, AC 3)

### Requirement 9: Real-Project Connection Sections

**User Story:** As a learner, I want to see how each chapter's concepts apply to a real Go project, so that I understand practical relevance.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce each real-project.md referencing a consistent example project across all chapters.
2. WHEN a real-project.md is generated, THE Curriculum_Generator SHALL map at least 3 concepts from the chapter to specific usage patterns in the example project.
3. THE Curriculum_Generator SHALL include concrete code structure references (file paths, function names, or architectural patterns) in each real-project.md.

#### Properties

- P9.1: All real-project.md files reference the same example project (Req 9, AC 1)
- P9.2: Every real-project.md maps at least 3 concepts to the example project (Req 9, AC 2)
- P9.3: Every real-project.md includes concrete code structure references (Req 9, AC 3)

### Requirement 10: Module Sequencing and Prerequisites

**User Story:** As a learner, I want chapters ordered from foundational to advanced with clear prerequisites, so that each chapter builds on knowledge from previous ones.

#### Acceptance Criteria

1. THE Curriculum SHALL sequence Tier 1 chapters (01–08) before Tier 2 chapters (09–15), and Tier 2 chapters before Tier 3 chapters (16–19).
2. THE Curriculum SHALL sequence Chapter 01 (Modules, Build & Project Structure) as the first chapter, covering project setup before any application-level concepts.
3. THE Curriculum SHALL sequence Chapter 02 (Types, Variables & Zero Values) before Chapter 03 (Structs, Methods & Receivers), as structs depend on the type system.
4. THE Curriculum SHALL sequence Chapter 04 (Slices, Arrays & Pointers) before Chapter 08 (Concurrency), as channels and goroutines require understanding of value semantics.
5. THE Curriculum SHALL sequence Chapter 05 (Error Handling) before Chapter 06 (Interfaces & Abstraction), as the error interface is a key motivating example.
6. THE Curriculum SHALL sequence Chapter 06 (Interfaces & Abstraction) before Chapter 12 (Test Doubles & Service Architecture), as test doubles depend on interface-based design.
7. THE Curriculum SHALL sequence Chapter 07 (Testing) before any Tier 2 chapter, as testing skills are required for practical development.
8. WHEN a Tier 3 chapter is generated, THE Curriculum_Generator SHALL note that the chapter can be consumed independently of other Tier 3 chapters.

#### Properties

- P10.1: Tier 1 chapters (01–08) precede Tier 2 (09–15), which precede Tier 3 (16–19) (Req 10, AC 1)
- P10.2: Chapter numbering preserves all prerequisite orderings (02→03, 04→08, 05→06, 06→12, 07→Tier 2) (Req 10, AC 2–7)
- P10.3: Tier 3 chapters note independent consumability (Req 10, AC 8)

### Requirement 11: Curriculum Index and Documentation

**User Story:** As a learner, I want a top-level index, a master guide, and a concept map, so that I can navigate the full curriculum and understand how concepts connect.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce a `curriculum/index.md` file listing all 19 chapters grouped by tier, with one-sentence summaries, relative links, and prerequisite chapters for each entry.
2. THE Curriculum_Generator SHALL produce a `curriculum/docs/master-guide.md` file providing a complete curriculum overview including learning path recommendations, time estimates, and tier descriptions.
3. THE Curriculum_Generator SHALL produce a `curriculum/docs/concept-map.md` file documenting how concepts link across chapters, showing which later chapters build on which earlier concepts.
4. THE Curriculum_Index SHALL include a "Prerequisites" note indicating the curriculum assumes basic programming experience in at least one language.
5. THE Curriculum_Index SHALL include a tier overview section describing the purpose and scope of each tier.

#### Properties

- P11.1: index.md lists all 19 chapters with links, summaries, and prerequisites (Req 11, AC 1)
- P11.2: master-guide.md exists with curriculum overview (Req 11, AC 2)
- P11.3: concept-map.md exists with cross-chapter concept links (Req 11, AC 3)
- P11.4: index.md includes prerequisites note and tier overview (Req 11, AC 4–5)

### Requirement 12: Content Generation Pipeline

**User Story:** As a curriculum author, I want a defined content generation pipeline, so that each piece of chapter content is produced systematically and validated for quality.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL define a lesson generation task that produces lesson.md (prose + analogy) and runnable .go files for each Lesson.
2. THE Curriculum_Generator SHALL define an exercise creation task that produces instructions.md, starter.go, and solution.go for each Exercise.
3. THE Curriculum_Generator SHALL define a quiz creation task that produces quiz.md for each Chapter.
4. THE Curriculum_Generator SHALL define a cheat sheet creation task that produces cheat-sheet.md for each Chapter.
5. THE Curriculum_Generator SHALL define a compilation validation task that verifies all .go files compile with `go build ./...` from their respective directories.
6. THE Curriculum_Generator SHALL define a run validation task that verifies all lesson main.go and solution.go files produce output when executed.
7. THE Curriculum_Generator SHALL define a consistency validation task that verifies all generated content within a Chapter references the same concepts, terminology, and code patterns.

#### Properties

- P12.1: Pipeline defines discrete tasks for lessons, exercises, quizzes, cheat sheets, compilation validation, run validation, and consistency (Req 12, AC 1–7)

### Requirement 13: Design Decisions — Tone, Depth, Scope, and Interactive-First Approach

**User Story:** As a learner, I want the curriculum to maintain a consistent tone, appropriate depth per tier, and an interactive-first approach, so that the learning experience is cohesive, trustworthy, and hands-on.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce all content in a beginner-friendly tone that is clear, structured, and uses real-world analogies to explain technical concepts.
2. THE Curriculum_Generator SHALL calibrate content depth so that Tier 1 chapters cover foundational concepts with detailed explanations, Tier 2 chapters cover practical patterns with moderate detail, and Tier 3 chapters cover advanced topics with concise expert-level treatment.
3. THE Curriculum_Generator SHALL target Go version 1.24 or later and use standard library packages as the primary reference.
4. WHERE third-party libraries are referenced, THE Curriculum_Generator SHALL limit references to Cobra for CLI construction and rapid for property-based testing.
5. THE Curriculum_Generator SHALL treat the Source_MD_Files as reference material for content quality, style, and topic coverage, not as structural constraints on the curriculum organization.
6. THE Curriculum_Generator SHALL prioritize runnable code over prose — every concept introduction must be accompanied by executable Go code, not just Markdown code blocks.

#### Properties

- P13.1: All content uses beginner-friendly tone with analogies (Req 13, AC 1)
- P13.2: Tier depth calibration is maintained (detailed → moderate → concise) (Req 13, AC 2)
- P13.3: All Go code targets Go 1.24+ standard library (Req 13, AC 3)
- P13.4: Only Cobra and rapid appear as third-party library references (Req 13, AC 4)
- P13.5: Every concept has accompanying runnable .go code (Req 13, AC 6)


### Requirement 14: Final Deliverable Completeness

**User Story:** As a learner, I want the final output to be a complete, self-contained interactive Go learning course, so that I can work through the entire curriculum without needing external resources.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL produce exactly 19 Chapter folders, one for each chapter defined in the curriculum structure.
2. THE Curriculum_Generator SHALL produce the top-level documentation files: index.md, docs/master-guide.md, and docs/concept-map.md.
3. THE Curriculum_Generator SHALL ensure the final deliverable totals approximately 200+ files across 19 chapter folders, 3 documentation files, and all lesson/exercise subfolders.
4. THE Curriculum_Generator SHALL ensure the final deliverable is self-contained, requiring no external files or dependencies beyond the generated content and Go 1.24+ toolchain.
5. WHEN all Chapter folders and documentation files are generated, THE Curriculum_Generator SHALL produce a deliverable ready for learner consumption without additional editing or assembly.

#### Properties

- P14.1: Exactly 19 chapter folders exist in the curriculum directory (Req 14, AC 1)
- P14.2: All 3 documentation files exist (index.md, master-guide.md, concept-map.md) (Req 14, AC 2)
- P14.3: Total file count is approximately 200+ files (Req 14, AC 3)
- P14.4: Deliverable is self-contained (Req 14, AC 4)

### Requirement 15: Chapter 01 — Modules, Build & Project Structure Content Coverage

**User Story:** As a learner, I want Chapter 01 to cover Go modules, project structure, building, testing, and documentation tooling through runnable lessons, so that I can set up and manage a Go project.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on go.mod and module declaration in Chapter 01.
2. THE Curriculum_Generator SHALL include a Lesson on go.sum and dependency verification in Chapter 01.
3. THE Curriculum_Generator SHALL include a Lesson on the internal/ directory pattern for compiler-enforced privacy in Chapter 01.
4. THE Curriculum_Generator SHALL include a Lesson on building Go binaries and cross-compilation (GOOS, GOARCH) in Chapter 01.
5. THE Curriculum_Generator SHALL include a Lesson on running tests with go test and the ./... glob in Chapter 01.
6. THE Curriculum_Generator SHALL include a Lesson on Makefile usage for common project tasks (build, test, vet, lint, clean) in Chapter 01.
7. THE Curriculum_Generator SHALL include a Lesson on Godoc comment conventions in Chapter 01.

#### Properties

- P15.1: Chapter 01 contains lessons covering go.mod, go.sum, internal/, cross-compilation, go test, Makefile, and Godoc (Req 15, AC 1–7)

### Requirement 16: Chapter 02 — Types, Variables & Zero Values Content Coverage

**User Story:** As a learner, I want Chapter 02 to cover Go's type system, variable declarations, and zero-value semantics through runnable lessons, so that I understand the foundation all other Go constructs build on.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on Go's type system and basic types (int, string, bool, float64, byte, rune) in Chapter 02.
2. THE Curriculum_Generator SHALL include a Lesson on variable declarations (var, short declaration :=, const) in Chapter 02.
3. THE Curriculum_Generator SHALL include a Lesson on zero-value semantics (every type has a defined zero value) in Chapter 02.
4. THE Curriculum_Generator SHALL include a Lesson on type conversions and type assertions in Chapter 02.
5. THE Curriculum_Generator SHALL include a Lesson on strings, runes, and UTF-8 encoding in Chapter 02.
6. THE Curriculum_Generator SHALL include a Lesson on control flow (if, for, switch) and Go's single loop construct in Chapter 02.

#### Properties

- P16.1: Chapter 02 contains lessons covering basic types, declarations, zero values, conversions, strings/runes, and control flow (Req 16, AC 1–6)

### Requirement 17: Chapter 03 — Structs, Methods & Receivers Content Coverage

**User Story:** As a learner, I want Chapter 03 to cover struct definitions, methods, receivers, and struct tags through runnable lessons, so that I can define and work with custom data types in Go.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on struct definitions and field access in Chapter 03.
2. THE Curriculum_Generator SHALL include a Lesson on methods and value receivers vs pointer receivers in Chapter 03.
3. THE Curriculum_Generator SHALL include a Lesson on struct embedding for composition in Chapter 03.
4. THE Curriculum_Generator SHALL include a Lesson on YAML and JSON struct tags for serialization control in Chapter 03.
5. THE Curriculum_Generator SHALL include a Lesson on zero values for structs and how uninitialized fields behave in Chapter 03.
6. THE Curriculum_Generator SHALL include a Lesson on constructor functions (NewXxx pattern) in Chapter 03.

#### Properties

- P17.1: Chapter 03 contains lessons covering struct definitions, receivers, embedding, tags, zero values, and constructors (Req 17, AC 1–6)

### Requirement 18: Chapter 04 — Slices, Arrays & Pointers Content Coverage

**User Story:** As a learner, I want Chapter 04 to cover slice internals, append behavior, and pointer mechanics through runnable lessons, so that I understand Go's memory model for collections and references.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on arrays (fixed-size) vs slices (dynamic) in Chapter 04.
2. THE Curriculum_Generator SHALL include a Lesson on slice internals (pointer, length, capacity) and the slice header in Chapter 04.
3. THE Curriculum_Generator SHALL include a Lesson on append behavior, reallocation, and the gotcha of append modifying the original backing array in Chapter 04.
4. THE Curriculum_Generator SHALL include a Lesson on pointer mechanics (& and * operators, nil pointers) in Chapter 04.
5. THE Curriculum_Generator SHALL include a Lesson on pass-by-value semantics and when to use pointers vs values in Chapter 04.
6. THE Curriculum_Generator SHALL include a Lesson on maps (creation, access, deletion, zero-value behavior) in Chapter 04.

#### Properties

- P18.1: Chapter 04 contains lessons covering arrays vs slices, slice internals, append, pointers, pass-by-value, and maps (Req 18, AC 1–6)

### Requirement 19: Chapter 05 — Error Handling Content Coverage

**User Story:** As a learner, I want Chapter 05 to cover Go's error handling philosophy and patterns through runnable lessons, so that I can write robust code that handles failures explicitly.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on errors as values and the if-err-nil pattern in Chapter 05.
2. THE Curriculum_Generator SHALL include a Lesson on error wrapping with fmt.Errorf and the %w verb in Chapter 05.
3. THE Curriculum_Generator SHALL include a Lesson on errors.Is for checking specific error values in the chain in Chapter 05.
4. THE Curriculum_Generator SHALL include a Lesson on errors.As for checking error types in the chain in Chapter 05.
5. THE Curriculum_Generator SHALL include a Lesson on custom error types implementing the error interface in Chapter 05.
6. THE Curriculum_Generator SHALL include a Lesson on sentinel errors and when to define package-level error variables in Chapter 05.

#### Properties

- P19.1: Chapter 05 contains lessons covering if-err-nil, wrapping, errors.Is, errors.As, custom errors, and sentinels (Req 19, AC 1–6)

### Requirement 20: Chapter 06 — Interfaces & Abstraction Content Coverage

**User Story:** As a learner, I want Chapter 06 to cover Go's implicit interface system and interface-driven design through runnable lessons, so that I can write flexible, testable, and decoupled code.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on implicit interface satisfaction (no implements keyword) in Chapter 06.
2. THE Curriculum_Generator SHALL include a Lesson on the "accept interfaces, return structs" design principle in Chapter 06.
3. THE Curriculum_Generator SHALL include a Lesson on compile-time interface checks using var _ Interface = (*Type)(nil) in Chapter 06.
4. THE Curriculum_Generator SHALL include a Lesson on interface-driven design and how interfaces enable testability in Chapter 06.
5. THE Curriculum_Generator SHALL include a Lesson on the empty interface (any) and type switches in Chapter 06.
6. THE Curriculum_Generator SHALL include a Lesson on common standard library interfaces (io.Reader, io.Writer, fmt.Stringer, error) in Chapter 06.

#### Properties

- P20.1: Chapter 06 contains lessons covering implicit satisfaction, accept-interfaces-return-structs, compile-time checks, testability, empty interface, and stdlib interfaces (Req 20, AC 1–6)

### Requirement 21: Chapter 07 — Testing: Table-Driven & Property-Based Content Coverage

**User Story:** As a learner, I want Chapter 07 to cover table-driven tests, subtests, test helpers, and property-based testing with rapid through runnable lessons, so that I can write comprehensive and maintainable test suites.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on table-driven tests with t.Run subtests in Chapter 07.
2. THE Curriculum_Generator SHALL include a Lesson on t.Helper() for clean test failure reporting in Chapter 07.
3. THE Curriculum_Generator SHALL include a Lesson on t.Cleanup() and t.TempDir() for test resource management in Chapter 07.
4. THE Curriculum_Generator SHALL include a Lesson on property-based testing concepts and the rapid library in Chapter 07.
5. THE Curriculum_Generator SHALL include a Lesson on rapid generators (rapid.String, rapid.IntRange, rapid.SliceOfN, rapid.SampledFrom) in Chapter 07.
6. THE Curriculum_Generator SHALL include a Lesson on common property patterns (idempotence, round-trip, invariant, no-crash) in Chapter 07.
7. THE Curriculum_Generator SHALL include a Lesson comparing table-driven and property-based testing approaches (when to use each) in Chapter 07.

#### Properties

- P21.1: Chapter 07 contains lessons covering table-driven tests, t.Helper, t.Cleanup/t.TempDir, rapid concepts, rapid generators, property patterns, and comparison (Req 21, AC 1–7)

### Requirement 22: Chapter 08 — Concurrency: Goroutines, Channels & Patterns Content Coverage

**User Story:** As a learner, I want Chapter 08 to cover goroutines, channels, select, and concurrency patterns through runnable lessons, so that I can write concurrent Go programs safely.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on goroutines (go keyword, lightweight threads) in Chapter 08.
2. THE Curriculum_Generator SHALL include a Lesson on channels (unbuffered and buffered, send/receive operations) in Chapter 08.
3. THE Curriculum_Generator SHALL include a Lesson on the select statement for multiplexing channel operations in Chapter 08.
4. THE Curriculum_Generator SHALL include a Lesson on the worker pool pattern in Chapter 08.
5. THE Curriculum_Generator SHALL include a Lesson on pipeline patterns (chaining goroutines via channels) in Chapter 08.
6. THE Curriculum_Generator SHALL include a Lesson on sync.WaitGroup for coordinating goroutine completion in Chapter 08.
7. THE Curriculum_Generator SHALL include a Lesson on common concurrency pitfalls (goroutine leaks, race conditions, deadlocks) in Chapter 08.

#### Properties

- P22.1: Chapter 08 contains lessons covering goroutines, channels, select, worker pools, pipelines, WaitGroup, and pitfalls (Req 22, AC 1–7)

### Requirement 23: Chapter 09 — Context, Defer & Resource Management Content Coverage

**User Story:** As a learner, I want Chapter 09 to cover context.Context, cancellation, timeouts, and defer through runnable lessons, so that I can manage resources and request lifecycles in Go applications.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on context.Context for carrying deadlines, cancellation signals, and values in Chapter 09.
2. THE Curriculum_Generator SHALL include a Lesson on context.WithTimeout and context.WithCancel in Chapter 09.
3. THE Curriculum_Generator SHALL include a Lesson on context propagation through function call chains in Chapter 09.
4. THE Curriculum_Generator SHALL include a Lesson on nested timeouts and context hierarchies in Chapter 09.
5. THE Curriculum_Generator SHALL include a Lesson on defer for resource cleanup (files, database connections, transactions) in Chapter 09.
6. THE Curriculum_Generator SHALL include a Lesson on defer execution order (LIFO) and argument capture semantics in Chapter 09.

#### Properties

- P23.1: Chapter 09 contains lessons covering context.Context, WithTimeout/WithCancel, propagation, nested timeouts, defer cleanup, and defer LIFO (Req 23, AC 1–6)

### Requirement 24: Chapter 10 — Encoding: JSON, YAML & IO Patterns Content Coverage

**User Story:** As a learner, I want Chapter 10 to cover marshaling, unmarshaling, custom types, and IO patterns through runnable lessons, so that I can serialize data and work with files and streams in Go.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on JSON marshaling and unmarshaling with encoding/json in Chapter 10.
2. THE Curriculum_Generator SHALL include a Lesson on YAML marshaling and unmarshaling in Chapter 10.
3. THE Curriculum_Generator SHALL include a Lesson on custom MarshalJSON/UnmarshalJSON for complex types in Chapter 10.
4. THE Curriculum_Generator SHALL include a Lesson on the io.Reader and io.Writer interfaces in Chapter 10.
5. THE Curriculum_Generator SHALL include a Lesson on bufio for buffered reading and scanning in Chapter 10.
6. THE Curriculum_Generator SHALL include a Lesson on file operations (os.Open, os.Create, os.ReadFile, os.WriteFile) in Chapter 10.
7. THE Curriculum_Generator SHALL include a Lesson on strings.NewReplacer for named placeholder substitution in Chapter 10.

#### Properties

- P24.1: Chapter 10 contains lessons covering JSON, YAML, custom marshal, io.Reader/Writer, bufio, file ops, and strings.NewReplacer (Req 24, AC 1–7)

### Requirement 25: Chapter 11 — database/sql & Data Persistence Content Coverage

**User Story:** As a learner, I want Chapter 11 to cover database/sql patterns, transactions, and nullable columns through runnable lessons, so that I can build data-persistent Go applications.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on sql.Open and blank imports for driver registration in Chapter 11.
2. THE Curriculum_Generator SHALL include a Lesson on parameterized queries for SQL injection prevention in Chapter 11.
3. THE Curriculum_Generator SHALL include a Lesson on Scan for reading query results (single row and multiple rows) in Chapter 11.
4. THE Curriculum_Generator SHALL include a Lesson on transactions (BeginTx, Commit, Rollback, defer Rollback pattern) in Chapter 11.
5. THE Curriculum_Generator SHALL include a Lesson on nullable columns with sql.NullString and sql.NullInt64 in Chapter 11.
6. THE Curriculum_Generator SHALL include a Lesson on ExecContext for INSERT, UPDATE, DELETE operations in Chapter 11.

#### Properties

- P25.1: Chapter 11 contains lessons covering sql.Open, parameterized queries, Scan, transactions, nullable columns, and ExecContext (Req 25, AC 1–6)

### Requirement 26: Chapter 12 — Test Doubles & Service Architecture Content Coverage

**User Story:** As a learner, I want Chapter 12 to cover test doubles, configurable failure doubles, and the service layer pattern through runnable lessons, so that I can architect testable Go applications.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on creating test doubles without mocking frameworks (implicit interface satisfaction) in Chapter 12.
2. THE Curriculum_Generator SHALL include a Lesson on stubs, spies, and fakes as test double variants in Chapter 12.
3. THE Curriculum_Generator SHALL include a Lesson on configurable failure doubles for error testing in Chapter 12.
4. THE Curriculum_Generator SHALL include a Lesson on the service layer architectural pattern in Chapter 12.
5. THE Curriculum_Generator SHALL include a Lesson on thin handlers and fat service design (CLI and HTTP handlers as thin wrappers) in Chapter 12.
6. THE Curriculum_Generator SHALL include a Lesson on dependency injection via interfaces for testable service construction in Chapter 12.

#### Properties

- P26.1: Chapter 12 contains lessons covering test doubles, stubs/spies/fakes, failure doubles, service layer, thin handlers, and DI via interfaces (Req 26, AC 1–6)

### Requirement 27: Chapter 13 — CLI Tools: Cobra, slog & Build Injection Content Coverage

**User Story:** As a learner, I want Chapter 13 to cover Cobra CLI construction, structured logging with slog, and build-time value injection through runnable lessons, so that I can build production CLI tools.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on Cobra command structure (root command, subcommands, RunE) in Chapter 13.
2. THE Curriculum_Generator SHALL include a Lesson on persistent vs local flags in Cobra in Chapter 13.
3. THE Curriculum_Generator SHALL include a Lesson on PersistentPreRunE for shared setup logic in Chapter 13.
4. THE Curriculum_Generator SHALL include a Lesson on log/slog structured logging with levels and key-value pairs in Chapter 13.
5. THE Curriculum_Generator SHALL include a Lesson on configuring slog handlers (TextHandler, JSONHandler) and log levels in Chapter 13.
6. THE Curriculum_Generator SHALL include a Lesson on ldflags for compile-time version injection in Chapter 13.
7. THE Curriculum_Generator SHALL include a Lesson on signal handling (signal.NotifyContext) and graceful shutdown in Chapter 13.

#### Properties

- P27.1: Chapter 13 contains lessons covering Cobra commands, flags, PersistentPreRunE, slog, slog handlers, ldflags, and signal handling (Req 27, AC 1–7)

### Requirement 28: Chapter 14 — HTTP Servers, Middleware & Templates Content Coverage

**User Story:** As a learner, I want Chapter 14 to cover go:embed, HTML templates, HTTP server construction, middleware, and SSE streaming through runnable lessons, so that I can build web-serving Go applications.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on go:embed for baking files into binaries in Chapter 14.
2. THE Curriculum_Generator SHALL include a Lesson on embed.FS for runtime file access and serving static files in Chapter 14.
3. THE Curriculum_Generator SHALL include a Lesson on html/template for server-side rendering with auto-escaping in Chapter 14.
4. THE Curriculum_Generator SHALL include a Lesson on template composition with base layouts, content blocks, and ParseFS in Chapter 14.
5. THE Curriculum_Generator SHALL include a Lesson on net/http ServeMux and Go 1.22+ method-aware routing in Chapter 14.
6. THE Curriculum_Generator SHALL include a Lesson on the middleware pattern for wrapping HTTP handlers in Chapter 14.
7. THE Curriculum_Generator SHALL include a Lesson on httptest for testing HTTP handlers without a real server in Chapter 14.
8. THE Curriculum_Generator SHALL include a Lesson on Server-Sent Events (SSE) for streaming updates to browsers in Chapter 14.

#### Properties

- P28.1: Chapter 14 contains lessons covering go:embed, embed.FS, html/template, template composition, ServeMux, middleware, httptest, and SSE (Req 28, AC 1–8)

### Requirement 29: Chapter 15 — Generics: Type Parameters & Constraints Content Coverage

**User Story:** As a learner, I want Chapter 15 to cover Go generics, type parameters, and constraints through runnable lessons, so that I can write reusable type-safe code.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on type parameters and generic function syntax in Chapter 15.
2. THE Curriculum_Generator SHALL include a Lesson on type constraints and the constraints package in Chapter 15.
3. THE Curriculum_Generator SHALL include a Lesson on generic data structures (helper collections, stacks, sets) in Chapter 15.
4. THE Curriculum_Generator SHALL include a Lesson on when to use generics vs interfaces in Chapter 15.
5. THE Curriculum_Generator SHALL include a Lesson on type inference and constraint satisfaction in Chapter 15.
6. THE Curriculum_Generator SHALL include a Lesson on the comparable and any constraints in Chapter 15.

#### Properties

- P29.1: Chapter 15 contains lessons covering type parameters, constraints, generic data structures, generics vs interfaces, type inference, and comparable/any (Req 29, AC 1–6)

### Requirement 30: Chapter 16 — Performance: Benchmarks, Fuzzing & Profiling Content Coverage

**User Story:** As a learner, I want Chapter 16 to cover benchmarking, fuzzing, and profiling tools through runnable lessons, so that I can measure and optimize Go program performance.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on writing benchmarks with testing.B in Chapter 16.
2. THE Curriculum_Generator SHALL include a Lesson on fuzz testing with testing.F in Chapter 16.
3. THE Curriculum_Generator SHALL include a Lesson on pprof for CPU and memory profiling in Chapter 16.
4. THE Curriculum_Generator SHALL include a Lesson on runtime tracing with go tool trace in Chapter 16.
5. THE Curriculum_Generator SHALL include a Lesson on interpreting benchmark results and avoiding common benchmarking mistakes in Chapter 16.
6. THE Curriculum_Generator SHALL include a Lesson on the race detector (go test -race) for concurrency bug detection in Chapter 16.

#### Properties

- P30.1: Chapter 16 contains lessons covering testing.B, testing.F, pprof, go tool trace, benchmark interpretation, and race detector (Req 30, AC 1–6)

### Requirement 31: Chapter 17 — Memory Model & Escape Analysis Content Coverage

**User Story:** As a learner, I want Chapter 17 to cover heap vs stack allocation, escape analysis, and the Go memory model through runnable lessons, so that I understand how Go manages memory under the hood.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on stack vs heap allocation in Go in Chapter 17.
2. THE Curriculum_Generator SHALL include a Lesson on escape analysis and compiler decisions (go build -gcflags="-m") in Chapter 17.
3. THE Curriculum_Generator SHALL include a Lesson on the Go memory model and happens-before relationships in Chapter 17.
4. THE Curriculum_Generator SHALL include a Lesson on garbage collection basics and GC tuning (GOGC) in Chapter 17.
5. THE Curriculum_Generator SHALL include a Lesson on sync.Pool for reducing allocation pressure in Chapter 17.
6. THE Curriculum_Generator SHALL include a Lesson on common patterns that cause unnecessary heap allocations in Chapter 17.

#### Properties

- P31.1: Chapter 17 contains lessons covering stack vs heap, escape analysis, memory model, GC tuning, sync.Pool, and allocation patterns (Req 31, AC 1–6)

### Requirement 32: Chapter 18 — Advanced Build, Release & Workspaces Content Coverage

**User Story:** As a learner, I want Chapter 18 to cover goreleaser, build tags, Docker optimization, and multi-module workspaces through runnable lessons, so that I can manage complex Go build and release pipelines.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on goreleaser for automated cross-platform releases in Chapter 18.
2. THE Curriculum_Generator SHALL include a Lesson on build tags (//go:build) for conditional compilation in Chapter 18.
3. THE Curriculum_Generator SHALL include a Lesson on Docker multi-stage builds for minimal Go container images in Chapter 18.
4. THE Curriculum_Generator SHALL include a Lesson on Go workspaces (go.work) for multi-module development in Chapter 18.
5. THE Curriculum_Generator SHALL include a Lesson on build caching and reproducible builds in Chapter 18.
6. THE Curriculum_Generator SHALL include a Lesson on CI/CD pipeline patterns for Go projects in Chapter 18.

#### Properties

- P32.1: Chapter 18 contains lessons covering goreleaser, build tags, Docker multi-stage, go.work, build caching, and CI/CD patterns (Req 32, AC 1–6)

### Requirement 33: Chapter 19 — Networking, Crypto & Protocol Design Content Coverage

**User Story:** As a learner, I want Chapter 19 to cover low-level networking, cryptography basics, and protocol design through runnable lessons, so that I can build network-aware Go applications.

#### Acceptance Criteria

1. THE Curriculum_Generator SHALL include a Lesson on net.TCPConn and TCP client/server patterns in Chapter 19.
2. THE Curriculum_Generator SHALL include a Lesson on UDP networking with net.UDPConn in Chapter 19.
3. THE Curriculum_Generator SHALL include a Lesson on protocol-level design (framing, message formats, state machines) in Chapter 19.
4. THE Curriculum_Generator SHALL include a Lesson on the crypto package basics (hashing, HMAC) in Chapter 19.
5. THE Curriculum_Generator SHALL include a Lesson on TLS primitives and tls.Config in Chapter 19.
6. THE Curriculum_Generator SHALL include a Lesson on net.Listener and building custom servers in Chapter 19.

#### Properties

- P33.1: Chapter 19 contains lessons covering TCP, UDP, protocol design, crypto/hashing, TLS, and net.Listener (Req 33, AC 1–6)

## Non-Functional Requirements

### Usability

- All content uses a beginner-friendly tone with clear explanations and real-world analogies
- Every chapter follows a consistent, predictable folder structure so learners know what to expect
- Content is scannable: lesson.md files are concise, cheat sheets use tables, quizzes use collapsible sections
- Exercises include starter.go with TODO comments so learners can begin immediately
- Every lesson is self-contained: `go run .` works from the lesson directory without setup beyond having Go installed
- Lesson folders are named with sequential numbers and descriptive slugs for easy navigation

### Performance

- N/A — this is a static content generation project, not a running system. There are no runtime performance requirements.

### Maintainability

- Every chapter follows the same folder template (README.md, lesson subfolders, exercises/, quiz.md, cheat-sheet.md, real-project.md), making it easy to add or modify chapters
- Validation property-based tests catch structural regressions (file presence, compilation, section counts)
- Content is derived from a defined topic ranking and source material, making updates traceable
- Each lesson is isolated in its own folder with its own go.mod, so changes to one lesson cannot break another

### Compatibility

- All Go code targets Go 1.24+ and uses standard library packages as the primary reference
- Output is standard Markdown format for .md files, renderable in any Markdown viewer (GitHub, VS Code, terminal)
- All .go files compile and run with the standard Go toolchain (`go build`, `go run`)
- Only two third-party libraries referenced: Cobra (CLI) and rapid (property-based testing)

## Constraints

1. **No Third-Party Libraries Beyond Cobra and rapid**: All Go code examples use standard library packages. The only third-party references allowed are Cobra (CLI construction) and pgregory.net/rapid (property-based testing).
2. **Source MD Files Are Reference Material**: The six Markdown files in `concept-learning-modules/` inform content quality and style but do not dictate chapter boundaries or curriculum structure.
3. **Go 1.24+ Target Version**: All code examples, API references, and feature mentions target Go 1.24 or later.
4. **Self-Contained Lessons**: Every lesson directory must be self-contained with its own go.mod so that `go run .` works from that directory without any external setup beyond having Go installed.
5. **Compilable Starter Code**: Every starter.go must compile even though it is incomplete. Incompleteness is expressed through TODO comments and placeholder return values, not syntax errors.
6. **Tier 3 Independence**: Tier 3 chapters (16–19) must be independently consumable with no inter-dependencies within the tier.
7. **Interactive-First**: No lesson should introduce a concept without runnable Go code demonstrating it. Markdown-only explanations without accompanying .go files are not permitted for lessons.
8. **Concept Reinforcement Required**: Later chapters must reference and build on earlier concepts in their code. Exercises in later chapters must use types, patterns, or techniques from prior chapters.

## Success Criteria

The curriculum is successfully implemented when:
1. All 19 chapter folders exist under the correct tier directories (tier1-core/, tier2-practical/, tier3-expert/)
2. The top-level documentation files exist (index.md, docs/master-guide.md, docs/concept-map.md)
3. Every chapter folder contains README.md, 6–10 lesson subfolders, exercises/ with 3–5 exercise subfolders, quiz.md, cheat-sheet.md, and real-project.md
4. Every lesson subfolder contains lesson.md (with analogy), at least one .go file, and a go.mod file
5. Every exercise subfolder contains instructions.md, starter.go, and solution.go
6. All lesson .go files and solution.go files compile and run with `go run .` from their directory
7. All starter.go files compile (even though incomplete)
8. Later chapters deliberately reuse concepts from earlier chapters in their code and exercises
9. Every chapter README lists which prior concepts are reinforced
10. Cross-chapter references are present where required (Req 6)
11. All Go code targets Go 1.24+ with only Cobra and rapid as third-party dependencies
12. Validation property tests pass for all structural correctness properties
13. Total file count is approximately 200+ files across the curriculum directory
