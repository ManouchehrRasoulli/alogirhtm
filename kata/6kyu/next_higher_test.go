package _kyu

import (
	"strconv"
	"testing"
)

/*
Your task is to find the next higher number (int) with the same number of '1'- Bits.

I.e. as many 1 bits as before and output next higher than input. Input is always an int between 1 and 1<<30 (possibly inclusive). No bad cases or special tricks...

Some easy examples:
Input: 129  => Output: 130 (10000001 => 10000010)
Input: 127 => Output: 191 (01111111 => 10111111)
Input: 1 => Output: 2 (01 => 10)
Input: 323423 => Output: 323439 (1001110111101011111 => 1001110111101101111)
First some static tests, later on many random tests too;-)!
*/

func TestNextHigher2(t *testing.T) {
	t.Log(strconv.FormatInt(int64(805306367), 2))
	t.Log(strconv.FormatInt(int64(-805306367), 2))
	t.Log(strconv.FormatInt(int64(939524095), 2))
	t.Log(strconv.ParseInt("0101111111111111111111111111111", 2, 64))
	t.Log(strconv.ParseInt("0110111111111111111111111111111", 2, 64))

	t.Log(strconv.FormatInt(int64(128), 2))
	t.Log(strconv.FormatInt(int64(256), 2))

	t.Log(strconv.FormatInt(int64(1), 2))
	t.Log(strconv.FormatInt(int64(2), 2))

	t.Log(strconv.FormatInt(int64(1022), 2))
	t.Log(strconv.FormatInt(int64(1279), 2))
}

func TestNumOne(t *testing.T) {
	t.Log(numOne("1100"))
}

func TestNextHigherFirstAttempt(t *testing.T) {
	arr := [...][2]int{
		{128, 256},
		{1, 2},
		{1022, 1279},
		{127, 191},
		{1253343, 1253359},
		{805306367, 939524095}, // didn't response
	}
	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		funcOut := NextHigherFirstAttempt(input)
		if funcOut != output {
			t.Log("invalid output", "   ---   ", input, "   ***   ", output, "   &&&   ", funcOut)
		}
	}
}

func TestNextHigher(t *testing.T) {
	arr := [...][2]int{
		{128, 256},
		{1, 2},
		{1022, 1279},
		{127, 191},
		{1253343, 1253359},
		{805306367, 939524095}, // didn't response
	}
	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		funcOut := NextHigher(input)
		if funcOut != output {
			t.Log("invalid output", "   ---   ", input, "   ***   ", output, "   &&&   ", funcOut)
		}
	}
}
