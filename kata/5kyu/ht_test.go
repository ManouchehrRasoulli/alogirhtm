package _kyu

import "testing"

func TestHumanReadableTime(t *testing.T) {
	t.Log(HumanReadableTime(0))
	t.Log(HumanReadableTime(59))
	t.Log(HumanReadableTime(90))
	t.Log(HumanReadableTime(86400))
	t.Log(HumanReadableTime(359999))
}
