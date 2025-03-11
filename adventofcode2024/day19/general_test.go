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

	tokens := Tokens(rp, dna)
	t.Log(tokens)
}
