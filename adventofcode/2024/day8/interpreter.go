package day8

import (
	"fmt"
)

type Location struct {
	X, Y int
}

func (l *Location) String() string {
	return fmt.Sprintf("(%d,%d)", l.X, l.Y)
}

type Antenna struct {
	Id       rune
	Location Location
}

func (a *Antenna) String() string {
	return fmt.Sprintf("^%s^(%d, %d)", string(a.Id), a.Location.X, a.Location.Y)
}

type AntiNode struct {
	Antennas string
}

func (a *AntiNode) String() string {
	return fmt.Sprintf("~V~[%s]", string(a.Antennas))
}

type City struct {
	Field     [][]rune
	Antennas  map[rune][]Antenna
	AntiNodes map[Location][]AntiNode
}

func (c *City) CalculateAntiNodes() {
	for _, antennas := range c.Antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				iTjAntiNode := AntiNodeLocation(antennas[i].Location, antennas[j].Location)
				jTiAntiNode := AntiNodeLocation(antennas[j].Location, antennas[i].Location)

				c.AntiNodes[iTjAntiNode] = append(c.AntiNodes[iTjAntiNode], AntiNode{
					Antennas: fmt.Sprintf("%s -> %s", antennas[i].String(), antennas[j].String()),
				})
				c.AntiNodes[jTiAntiNode] = append(c.AntiNodes[jTiAntiNode], AntiNode{
					Antennas: fmt.Sprintf("%s -> %s", antennas[i].String(), antennas[j].String()),
				})
			}
		}
	}
}

func (c *City) AntiNodeCount() int {
	count := 0

	for location, _ := range c.AntiNodes {
		if location.X < 0 || location.Y < 0 || location.X >= len(c.Field) || location.Y >= len(c.Field) {
			continue
		}

		count++
	}

	return count
}

func (c *City) PlaceAntiNotesOnMap() {
	for location, _ := range c.AntiNodes {
		if location.X < 0 || location.Y < 0 || location.X >= len(c.Field) || location.Y >= len(c.Field) {
			continue
		}

		if c.Field[location.X][location.Y] == '.' {
			c.Field[location.X][location.Y] = '#'
		} else {
			c.Field[location.X][location.Y] = '$' // jam collied with another antenna
		}
	}
}

func (c *City) String() string {
	field := "\n"
	for _, row := range c.Field {
		field += string(row)
		field += "\n"
	}

	field += "\n ANTENNAS \n"
	for _, row := range c.Antennas {
		for _, an := range row {
			field += an.String() + " "
		}
		field += "\n"
	}

	field += "\n ANTI-NODES \n"
	for key, row := range c.AntiNodes {
		field += key.String()
		for _, item := range row {
			field += item.String() + " "
		}
		field += "\n"
	}

	return field
}

// AntiNodeLocation
// calculate anti-node for location a
func AntiNodeLocation(a, b Location) Location {
	dv := DirectionVector(b, a)
	return Location{
		X: a.X + dv.X,
		Y: a.Y + dv.Y,
	}
}

func DirectionVector(a, b Location) Location {
	return Location{
		X: b.X - a.X,
		Y: b.Y - a.Y,
	}
}
