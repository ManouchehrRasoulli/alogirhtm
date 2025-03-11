package day1

import "testing"

func TestInit(t *testing.T) {
	t.Log("initiated ...")

	t.Log(sources)
	t.Log(syncs)

	totalDiff := int64(0)

	for i := 0; i < len(sources); i++ {
		totalDiff += diff(sources[i], syncs[i])
	}

	t.Log(totalDiff)
}
