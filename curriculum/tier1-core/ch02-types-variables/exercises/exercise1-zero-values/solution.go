//go:build ignore

// Exercise 1: Explore Zero Values — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

func main() {
	fmt.Println("Zero Value Explorer")
	fmt.Println("===================")
	fmt.Println("Type       | Zero Value")
	fmt.Println("-----------|----------")

	var i int
	var f float64
	var b bool
	var s string
	var by byte
	var r rune

	fmt.Printf("int        | %d\n", i)
	fmt.Printf("float64    | %g\n", f)
	fmt.Printf("bool       | %t\n", b)
	fmt.Printf("string     | %q\n", s)
	fmt.Printf("byte       | %d\n", by)
	fmt.Printf("rune       | %d\n", r)

	fmt.Println("-----------|----------")
	fmt.Printf("Checked %d types — all have defined zero values!\n", 6)
}
