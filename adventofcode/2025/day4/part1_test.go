package day4

import "testing"

func TestAccessAbleRolesSample1(t *testing.T) {
	t.Log(AccessAbleRoles(sample1)) // 13
}

func TestAccessAbleRolesPuzzle(t *testing.T) {
	t.Log(AccessAbleRoles(puzzle)) // 1393
}
