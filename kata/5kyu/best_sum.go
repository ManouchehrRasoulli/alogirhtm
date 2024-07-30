package _kyu

func ChooseBestSum(t, k int, ls []int) int {
	outerbest := -1
	for i, d := range ls {
		if len(ls) < k {
			continue
		}
		if k > 1 {
			innerbest := ChooseBestSum(t-d, k-1, ls[i+1:])
			if innerbest < 0 {
				continue
			}
			d += innerbest
		}
		if d <= t && d > outerbest {
			outerbest = d
		}
	}
	return outerbest
}
