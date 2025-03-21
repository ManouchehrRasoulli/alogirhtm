package day17

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestTest1Part1(t *testing.T) {
	data, err := helper.ReadAll("test1.part1")
	if err != nil {
		t.Fatal(err)
	}

	pc := ReadPc(data)
	t.Logf("%+v\n", pc)
	pc.Execute()
	t.Logf("%+v\n", pc)
	pc.LogOutputs() // 4,6,3,5,6,3,5,2,1,0
}

func TestPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	pc := ReadPc(data)
	t.Logf("%+v\n", pc)
	pc.Execute()
	t.Logf("%+v\n", pc)
	pc.LogOutputs() // 4,6,1,4,2,1,3,1,6
}

func TestTest1Part2(t *testing.T) {
	data, err := helper.ReadAll("test1.part2")
	if err != nil {
		t.Fatal(err)
	}

	pc := ReadPc(data)
	t.Logf("%+v\n", pc)
	t.Log(pc.FindRegisterAValue()) // 117440
	pc.LogOutputs()                // 0,3,5,4,3,0
	t.Logf("%+v\n", pc)
}

func TestPuzzlePart2(t *testing.T) {
	data, err := helper.ReadAll("puzzle.part1")
	if err != nil {
		t.Fatal(err)
	}

	pc := ReadPc(data)
	t.Logf("%+v\n", pc)
	t.Log(pc.FindRegisterAValue()) // 202366627359274
	pc.LogOutputs()                // 2,4,1,1,7,5,4,6,1,4,0,3,5,5,3,0
	t.Logf("%+v\n", pc)
}
