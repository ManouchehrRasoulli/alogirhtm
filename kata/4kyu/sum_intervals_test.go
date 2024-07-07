package _kyu

import "testing"

func Test_merge(t *testing.T) {
	t.Log(merge([2]int{1, 2}, [2]int{3, 4}))
	t.Log(merge([2]int{1, 3}, [2]int{3, 4}))
	t.Log(merge([2]int{1, 4}, [2]int{3, 6}))
	t.Log(merge([2]int{4, 10}, [2]int{3, 4}))
	t.Log(merge([2]int{7, 10}, [2]int{5, 6}))
	t.Log(merge([2]int{1, 5}, [2]int{1, 6}))
}

func Test_deleteFromInterval(t *testing.T) {
	v := [][2]int{
		{1, 5},
		{10, 20},
		{1, 6},
		{16, 19},
		{5, 11}}

	v = deleteFromInterval(v, [2]int{5, 11})
	t.Log(v)
}

func Test_reduce(t *testing.T) {
	t.Log(reduceMost([][2]int{
		{1, 5},
		{10, 20},
		{1, 6},
		{16, 19},
		{5, 11}}))
}

func Test_reduceMost(t *testing.T) {
	t.Log(reduceMost([][2]int{
		{1, 4},
		{7, 10},
		{3, 5},
	}))
}

func Test_abs(t *testing.T) {
	t.Log(dif(-100000000, 20))
	t.Log(abs(-100000000))
}

func TestSumOfIntervals(t *testing.T) {
	t.Log(SumOfIntervals([][2]int{
		{1, 5},
		{10, 20},
		{1, 6},
		{16, 19},
		{5, 11}}))

	t.Log(SumOfIntervals([][2]int{
		{1, 4},
		{7, 10},
		{3, 5},
	}))

	t.Log(SumOfIntervals([][2]int{
		{0, 20},
		{-100000000, 10},
		{30, 40},
	}))

	t.Log(SumOfIntervals([][2]int{
		{1, 5},
	}))

	t.Log(SumOfIntervals([][2]int{
		{1, 5},
		{6, 10},
	}))
}
