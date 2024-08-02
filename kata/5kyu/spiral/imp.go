package kata

type position struct {
	x int
	y int
}

type direction int

const (
	right direction = iota
	down
	left
	up
)

func changeDirection(dir direction) direction {
	d := dir + 1
	if d > up {
		return right
	}
	return d
}

func nextPosition(pos position, d direction) position {
	switch d {
	case right:
		pos.y++
		return pos
	case down:
		pos.x++
		return pos
	case left:
		pos.y--
		return pos
	default:
		pos.x--
		return pos
	}
}

func CreateSpiral(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	pos := position{
		x: 0,
		y: 0,
	}
	d := right

	for i := 1; i <= n*n; i++ {
		matrix[pos.x][pos.y] = i
		nextPos := nextPosition(pos, d)
		if nextPos.x > n-1 || nextPos.y > n-1 || nextPos.x < 0 || nextPos.y < 0 {
			d = changeDirection(d)
			nextPos = nextPosition(pos, d)
		}
		if matrix[nextPos.x][nextPos.y] != 0 {
			d = changeDirection(d)
			nextPos = nextPosition(pos, d)
		}
		pos = nextPos
	}

	return matrix
}
