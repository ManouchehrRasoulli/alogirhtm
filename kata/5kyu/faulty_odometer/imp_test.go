package kata

import "testing"

func TestFaultyOdometerNaive(t *testing.T) {
	t.Log(FaultyOdometerNaive(13))        // 12
	t.Log(FaultyOdometerNaive(15))        // 13
	t.Log(FaultyOdometerNaive(55))        // 40
	t.Log(FaultyOdometerNaive(2005))      // 1462
	t.Log(FaultyOdometerNaive(165826622)) // 69517865
}

func TestFaultyOdometerNaiveSecond(t *testing.T) {
	t.Log(FaultyOdometerNaiveSecond(13))        // 12
	t.Log(FaultyOdometerNaiveSecond(15))        // 13
	t.Log(FaultyOdometerNaiveSecond(55))        // 40
	t.Log(FaultyOdometerNaiveSecond(2005))      // 1462
	t.Log(FaultyOdometerNaiveSecond(165826622)) // 69517865
}

func TestFaultyOdometer(t *testing.T) {
	t.Log(FaultyOdometer(13))        // 12
	t.Log(FaultyOdometer(15))        // 13
	t.Log(FaultyOdometer(55))        // 40
	t.Log(FaultyOdometer(2005))      // 1462
	t.Log(FaultyOdometer(165826622)) // 69517865
}
