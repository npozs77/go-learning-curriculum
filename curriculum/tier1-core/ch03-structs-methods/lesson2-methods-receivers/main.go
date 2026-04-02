// Lesson 2: Methods and Value Receivers vs Pointer Receivers
// Demonstrates: defining methods, value vs pointer receivers, and mutation behavior
//
// Run: go run .
package main

import "fmt"

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
// Value receiver — works on a copy, does not modify the original.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the rectangle.
// Value receiver — only reads data.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale multiplies both dimensions by the given factor.
// Pointer receiver — modifies the original struct.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	fmt.Println("=== Methods and Receivers ===")
	fmt.Println()

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Original:  %+v\n", rect)
	fmt.Printf("Area:      %.1f\n", rect.Area())
	fmt.Printf("Perimeter: %.1f\n", rect.Perimeter())

	fmt.Println()

	// Pointer receiver modifies the original
	rect.Scale(2)
	fmt.Printf("After Scale(2): %+v\n", rect)
	fmt.Printf("New Area:       %.1f\n", rect.Area())

	fmt.Println()

	// Value receiver does NOT modify the original
	// Even if we tried to assign inside a value receiver method,
	// it would only change the copy.
	fmt.Println("Value receiver = copy (read-only effect)")
	fmt.Println("Pointer receiver = original (can mutate)")

	fmt.Println()
	fmt.Println("Key insight: use pointer receivers when you need to modify the struct.")
}
