package day19

import "testing"

func TestReadInputTest(t *testing.T) {
	data, err := ReadInput("input_test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(data)

	rp, dna := Parse(data)
	t.Log(rp, dna)

	t.Log(Diffs(rp, dna))
}

func TestReadInputPuzzle(t *testing.T) {
	data, err := ReadInput("input_puzzle.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(data)

	rp, dna := Parse(data)
	t.Log(rp, dna)

	t.Log(Diffs(rp, dna))
}

func TestFindSubstringFromIndex(t *testing.T) {
	index := FindSubstringFromIndex("HOH", "H", 0)
	t.Log(index)
	index = FindSubstringFromIndex("HOH", "H", index+1)
	t.Log(index)

	newStr := ReplaceSubstringAt("HOH", "H", "HO", 0)
	t.Log(newStr)
	newStr = ReplaceSubstringAt("HOH", "H", "HO", 2)
	t.Log(newStr)
}

func TestReduceTest(t *testing.T) {
	data, err := ReadInput("input_test_part2.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	rp, dna := Parse(data)
	t.Log(Reduce(rp, dna, "e"))
}

func TestReducePuzzle(t *testing.T) {
	data, err := ReadInput("input_puzzle_part2.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	rp, dna := Parse(data)
	t.Log(Reduce(rp, dna, "e"))
}
