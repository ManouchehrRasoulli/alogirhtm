package day21

import (
	"log"
	"testing"
)

func TestKeypad(t *testing.T) {
	keypad := NewKeypad("029A")
	path := keypad.CharToDirections()
	log.Println(string(path)) // <A^A>^^AvvvA
	if string(path) != "<A^A>^^AvvvA" {
		log.Fatal("invalid output !!")
	}
}
