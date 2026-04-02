# Chapter 01 Quiz: Modules, Build & Project Structure

Test your understanding of Go modules, building, and project structure.

---

**Question 1:** What file declares your module's unique path and its dependencies?

- a) `go.sum`
- b) `go.mod`
- c) `package.json`
- d) `Makefile`

---

**Question 2:** What does `go.sum` store for each dependency?

- a) The source code of each dependency
- b) The download URL for each dependency
- c) Cryptographic hashes for tamper detection
- d) The license type of each dependency

---

**Question 3:** What happens if an external module tries to import a package from your `internal/` directory?

- a) It works but prints a deprecation warning
- b) It works in test mode only
- c) The Go compiler rejects the import
- d) It works if the package is exported

---

**Question 4:** Which environment variables control cross-compilation in Go?

- a) `OS` and `ARCH`
- b) `GOOS` and `GOARCH`
- c) `TARGET_OS` and `TARGET_ARCH`
- d) `BUILD_OS` and `BUILD_ARCH`

---

**Question 5:** How should a Godoc comment for a function called `LoadConfig` begin?

- a) `// This function loads the configuration`
- b) `// @func LoadConfig`
- c) `// LoadConfig reads...`
- d) `/* LoadConfig */`

---

<details>
<summary>Answers</summary>

1. **b)** — `go.mod` declares the module path and dependencies; `go.sum` stores verification hashes.
2. **c)** — `go.sum` stores cryptographic hashes so Go can verify dependencies haven't been tampered with.
3. **c)** — The Go compiler enforces `internal/` privacy at build time, rejecting imports from external modules.
4. **b)** — `GOOS` sets the target operating system and `GOARCH` sets the target CPU architecture.
5. **c)** — Godoc convention requires comments to start with the symbol name followed by a description.

</details>
