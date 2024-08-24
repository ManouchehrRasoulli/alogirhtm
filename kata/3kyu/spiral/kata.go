package kata

type direction int

const (
	right direction = iota
	down
	left
	up
)

type snake struct {
	d     direction
	loc   [2]int
	steps int
}

func newSnake(loc [2]int, steps int) *snake {
	return &snake{
		d:     up,
		loc:   loc,
		steps: steps,
	}
}

func (s *snake) move(maze [][]int) bool {
	if s.steps < 1 {
		return false
	}

	for i := 0; i <= s.steps; i++ {
		switch s.d {
		case up:
			maze[s.loc[0]][s.loc[1]] = 1
			if i == s.steps {
				s.d = right
				break
			}
			s.loc[0]--
			break
		case right:
			if i != 0 && maze[s.loc[0]+1][s.loc[1]] == 1 {
				s.steps -= 2
				break
			}

			maze[s.loc[0]][s.loc[1]] = 1
			if i == s.steps {
				s.d = down
				s.steps -= 2
				break
			}
			s.loc[1]++
			break
		case down:
			maze[s.loc[0]][s.loc[1]] = 1
			if i == s.steps {
				s.d = left
				break
			}
			s.loc[0]++
			break
		case left:
			if i != 0 && maze[s.loc[0]-1][s.loc[1]] == 1 {
				s.steps -= 2
				break
			}

			maze[s.loc[0]][s.loc[1]] = 1
			if i == s.steps {
				s.d = up
				s.steps -= 2
				break
			}
			s.loc[1]--
			break
		}
	}

	return true
}

func Spiralize(size int) [][]int {
	maze := make([][]int, size)
	for i := 0; i < size; i++ {
		maze[i] = make([]int, size)
		if i == 0 || i == size-1 {
			for j := 0; j < size; j++ {
				maze[i][j] = 1
			}
		}
		maze[i][size-1] = 1
	}

	s := newSnake([2]int{size - 1, 0}, size-3)
	for s.move(maze) {

	}

	return maze
}
