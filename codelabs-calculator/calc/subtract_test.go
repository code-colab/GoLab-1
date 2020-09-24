package calc

import "testing"

func TestSubtract(t *testing.T) {
	result, err := Subtract(5, 2)
	if result != 3 || err != nil {
		t.Errorf("An error has been returned or the subtraction result was incorrect, got: %d, want: %d.", result, 3)
	}
}
