package kata

import "testing"

func TestDecoder(t *testing.T) {
	t.Log(Decode("XXI"))     // 21
	t.Log(Decode("MMVIII"))  // 2008
	t.Log(Decode("MDCLXVI")) // 1666
	t.Log(Decode("IV"))      // 4
}
