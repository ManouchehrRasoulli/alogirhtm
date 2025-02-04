package kata

import "testing"

func TestInit(t *testing.T) {
	t.Log(encodeStr)
	t.Logf("%+v\n", encode)
	t.Logf("%+v\n", decodeMap)
}

func TestEncode(t *testing.T) {
	v := Encode([]byte("1710319735193131327417641117654495650015916157913185"))
	t.Log(string(v))
}

func TestDecode(t *testing.T) {
	v := Decode([]byte("dr.J"))
	t.Log(string(v))
}
