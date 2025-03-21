package day17

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRegex = regexp.MustCompile(`[-]{0,1}[0-9]+`)
)

func ReadPc(data string) *Pc {
	pc := Pc{
		oldPc:             0,
		pc:                0,
		oldRegisters:      [3]int{0, 0, 0},
		registers:         [3]int{0, 0, 0},
		instructionString: []string{},
		instructions:      make([]Instruction, 0),
		outputs:           nil,
	}

	lines := strings.Split(data, "\n")
	for i, line := range lines {
		if i < 3 { // register data
			rawData := numberRegex.FindAllString(line, -1)[0]
			pc.oldRegisters[i], _ = strconv.Atoi(rawData)
			pc.registers[i], _ = strconv.Atoi(rawData)

			continue
		}

		if len(line) == 0 {
			continue
		}

		instructions := numberRegex.FindAllString(line, -1)
		for j := 0; j < len(instructions); j += 2 {
			opcode, _ := strconv.Atoi(instructions[j])
			combo, _ := strconv.Atoi(instructions[j+1])

			pc.instructionString = append(pc.instructionString, instructions[j], instructions[j+1])
			pc.instructions = append(pc.instructions, InstructionFromCode(Opcode(opcode), Combo(combo)))
		}
	}

	return &pc
}
