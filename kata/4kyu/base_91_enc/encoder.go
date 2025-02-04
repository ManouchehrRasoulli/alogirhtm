package kata

import "math"

func encodedLen(n int) int {
	return int(math.Ceil(float64(n) * 16.0 / 13.0))
}

func encodeByte(dst, src []byte) int {
	var queue, numBits uint

	n := 0
	for i := 0; i < len(src); i++ {
		queue |= uint(src[i]) << numBits
		numBits += 8
		if numBits > 13 {
			v := queue & 8191

			queue >>= 13
			numBits -= 13

			dst[n] = encode[v%91]
			n++
			dst[n] = encode[v/91]
			n++
		}
	}

	if numBits > 0 {
		dst[n] = encode[queue%91]
		n++

		if numBits > 7 || queue > 90 {
			dst[n] = encode[queue/91]
			n++
		}
	}

	return n
}

func Encode(d []byte) []byte {
	dst := make([]byte, encodedLen(len(d)))
	n := encodeByte(dst, d)

	return dst[:n]
}
