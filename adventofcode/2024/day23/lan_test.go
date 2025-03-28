package day23

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")

	lan := NewLan()
	lan.MakeConnections(lines)
	t.Log(lan.ChiefComputerTriples()) // 7
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")

	lan := NewLan()
	lan.MakeConnections(lines)
	t.Log(lan.ChiefComputerTriples()) // 1119
}
