package distinct

import (
	"testing"
)

func Test_Process(t *testing.T) {
	m := process("rabbbit")
	t.Log(m)
	for _, v := range m['r'] {
		t.Log(v)
	}
}

func Test_NumDistinct_(t *testing.T) {
	t.Log(numDistinct_("leet", "e"))         // 2
	t.Log(numDistinct_("leet", "let"))       // 2
	t.Log(numDistinct_("leet", "leet"))      // 1
	t.Log(numDistinct_("leet", "leeet"))     // 0
	t.Log(numDistinct_("rabbbit", "rabbit")) // 3
	t.Log(numDistinct_("babgbag", "bag"))    // 5
	t.Log(numDistinct_("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaae"+
		"deeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc")) // 700531452
}

func Test_NumDistinct(t *testing.T) {
	t.Log(numDistinct("leet", "e"))         // 2
	t.Log(numDistinct("leet", "let"))       // 2
	t.Log(numDistinct("leet", "leet"))      // 1
	t.Log(numDistinct("leet", "leeet"))     // 0
	t.Log(numDistinct("rabbbit", "rabbit")) // 3
	t.Log(numDistinct("babgbag", "bag"))    // 5
	t.Log(numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaae"+
		"deeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc")) // 700531452
}
