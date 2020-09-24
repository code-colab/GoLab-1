package calc

import "testing"

func TestMultiply(t *testing.T) {
	result, err := Multiply(5, 2)
	if result != 10 || err != nil {
		t.Errorf("An error has been returned or the multiply result was incorrect, got: %d, want: %d.", result, 10)
	}
}
