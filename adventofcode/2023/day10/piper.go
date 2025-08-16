package day10

import (
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
	"log"
	"strings"
)

type movement rune

const (
	NorthSouth movement = '|'
	EastWest   movement = '-'
	NorthEast  movement = 'L'
	NorthWest  movement = 'J'
	SouthEast  movement = 'F'
	SouthWest  movement = '7'
	Ground     movement = '.'
	Start      movement = 'S'
)

var compatibilities = map[movement]map[helper.Direction][]movement{
	NorthSouth: {
		helper.Top:    {SouthEast, SouthWest, NorthSouth},
		helper.Bottom: {NorthEast, NorthWest, NorthSouth},
	},
	EastWest: {
		helper.Left:  {NorthEast, SouthEast, EastWest},
		helper.Right: {NorthWest, SouthWest, EastWest},
	},
	NorthEast: {
		helper.Top:   {NorthSouth, SouthEast, SouthWest},
		helper.Right: {EastWest, SouthWest, NorthWest},
	},
	NorthWest: {
		helper.Left: {EastWest, NorthEast, SouthEast},
		helper.Top:  {NorthSouth, SouthEast, SouthWest},
	},
	SouthEast: {
		helper.Right:  {NorthWest, SouthWest, EastWest},
		helper.Bottom: {NorthWest, NorthEast, NorthSouth},
	},
	SouthWest: {
		helper.Left:   {SouthEast, NorthEast, EastWest},
		helper.Bottom: {NorthSouth, NorthEast, NorthWest},
	},
}

var nextDirection = map[movement]map[helper.Direction]helper.Direction{
	NorthSouth: {
		helper.Bottom: helper.Bottom,
		helper.Top:    helper.Top,
	},
	EastWest: {
		helper.Left:  helper.Left,
		helper.Right: helper.Right,
	},
	NorthEast: {
		helper.Bottom: helper.Right,
		helper.Left:   helper.Top,
	},
	NorthWest: {
		helper.Right:  helper.Top,
		helper.Bottom: helper.Left,
	},
	SouthEast: {
		helper.Top:  helper.Right,
		helper.Left: helper.Bottom,
	},
	SouthWest: {
		helper.Top:   helper.Left,
		helper.Right: helper.Bottom,
	},
}

func ReadMap(data string) ([][]movement, *helper.Location) {
	// read in the map and return the starting location
	var (
		field = make([][]movement, 0)
		st    = &helper.Location{}
	)

	for i, line := range strings.Split(data, "\n") {
		chars := strings.Split(line, "")
		movements := make([]movement, 0)
		for j, ch := range chars {
			mv := movement(ch[0])
			if mv == Start {
				st = helper.NewLocation(i, j)
			}
			movements = append(movements, movement(ch[0]))
		}
		field = append(field, movements)
	}

	return field, st
}

type directionLocation struct {
	direction helper.Direction
	location  *helper.Location
}

func part1(field [][]movement, d1 directionLocation, d2 directionLocation) int {
	if d1.location.Compare(d2.location) == 0 {
		return 1
	}

	d1x, d1y := d1.location.Get()
	d1nd, ok := nextDirection[field[d1x][d1y]][d1.direction]
	if !ok {
		log.Println(d1, d2)
		panic("could not find next direction d1 issue")
	}
	d1nl := d1.location.Move(d1nd)

	d2x, d2y := d2.location.Get()
	d2nd, ok := nextDirection[field[d2x][d2y]][d2.direction]
	if !ok {
		log.Println(d1, d2)
		panic("could not find next direction d2 issue")
	}
	d2nl := d2.location.Move(d2nd)

	return 1 + part1(field, directionLocation{
		direction: d1nd,
		location:  d1nl,
	}, directionLocation{
		direction: d2nd,
		location:  d2nl,
	})
}

func Part1(field [][]movement, start *helper.Location) int {
	validMoves := make([]directionLocation, 0)
	for _, dir := range helper.PathDirections {
		nl := start.Move(dir)
		// check if valid
		i, j := nl.Get()
		if i < 0 || j < 0 || i >= len(field) || j >= len(field[i]) {
			continue
		}
		mv := field[i][j]
		if _, ok := nextDirection[mv][dir]; ok {
			validMoves = append(validMoves, directionLocation{
				direction: dir,
				location:  nl,
			})
		}
	}

	if len(validMoves) != 2 {
		panic("invalid number of movement !!")
	}

	return part1(field, validMoves[0], validMoves[1])
}

func part2(field [][]movement, d1 directionLocation, d2 directionLocation) int {
	if d1.location.Compare(d2.location) == 0 {
		return 1
	}

	d1x, d1y := d1.location.Get()
	d1nd, ok := nextDirection[field[d1x][d1y]][d1.direction]
	if !ok {
		log.Println(d1, d2)
		panic("could not find next direction d1 issue")
	}
	d1nl := d1.location.Move(d1nd)
	field[d1x][d1y] = movement(helper.DirectionToChar[d1.direction])

	return 1 + part2(field, directionLocation{
		direction: d1nd,
		location:  d1nl,
	}, d2)
}

func Part2(field [][]movement, start *helper.Location) int {
	validMoves := make([]directionLocation, 0)
	for _, dir := range helper.PathDirections {
		nl := start.Move(dir)
		// check if valid
		i, j := nl.Get()
		if i < 0 || j < 0 || i >= len(field) || j >= len(field[i]) {
			continue
		}
		mv := field[i][j]
		if _, ok := nextDirection[mv][dir]; ok {
			validMoves = append(validMoves, directionLocation{
				direction: dir,
				location:  nl,
			})
		}
	}

	if len(validMoves) != 2 {
		panic("invalid number of movement !!")
	}

	return part2(field, validMoves[0], validMoves[1])
}

func printField(field [][]movement) {
	for _, row := range field {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}
