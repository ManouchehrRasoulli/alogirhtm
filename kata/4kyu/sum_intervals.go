package _kyu

import "fmt"

func merge(a, b [2]int) (bool, [2]int) {
	if a[0] >= b[0] && a[1] >= b[0] && a[1] <= b[1] {
		return true, b
	}
	if b[0] >= a[0] && b[1] >= a[0] && b[1] <= a[1] {
		return true, a
	}
	if a[0] <= b[0] && a[1] >= b[0] && a[1] <= b[1] {
		return true, [2]int{a[0], b[1]}
	}
	if b[0] <= a[0] && b[1] >= a[0] && b[1] <= a[1] {
		return true, [2]int{b[0], a[1]}
	}
	return false, [2]int{}
}

func deleteFromInterval(interval [][2]int, x [2]int) [][2]int {
	n := make([][2]int, 0)
	for _, v := range interval {
		if v[0] == x[0] && v[1] == x[1] {
			continue
		}
		n = append(n, v)
	}
	return n
}

func reduce(intervals [][2]int) [][2]int {
	for i, v := range intervals {
		for j, vv := range intervals {
			if i != j {
				if t, r := merge(v, vv); t {
					intervals = deleteFromInterval(intervals, v)
					intervals = deleteFromInterval(intervals, vv)
					intervals = append(intervals, r)
				}
			}
		}
	}

	return intervals
}

func reduceMost(interval [][2]int) [][2]int {
	m := reduce(interval)
	if len(m) == len(interval) {
		return m
	}
	return reduceMost(m)
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

func dif(a, b int) int {
	if a < 0 && b < 0 {
		return dif(abs(a), abs(b))
	}

	if a < 0 && b > 0 {
		return abs(a) + b
	}
	if b < 0 && a > 0 {
		return abs(b) + a
	}

	if a > b {
		return a - b
	}
	return b - a
}

func SumOfIntervals(intervals [][2]int) int {
	// your code here

	fmt.Printf("%+v\n", intervals)
	intervals = reduceMost(intervals)
	fmt.Printf("after %+v\n", intervals)

	sm := 0
	for _, v := range intervals {
		sm += dif(v[1], v[0])
	}

	return sm
}
