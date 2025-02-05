package kata

import (
	"math/rand"
	"testing"
	"time"
)

func TestLearningMachine(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	mach := NewMachine()

	type test struct {
		action int
		num    int
		result int
		msg    string
	}

	actions := Actions()

	for i := 0; i < 200; i++ {
		number := rnd.Intn(100)
		cmd := i % 5

		num := mach.Command(cmd, number)

		mach.Response(actions[cmd](number) == num)
	}

	tests := []test{{0, 100, 101, "Should apply the num + 1 action to the command 0"},
		{1, 100, 0, "Should apply the num * 0 action to the command 1"},
		{2, 100, 50, "Should apply the num / 2 action to the command 2"},
		{3, 1, 100, "Should apply the num * 100 action to the command 3"},
		{4, 100, 0, "Should apply the num % 2 action to the command 4"}}

	for _, test := range tests {
		t.Run(test.msg, func(t *testing.T) {
			if mach.Command(test.action, test.num) != test.result {
				t.Fatalf("\nExpected: %d\nActual: %d", test.result, mach.Command(test.action, test.num))
			}
		})
	}
}
