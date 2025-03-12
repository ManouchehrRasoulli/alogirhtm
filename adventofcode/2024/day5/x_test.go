package day5

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	rules, updates := Parse(data)

	precedenceMap := CreatePrecedenceMap(rules)
	Interpret(precedenceMap, updates)

	for _, update := range updates {
		log.Println(update)
	}

	t.Log(ValidSum(updates)) // 143
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	rules, updates := Parse(data)

	precedenceMap := CreatePrecedenceMap(rules)
	Interpret(precedenceMap, updates)

	for _, update := range updates {
		log.Println(update)
	}

	t.Log(ValidSum(updates)) // 6505
}
