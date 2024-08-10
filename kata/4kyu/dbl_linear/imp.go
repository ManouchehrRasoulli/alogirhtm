package kata

func DblLinear(n int) int {

	// Code
	u := []int{1}
	i := 0
	j := 0

	var y int
	var z int

	for len(u) <= n {
		y = 2*u[i] + 1
		z = 3*u[j] + 1

		if y < z {
			u = append(u, y)
			i++
		} else if y > z {
			u = append(u, z)
			j++
		} else {
			u = append(u, y)
			i++
			j++
		}
	}
	return u[len(u)-1]

}
