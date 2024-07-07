package next_bigger

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(nsa []string, begin, end int) {
	for begin <= end {
		nsa[begin], nsa[end] = nsa[end], nsa[begin]
		begin++
		end--
	}
}

func NextBigger(n int) int {
	ns := fmt.Sprintf("%d", n)
	nsa := strings.Split(ns, "")
	if len(nsa) == 1 {
		return -1
	}

	var i int
	for i = len(nsa) - 1; i > 0; i-- {
		if nsa[i] > nsa[i-1] {
			break
		}
	}
	if i != 0 {
		for j := len(nsa) - 1; j >= i; j-- {
			if nsa[i-1] < nsa[j] {
				nsa[j], nsa[i-1] = nsa[i-1], nsa[j]
				break
			}
		}
	}

	reverse(nsa, i, len(nsa)-1)

	v, _ := strconv.Atoi(strings.Join(nsa, ""))
	if v > n {
		return v
	}

	return -1
}
