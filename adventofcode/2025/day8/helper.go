package day8

import (
	"fmt"
	"log"
	"math"
	"sort"
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
				math.Pow(b.Z-a.Z, 2))*10000) / 10000
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

func (dc *distanceCollection) insert(a, b Location) {
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
}

func (dc *distanceCollection) next() *distanceNode {
	if dc.head == nil {
		return nil
	}

	c := dc.head
	dc.head = c.next
	return c
}

type cluster struct {
	clusters []map[Location]struct{}
}

func newCluster() *cluster {
	return &cluster{
		clusters: make([]map[Location]struct{}, 0),
	}
}

func (c *cluster) count() int {
	return len(c.clusters)
}

func (c *cluster) insert(a Location) {
	for _, cl := range c.clusters {
		_, hasA := cl[a]
		if hasA {
			return
		}
	}

	c.clusters = append(c.clusters, map[Location]struct{}{a: {}})
}

func (c *cluster) merge(a, b Location) {
	var foundIdx []int
	for i, cl := range c.clusters {
		_, hasA := cl[a]
		_, hasB := cl[b]
		if hasA && hasB {
			return
		}
		if hasA || hasB {
			foundIdx = append(foundIdx, i)
		}
	}

	if len(foundIdx) == 0 {
		cl := map[Location]struct{}{a: {}, b: {}}
		c.clusters = append(c.clusters, cl)
		return
	}

	// merge all relevant clusters plus a, b
	base := c.clusters[foundIdx[0]]
	base[a] = struct{}{}
	base[b] = struct{}{}

	for k := 1; k < len(foundIdx); k++ {
		idx := foundIdx[k]
		for loc := range c.clusters[idx] {
			base[loc] = struct{}{}
		}
		c.clusters[idx] = nil
	}

	res := c.clusters[:0]
	for _, cl := range c.clusters {
		if cl != nil {
			res = append(res, cl)
		}
	}
	c.clusters = res
}

func (c *cluster) topProduct(x int) int {
	if x <= 0 || len(c.clusters) == 0 {
		return 0
	}

	sizes := make([]int, 0, len(c.clusters))
	for _, cl := range c.clusters {
		sizes = append(sizes, len(cl))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	limit := x
	if len(sizes) < x {
		limit = len(sizes)
	}

	prod := 1
	for i := 0; i < limit; i++ {
		prod *= sizes[i]
	}
	return prod
}
