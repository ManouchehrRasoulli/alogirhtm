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
	log.Println(precedenceMap)

	Interpret(precedenceMap, updates)

	for _, update := range updates {
		log.Println(update)
	}

	first := ValidSum(updates)
	t.Log(first) // 143

	log.Println("fixing...")

	Fix(updates, precedenceMap)
	Interpret(precedenceMap, updates)
	for _, update := range updates {
		log.Println(update)
	}

	second := ValidSum(updates)
	t.Log(second - first) // 123
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

	first := ValidSum(updates)
	t.Log(first) // 6505

	log.Println("fixing...")

	Fix(updates, precedenceMap)
	Interpret(precedenceMap, updates)
	for _, update := range updates {
		log.Println(update)
	}

	second := ValidSum(updates)
	t.Log(second - first) // 6897
}
