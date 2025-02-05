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

type operation struct {
	doneForOperations []int
	fn                func(int) int
}

func (o *operation) doesDoneForThatCommand(cmd int) bool {
	for _, doneForOperation := range o.doneForOperations {
		if doneForOperation == cmd {
			return true
		}
	}

	return false
}

type Machine struct {
	operations []operation

	// state
	runningCommand int
	fn             func(int) int

	learned []func(int) int
}

func NewMachine() Machine {
	m := Machine{
		operations: make([]operation, 0),
	}

	functions := Actions()
	for _, fn := range functions {
		m.operations = append(m.operations, operation{
			doneForOperations: make([]int, 0),
			fn:                fn,
		})
	}

	m.learned = make([]func(int) int, len(m.operations))

	return m
}

func (m *Machine) Command(cmd int, num int) int {
	cmd = cmd % len(m.learned)

	if fn := m.learned[cmd]; fn != nil {
		m.fn = fn
		m.runningCommand = cmd

		return fn(num)
	}

	for i, op := range m.operations {
		if !op.doesDoneForThatCommand(cmd) {

			m.fn = op.fn
			m.runningCommand = cmd
			m.operations[i].doneForOperations = append(m.operations[i].doneForOperations, cmd)

			break
		}
	}

	if m.fn == nil {
		for i, _ := range m.operations {
			m.operations[i].doneForOperations = make([]int, 0)
		}

		return m.Command(cmd, num)
	}

	return m.fn(num)
}

func (m *Machine) Response(res bool) {
	if res {
		m.learned[m.runningCommand] = m.fn
	}

	m.fn = nil
	m.runningCommand = 0
}
