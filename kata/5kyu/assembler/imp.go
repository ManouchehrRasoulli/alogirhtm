package kata

import (
	"log"
	"strconv"
	"strings"
)

const (
	IncrementOperation   = "inc"
	DecrementOperation   = "dec"
	MoveOperation        = "mov"
	JumpNotZeroOperation = "jnz"
)

type instruction interface {
	execute(variables map[string]int, pc int) (npc int)
}

type inc struct {
	op       string
	variable string
}

func (i inc) execute(variables map[string]int, pc int) (npc int) {
	npc = pc + 1
	variables[i.variable]++
	return
}

type dec struct {
	op       string
	variable string
}

func (d dec) execute(variables map[string]int, pc int) (npc int) {
	npc = pc + 1
	variables[d.variable]--
	return
}

type move struct {
	op   string
	to   string
	from string
}

func (m move) execute(variables map[string]int, pc int) (npc int) {
	npc = pc + 1

	if v, ok := variables[m.from]; !ok {
		// it's not a variable
		variables[m.to], _ = strconv.Atoi(m.from)
	} else {
		variables[m.to] = v
	}

	return
}

type jumpNotZero struct {
	op        string
	condition string
	index     int
}

func (j jumpNotZero) execute(variables map[string]int, pc int) (npc int) {
	if v, ok := variables[j.condition]; !ok {
		npc = pc + j.index
	} else {
		if v == 0 {
			npc = pc + 1
		} else {
			npc = pc + j.index
		}
	}

	return
}

func decodeStringToInstruction(instruction string) instruction {
	instructionArray := strings.Split(instruction, " ")
	command := instructionArray[0]
	switch command {
	case IncrementOperation:
		variable := instructionArray[1]
		return inc{
			op:       IncrementOperation,
			variable: variable,
		}
	case DecrementOperation:
		variable := instructionArray[1]
		return dec{
			op:       DecrementOperation,
			variable: variable,
		}
	case MoveOperation:
		to := instructionArray[1]
		from := instructionArray[2]
		return move{
			op:   MoveOperation,
			to:   to,
			from: from,
		}
	case JumpNotZeroOperation:
		variable := instructionArray[1]
		switch variable {
		case IncrementOperation, DecrementOperation, MoveOperation:
			return jumpNotZero{
				op:        JumpNotZeroOperation,
				condition: "1",
				index:     1,
			}
		}
		index, _ := strconv.Atoi(instructionArray[2])
		return jumpNotZero{
			op:        JumpNotZeroOperation,
			condition: variable,
			index:     index,
		}
	}

	return nil
}

type assembler struct {
	variables    map[string]int
	pc           int
	instructions []instruction
}

func newAssembler(program []string) *assembler {
	a := &assembler{
		variables:    make(map[string]int),
		pc:           0,
		instructions: nil,
	}

	instructions := make([]instruction, 0)
	for _, p := range program {
		i := decodeStringToInstruction(p)
		if i == nil {
			log.Fatal("invalid instruction", p)
		}

		instructions = append(instructions, i)
	}

	a.instructions = instructions
	return a
}

func (a *assembler) execute() (variables map[string]int) {
	for a.pc < len(a.instructions) {
		a.pc = a.instructions[a.pc].execute(a.variables, a.pc)
	}

	return a.variables
}

func SimpleAssembler(program []string) map[string]int {
	log.Printf("[]string{\"%s\"}", strings.Join(program, "\",\""))
	a := newAssembler(program)
	return a.execute()
}
