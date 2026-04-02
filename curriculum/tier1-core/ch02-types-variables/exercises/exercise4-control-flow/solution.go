//go:build ignore

// Exercise 4: Implement Control Flow Patterns — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz (1–20)")
	fmt.Println("===============")

	fizz, buzz, fizzBuzz := 0, 0, 0

	for i := 1; i <= 20; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("  FizzBuzz")
			fizzBuzz++
		case i%3 == 0:
			fmt.Println("  Fizz")
			fizz++
		case i%5 == 0:
			fmt.Println("  Buzz")
			buzz++
		default:
			fmt.Printf("  %d\n", i)
		}
	}

	fmt.Printf("\nSummary: %d Fizz, %d Buzz, %d FizzBuzz\n", fizz, buzz, fizzBuzz)
}
