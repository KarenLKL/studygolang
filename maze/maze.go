package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	value := 0
	if p.i < 0 || p.i >= len(grid) {
		return value, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return value, false
	}

	return grid[p.i][p.j], true
}

// 左上右下，方向
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	fmt.Println()
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
}

func walk(maze [][]int, start, end point) [][]int {
	q := []point{start}
	//初始为大小与maze相同的数组，用于存放走过的路径
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	for len(q) > 0 {
		current := q[0]
		if current == end {
			break
		}
		q = q[1:]
		for _, d := range dirs {
			next := current.add(d)
			//maze的下个点不在边界或者等于一，跳过
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			//steps的下个点不在边界或者已经走过，跳过
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			// 如果等于起点，跳过
			if next == start {
				continue
			}

			currentSteps, _ := current.at(steps)
			steps[next.i][next.j] = currentSteps + 1

			q = append(q, next)
		}
	}
	return steps
}

func readMaze(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		panic(err)
	}
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze

}
