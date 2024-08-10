package kata

func MaxDownSum(pyramid [][]int, x, y int) int {
	s := pyramid[x][y]
	ls := s + pyramid[x+1][y]
	rs := s + pyramid[x+1][y+1]
	if ls > rs {
		return ls
	}
	return rs
}

func LongestSlideDown(pyramid [][]int) int {
	n := len(pyramid) - 1
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			pyramid[i][j] = MaxDownSum(pyramid, i, j)
		}
	}
	return pyramid[0][0]
}
