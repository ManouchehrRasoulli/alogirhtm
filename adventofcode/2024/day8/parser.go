package day8

import "strings"

func ParseCityPart1(input string) *City {
	c := City{
		Field:     make([][]rune, 0),
		Antennas:  make(map[rune][]Antenna),
		AntiNodes: make(map[Location][]AntiNode),
	}
	for i, line := range strings.Split(input, "\n") {
		c.Field = append(c.Field, []rune(line))
		c.Field[i] = []rune(line)
		for j, item := range c.Field[i] {
			if item != '.' {
				c.Antennas[item] = append(c.Antennas[item], Antenna{
					Id:       item,
					Location: Location{X: i, Y: j},
				})
			}
		}
	}

	return &c
}
