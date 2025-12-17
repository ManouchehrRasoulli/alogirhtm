package day8

import "testing"

func TestConnectedCircuitsAndReturnLastOneSample1(t *testing.T) {
	t.Log(ConnectedCircuitsAndReturnLastOne(sample1)) // 25272
}

func TestConnectedCircuitsAndReturnLastOnePuzzle(t *testing.T) {
	t.Log(ConnectedCircuitsAndReturnLastOne(puzzle)) // 4884971896
}
