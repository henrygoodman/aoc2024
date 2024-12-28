package day18

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(18)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(input))
	})

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(input))
	})
}

func parseGrid(width, height, iterations int, input []string) [][]rune {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for i, pair := range input {
		if i >= iterations {
			break
		}
		parts := strings.Split(pair, ",")

		if len(parts) != 2 {
			fmt.Println("Unexpected input format")
			return nil
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		grid[y][x] = '#'
	}
	return grid
}

type Coordinate struct {
	y, x int
}

func bfs(grid [][]rune, start, end Coordinate) int {
	height := len(grid)
	width := len(grid[0])
	directions := []Coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := []struct {
		coord    Coordinate
		distance int
	}{{start, 0}}

	visited := make(map[Coordinate]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.coord == end {
			return current.distance
		}

		for _, d := range directions {
			neighbor := Coordinate{current.coord.y + d.y, current.coord.x + d.x}

			if neighbor.y < 0 || neighbor.y >= height || neighbor.x < 0 || neighbor.x >= width {
				continue
			}
			if grid[neighbor.y][neighbor.x] == '#' || visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			queue = append(queue, struct {
				coord    Coordinate
				distance int
			}{neighbor, current.distance + 1})
		}
	}

	return -1
}

func solvePart1(input []string) int {
	height := 71
	width := 71
	iterations := 1024

	grid := parseGrid(width, height, iterations, input)

	start := Coordinate{0, 0}
	end := Coordinate{height - 1, width - 1}

	return bfs(grid, start, end)
}

func isPathPossible(grid [][]rune, start, end Coordinate) bool {
	height := len(grid)
	width := len(grid[0])
	directions := []Coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	visited := make(map[Coordinate]bool)
	queue := []Coordinate{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		for _, d := range directions {
			neighbor := Coordinate{current.y + d.y, current.x + d.x}

			if neighbor.y < 0 || neighbor.y >= height || neighbor.x < 0 || neighbor.x >= width {
				continue
			}
			if grid[neighbor.y][neighbor.x] == '#' || visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}

	return false
}


func solvePart2(input []string) int {
	height := 71
	width := 71

	// Binary search with lower bound from p1, upper bound as all obstacles, return lowest where no valid path exists
	start := Coordinate{0, 0}
	end := Coordinate{height - 1, width - 1}
	lowerBound := solvePart1(input)

	upperBound := len(input)
	result := -1

	for lowerBound <= upperBound {
		mid := (lowerBound + upperBound) / 2
		grid := parseGrid(width, height, mid, input)

		if isPathPossible(grid, start, end) {
			lowerBound = mid + 1
		} else {
			result = mid
			upperBound = mid - 1
		}
	}

	if result != -1 {
		fmt.Printf("Input: %s\n", input[result-1])
	}

	return result
}


