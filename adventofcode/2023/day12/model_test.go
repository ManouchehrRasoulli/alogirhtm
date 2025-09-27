package day12

import "testing"

func TestModel(t *testing.T) {
	model := dataLines{
		Data:  []value{operational, damaged, damaged, unknown, operational},
		Match: []int{2},
	}

	t.Log(model)
	t.Log(model.DataMatches())
}
