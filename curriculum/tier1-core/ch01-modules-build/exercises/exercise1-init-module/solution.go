//go:build ignore

// Exercise 1: Initialize a Go Module — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println("Error reading go.mod:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			name := strings.TrimPrefix(line, "module ")
			fmt.Println("Module name:", name)
		}
		if strings.HasPrefix(line, "go ") {
			version := strings.TrimPrefix(line, "go ")
			fmt.Println("Go version: ", version)
		}
	}
}
