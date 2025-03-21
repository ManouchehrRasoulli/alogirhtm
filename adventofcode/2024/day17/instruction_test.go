package day17

import "testing"

func TestTwoOfPower(t *testing.T) {
	t.Log(twoOfPower(0)) // 1
	t.Log(twoOfPower(3)) // 8
}

func TestAdv_Execute(t *testing.T) {
	registers := [3]int{729, 4, 0}
	t.Log(registers) // [729 4 0]
	adv := Adv{
		combo: LiteralComboValue1,
	}
	t.Log(adv.Execute(0, &registers)) // 1, -1
	t.Log(registers)                  // [364 4 0]
	adv = Adv{
		combo: ComboValueRegisterB,
	}
	t.Log(adv.Execute(0, &registers)) // 1, -1
	t.Log(registers)                  // [22 4 0]
}

func TestBxl_Execute(t *testing.T) {
	registers := [3]int{729, 4, 0}
	t.Log(registers) // [729 4 0]
	bxl := Bxl{
		literal: int(LiteralComboValue3),
	}
	t.Log(bxl.Execute(0, &registers)) // 1, -1
	t.Log(registers)                  // [729 7 0] --> 4 xor 3
}
