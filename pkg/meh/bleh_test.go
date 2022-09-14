package meh

import "testing"

func TestSum(t *testing.T) {
	expected := "hello pants"
	greeting := Hello("pants")
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, expected)
	}
}
