package model2

import "testing"

func TestTerminal(t *testing.T) {
	terminal := NewTerminal("default")
	terminal.TypeInCommand("echo abc")
}
