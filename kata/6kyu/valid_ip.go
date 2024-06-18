package _kyu

import (
	"net"
	"strconv"
	"strings"
)

func BestIs_valid_ip(ip string) bool {
	res := net.ParseIP(ip)
	if res == nil {
		return false
	}
	return true
}

func Is_valid_ip(ip string) bool {
	dotDecimal := strings.Split(ip, ".")
	if len(dotDecimal) != 4 {
		return false
	}
	for _, d := range dotDecimal {
		if len(d) > 1 && strings.HasPrefix(d, "0") {
			return false
		}
		di, err := strconv.Atoi(d)
		if err != nil || di > 256 || di < 0 {
			return false
		}
	}
	return true
}
