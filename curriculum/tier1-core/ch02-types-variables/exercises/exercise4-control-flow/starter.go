// Exercise 4: Implement Control Flow Patterns
// Task: Write FizzBuzz using a conditionless switch
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz (1–20)")
	fmt.Println("===============")

	// TODO: Declare counters for fizz, buzz, and fizzBuzz
	// Hint: fizz, buzz, fizzBuzz := 0, 0, 0

	// TODO: Loop from 1 to 20
	// Hint: for i := 1; i <= 20; i++ { ... }

	// TODO: Inside the loop, use a conditionless switch to check:
	//   - i%15 == 0 → print "  FizzBuzz" and increment fizzBuzz
	//   - i%3 == 0  → print "  Fizz" and increment fizz
	//   - i%5 == 0  → print "  Buzz" and increment buzz
	//   - default    → print the number with fmt.Printf("  %d\n", i)

	// TODO: Print the summary line
	// Hint: fmt.Printf("\nSummary: %d Fizz, %d Buzz, %d FizzBuzz\n", fizz, buzz, fizzBuzz)

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
