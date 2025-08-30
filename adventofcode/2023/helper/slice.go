package helper

func Insert[T any](s []T, index int, value T) []T {
	if index < 0 || index > len(s) {
		panic("index out of range")
	}
	// append part before index, then value, then the rest
	s = append(s[:index], append([]T{value}, s[index:]...)...)
	return s
}
