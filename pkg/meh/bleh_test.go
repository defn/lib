package meh

import "testing"

func TestHello(t *testing.T) {
	expected := "hello pants"
	greeting := Hello("pants")
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, expected)
	}
}
