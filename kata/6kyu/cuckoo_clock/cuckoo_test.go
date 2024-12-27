package kata

import "testing"

func Test_CuckooClock(t *testing.T) {
	t.Log("Starting ...")

	t.Log(CuckooClock("12:46", 10)) // 03:00
	t.Log(CuckooClock("03:38", 19)) // 06:00
	t.Log(CuckooClock("10:00", 1))  // 10:00
	t.Log(CuckooClock("10:00", 10)) // 10:00
	t.Log(CuckooClock("10:00", 11)) // 10:15
	t.Log(CuckooClock("10:00", 13)) // 10:45
	t.Log(CuckooClock("10:00", 20)) // 11:00
	t.Log(CuckooClock("09:53", 50)) // 02:30

	t.Log("Finished !")
}
