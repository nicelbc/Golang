package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("maze/maze.txt")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}

}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	file.Close()

	return maze
}

type point struct {
	i, j int
}

//方向 上左下右
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//point的加法
func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

//点->数组中的值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//队列
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)
			//越界 或 撞墙
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			//越界 或 已经走过的点
			if !ok || val != 0 {
				continue
			}

			//回到原点
			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)
		}
	}

	return steps
}
