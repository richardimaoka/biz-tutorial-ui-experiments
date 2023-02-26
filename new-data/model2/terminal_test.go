package model2

import "testing"

func TestTerminal(t *testing.T) {
	terminal := NewTerminal("default")
	compareAfterMarshal(t, "testdata/terminal/new-terminal.json", terminal)
}
