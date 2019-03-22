package helperFunc

import "testing"

func TestCalculateFibonacciWithoutHash(t *testing.T) {
	actual := CalculateFibonacciWithoutHash(7)
	expected := 8

	if actual != expected {
		t.Error("Failed")
	}
}

func TestCalculateFibonacciHash(t *testing.T) {
	actual := CalculateFibonacciWithoutHash(7)
	expected := 8

	if actual != expected {
		t.Error("Failed")
	}
}

func TestCalculateFibonacciLoop(t *testing.T) {
	actual := CalculateFibonacciWithoutHash(7)
	expected := 8

	if actual != expected {
		t.Error("Failed")
	}
}

func TestCalculateFibonacciHash2(t *testing.T) {
	actual := CalculateFibonacciWithoutHash(7)
	expected := 8

	if actual != expected {
		t.Error("Failed")
	}
}
