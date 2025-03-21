package day17

import (
	"fmt"
	"log"
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

const (
	RegisterA = int(ComboValueRegisterA) % 4
	RegisterB = int(ComboValueRegisterB) % 4
	RegisterC = int(ComboValueRegisterC) % 4
)

type Instruction interface {
	// Execute
	// give pc into the instruction and then
	// the next value for pc calculated and will return
	// with that instruction
	Execute(pc int, registers []int) (npc int)
}

func InstructionFromCode(code Opcode, combo Combo) Instruction {
	switch code {
	case AdvInstruction:
		return Adv{
			combo: combo,
		}
	case BxlInstruction:
		return Bxl{
			combo: combo,
		}
	case BstInstruction:
		return Bst{
			combo: combo,
		}
	case JnzInstruction:
		return Jnz{
			combo: combo,
		}
	case BxcInstruction:
		return Bxc{
			combo: combo,
		}
	case OutInstruction:
		return Out{
			combo: combo,
		}
	case BdvInstruction:
		return Bdv{
			combo: combo,
		}
	case CdvInstruction:
		return Cdv{
			combo: combo,
		}
	}

	panic(fmt.Sprintf("invalid instruction code: %d", code))
}

type Adv struct {
	combo Combo
}

func (a Adv) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterA] /= ComboToValue(a.combo, registers) << 1
	return
}

type Bxl struct {
	combo Combo
}

func (b Bxl) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterB] ^= ComboToValue(b.combo, registers)
	return
}

type Bst struct {
	combo Combo
}

func (b Bst) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterB] = ComboToValue(b.combo, registers) % 8
	return
}

type Jnz struct {
	combo Combo
}

func (j Jnz) Execute(pc int, registers []int) (npc int) {
	if registers[RegisterA] == 0 {
		npc = pc + 1
		return npc
	}

	return int(j.combo / 2) // jump into the instruction location
}

type Bxc struct {
	combo Combo
}

func (b Bxc) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterB] ^= registers[RegisterC]
	return
}

type Out struct {
	combo Combo
}

func (o Out) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	log.Printf("%d ", ComboToValue(o.combo, registers)%8)
	return
}

type Bdv struct {
	combo Combo
}

func (b Bdv) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterB] = registers[RegisterA] / ComboToValue(b.combo, registers) << 1
	return
}

type Cdv struct {
	combo Combo
}

func (c Cdv) Execute(pc int, registers []int) (npc int) {
	npc = pc + 1
	registers[RegisterC] = registers[RegisterA] / ComboToValue(c.combo, registers) << 1
	return
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
