package kata

import (
	"fmt"
	"strings"
)

func slidingPuzzle(board [][]int) int {
	// Flatten the board into a string
	start := ""
	for _, row := range board {
		for _, num := range row {
			start += fmt.Sprintf("%d", num)
		}
	}
	goal := "123450"

	// Define the neighbors for each index in the 2x3 grid
	neighbors := [][]int{
		{1, 3},    // 0 can swap with 1 or 3
		{0, 2, 4}, // 1 can swap with 0, 2, or 4
		{1, 5},    // 2 can swap with 1 or 5
		{0, 4},    // 3 can swap with 0 or 4
		{1, 3, 5}, // 4 can swap with 1, 3, or 5
		{2, 4},    // 5 can swap with 2 or 4
	}

	// BFS queue: (state string, zero index, moves count)
	queue := []struct {
		state   string
		zeroIdx int
		moveCnt int
	}{{start, strings.Index(start, "0"), 0}}

	// Visited states set
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		// Pop the front element from the queue
		cur := queue[0]
		queue = queue[1:]

		// Check if we've reached the goal
		if cur.state == goal {
			return cur.moveCnt
		}

		// Try all possible swaps for the current zero index
		for _, neighbor := range neighbors[cur.zeroIdx] {
			// Swap the zero with the neighbor to generate the new state
			newState := []rune(cur.state)
			newState[cur.zeroIdx], newState[neighbor] = newState[neighbor], newState[cur.zeroIdx]
			nextState := string(newState)

			// Add the new state to the queue if not visited
			if !visited[nextState] {
				visited[nextState] = true
				queue = append(queue, struct {
					state   string
					zeroIdx int
					moveCnt int
				}{nextState, neighbor, cur.moveCnt + 1})
			}
		}
	}

	// If we exhaust the queue without finding the goal
	return -1
}
