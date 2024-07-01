package _kyu

import (
	"fmt"
)

type format string

const (
	year   format = "year"
	day    format = "day"
	hour   format = "hour"
	minute format = "minute"
	second format = "second"
)

type timeSector struct {
	value  int64
	format format
}

func (t timeSector) String() string {
	if t.value == 0 {
		return ""
	}

	v := fmt.Sprintf("%d %s", t.value, t.format)
	if t.value > 1 {
		v += "s"
	}

	return v
}

func parseSectors(seconds int64) []string {
	sectors := []string{
		timeSector{
			value:  seconds / 31536000,
			format: year,
		}.String(),
		timeSector{
			value:  seconds / 86400 % 365,
			format: day,
		}.String(),
		timeSector{
			value:  seconds / 3600 % 24,
			format: hour,
		}.String(),
		timeSector{
			value:  seconds / 60 % 60,
			format: minute,
		}.String(),
		timeSector{
			value:  seconds % 60,
			format: second,
		}.String(),
	}

	sc := make([]string, 0)
	for _, v := range sectors {
		if len(v) == 0 {
			continue
		}
		sc = append(sc, v)
	}

	return sc
}

func FormatDuration(seconds int64) string {
	fmt.Println(seconds)

	if seconds == 0 {
		return "now"
	}

	s := parseSectors(seconds)
	if len(s) == 1 {
		return s[0]
	}

	mx := len(s) - 1
	formated := ""
	for i := mx; i >= 0; i-- {
		if i == mx {
			formated = fmt.Sprintf(" and %s", s[i])
			continue
		}

		if i == 0 {
			formated = fmt.Sprintf("%s%s", s[i], formated)
			continue
		}

		formated = fmt.Sprintf(", %s%s", s[i], formated)
	}

	return formated
}
