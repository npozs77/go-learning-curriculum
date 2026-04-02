package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d, want %d", got, want)
	}
}

func TestAddNegative(t *testing.T) {
	got := Add(-1, 1)
	want := 0
	if got != want {
		t.Errorf("Add(-1, 1) = %d, want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(4, 5)
	want := 20
	if got != want {
		t.Errorf("Multiply(4, 5) = %d, want %d", got, want)
	}
}
