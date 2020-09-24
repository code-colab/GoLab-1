package calc

import "testing"

func TestSum(t *testing.T) {
	total, err := Sum(5, 5)
	if total != 10 || err != nil {
		t.Errorf("An error has been returned or the sum result was incorrect, got: %d, want: %d.", total, 10)
	}
}