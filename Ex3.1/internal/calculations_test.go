package calculations

import "testing"

func TestFactorial(t *testing.T) {
	result := Calculate(5, false)
	if result != 120 {
		t.Error("Expected 120, got ", result)
	}
}
