# Chapter 03 Quiz: Structs, Methods & Receivers

Test your understanding of struct definitions, methods, receivers, embedding, and constructor patterns.

---

**Question 1:** What happens when you assign one struct variable to another in Go?

- a) Both variables point to the same data (shared reference)
- b) The entire struct is copied — changes to one don't affect the other
- c) Only the pointer fields are copied; value fields are shared
- d) It causes a compile error unless you use the `copy()` function

---

**Question 2:** When should you use a pointer receiver instead of a value receiver?

- a) Always — pointer receivers are faster in every case
- b) When the method needs to modify the struct's fields
- c) Only when the struct has more than 3 fields
- d) When the method returns a value

---

**Question 3:** What does struct embedding provide in Go?

- a) Classical inheritance with an "is-a" relationship
- b) Composition — the embedded struct's fields and methods are promoted
- c) Automatic interface implementation
- d) Runtime polymorphism through virtual method dispatch

---

**Question 4:** What does the struct tag `json:",omitempty"` do?

- a) Makes the field required in JSON input
- b) Skips the field during JSON encoding if it has its zero value
- c) Sets the field to an empty string in JSON output
- d) Causes a runtime error if the field is empty

---

**Question 5:** What is the Go convention for a constructor function that creates a `Server` struct?

- a) `Server.New()`
- b) `CreateServer()`
- c) `NewServer()`
- d) `Server{}.Init()`

---

<details>
<summary>Answers</summary>

1. **b)** — Structs are value types in Go, so assignment copies all fields; the two variables are completely independent after the copy.
2. **b)** — Use a pointer receiver when the method needs to mutate the struct's fields; a value receiver works on a copy and can't change the original.
3. **b)** — Embedding provides composition: the embedded struct's fields and methods are promoted to the outer struct, but there's no inheritance hierarchy.
4. **b)** — The `omitempty` option tells `encoding/json` to skip the field in the JSON output when the field holds its zero value (0, "", false, nil).
5. **c)** — Go convention is `NewXxx()` for constructor functions — a plain function that returns an initialized instance, often as a pointer.

</details>
