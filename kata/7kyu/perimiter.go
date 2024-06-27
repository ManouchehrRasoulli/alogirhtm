package _kyu

func EquableTriangle(a, b, c int) bool {
	if a+b <= c || a+c <= b || b+c <= a {
		return false
	}
	return a*b/2 == a+b+c
}
