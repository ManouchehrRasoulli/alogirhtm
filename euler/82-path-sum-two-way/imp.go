package _2_path_sum_two_way

import (
	_ "embed"
	"fmt"
	"github.com/ManouchehrRasoulli/alogirhtm/euler/helper"
	"math"
	"strconv"
	"strings"
)

type bfs struct {
	locations []costPoint
	visited   map[string]struct{} // will not be used
}

func newBfs() *bfs {
	return &bfs{}
}

func (bfs *bfs) minLeftToRight(matrix [][]int) int {
	minCost := math.MaxInt
	for i := 0; i < len(matrix); i++ {
		iCost := bfs.search(i, 0, matrix)

		if iCost < minCost {
			minCost = iCost
		}
	}

	return minCost
}

func (bfs *bfs) search(row, col int, matrix [][]int) int {
	bfs.locations = make([]costPoint, 0)
	bfs.visited = make(map[string]struct{})

	bfs.locations = append(bfs.locations, costPoint{
		cost: matrix[row][col],
		points: point{
			row: row,
			col: col,
		},
	})

	for len(bfs.locations) != 0 {
		loc := bfs.nextPoint()
		bfs.visited[loc.points.String()] = struct{}{}
		ngs := bfs.neighbors(matrix, loc.points)

		if len(ngs) == 0 {
			return loc.cost
		}

		for _, n := range ngs {
			if _, visited := bfs.visited[n.String()]; visited {
				continue
			}

			bfs.locations = append(bfs.locations, costPoint{
				cost: loc.cost + matrix[n.row][n.col],
				points: point{
					row: n.row,
					col: n.col,
				},
			})
		}
	}

	// no path !!!
	return -1
}

func (bfs *bfs) nextPoint() costPoint {
	var (
		minIndex int
		minCost  = math.MaxInt
	)

	for i, location := range bfs.locations {
		if location.cost < minCost {
			minIndex = i
			minCost = location.cost
		}
	}

	var loc costPoint
	bfs.locations, loc = helper.RemoveAndReturnIndex[costPoint](bfs.locations, minIndex)

	return loc
}

func (bfs *bfs) neighbors(matrix [][]int, p point) []point {
	ng := make([]point, 0)

	if p.col == len(matrix)-1 {
		return ng
	}

	if p.row-1 >= 0 {
		ng = append(ng, point{row: p.row - 1, col: p.col})
	}

	if p.row+1 < len(matrix) {
		ng = append(ng, point{row: p.row + 1, col: p.col})
	}

	if p.col+1 < len(matrix) {
		ng = append(ng, point{row: p.row, col: p.col + 1})
	}

	return ng
}

type (
	point struct {
		row, col int
	}

	costPoint struct {
		cost   int
		points point
	}
)

func (c point) String() string {
	return fmt.Sprintf("%d-%d", c.row, c.col)
}

var (
	//go:embed input.txt
	fileContent string

	matrix [][]int
)

func init() {
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		items := strings.Split(line, ",")
		row := make([]int, len(items))
		for i, item := range items {
			v, _ := strconv.ParseInt(item, 10, 64)
			row[i] = int(v)
		}
		matrix = append(matrix, row)
	}
}
