package _kyu

import (
	"log"
	"testing"
)

func TestDirReduc(t *testing.T) {
	a := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	log.Println(DirReduc(a))

	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	log.Println(DirReduc(a))

	a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
	log.Println(DirReduc(a))

	a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
	log.Println(DirReduc(a))
}
