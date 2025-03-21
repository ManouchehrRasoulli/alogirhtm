package day17

import (
	"fmt"
	"log"
	"runtime/debug"
	"slices"
	"strings"
)

type Pc struct {
	oldPc             int
	pc                int
	oldRegisters      [3]int
	registers         [3]int
	instructionString []string
	instructions      []Instruction
	outputs           []string
}

func (pc *Pc) Reset() {
	pc.pc = pc.oldPc
	pc.registers = pc.oldRegisters
	pc.outputs = nil
}

func (pc *Pc) String() string {
	data := "\n🖳\n"

	data += fmt.Sprintf("registers:: %+v\n", pc.registers)

	for _, i := range pc.instructions {
		data += i.String()
		data += "\n"
	}

	if len(pc.outputs) > 0 {
		data += fmt.Sprintf("outputs:\n%+v", pc.outputs)
	}

	return data
}

func (pc *Pc) Execute() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Printf("%+v\n", pc)
		}
	}()

	var (
		output int
		run    int
	)
	for pc.pc < len(pc.instructions) {
		run++
		pc.pc, output = pc.instructions[pc.pc].Execute(pc.pc, &pc.registers)
		if output != NoValue {
			pc.outputs = append(pc.outputs, fmt.Sprintf("%d", output))
		}
	}
}

func (pc *Pc) FindRegisterAValue() int {
	var (
		a = 0
		d = 1
	)

	for {
		pc.Reset()
		pc.registers[RegisterA] = a

		pc.Execute()

		oc := len(pc.outputs)
		ic := len(pc.instructionString)

		if slices.Equal(pc.instructionString, pc.outputs) || oc > ic {
			break
		}

		if oc-d >= 0 && ic-d >= 0 &&
			pc.outputs[oc-d] == pc.instructionString[ic-d] &&
			(d < 2 || pc.outputs[oc-d+1] == pc.instructionString[ic-d+1]) {
			d++ // if each of these instructions matched then we can move into next instruction output match and shift a with 3 bits
			a <<= 3
		} else {
			a++
		}
	}

	return a
}

func (pc *Pc) LogOutputs() {
	log.Println(strings.Join(pc.outputs, ","))
}
