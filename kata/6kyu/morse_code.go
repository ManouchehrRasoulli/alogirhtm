package _kyu

import (
	"strings"
)

var MORSE_CODE = map[string]string{
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

func DecodeMorse(morseCode string) string {
	MORSE_CODE[""] = " "
	morseCode = strings.TrimRight(strings.TrimLeft(morseCode, " "), " ")

	codes := strings.Split(morseCode, " ")
	var result strings.Builder

	for _, c := range codes {
		result.WriteString(MORSE_CODE[c])
	}

	return strings.ReplaceAll(result.String(), "  ", " ")
}
