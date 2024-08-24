package kata

import "testing"

func TestSpiralize(t *testing.T) {
	maze := Spiralize(6)
	for i := 0; i < len(maze); i++ {
		t.Log(maze[i])
	}

	maze = Spiralize(8)
	for i := 0; i < len(maze); i++ {
		t.Log(maze[i])
	}

	maze = Spiralize(9)
	for i := 0; i < len(maze); i++ {
		t.Log(maze[i])
	}
}
