package kata

var cmp []int
var sum []int

func init() {
	cmp = make([]int, 3)
	cmp[0] = 1
	cmp[1] = 1
	cmp[2] = 2

	sum = make([]int, 3)
	sum[0] = 1
	sum[1] = 2
	sum[2] = 4
}

func Perimeter(n int) int {
	if n < len(sum) {
		return 4 * sum[n]
	}

	for i := len(cmp); i <= n; i++ {
		cmp = append(cmp, cmp[i-1]+cmp[i-2])
		sum = append(sum, sum[i-1]+cmp[i])
	}

	return 4 * sum[n]
}
