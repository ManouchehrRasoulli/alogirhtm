package _kyu

import (
	"fmt"
	"time"
)

func HumanReadableTime(seconds int) string {
	x := time.Duration(seconds) * time.Second

	h := x / time.Hour
	x = x - h*time.Hour

	m := x / time.Minute
	x = x - m*time.Minute

	s := x / time.Second
	x = x - s*time.Second

	return fmt.Sprintf("%02d:%02d:%02d", int(h), int(m), int(s))
}
