package kata

func Parse(data string) []int {
	output := make([]int, 0)
	val := 0

	for _, op := range data {
		switch op {
		case 'i':
			val++
			break
		case 'd':
			val--
			break
		case 's':
			val *= val
			break
		case 'o':
			output = append(output, val)
			break
		}
	}

	return output
}
