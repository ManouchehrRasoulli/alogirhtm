package kata

import "testing"

func TestBigPrimefacDiv(t *testing.T) {
	t.Log(BigPrimefacDiv(1969))      // [179, 179]
	t.Log(BigPrimefacDiv(-1800))     // [5, 900]
	t.Log(BigPrimefacDiv(997))       // [179, 179]
	t.Log(BigPrimefacDiv(734897881)) // [27109 27109]
}
