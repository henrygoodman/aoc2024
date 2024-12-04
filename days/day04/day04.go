package day04

import (
	"aoc2024/common"
	"fmt"
)

type Coordinate struct {
	x, y int
}

func Solve() {
	input, err := common.ReadInput(4)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	grid := parseGrid(input)

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(grid))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(grid))
	})
}

func parseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}

var directions = []Coordinate{
	{0, 1},  {1, 0},  {0, -1},  {-1, 0},
	{1, 1},  {1, -1}, {-1, 1}, {-1, -1},
}

func countXMAS(grid [][]rune, start Coordinate) int {
	target := "XMAS"
	rows, cols := len(grid), len(grid[0])
	count := 0

	// Can short circuit well here since the target is known.
	// - Bounds check
	// - Final character check (if not 'S' then no possible match)
	for _, dir := range directions {
		nx, ny := start.x+(len(target)-1)*dir.x, start.y+(len(target)-1)*dir.y
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] != rune(target[len(target)-1]) {
			continue
		}

		match := true
		for k := 0; k < len(target); k++ {
			tx, ty := start.x+k*dir.x, start.y+k*dir.y
			if grid[tx][ty] != rune(target[k]) {
				match = false
				break
			}
		}

		if match {
			count++
		}
	}

	return count
}


func solvePart1(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	totalCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				totalCount += countXMAS(grid, Coordinate{i, j})
			}
		}
	}

	return totalCount
}


func solvePart2(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	totalCount := 0

	// Inlining this actually is significantly faster than making a new function
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				tl, tr := grid[i-1][j-1], grid[i-1][j+1]
				bl, br := grid[i+1][j-1], grid[i+1][j+1]
				if (tl == 'M' || tl == 'S') &&
					(tr == 'M' || tr == 'S') &&
					(bl == 'M' || bl == 'S') &&
					(br == 'M' || br == 'S') &&
					tl != br &&
					tr != bl {
					totalCount++
				}
			}
		}
	}

	return totalCount
}

