package day21

import "testing"

func TestKeyPad(t *testing.T) {
	keypad := NewKeypad(WithNumPad("029A"))
	path := keypad.KeyToDirections()
	t.Log(string(path)) // v<<A>>^A<A>A<AAv>A^A<vAAA^>A
	if string(path) != "v<<A>>^A<A>A<AAv>A^A<vAAA^>A" {
		t.Fatal("invalid keypad output")
	}
}

func TestKeyPadWithOtherKeypad(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("029A"))
	path := keypad.KeyToDirections()
	t.Log(string(path))
	// <vA<AA>>^AvAA<^A>Av<<A>>^AvA^Av<<A>>^AA<vA>A^A<A>Av<<A>A^>AAA<Av>A^A
	if string(path) != "<vA<AA>>^AvAA<^A>Av<<A>>^AvA^Av<<A>>^AA<vA>A^A<A>Av<<A>A^>AAA<Av>A^A" {
		t.Fatal("invalid keypad output")
	}
}

func TestKeyPadWithOtherKeypad2(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("179A"))
	path := keypad.KeyToDirections()
	t.Log(string(path))
	// v<<A>>^A<vA<A>>^AAvAA<^A>Av<<A>>^AAvA^A<vA^>AA<A>Av<<A>A^>AAA<Av>A^A
	// <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
	if string(path) != "v<<A>>^A<vA<A>>^AAvAA<^A>Av<<A>>^AAvA^A<vA^>AA<A>Av<<A>A^>AAA<Av>A^A" {
		t.Fatal("invalid keypad output")
	}
}

func TestKeyPadWithOtherKeypad3(t *testing.T) {
	keypad := NewKeypad(WithKeyPad("980A"))
	path := keypad.KeyToDirections()
	t.Log(string(path))
	// v<<A>>^AAAvA^A<vA<AA>>^AvAA<^A>Av<<A>A^>AAA<Av>A^A<vA^>A<A>A
	// <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
	if string(path) != "v<<A>>^AAAvA^A<vA<AA>>^AvAA<^A>Av<<A>A^>AAA<Av>A^A<vA^>A<A>A" {
		t.Fatal("invalid keypad output")
	}
}
