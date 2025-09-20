package change

import "testing"

func Test_greedyChangeMoney(t *testing.T) {
	buckets := greedyChangeMoney(200, 50, 200)
	t.Log(buckets) // [200]

	buckets = greedyChangeMoney(223, 50, 500)
	t.Log(buckets) // [223]

	buckets = greedyChangeMoney(223, 50, 200)
	t.Log(buckets) // [177, 50]

	buckets = greedyChangeMoney(279, 50, 200)
	t.Log(buckets) // [200, 79]

	buckets = greedyChangeMoney(425, 50, 200)
	t.Log(buckets) // [200, 175, 50]
}
