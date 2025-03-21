package day17

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
)

type Pc struct {
	pc           int
	registers    [3]int
	instructions []Instruction
	outputs      []string
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

	log.Printf("pc: %d, outputs: %+v, run: %d\n", pc.pc, pc.outputs, run)
}

func (pc *Pc) LogOutputs() {
	log.Println(strings.Join(pc.outputs, ","))
}
