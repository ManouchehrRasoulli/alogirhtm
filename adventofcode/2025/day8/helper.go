package day8

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func ParseLocations(input string) []Location {
	var (
		locations []Location
		lines     = strings.Split(input, "\n")
	)

	for _, line := range lines {
		var x, y, z float64
		_, err := fmt.Sscanf(line, "%f,%f,%f", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		locations = append(locations, Location{x, y, z})
	}

	return locations
}

type Location struct {
	X float64
	Y float64
	Z float64
}

func (loc Location) String() string {
	return fmt.Sprintf("%d,%d,%d", int64(loc.X), int64(loc.Y), int64(loc.Z))
}

func distance(a, b Location) float64 {
	// D = √[(x2 - x1)² + (y2 - y1)² +(z2 - z1)²]
	// then round to 4 decimal points
	return math.Round(
		math.Sqrt(
			math.Pow(b.X-a.X, 2)+
				math.Pow(b.Y-a.Y, 2)+
				math.Pow(b.Z-a.Z,
					2))*10000) / 10000
}

type distanceNode struct {
	distance float64
	a        Location
	b        Location
	next     *distanceNode
}

func (d *distanceNode) insert(newNode *distanceNode) {
	if d.distance > newNode.distance {
		log.Println("we shouldn't have a new distance")
		return
	}

	if d.next == nil {
		d.next = newNode
		return
	}

	if d.next.distance > newNode.distance {
		newNode.next = d.next
		d.next = newNode
		return
	}

	d.next.insert(newNode)
}

func (d *distanceNode) String() string {
	if d.next == nil {
		return fmt.Sprintf("{distance:%f, a:%s, b:%s}end", d.distance, d.a.String(), d.b.String())
	}

	return fmt.Sprintf("{distance:%f, a:%s, b:%s} ->\n%s", d.distance, d.a.String(), d.b.String(), d.next.String())
}

type distanceCollection struct {
	head *distanceNode
}

func newDistanceCollection() *distanceCollection {
	return &distanceCollection{
		head: nil,
	}
}

func (dc *distanceCollection) inser(a, b Location) {
	newNode := &distanceNode{
		distance: distance(a, b),
		a:        a,
		b:        b,
		next:     nil,
	}

	if dc.head == nil {
		dc.head = newNode
		return
	}

	if newNode.distance < dc.head.distance {
		newNode.next = dc.head
		dc.head = newNode
		return
	}

	dc.head.insert(newNode)

	p := dc.head
	for i := 0; i < 999; i++ {
		if p.next != nil {
			p = p.next
		}
	}
	p.next = nil
}

func (dc *distanceCollection) next() *distanceNode {
	if dc.head == nil {
		return nil
	}

	c := dc.head
	dc.head = c.next
	return c
}
