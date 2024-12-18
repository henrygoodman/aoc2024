package day12

import (
	"aoc2024/common"
	"fmt"
)

func Solve() {
	input, err := common.ReadInput(12)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	grid := common.ParseGrid(input)

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(grid))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(grid))
	})
}

func floodFill(grid [][]rune, visited [][]bool, x, y int) (area, perimeter int) {
	rows := len(grid)
	cols := len(grid[0])
	letter := grid[x][y]

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	stack := [][2]int{{x, y}}
	visited[x][y] = true

	for len(stack) > 0 {
		// Pop and process queue
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cx, cy := top[0], top[1]
		
		area++

		for _, d := range directions {
			nx, ny := cx+d[0], cy+d[1]

			// Out of bounds -> increase perimeter
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				perimeter++
				continue
			}

			// Different letter or already visited
			if grid[nx][ny] != letter {
				perimeter++
				continue
			}

			if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return area, perimeter
}

func solvePart1(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	total := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				area, perimeter := floodFill(grid, visited, i, j)
				total += area * perimeter
			}
		}
	}

	return total
}

func solvePart2(grid [][]rune) int {
	return 5
}
