package day25

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
	keys, locks := ReadCombinations(lines)

	t.Log("keys")

	for _, k := range keys {
		t.Log(k)
	}

	t.Log("locks")

	for _, l := range locks {
		t.Log(l)
	}

	t.Log(" -- testing -- ")
	total := 0
	for _, l := range locks {
		total += l.TryFits(keys)
	}

	t.Log("total", total) // 3
}

func TestInputPart1(t *testing.T) {
	data, err := helper.ReadAll("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")
	keys, locks := ReadCombinations(lines)

	t.Log("keys")

	for _, k := range keys {
		t.Log(k)
	}

	t.Log("locks")

	for _, l := range locks {
		t.Log(l)
	}

	t.Log(" -- testing -- ")
	total := 0
	for _, l := range locks {
		total += l.TryFits(keys)
	}

	t.Log("total", total) // 3021
}
