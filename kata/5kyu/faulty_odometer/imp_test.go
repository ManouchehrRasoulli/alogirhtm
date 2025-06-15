package kata

import "testing"

func TestFaultyOdometerNaive(t *testing.T) {
	t.Log(FaultyOdometerNaive(13))        // 12
	t.Log(FaultyOdometerNaive(15))        // 13
	t.Log(FaultyOdometerNaive(2005))      // 1462
	t.Log(FaultyOdometerNaive(165826622)) // 69517865
}
