package kata

import (
	"fmt"
	"strings"
)

const (
	wordSpliter = "00000000000000"
	charSpliter = "000000"
	dotSpliter  = "00"
	dotBit      = "11"
	dashBit     = "111111"
)

var codes map[string]string = map[string]string{
	"-": "T", "--": "M",
	"---":   "O",
	"-----": "0", "----.": "9", "---..": "8",
	"---...": ":", "--.": "G",
	"--.-": "Q", "--..": "Z",
	"--..--": ",", "--...": "7",
	"-.": "N", "-.-": "K", "-.--": "Y",
	"-.--.":  "(",
	"-.--.-": ")", "-.-.": "C",
	"-.-.--": "!",
	"-.-.-.": ";", "-..": "D",
	"-..-": "X", "-..-.": "/",
	"-...": "B", "-...-": "=", "-....": "6",
	"-....-": "-", ".": "E",
	".-": "A", ".--": "W",
	".---": "J", ".----": "1", ".----.": "'", ".--.": "P",
	".--.-.": "@", ".-.": "R", ".-.-.": "+",
	".-.-.-": ".", ".-..": "L", ".-..-.": "\"",
	".-...": "&", "..": "I", "..-": "U", "..---": "2", "..--.-": "_", "..--..": "?", "..-.": "F",
	"...": "S", "...-": "V", "...--": "3",
	"...---...": "SOS", "...-..-": "$",
	"....": "H", "....-": "4", ".....": "5",
	"   ": " ", " ": "",
}

func morseCodes(morse string) string {
	return codes[morse]
}

/* Perfect bits without noise */

func PerfectDecodeBits(bits string) string {
	fmt.Println(bits)

	morseCode := strings.Builder{}
	bitString := strings.TrimRight(strings.TrimLeft(bits, "0"), "0")
	words := strings.Split(bitString, wordSpliter)
	for wi, w := range words {
		chars := strings.Split(w, charSpliter)
		for ci, c := range chars {
			dots := strings.Split(c, dotSpliter)
			for _, d := range dots {
				switch d {
				case dotBit:
					_, _ = morseCode.WriteString(".")
					continue
				case dashBit:
					_, _ = morseCode.WriteString("-")
					continue
				}
			}

			if ci != len(chars)-1 {
				_, _ = morseCode.WriteString(" ")
			}
		}

		if wi != len(words)-1 {
			_, _ = morseCode.WriteString("   ")
		}
	}

	return morseCode.String()
}

func PerfectDecodeMorse(morseCode string) string {
	fmt.Println(morseCode)

	morseCode = strings.TrimRight(strings.TrimLeft(morseCode, " "), " ")

	codes := strings.Split(morseCode, " ")
	var result strings.Builder

	for _, c := range codes {
		if c == "" {
			_, _ = result.WriteString(" ")
			continue
		}
		_, _ = result.WriteString(morseCodes(c))
	}

	message := strings.ReplaceAll(result.String(), "  ", " ")
	fmt.Println(message)
	return message
}
