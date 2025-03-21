package day17

import "testing"

func TestAdv_Execute(t *testing.T) {
	registers := []int{729, 4, 0}
	t.Log(registers) // [729 4 0]
	adv := Adv{
		combo: LiteralComboValue1,
	}
	t.Log(adv.Execute(0, registers)) // 1
	t.Log(registers)                 // [364 4 0]
	adv = Adv{
		combo: ComboValueRegisterB,
	}
	t.Log(adv.Execute(0, registers)) // 1
	t.Log(registers)                 // [45 4 0]
}

func TestBxl_Execute(t *testing.T) {
	registers := []int{729, 4, 0}
	t.Log(registers) // [729 4 0]
	bxl := Bxl{
		combo: LiteralComboValue3,
	}
	t.Log(bxl.Execute(0, registers)) // 1
	t.Log(registers)                 // [729 7 0] --> 4 xor 3
}
