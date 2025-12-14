package day5

import "testing"

func TestFreshIngredientsSample1(t *testing.T) {
	t.Log(FreshIngredients(sample1)) // 3, 3
}

func TestFreshIngredientsPuzzle(t *testing.T) {
	t.Log(FreshIngredients(puzzle)) // 563, 437
}
