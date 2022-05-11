package isEven

import "testing"

func TestEven(t *testing.T) {
	if !IsEven(4) {
		t.Fatal("Test Failed")
	}
}

func TestNotEven(t *testing.T) {
	if IsEven(3) {
		t.Fatal("Test Failed")
	}
}
