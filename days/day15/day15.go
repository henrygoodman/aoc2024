package day15

import (
	"aoc2024/common"
	"fmt"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(15)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sections := strings.Split(strings.Join(input, "\n"), "\n\n")
	if len(sections) < 2 {
		fmt.Println("Input format is incorrect. Missing grid or instructions.")
		return
	}

	gridSection := sections[0]
	grid := parseGrid(gridSection)

	instructionSection := sections[1]
	instructions := parseInstructions(instructionSection)

	playerPos := findPlayerPosition(grid)

	common.Time("Part 1", func() {
		fmt.Println("Part 1 Answer:", solvePart1(grid, instructions, playerPos))
	})
	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(grid, instructions, playerPos))
	})
}

type Coordinate struct {
	x, y int
}

func parseGrid(gridSection string) [][]rune {
	lines := strings.Split(strings.TrimSpace(gridSection), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func parseInstructions(instructionSection string) string {
	return strings.ReplaceAll(instructionSection, "\n", "")
}

func findPlayerPosition(grid [][]rune) Coordinate {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				return Coordinate{x: x, y: y}
			}
		}
	}
	return Coordinate{-1, -1}
}

func updateGrid(grid [][]rune, instruction rune, playerPos *Coordinate) [][]rune {
	switch instruction {
	case '>':
		tryMove(grid, playerPos, 0, 1)
	case '<':
		tryMove(grid, playerPos, 0, -1)
	case '^':
		tryMove(grid, playerPos, -1, 0)
	case 'v':
		tryMove(grid, playerPos, 1, 0)
	}
	return grid
}

func tryMove(grid [][]rune, playerPos *Coordinate, dy, dx int) {
	height := len(grid)
	width := len(grid[0])
	x, y := playerPos.x, playerPos.y

	nextX, nextY := x+dx, y+dy
	if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height || grid[nextY][nextX] == '#' {
		// Blocked by wall
		return
	}

	// Check if the move will push all circles
	if grid[nextY][nextX] == 'O' {
		if !canPushCircles(grid, nextY, nextX, dy, dx) {
			return
		}
		moveAllCircles(grid, y, x, dy, dx)
	}

	// Move the player
	grid[y][x] = '.'
	grid[nextY][nextX] = '@'
	playerPos.x, playerPos.y = nextX, nextY
}

func canPushCircles(grid [][]rune, startY, startX, dy, dx int) bool {
	height := len(grid)
	width := len(grid[0])
	y, x := startY, startX

	// Traverse in the direction to check if all circles can move
	for y >= 0 && y < height && x >= 0 && x < width {
		if grid[y][x] == '#' {
			return false
		}
		if grid[y][x] == '.' {
			return true
		}
		y += dy
		x += dx
	}
	return false
}

func moveAllCircles(grid [][]rune, playerY, playerX, dy, dx int) {
	height := len(grid)
	width := len(grid[0])
	y, x := playerY+dy, playerX+dx // Start moving from the position after the player

	// Traverse in the target direction
	for y >= 0 && y < height && x >= 0 && x < width {
		if grid[y][x] == '#' {
			break
		}
		if grid[y][x] == '.' {
			moveCircleChain(grid, y, x, -dy, -dx)
			return
		}
		y += dy
		x += dx
	}
}

// Move circles over into the free space
// (only useful if the current move is pushing a circle)
func moveCircleChain(grid [][]rune, freeY, freeX, dy, dx int) {
	y, x := freeY+dy, freeX+dx

	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0]) {
		if grid[y][x] != 'O' {
			break
		}
		grid[y-dy][x-dx] = 'O' // Move circle to the free space
		grid[y][x] = '.'       // Clear the original position
		y += dy
		x += dx
	}
}

func solvePart1(grid [][]rune, instructions string, playerPos Coordinate) int {
	for _, instruction := range instructions {
		grid = updateGrid(grid, instruction, &playerPos)
	}
	ret := 0

	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == 'O' {
				ret += 100 * y + x
			}
		}
	}
	return ret
}

func solvePart2(grid [][]rune, instructions string, playerPos Coordinate) int {
	return 5
}
