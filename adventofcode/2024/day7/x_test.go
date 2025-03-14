package day7

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Log(GenerateCombinationOfOperationInLen([]string{"+", "*"}, 3, ""))      // [[+++] [++*] [+*+] [+**] [*++] [*+*] [**+] [***]]
	t.Log(GenerateCombinationOfOperationInLen([]string{"+", "*"}, 4, ""))      // [[++++] [+++*] [++*+] [++**] [+*++] [+*+*] [+**+] [+***] [*+++] [*++*] [*+*+] [*+**] [**++] [**+*] [***+] [****]]
	t.Log(GenerateCombinationOfOperationInLen([]string{"+", "*", "|"}, 2, "")) // [++ +* +| *+ ** *| |+ |* ||]
}

func TestDoOpsWithNumbers(t *testing.T) {
	// t.Log(DoOpsWithNumbers("+*", []int{81, 40, 27})) // 3267
	t.Log(DoOpsWithNumbers("|", []int{15, 6})) // 156
}

func TestOperatorsTestPart1(t *testing.T) {
	input, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	records := ReadInRecordPart1(input)
	for _, op := range records {
		t.Log(op)
	}

	t.Log("Computing ---")
	MatchOperatorsForRecordPart1(records)

	sum := 0
	for _, op := range records {
		t.Log(op)
		if op.IsReachable {
			sum += op.FinalValue
		}
	}

	t.Log(sum) // 3749
}

func TestOperatorsPuzzlePart1(t *testing.T) {
	input, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	records := ReadInRecordPart1(input)
	for _, op := range records {
		t.Log(op)
	}

	t.Log("Computing ---")
	MatchOperatorsForRecordPart1(records)

	sum := 0
	for _, op := range records {
		t.Log(op)
		if op.IsReachable {
			sum += op.FinalValue
		}
	}

	t.Log(sum) // 1985268524462
}

func TestOperatorsTestPart2(t *testing.T) {
	input, err := helper.ReadAll("input_test_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	records := ReadInRecordPart1(input)
	for _, op := range records {
		t.Log(op)
	}

	t.Log("Computing ---")
	MatchOperatorsForRecordPart2(records)

	sum := 0
	for _, op := range records {
		t.Log(op)
		if op.IsReachable {
			sum += op.FinalValue
		}
	}

	t.Log(sum) // 11387
}

func TestOperatorsPuzzlePart2(t *testing.T) {
	input, err := helper.ReadAll("input_puzzle_part1.txt")
	if err != nil {
		t.Fatal(err)
	}

	records := ReadInRecordPart1(input)
	for _, op := range records {
		t.Log(op)
	}

	t.Log("Computing ---")
	MatchOperatorsForRecordPart2(records)

	sum := 0
	for _, op := range records {
		t.Log(op)
		if op.IsReachable {
			sum += op.FinalValue
		}
	}

	t.Log(sum) // 150077710195188
}
