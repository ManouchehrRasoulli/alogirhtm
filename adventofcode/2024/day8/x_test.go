package day8

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"log"
	"testing"
)

func TestAntiNodeLocation(t *testing.T) {
	testTable := []struct {
		Title  string
		A      Location
		B      Location
		Expect Location
	}{
		{
			Title: "a->b",
			A: Location{
				X: 3,
				Y: 4,
			},
			B: Location{
				X: 5,
				Y: 5,
			},
			Expect: Location{
				X: 1,
				Y: 3,
			},
		},
		{
			Title: "b->a",
			A: Location{
				X: 5,
				Y: 5,
			},
			B: Location{
				X: 3,
				Y: 4,
			},
			Expect: Location{
				X: 7,
				Y: 6,
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.Title, func(t *testing.T) {
			an := AntiNodeLocation(tt.A, tt.B)
			if an != tt.Expect {
				t.Fatalf("NOK -- (%d, %d) want %d, got %d", tt.A, tt.B, tt.Expect, an)
			} else {
				t.Logf("OK -- (%d, %d) want %d, got %d", tt.A, tt.B, tt.Expect, an)
			}
		})
	}
}

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	city := ParseCityPart1(data)
	log.Println(city)

	city.CalculateAntiNodes()

	log.Println(city)

	log.Println(city.AntiNodeCount()) // 14

	city.PlaceAntiNotesOnMap()

	log.Println(city)
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	city := ParseCityPart1(data)
	log.Println(city)

	city.CalculateAntiNodes()

	log.Println(city)

	log.Println(city.AntiNodeCount()) // 359

	city.PlaceAntiNotesOnMap()

	log.Println(city)
}
