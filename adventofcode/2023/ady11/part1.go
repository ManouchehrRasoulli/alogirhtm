package ady11

import (
	"github.com/ManouchehrRasoulli/alogirhtm/adventofcode/2023/helper"
)

func Part1(universe [][]rune) int {
	var (
		galaxies = make(map[*helper.Location]struct{})
	)

	for i, row := range universe {
		for j, col := range row {
			if col == galaxy {
				galaxies[helper.NewLocation(i, j)] = struct{}{}
			}
		}
	}

	var (
		total = 0
		sc    = make(map[helper.Location]struct{})
	)
	for glx := range galaxies {
		distance := bfs(universe, glx)
		for k, d := range distance {
			if _, ok := sc[k]; !ok {
				total += d
			}
		}
		sc[*glx] = struct{}{}
	}

	return total
}

func bfs(universe [][]rune, start *helper.Location) map[helper.Location]int {

	distance := make(map[helper.Location]int)
	visited := make(map[helper.Location]struct{})
	queue := []QueueItem{{loc: start, distance: 0}}

	for len(queue) > 0 {
		// Pop front
		item := queue[0]
		queue = queue[1:]

		x, y := item.loc.Get()
		if x < 0 || y < 0 || x >= len(universe) || y >= len(universe[0]) {
			continue
		}
		if _, ok := visited[*item.loc]; ok {
			continue
		}
		visited[*item.loc] = struct{}{}

		if universe[x][y] == galaxy {
			if _, ok := distance[*item.loc]; !ok {
				distance[*item.loc] = item.distance
			}
		}

		for _, dir := range helper.PathDirections {
			next := item.loc.Move(dir)
			if _, seen := visited[*next]; !seen {
				queue = append(queue, QueueItem{loc: next, distance: item.distance + 1})
			}
		}
	}

	return distance
}

type QueueItem struct {
	loc      *helper.Location
	distance int
}
