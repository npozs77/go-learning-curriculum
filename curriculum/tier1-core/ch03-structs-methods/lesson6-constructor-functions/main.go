// Lesson 6: Constructor Functions (NewXxx Pattern)
// Demonstrates: the NewXxx convention for creating validated struct instances
//
// Run: go run .
package main

import (
	"errors"
	"fmt"
)

// Server represents a network server configuration.
type Server struct {
	Host string
	Port int
}

// NewServer creates a Server with validation and defaults.
// Returns an error if the port is out of range.
func NewServer(host string, port int) (*Server, error) {
	if port < 1 || port > 65535 {
		return nil, errors.New("port must be between 1 and 65535")
	}
	if host == "" {
		host = "localhost"
	}
	return &Server{Host: host, Port: port}, nil
}

// Address returns the server's address as host:port.
func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Temperature represents a temperature in Celsius.
type Temperature struct {
	celsius float64
}

// NewTemperature creates a Temperature, rejecting values below absolute zero.
func NewTemperature(celsius float64) (*Temperature, error) {
	if celsius < -273.15 {
		return nil, fmt.Errorf("%.2f°C is below absolute zero (-273.15°C)", celsius)
	}
	return &Temperature{celsius: celsius}, nil
}

// Celsius returns the temperature in Celsius.
func (t *Temperature) Celsius() float64 {
	return t.celsius
}

// Fahrenheit returns the temperature converted to Fahrenheit.
func (t *Temperature) Fahrenheit() float64 {
	return t.celsius*9/5 + 32
}

func main() {
	fmt.Println("=== Constructor Functions (NewXxx Pattern) ===")
	fmt.Println()

	// Successful construction with default host
	srv, err := NewServer("", 8080)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Server: %s (default host applied)\n", srv.Address())
	}

	// Successful construction with explicit host
	srv2, err := NewServer("api.example.com", 443)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Server: %s\n", srv2.Address())
	}

	// Failed construction — invalid port
	_, err = NewServer("localhost", 99999)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()

	// Temperature with validation
	temp, err := NewTemperature(100)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Water boils at %.1f°C (%.1f°F)\n", temp.Celsius(), temp.Fahrenheit())
	}

	// Temperature below absolute zero — rejected
	_, err = NewTemperature(-300)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("Key insight: use NewXxx when zero values aren't enough or inputs need validation.")
}
