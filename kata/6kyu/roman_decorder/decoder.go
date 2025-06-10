package kata

var numbers = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Decode(roman string) int {
	sum := 0
	for i := 1; i <= len(roman); i++ {
		var current int
		if i < len(roman) {
			current = numbers[roman[i]]
		}
		prev := numbers[roman[i-1]]

		if current > prev {
			sum += current - prev
			i++
		} else {
			sum += prev
		}
	}

	return sum
}
