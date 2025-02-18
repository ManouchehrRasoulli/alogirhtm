package _kyu

func FindHeight(cubes int) int {
	var (
		level       = 0
		cubeOnLayer = 0
	)

	for i := 1; ; i++ {
		cubeOnLayer += i
		cubes -= cubeOnLayer
		if cubes < 0 {
			return level
		}
		level++
	}
}
