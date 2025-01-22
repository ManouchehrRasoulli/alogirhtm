package kata

import "testing"

func Test_Hammer3_4(t *testing.T) {
	v := Hammer(3)
	if v != 3 {
		t.Logf("Hamming %+v\n", hamming)
		t.Fatal("Expected 3, got ", v)
	}

	v = Hammer(4)
	if v != 4 {
		t.Logf("Hamming %+v\n", hamming)
		t.Fatal("Expected 4, got ", v)
	}

	t.Log(hamming)
}

func Test_Hammer19_11(t *testing.T) {
	v := Hammer(19)
	if v != 32 {
		t.Logf("Hamming %+v\n", hamming)
		t.Fatal("Expected 32, got ", v)
	}

	v = Hammer(11)
	if v != 15 {
		t.Logf("Hamming %+v\n", hamming)
		t.Fatal("Expected 15, got ", v)
	}

	t.Log(hamming)
}
