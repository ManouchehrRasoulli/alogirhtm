package day17

import (
	"fmt"
)

const (
	registerA = iota
	registerB
	registerC
)

type Opcode int

const (
	AdvInstruction Opcode = iota
	BxlInstruction
	BstInstruction
	JnzInstruction
	BxcInstruction
	OutInstruction
	BdvInstruction
	CdvInstruction
)

type Combo int

const (
	LiteralComboValue0 Combo = iota
	LiteralComboValue1
	LiteralComboValue2
	LiteralComboValue3
	ComboValueRegisterA
	ComboValueRegisterB
	ComboValueRegisterC
	NotValidComboOperand
)

type Instruction interface {
	// Execute
	// give pc into the instruction and then
	// the next value for pc calculated and will return
	// with that instruction
	Execute(pc int, registers []int) (npc int)
}

func ComboToValue(combo Combo, registers []int) int {
	switch combo {
	case LiteralComboValue0, LiteralComboValue1, LiteralComboValue2, LiteralComboValue3:
		return int(combo)
	case ComboValueRegisterA, ComboValueRegisterB, ComboValueRegisterC:
		return registers[int(combo)%4] // return the register value for combo operand
	case NotValidComboOperand:
		panic(fmt.Sprintf("invalid combo operand 7 !!"))
	default:
		panic(fmt.Sprintf("invalid combo operand: %d", combo))
	}
}

type Adv struct {
	Combo Combo
}

func (a Adv) Execute(pc int, registers []int) (npc int) {
	panic("implement me")
}

var _ Instruction = &Adv{}
