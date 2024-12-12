package day10

import (
	"aoc2024/common"
	"fmt"
)

type Coordinate struct {
	y, x int
}

func Solve() {
	input, err := common.ReadInput(10)
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

func countUniqueEndpointsDFS(grid [][]rune, start Coordinate, pathLength int, endpoints map[Coordinate]struct{}) {
	rows, cols := len(grid), len(grid[0])
	currentValue := grid[start.y][start.x]

	if pathLength == 10 && currentValue == '9' {
		endpoints[start] = struct{}{}
		return
	}

	directions := []Coordinate{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, d := range directions {
		nx, ny := start.x+d.x, start.y+d.y
		if nx >= 0 && nx < cols && ny >= 0 && ny < rows && grid[ny][nx] == currentValue+1 {
			countUniqueEndpointsDFS(grid, Coordinate{ny, nx}, pathLength+1, endpoints)
		}
	}
}

func countPathsDFS(grid [][]rune, start Coordinate, pathLength int) int {
	rows, cols := len(grid), len(grid[0])
	currentValue := grid[start.y][start.x]

	if pathLength == 10 && currentValue == '9' {
		return 1
	}

	directions := []Coordinate{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	validPaths := 0

	for _, d := range directions {
		nx, ny := start.x+d.x, start.y+d.y
		if nx >= 0 && nx < cols && ny >= 0 && ny < rows && grid[ny][nx] == currentValue+1 {
			validPaths += countPathsDFS(grid, Coordinate{ny, nx}, pathLength+1)
		}
	}

	return validPaths
}

func solvePart1(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	totalUniqueEndpoints := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				endpoints := make(map[Coordinate]struct{})
				countUniqueEndpointsDFS(grid, Coordinate{i, j}, 1, endpoints)
				totalUniqueEndpoints += len(endpoints)
			}
		}
	}

	return totalUniqueEndpoints
}

func solvePart2(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	totalUniquePaths := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				totalUniquePaths += countPathsDFS(grid, Coordinate{i, j}, 1)
			}
		}
	}

	return totalUniquePaths
}
