// Exercise 2: Add a Dependency
// Task: Track dependencies with no duplicates
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

// Dependency represents a Go module dependency.
type Dependency struct {
	Name    string
	Version string
}

// TODO: Implement addDependency that adds a dep to the slice if not already present.
// It should print "Adding: name version" or "Adding: name version (already exists, skipping)"
// Hint: Return the (possibly updated) slice.
func addDependency(deps []Dependency, name, version string) []Dependency {
	// TODO: Check if dependency already exists by name
	// TODO: If not, append and print "Adding: ..."
	// TODO: If yes, print "Adding: ... (already exists, skipping)"
	return deps
}

// TODO: Implement listDependencies that prints all deps in a formatted list.
func listDependencies(deps []Dependency) {
	// TODO: Print each dependency name and version
	// TODO: Print the total count
}

func main() {
	var deps []Dependency

	deps = addDependency(deps, "github.com/spf13/cobra", "v1.10.2")
	deps = addDependency(deps, "pgregory.net/rapid", "v1.2.0")
	deps = addDependency(deps, "github.com/spf13/cobra", "v1.10.2")

	fmt.Println()
	listDependencies(deps)

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
