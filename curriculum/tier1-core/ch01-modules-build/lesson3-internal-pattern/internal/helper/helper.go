// Package helper provides internal utility functions.
// This package is inside internal/ so only this module can import it.
package helper

import "fmt"

// Greet returns a formatted greeting message.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! (from internal/helper)", name)
}

// ModuleInfo returns information about the internal package pattern.
func ModuleInfo() string {
	return "This function lives in internal/helper — external modules cannot call it."
}
