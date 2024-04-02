package regular_expression

import "testing"

func TestNaiveMatch(t *testing.T) {
	t.Log(NaiveMatch("aa", "a"), " should be false")
	t.Log(NaiveMatch("aa", "a*"), " should be true")
	t.Log(NaiveMatch("ab", ".*"), " should be true")
	t.Log(NaiveMatch("aab", "a*b"), " should be true")
	t.Log(NaiveMatch("aab", "*b"), " should be true")
	t.Log(NaiveMatch("aaa", ".b."), " should be false")
}
