package kata

var pairs = [][2]byte{
	{'0', '0'},
	{'1', '1'},
	{'6', '9'},
	{'8', '8'},
	{'9', '6'},
}

func isInRange(s, low, high string) bool {
	if len(s) < len(low) || (len(s) == len(low) && s < low) {
		return false
	}
	if len(s) > len(high) || (len(s) == len(high) && s > high) {
		return false
	}
	return true
}

func generateStrobos(n, length int, current []byte, low, high string, result *uint64) {
	if n == 0 {
		num := string(current)
		if isInRange(num, low, high) {
			*result++
		}
		return
	}
	if n == 1 {
		for _, mid := range []byte{'0', '1', '8'} {
			current[(length-n)/2] = mid
			if isInRange(string(current), low, high) {
				*result++
			}
		}
		return
	}
	for _, p := range pairs {
		// no leading zeros unless length is 1
		if (length == n) && p[0] == '0' {
			continue
		}
		current[(length-n)/2] = p[0]
		current[length-1-(length-n)/2] = p[1]
		generateStrobos(n-2, length, current, low, high, result)
	}
}

func __UpsideDown(n1, n2 string) uint64 {
	var result uint64
	for length := len(n1); length <= len(n2); length++ {
		current := make([]byte, length)
		generateStrobos(length, length, current, n1, n2, &result)
	}
	return result
}
