package _kyu

import "testing"

func TestIpsBetween(t *testing.T) {
	t.Log(IpsBetween("10.0.0.0", "10.0.0.50"))
	t.Log(IpsBetween("20.0.0.10", "20.0.1.0"))
	t.Log(IpsBetween("10.0.0.0", "10.0.1.0"))
}
