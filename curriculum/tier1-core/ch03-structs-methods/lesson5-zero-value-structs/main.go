// Lesson 5: Zero Values for Structs
// Demonstrates: zero-value struct behavior and designing useful zero values
//
// Run: go run .
package main

import "fmt"

// Config represents application settings.
// Designed so the zero value is a useful default.
type Config struct {
	Host    string
	Port    int
	Verbose bool
}

// HostOrDefault returns the host, falling back to "localhost" if empty.
func (c Config) HostOrDefault() string {
	if c.Host == "" {
		return "localhost"
	}
	return c.Host
}

// PortOrDefault returns the port, falling back to 8080 if zero.
func (c Config) PortOrDefault() int {
	if c.Port == 0 {
		return 8080
	}
	return c.Port
}

// Counter tracks a running count.
// Its zero value (count=0) is immediately usable.
type Counter struct {
	count int
}

// Increment adds one to the counter.
func (c *Counter) Increment() {
	c.count++
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.count
}

func main() {
	fmt.Println("=== Zero Values for Structs ===")
	fmt.Println()

	// Declare a Config without initialization — every field is its zero value
	var cfg Config
	fmt.Println("Zero-value Config:")
	fmt.Printf("  Host:    %q (zero value for string)\n", cfg.Host)
	fmt.Printf("  Port:    %d  (zero value for int)\n", cfg.Port)
	fmt.Printf("  Verbose: %t (zero value for bool)\n", cfg.Verbose)

	fmt.Println()

	// Methods can provide sensible defaults for zero values
	fmt.Println("With default fallbacks:")
	fmt.Printf("  Host: %s\n", cfg.HostOrDefault())
	fmt.Printf("  Port: %d\n", cfg.PortOrDefault())

	fmt.Println()

	// Counter works immediately at its zero value
	var c Counter
	fmt.Printf("Counter starts at: %d\n", c.Value())
	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Printf("After 3 increments: %d\n", c.Value())

	fmt.Println()
	fmt.Println("Key insight: design structs so zero values are useful defaults.")
}
