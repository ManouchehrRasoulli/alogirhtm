package kata

import (
	"fmt"
	"sort"
	"strings"
)

type mx struct {
	s1c int
	s2c int
}

type fx struct {
	i string
	v string
}

func (x *fx) String() string {
	return fmt.Sprintf("%s:%s", x.i, x.v)
}

func isLower(v rune) bool {
	return v >= 'a' && v <= 'z'
}

func Mix(s1, s2 string) string {
	x := make(map[rune]mx)

	for _, v := range s1 {
		if isLower(v) {
			x[v] = mx{
				s1c: x[v].s1c + 1,
				s2c: x[v].s2c,
			}
		}
	}

	for _, v := range s2 {
		if isLower(v) {
			x[v] = mx{
				s1c: x[v].s1c,
				s2c: x[v].s2c + 1,
			}
		}
	}

	f := make([]fx, 0)
	for k, v := range x {
		if !(v.s1c > 1) && !(v.s2c > 1) {
			continue
		}

		sight := "="
		count := v.s1c

		if v.s1c > v.s2c {
			sight = "1"
			count = v.s1c
		} else if v.s2c > v.s1c {
			sight = "2"
			count = v.s2c
		}

		l := fx{
			i: sight,
			v: strings.Repeat(string(k), count),
		}
		f = append(f, l)
	}

	sort.Slice(f, func(i, j int) bool {
		if len(f[i].v) > len(f[j].v) {
			return true
		}
		if len(f[i].v) == len(f[j].v) {
			if f[i].i == f[j].i {
				return f[i].v < f[j].v
			}
			if f[i].i == "1" {
				return true
			}
			if f[i].i == "2" && f[j].i == "=" {
				return true
			}
		}
		return false
	})

	fs := make([]string, 0)
	for _, v := range f {
		fs = append(fs, v.String())
	}

	return strings.Join(fs, "/")
}
