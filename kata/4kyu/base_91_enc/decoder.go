package kata

import "math"

func decodedLen(n int) int {
	return int(math.Ceil(float64(n) * 14.0 / 16.0))
}

func decodeBytes(dst, src []byte) int {
	var queue, numBits uint
	var v int = -1

	n := 0
	for i := 0; i < len(src); i++ {
		if decodeMap[src[i]] == 0xff {
			// The character is not in the encoding alphabet.
			return n
		}

		if v == -1 {
			// Start the next value.
			v = int(decodeMap[src[i]])
		} else {
			v += int(decodeMap[src[i]]) * 91
			queue |= uint(v) << numBits

			numBits += 13

			for ok := true; ok; ok = numBits > 7 {
				dst[n] = byte(queue)
				n++

				queue >>= 8
				numBits -= 8
			}

			// Mark this value complete.
			v = -1
		}
	}

	if v != -1 {
		dst[n] = byte(queue | uint(v)<<numBits)
		n++
	}

	return n
}

func Decode(d []byte) []byte {
	dst := make([]byte, decodedLen(len(d)))
	n := decodeBytes(dst, d)
	return dst[:n]
}
