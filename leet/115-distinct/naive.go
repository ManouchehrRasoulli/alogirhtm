package distinct

func process(s string) map[rune][]int {
	m := make(map[rune][]int)
	for i, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = []int{i}
		} else {
			m[v] = append(m[v], i)
		}
	}

	return m
}

func count(sm map[rune][]int, t string, lasPosS int, tPos int) int {
	if tPos >= len(t) {
		return 1
	}

	x := t[tPos]
	cnt := 0

	for _, v := range sm[rune(x)] {
		if v >= lasPosS { // count-in
			cnt += count(sm, t, v+1, tPos+1)
		}
	}

	return cnt
}

func numDistinct_(s string, t string) int {
	sm := process(s)
	return count(sm, t, 0, 0)
}
