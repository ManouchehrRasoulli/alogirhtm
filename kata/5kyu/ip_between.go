package _kyu

import (
	"strconv"
	"strings"
)

func ipToInt(ip string) [4]int {
	ipb := strings.Split(ip, ".")
	if len(ipb) != 4 {
		return [4]int{}
	}

	ipv := [4]int{}
	for i, s := range ipb {
		iv, _ := strconv.Atoi(s)
		ipv[i] = iv
	}

	return ipv
}

func IpsBetween(start, end string) int {
	str := ipToInt(start)
	en := ipToInt(end)

	var p, q int
	for i := 0; i < 4; i++ {
		p = p*256 + str[i]
		q = q*256 + en[i]
	}

	return q - p
}
