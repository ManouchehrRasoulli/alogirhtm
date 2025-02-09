package kata

import (
	"testing"
)

func Test_repairSequence(t *testing.T) {
	t.Logf("{%s}", joinInts(repairSequence([]int{5, 4, 6, 8, 10}), ","))
	t.Logf("{%s}", joinInts(repairSequence([]int{2, 4, 6, 8, 11}), ","))
	t.Logf("{%s}", joinInts(repairSequence([]int{2, 4, 6, 8, 15, 12}), ","))
	t.Logf("{%s}", joinInts(repairSequence([]int{1, 13, 15, 17}), ","))
}
