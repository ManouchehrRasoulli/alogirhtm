package kata

import (
	"log"
)

func isAllOne(n uint64, base uint64) bool {
	for n > 0 {
		if n%base != 1 {
			return false
		}

		n /= base
	}

	return true
}

func GetMinBase(n uint64) uint64 {
	log.Print(n)

	if n <= 1 {
		return 0
	}

	for base := uint64(2); base <= 1_000_000; base++ {
		if isAllOne(n, base) {
			return base
		}
	}

	return n - 1
}
