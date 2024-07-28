package _kyu

import "testing"

func TestValidISBN10(t *testing.T) {
	t.Log(ValidISBN10("1112223339")) // true
	t.Log(ValidISBN10("048665088X")) // true
	t.Log(ValidISBN10("1293000000")) // true
	t.Log(ValidISBN10("1234554321")) // true

	t.Log(ValidISBN10("1234512345")) // false
	t.Log(ValidISBN10("1293"))       // false
	t.Log(ValidISBN10("X123456788")) // false
	t.Log(ValidISBN10("ABCDEFGHIJ")) // false
	t.Log(ValidISBN10("XXXXXXXXXX")) // false
}
