package day4

import "testing"

func TestRemoveAccessAbleRolesSample1(t *testing.T) {
	t.Log(RemoveAccessAbleRoles(sample1)) // 43
}

func TestRemoveAccessAbleRolesPuzzle(t *testing.T) {
	t.Log(RemoveAccessAbleRoles(puzzle)) // 8643
}
