package day5

import "testing"

func TestFreshIngredientRangesSample1(t *testing.T) {
	t.Log(FreshIngredientRanges(sample1)) // 14
}

func TestFreshIngredientRangesPuzzle(t *testing.T) {
	t.Log(FreshIngredientRanges(puzzle)) // 338693411431456
}
