package _1_path_sum

import "testing"

func TestEmbed(t *testing.T) {
	t.Log(fileContent)
	t.Log(matrix)
	t.Log(len(matrix), len(matrix[1]))
}

func TestIsPathSum(t *testing.T) {
	t.Log(MinPathSum(matrix)) // 427337
}

func TestIsPathSum2(t *testing.T) {
	mx := [][]int{
		{131, 673, 234, 103, 18},
		{201, 96, 342, 965, 150},
		{630, 803, 746, 422, 111},
		{537, 699, 497, 121, 956},
		{805, 732, 524, 37, 331},
	}

	t.Log(MinPathSum(mx)) // 2427
}
