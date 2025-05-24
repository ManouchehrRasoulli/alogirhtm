package day8

import (
	"log"
	"regexp"
	"strings"
)

type Movement struct {
	NextRight string
	NextLeft  string
}

const (
	Left  rune = 'L'
	Right rune = 'R'
	Start      = "AAA"
	End        = "ZZZ"
)

type Input struct {
	Instructions []rune
	Movements    map[string]Movement
}

func ParseInput(in string) *Input {
	var (
		lines = strings.Split(in, "\n")
		input = Input{
			Instructions: make([]rune, 0),
			Movements:    make(map[string]Movement),
		}
		rgx = regexp.MustCompile(`[A-Z]{3}`)
	)

	for i, line := range lines {
		if i == 0 { // instructions
			for _, c := range line {
				input.Instructions = append(input.Instructions, c)
			}
			continue
		}
		if len(line) == 0 {
			continue
		}

		matches := rgx.FindAllString(line, -1)
		if len(matches) != 3 {
			log.Fatal("invalid movement !!", line)
		}

		if _, ok := input.Movements[matches[0]]; ok {
			log.Fatal("duplicated movement !!", line)
		}

		input.Movements[matches[0]] = Movement{
			NextLeft:  matches[1],
			NextRight: matches[2],
		}
	}

	return &input
}

func (i *Input) Part1() int {
	var (
		index = 0
		steps = 0
		next  = func(index int) int {
			index++
			if index >= len(i.Instructions) {
				return 0
			}
			return index
		}
		current = Start
	)

	for current != End {
		switch i.Instructions[index] {
		case Left:
			current = i.Movements[current].NextLeft
		case Right:
			current = i.Movements[current].NextRight
		default:
			log.Fatal("invalid instruction !!", i.Instructions[index])
		}
		steps++
		index = next(index)
	}

	return steps
}

func (i *Input) Part2() int {
	panic("not implemented yet")
}
