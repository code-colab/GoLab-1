package main
import "testing"

func TestSum(t *testing.T) {
    total, err := sum(5, 5)
    if total != 10 || err != nil {
       t.Errorf("An error has been returned or the sum result was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestDivide(t *testing.T) {
    result, err := divide(10, 5)
    if result != 2 || err != nil {
       t.Errorf("An error has been returned or the division result was incorrect, got: %d, want: %d.", result, 2)
    }
}

func TestDivideByZero(t *testing.T) {
    result, err := divide(10, 0)
    if result != 0 || err == nil {
       t.Errorf("Division didn't fail. An error was expected when dividin by zero.")
    }
}

func TestSubtract(t *testing.T) {
	result, err := subtract(5, 2)
	if result != 3 || err != nil {
		t.Errorf("An error has been returned or the subtraction result was incorrect, got: %d, want: %d.", result, 3)
	}
}

func TestMultiply(t *testing.T) {
	result, err := multiply(5, 2)
	if result != 10 || err != nil {
		t.Errorf("An error has been returned or the multiply result was incorrect, got: %d, want: %d.", result, 10)
	}
}