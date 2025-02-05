package kata

func Actions() []func(int) int {
	return []func(int) int{
		func(num int) int { // num + 1
			return num + 1
		},
		func(num int) int { // zero
			return num * 0
		},
		func(num int) int {
			return num / 2
		},
		func(num int) int {
			return num * 100
		},
		func(num int) int {
			return num % 2
		},
	}
}

type Machine struct {
	Cmd     int
	Mapping map[int]int
}

func NewMachine() Machine {
	return Machine{Mapping: make(map[int]int)}
}

func (m *Machine) Command(cmd int, num int) int {
	m.Cmd = cmd
	return Actions()[m.Mapping[cmd]](num)
}

func (m *Machine) Response(res bool) {
	if !res {
		if m.Mapping[m.Cmd] == 4 {
			m.Mapping[m.Cmd] = 0
		} else {
			m.Mapping[m.Cmd] += 1
		}
	}
}
