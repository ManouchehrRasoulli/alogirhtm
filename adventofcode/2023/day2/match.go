package day2

import (
	"log"
	"regexp"
	"strconv"
)

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Id     int
	Rounds []Round
	Red    int
	Blue   int
	Green  int
}

func ReadGame(line string) *Game {
	var (
		re = regexp.MustCompile(`(red|green|blue|\d|;)`)
		g  = Game{
			Id:     0,
			Rounds: make([]Round, 0),
			Red:    0,
			Blue:   0,
			Green:  0,
		}
		round Round
		err   error
	)

	items := re.FindAllString(line, -1)
	log.Println(items)
	g.Id, err = strconv.Atoi(items[0])
	if err != nil {
		log.Fatal(line, err, "game-id")
	}

	for i := 1; i < len(items); {
		if items[i] == ";" {
			i++
			g.Rounds = append(g.Rounds, round)
			round = Round{}
			continue
		}

		var (
			count int
			color string
		)

		count, err = strconv.Atoi(items[i])
		if err != nil {
			log.Fatal(line, err, "count", i)
		}
		color = items[i+1]
		switch color {
		case "red":
			round.Red = count
			g.Red = max(g.Red, count)
		case "blue":
			round.Blue = count
			g.Blue = max(g.Blue, count)
		case "green":
			round.Green = count
			g.Green = max(g.Green, count)
		default:
			log.Fatal(line, "invalid color !", color)
		}

		i += 2
	}

	return &g
}
