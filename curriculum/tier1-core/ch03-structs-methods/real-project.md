# Chapter 03 Real-Project Connection: Structs, Methods & Receivers

## How These Concepts Apply: Vocabulary Generator

In a vocabulary generator CLI application, the concepts from this chapter appear in these ways:

### Struct Definitions and Methods

The vocabulary generator's core data model is built on structs. In `internal/model/word.go`, the `WordEntry` struct groups a word with its translations, example sentences, and difficulty level: `type WordEntry struct { Word string; Translations []string; Difficulty int }`. A `Summary()` value receiver method returns a formatted one-line description for display. A `MarkLearned()` pointer receiver method updates the entry's internal state to track progress — pointer receiver because it mutates the struct.

### Struct Embedding for Composition

The vocabulary generator uses embedding to compose specialized types from simpler ones. In `internal/model/quiz.go`, a `QuizQuestion` struct embeds `WordEntry` to inherit all word data, then adds quiz-specific fields like `Options []string` and `CorrectIndex int`. This means `QuizQuestion` can call `Summary()` directly (promoted from `WordEntry`) while also having its own `CheckAnswer(index int) bool` method. The composition avoids duplicating word fields across multiple types.

### Struct Tags for JSON Serialization

The vocabulary generator saves and loads word lists as JSON files. In `internal/model/word.go`, the `WordEntry` struct uses JSON tags to control the output format: `` `json:"word"` `` for the word field, `` `json:"translations"` `` for the translations slice, and `` `json:"difficulty,omitempty"` `` to skip the difficulty field when it's zero (unrated words). The `internal/storage/json.go` file uses `json.Marshal` and `json.Unmarshal` with these tagged structs to persist the word collection to disk.

### Constructor Functions

The `NewWordEntry` constructor in `internal/model/word.go` validates that the word string is non-empty and that at least one translation is provided before returning a `*WordEntry`. This prevents invalid entries from entering the system. Similarly, `NewQuizSession` in `internal/service/quiz.go` takes a word list and difficulty filter, validates the inputs, shuffles the entries, and returns a ready-to-use `*QuizSession` — ensuring every quiz starts in a valid state.
