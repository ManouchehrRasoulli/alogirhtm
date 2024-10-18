package courseschedule

import "testing"

func Test_canFinish(t *testing.T) {
	t.Log(canFinish(2, [][]int{{1, 0}}))
	t.Log(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	t.Log(canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}))
	t.Log(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
