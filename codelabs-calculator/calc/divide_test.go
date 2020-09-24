package calc

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(10, 5)
	if result != 2 || err != nil {
		t.Errorf("An error has been returned or the division result was incorrect, got: %d, want: %d.", result, 2)
	}
}

func TestDivideByZero(t *testing.T) {
	result, err := Divide(10, 0)
	if result != 0 || err == nil {
		t.Errorf("Division didn't fail. An error was expected when dividin by zero.")
	}
}
