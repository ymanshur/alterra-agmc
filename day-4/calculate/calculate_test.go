package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Addition(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}
