package day24

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestInputSimplePart1(t *testing.T) {
	data, err := helper.ReadAll("simple.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	t.Log(DoOperations(lines)) // 4
}

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	t.Log(DoOperations(lines)) // 2024
}

func TestInputInputPart1(t *testing.T) {
	data, err := helper.ReadAll("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	t.Log(DoOperations(lines)) // 59619940979346
}

func TestInputSimplePart2(t *testing.T) {
	data, err := helper.ReadAll("simple.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	FixCircuit(lines)
}

func TestInputTestPart2(t *testing.T) {
	data, err := helper.ReadAll("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	FixCircuit(lines)
}
