//go:build ignore

// Exercise 2: Practice Type Conversions — SOLUTION
// This is the reference solution. Try to complete starter.go before looking!
//
// Run: go run solution.go
package main

import "fmt"

func main() {
	fmt.Println("Temperature Converter")
	fmt.Println("=====================")

	celsius := 36.6
	fahrenheitF := celsius*9.0/5.0 + 32.0
	fahrenheitI := int(fahrenheitF)

	fmt.Printf("Celsius (float64):    %.1f\n", celsius)
	fmt.Printf("Fahrenheit (float64): %.2f\n", fahrenheitF)
	fmt.Printf("Fahrenheit (int):     %d (truncated)\n", fahrenheitI)

	fmt.Println()
	backToC := (float64(fahrenheitI) - 32.0) * 5.0 / 9.0
	fmt.Printf("Converting back: %d°F → %.2f°C (float64)\n", fahrenheitI, backToC)
	fmt.Printf("Precision lost:  %.1f°C → %d°F → %.2f°C\n", celsius, fahrenheitI, backToC)
}
