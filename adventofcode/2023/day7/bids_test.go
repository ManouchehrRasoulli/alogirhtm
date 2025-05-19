package day7

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestResolveType(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected CardType
	}{
		{
			name:     "1",
			value:    "32T3K",
			expected: OnePair,
		},
		{
			name:     "2",
			value:    "KK677",
			expected: TwoPair,
		},
		{
			name:     "3",
			value:    "KTJJT",
			expected: TwoPair,
		},
		{
			name:     "4",
			value:    "T55J5",
			expected: ThreeOfAKind,
		},
		{
			name:     "5",
			value:    "QQQJA",
			expected: ThreeOfAKind,
		},
		{
			name:     "6",
			value:    "AAAAA",
			expected: FiveOfAKind,
		},
		{
			name:     "7",
			value:    "AA8AA",
			expected: FourOfAKind,
		},
		{
			name:     "8",
			value:    "TTT98",
			expected: ThreeOfAKind,
		},
		{
			name:     "10",
			value:    "23332",
			expected: FullHouse,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tp := ResolveType(test.value)
			if test.expected != tp {
				t.Fatal("Invalid", test.value, " Expected", test.expected, " Got", tp)
			}
		})
	}
}

func TestInputTestPart1(t *testing.T) {
	input, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	bids := ReadBids(input)
	t.Log(Part1(bids)) // 6440
}

func TestInputPuzzlePart1(t *testing.T) {
	input, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	bids := ReadBids(input)
	t.Log(Part1(bids)) // 248396258
}

func TestInputTestPart2(t *testing.T) {
	input, err := helper.ReadAll("test1.txt")
	if err != nil {
		t.Fatal(err)
	}

	bids := ReadBids(input)
	t.Log(Part2(bids)) // 5905
}

func TestInputPuzzlePart2(t *testing.T) {
	input, err := helper.ReadAll("puzzle.txt")
	if err != nil {
		t.Fatal(err)
	}

	bids := ReadBids(input)
	t.Log(Part2(bids)) // 246436046
}
