package _kyu

import "testing"

func TestIs_valid_ip(t *testing.T) {
	t.Log(Is_valid_ip("192.168.1.1"))
	t.Log(Is_valid_ip("123.456.78.90"))
	t.Log(Is_valid_ip("123.045.067.089"))
	t.Log(Is_valid_ip("0.34.82.53"))
	t.Log(Is_valid_ip("12.34.56"))
	t.Log(Is_valid_ip("12.34.56.-1"))
}
