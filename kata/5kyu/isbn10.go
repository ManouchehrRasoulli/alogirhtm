package _kyu

import (
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	if len(isbn) < 10 {
		return false
	}

	isbn = strings.ToUpper(isbn)

	isbns := strings.Split(isbn, "")
	isbnsd := make([]int, 0)

	for i, x := range isbns {
		xi, err := strconv.Atoi(x)
		if err != nil {
			if x == "X" && i == 9 {
				isbnsd = append(isbnsd, 10)
				continue
			}
			return false
		}
		isbnsd = append(isbnsd, xi)
	}

	count := 0
	for x, i := range isbnsd {
		count += i * (x + 1)
	}

	return count%11 == 0 // implement proper solution here
}
