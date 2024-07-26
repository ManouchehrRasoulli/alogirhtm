package _kyu

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	t.Log(FirstNonRepeating("stress"))
	t.Log(FirstNonRepeating("sTress"))
	t.Log(FirstNonRepeating("moonmen"))
}
