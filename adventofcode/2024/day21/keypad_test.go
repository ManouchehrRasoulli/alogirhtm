package day21

import "testing"

func TestKeyPad(t *testing.T) {
	keypad := NewKeypad(WithNumPad("029A"))
	seq, paths := keypad.KeyToDirections()
	for _, path := range paths {
		t.Log(seq, string(path), len(path))
	}
}

func TestKeyPadWithOtherKeypad(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("029A"))
	seq, paths := keypad.KeyToDirections()
	for _, path := range paths {
		t.Log(seq, string(path), len(path))
	}
}

func TestKeyPadWithOtherKeypad2(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("179A"))
	seq, paths := keypad.KeyToDirections()
	for _, path := range paths {
		t.Log(seq, string(path), len(path))
	}
}

func TestKeyPadWithOtherKeypad3(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("980A"))
	seq, paths := keypad.KeyToDirections()
	for _, path := range paths {
		t.Log(seq, string(path), len(path))
	}
}
