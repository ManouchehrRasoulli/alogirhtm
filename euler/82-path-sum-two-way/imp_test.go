package _2_path_sum_two_way

import "testing"

func TestIsPathSum(t *testing.T) {
	bfsBot := newBfs()
	cost := bfsBot.minLeftToRight(matrix)
	t.Log(cost)
}

func TestIsPathSum2(t *testing.T) {
	mx := [][]int{
		{131, 673, 234, 103, 18},
		{201, 96, 342, 965, 150},
		{630, 803, 746, 422, 111},
		{537, 699, 497, 121, 956},
		{805, 732, 524, 37, 331},
	}

	bfsBot := newBfs()
	cost := bfsBot.minLeftToRight(mx)
	t.Log(cost) // 994
}
