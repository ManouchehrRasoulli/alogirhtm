package day14

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2024/helper"
	"testing"
)

func TestInputTestPart1(t *testing.T) {
	data, err := helper.ReadAll("input.test1")
	if err != nil {
		t.Fatal(err)
	}

	bots := ParseBots(data)
	var (
		sec   = 5
		x     = 7
		y     = 11
		field = NewField(x, y)
	)

	for _, bot := range bots {
		t.Log(bot)
	}

	t.Log("started ---- with")
	field.DrawBots(bots)
	t.Log(field.String())

	for i := 0; i < sec; i++ {
		for _, b := range bots {
			b.MoveInGrid(x, y)

			field.DrawBots(bots)

			t.Log(i, field.String())
		}
	}
}

func TestInputTestPart2(t *testing.T) {
	data, err := helper.ReadAll("input.test2")
	if err != nil {
		t.Fatal(err)
	}

	bots := ParseBots(data)
	var (
		sec          = 100
		x            = 7
		y            = 11
		safetyFactor = 0
	)

	for i := 0; i < sec; i++ {
		for _, bot := range bots {
			bot.MoveInGrid(x, y)
		}
	}

	tl, tr, bl, br := CountQuadrants(bots, x, y)
	safetyFactor = tl * tr * bl * br
	t.Log(safetyFactor, " -- [", tl, tr, bl, br, "]") // 12 -- [1 3 4 1]
}

func TestInputPuzzlePart1(t *testing.T) {
	data, err := helper.ReadAll("input.puzzle")
	if err != nil {
		t.Fatal(err)
	}

	bots := ParseBots(data)
	var (
		sec          = 100
		x            = 103
		y            = 101
		safetyFactor = 0
	)

	for i := 0; i < sec; i++ {
		for _, bot := range bots {
			bot.MoveInGrid(x, y)
		}
	}

	tl, tr, bl, br := CountQuadrants(bots, x, y)
	safetyFactor = tl * tr * bl * br
	t.Log(safetyFactor, " -- [", tl, tr, bl, br, "]") // 12 -- [1 3 4 1]
}
