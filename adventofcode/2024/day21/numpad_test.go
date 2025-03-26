package day21

import (
	"log"
	"testing"
)

func TestNumpad(t *testing.T) {
	numpad := NewNumpad("029A")
	path := numpad.NumsToDirections()
	log.Println(string(path)) // <A^A>^^AvvvA
	if string(path) != "<A^A>^^AvvvA" {
		log.Fatal("invalid output !!")
	}
}
