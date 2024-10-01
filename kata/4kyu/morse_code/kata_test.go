package kata

import (
	"testing"
)

func Test_Kata(t *testing.T) {
	code := morseCodes("-.--.-")
	t.Log("-- output --", code)
}

func Test_DecodeBits(t *testing.T) {
	code := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	morseCode := PerfectDecodeBits(code)
	t.Log(morseCode == ".... . -.--   .--- ..- -.. .")
}

func Test_DecodeMorse(t *testing.T) {
	morseCode := ".... . -.--   .--- ..- -.. ."
	message := PerfectDecodeMorse(morseCode)
	t.Log(message)
}
