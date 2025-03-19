package day6

type Room struct {
	ground [][]rune
}

func (r *Room) WithinBound(x, y int) bool {
	if x < 0 || y < 0 || x >= len(r.ground) || y >= len(r.ground[x]) {
		return false
	}

	return true
}

func (r *Room) Size() (int, int) {
	return len(r.ground), len(r.ground[0])
}

func (r *Room) MakeObstacle(x, y int) {
	r.Change(x, y, '#')
}

func (r *Room) MakeFree(x, y int) {
	r.Change(x, y, '.')
}

func (r *Room) Change(x, y int, char rune) {
	r.ground[x][y] = char
}

func (r *Room) CanIMoveTo(x, y int) bool {
	if r.ground[x][y] == '#' {
		return false
	}

	return true
}

func (r *Room) String() string {
	var (
		s string
	)

	s += "\n"

	for _, row := range r.ground {
		for _, char := range row {
			s += string(char)
		}
		s += "\n"
	}

	return s
}
