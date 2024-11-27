package leet

import (
	"container/list"
)

// Helper function to perform BFS and find the shortest path
func bfs(n int, graph map[int][]int) int {
	queue := list.New()
	visited := make([]bool, n)
	queue.PushBack(0)
	visited[0] = true
	distance := 0

	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(int)
			if node == n-1 {
				return distance
			}
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue.PushBack(neighbor)
				}
			}
		}
		distance++
	}
	return -1 // If no path is found
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}
	graph[n-1] = []int{}

	results := []int{}
	for _, query := range queries {
		u, v := query[0], query[1]
		graph[u] = append(graph[u], v)
		results = append(results, bfs(n, graph))
	}
	return results
}
