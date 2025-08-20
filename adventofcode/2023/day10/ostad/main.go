package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
)

type Point struct {
	x, y int
}

// pipe connections: from a tile, what directions are valid
var directions = map[rune][]Point{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'L': {{0, -1}, {1, 0}},                  // up, right
	'J': {{0, -1}, {-1, 0}},                 // up, left
	'7': {{0, 1}, {-1, 0}},                  // down, left
	'F': {{0, 1}, {1, 0}},                   // down, right
	'S': {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}, // start may connect any
}

// add two points
func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

// check if direction is valid for a given pipe
func validMove(ch rune, dir Point) bool {
	for _, d := range directions[ch] {
		if d == dir {
			return true
		}
	}
	return false
}

//go:embed input.txt
var file string

func main() {
	// read input

	var grid [][]rune
	for _, line := range strings.Split(file, "\n") {
		grid = append(grid, []rune(line))
	}

	h := len(grid)
	w := len(grid[0])

	// find start S
	var start Point
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 'S' {
				start = Point{x, y}
			}
		}
	}

	// traverse loop
	path := []Point{start}
	visited := map[Point]bool{start: true}

	var prev Point
	cur := start
	foundNext := false
	// find a valid neighbor to start
	for _, d := range directions['S'] {
		n := add(cur, d)
		if n.x < 0 || n.x >= w || n.y < 0 || n.y >= h {
			continue
		}
		ch := grid[n.y][n.x]
		if ch == '.' {
			continue
		}
		// check if neighbor connects back
		back := Point{-d.x, -d.y}
		if validMove(ch, back) {
			prev = cur
			cur = n
			path = append(path, cur)
			visited[cur] = true
			foundNext = true
			break
		}
	}
	if !foundNext {
		log.Fatal("no valid start found")
	}

	// follow until we return to start
	for cur != start {
		ch := grid[cur.y][cur.x]
		moves := directions[ch]
		var next Point
		for _, d := range moves {
			n := add(cur, d)
			if n != prev {
				next = n
				break
			}
		}
		prev, cur = cur, next
		if visited[cur] {
			break
		}
		path = append(path, cur)
		visited[cur] = true
	}

	// Part 1: farthest distance (half the loop length)
	part1 := len(path) / 2
	fmt.Println("Part 1:", part1)

	// Part 2: shoelace formula + Pick's theorem
	area := 0
	for i := 0; i < len(path); i++ {
		j := (i + 1) % len(path)
		area += path[i].x*path[j].y - path[j].x*path[i].y
	}
	area = int(math.Abs(float64(area)) / 2)

	boundary := len(path)
	interior := area - boundary/2 + 1

	fmt.Println("Part 2:", interior)
}
