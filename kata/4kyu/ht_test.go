package _kyu

import "testing"

func TestHumanReadableTime(t *testing.T) {
	t.Log(FormatDuration(3662))     // 1h:1m:2s
	t.Log(FormatDuration(94608000)) // 3y
	t.Log(FormatDuration(99042673)) // 3y:51d:7h:51m:13s
}
