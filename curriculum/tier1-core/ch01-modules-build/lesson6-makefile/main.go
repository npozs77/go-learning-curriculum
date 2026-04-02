// Lesson 6: Makefile Usage for Common Project Tasks
// Demonstrates: typical Makefile targets for a Go project
//
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("=== Makefile for Go Projects ===")
	fmt.Println()

	// Print a typical Go Makefile
	fmt.Println("A typical Go project Makefile:")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println(`APP_NAME := myapp`)
	fmt.Println(`BUILD_DIR := ./build`)
	fmt.Println()
	fmt.Println(`.PHONY: build test vet fmt lint clean ci`)
	fmt.Println()
	fmt.Println(`build:`)
	fmt.Println(`	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/myapp`)
	fmt.Println()
	fmt.Println(`test:`)
	fmt.Println(`	go test -race ./...`)
	fmt.Println()
	fmt.Println(`vet:`)
	fmt.Println(`	go vet ./...`)
	fmt.Println()
	fmt.Println(`fmt:`)
	fmt.Println(`	gofmt -w .`)
	fmt.Println()
	fmt.Println(`lint: vet`)
	fmt.Println(`	staticcheck ./...`)
	fmt.Println()
	fmt.Println(`clean:`)
	fmt.Println(`	rm -rf $(BUILD_DIR)`)
	fmt.Println()
	fmt.Println(`ci: fmt vet test build`)
	fmt.Println(`	@echo "All checks passed!"`)
	fmt.Println("─────────────────────────────────────────")
	fmt.Println()

	// Show what each target does
	targets := []struct {
		name, command, desc string
	}{
		{"make build", "go build -o ./build/myapp ./cmd/myapp", "Compile the binary"},
		{"make test", "go test -race ./...", "Run all tests with race detector"},
		{"make vet", "go vet ./...", "Catch common mistakes"},
		{"make fmt", "gofmt -w .", "Format all Go files"},
		{"make clean", "rm -rf ./build", "Delete build artifacts"},
		{"make ci", "fmt + vet + test + build", "Full CI pipeline"},
	}

	fmt.Println("Target summary:")
	for _, t := range targets {
		fmt.Printf("  %-12s → %s\n", t.name, t.desc)
	}
}
