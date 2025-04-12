package day2

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"strings"
	"testing"
)

func TestInputTest1(t *testing.T) {
	data, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(data, "\n")
	t.Log(ReadGame(lines[0]))
}
