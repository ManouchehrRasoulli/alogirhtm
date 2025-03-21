package day17

import (
	"fmt"
)

const (
	NoValue = -1
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
	Execute(pc int, registers *[3]int) (npc int, output int)
	fmt.Stringer
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

func (a Adv) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterA] /= twoOfPower(ComboToValue(a.combo, registers))
	return npc, NoValue
}

func (a Adv) String() string {
	return fmt.Sprintf("Adv(%d,%d)", AdvInstruction, a.combo)
}

type Bxl struct {
	combo Combo
}

func (b Bxl) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterB] ^= int(b.combo)
	return npc, NoValue
}

func (b Bxl) String() string {
	return fmt.Sprintf("Bxl(%d,%d)", BxlInstruction, b.combo)
}

type Bst struct {
	combo Combo
}

func (b Bst) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterB] = ComboToValue(b.combo, registers) % 8
	return npc, NoValue
}

func (b Bst) String() string {
	return fmt.Sprintf("Bst(%d,%d)", BstInstruction, b.combo)
}

type Jnz struct {
	combo Combo
}

func (j Jnz) Execute(pc int, registers *[3]int) (npc int, output int) {
	if registers[RegisterA] == 0 {
		npc = pc + 1
		return npc, NoValue
	}

	return int(j.combo / 2), NoValue // jump into the instruction location
}

func (j Jnz) String() string {
	return fmt.Sprintf("Jnz(%d,%d)", JnzInstruction, j.combo)
}

type Bxc struct {
	combo Combo
}

func (b Bxc) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterB] ^= registers[RegisterC]
	return npc, NoValue
}

func (b Bxc) String() string {
	return fmt.Sprintf("Bxc(%d,%d)", BxcInstruction, b.combo)
}

type Out struct {
	combo Combo
}

func (o Out) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	return npc, ComboToValue(o.combo, registers) % 8
}

func (o Out) String() string {
	return fmt.Sprintf("Out(%d,%d)", OutInstruction, o.combo)
}

type Bdv struct {
	combo Combo
}

func (b Bdv) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterB] = registers[RegisterA] / twoOfPower(ComboToValue(b.combo, registers))
	return npc, NoValue
}

func (b Bdv) String() string {
	return fmt.Sprintf("Bdv(%d,%d)", BdvInstruction, b.combo)
}

type Cdv struct {
	combo Combo
}

func (c Cdv) Execute(pc int, registers *[3]int) (npc int, output int) {
	npc = pc + 1
	registers[RegisterC] = registers[RegisterA] / twoOfPower(ComboToValue(c.combo, registers))
	return npc, NoValue
}

func (c Cdv) String() string {
	return fmt.Sprintf("Cdv(%d,%d)", CdvInstruction, c.combo)
}

func ComboToValue(combo Combo, registers *[3]int) int {
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

func twoOfPower(i int) int {
	return 1 << i
}
