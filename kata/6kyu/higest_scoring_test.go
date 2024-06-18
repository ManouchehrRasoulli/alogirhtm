package _kyu

import (
	"testing"
)

/*
Given a string of words, you need to find the highest scoring word.

Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

For example, the score of abad is 8 (1 + 2 + 1 + 4).

You need to return the highest scoring word as a string.

If two words score the same, return the word that appears earliest in the original string.

All letters will be lowercase and all inputs will be valid.
*/

func TestSum(t *testing.T) {
	t.Log(score("a"))
	t.Log(score("b"))
	t.Log(score("abb"))
	t.Log(score("abc"))
	t.Log(score("taxi"))
	t.Log(score("ubud"))
}

func TestHighWithSort(t *testing.T) {
	arr := [...][2]string{
		{"man i need a taxi up to ubud", "taxi"},
		{"what time are we climbing up the volcano", "volcano"},
		{"take me to semynak", "semynak"},
		{"aa b", "aa"},
		{"b aa", "b"},
		{"bb d", "bb"},
		{"d bb", "d"},
		{"aaa b", "aaa"},
	}
	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		funcOut := HighWithSort(input)
		if funcOut != output {
			t.Log("invalid output", "   ---   ", input, "   ***   ", output, "   &&&   ", funcOut)
		}
	}
}

func TestHigh(t *testing.T) {
	arr := [...][2]string{
		{"man i need a taxi up to ubud", "taxi"},
		{"what time are we climbing up the volcano", "volcano"},
		{"take me to semynak", "semynak"},
		{"aa b", "aa"},
		{"b aa", "b"},
		{"bb d", "bb"},
		{"d bb", "d"},
		{"aaa b", "aaa"},
	}
	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		funcOut := High(input)
		if funcOut != output {
			t.Log("invalid output", "   ---   ", input, "   ***   ", output, "   &&&   ", funcOut)
		}
	}
}
