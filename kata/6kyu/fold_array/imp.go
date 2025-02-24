package kata

func fold(arr []int) []int {
	var (
		mid = len(arr) / 2
		odd = len(arr) % 2
	)

	for i := 0; i < mid; i++ {
		arr[i] += arr[len(arr)-i-1]
	}

	return arr[:mid+odd]
}

func FoldArray(arr []int, runs int) []int {
	for i := 0; i < runs; i++ {
		arr = fold(arr)
	}

	return arr
}
