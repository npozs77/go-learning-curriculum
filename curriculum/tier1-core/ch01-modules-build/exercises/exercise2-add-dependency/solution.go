//go:build ignore

// Exercise 2: Add a Dependency — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

// Dependency represents a Go module dependency.
type Dependency struct {
	Name    string
	Version string
}

func addDependency(deps []Dependency, name, version string) []Dependency {
	for _, d := range deps {
		if d.Name == name {
			fmt.Printf("Adding: %s %s (already exists, skipping)\n", name, version)
			return deps
		}
	}
	fmt.Printf("Adding: %s %s\n", name, version)
	return append(deps, Dependency{Name: name, Version: version})
}

func listDependencies(deps []Dependency) {
	fmt.Println("Dependencies:")
	for _, d := range deps {
		fmt.Printf("  %-26s %s\n", d.Name, d.Version)
	}
	fmt.Printf("\nTotal: %d dependencies\n", len(deps))
}

func main() {
	var deps []Dependency

	deps = addDependency(deps, "github.com/spf13/cobra", "v1.10.2")
	deps = addDependency(deps, "pgregory.net/rapid", "v1.2.0")
	deps = addDependency(deps, "github.com/spf13/cobra", "v1.10.2")

	fmt.Println()
	listDependencies(deps)
}
