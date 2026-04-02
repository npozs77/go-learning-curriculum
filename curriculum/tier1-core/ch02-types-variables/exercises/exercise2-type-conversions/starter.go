// Exercise 2: Practice Type Conversions
// Task: Convert temperatures between Celsius and Fahrenheit using different types
//
// Instructions: Fill in the TODO sections below.
// Run: go run .
package main

import "fmt"

func main() {
	fmt.Println("Temperature Converter")
	fmt.Println("=====================")

	// TODO: Define a float64 temperature in Celsius
	// Hint: celsius := 36.6

	// TODO: Convert Celsius to Fahrenheit as float64
	// Formula: F = C × 9/5 + 32
	// Hint: fahrenheitF := celsius*9.0/5.0 + 32.0

	// TODO: Convert the float64 Fahrenheit to int (truncated)
	// Hint: fahrenheitI := int(fahrenheitF)

	// TODO: Print the Celsius, float64 Fahrenheit, and int Fahrenheit values
	// Hint: fmt.Printf("Celsius (float64):    %.1f\n", celsius)

	// TODO: Convert the int Fahrenheit back to Celsius as float64
	// Hint: backToC := (float64(fahrenheitI) - 32.0) * 5.0 / 9.0

	// TODO: Print the back-conversion and show precision loss
	// Hint: fmt.Printf("Converting back: %d°F → %.2f°C (float64)\n", fahrenheitI, backToC)

	fmt.Println("Exercise not yet complete — fill in the TODOs above")
}
