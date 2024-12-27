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
		fmt.Println("Part 1 Answer:", solvePart1(copyGrid(grid), instructions, playerPos))
	})

	part2Grid := transformGrid(copyGrid(grid))
	playerPos = findPlayerPosition(part2Grid)

	common.Time("Part 2", func() {
		fmt.Println("Part 2 Answer:", solvePart2(part2Grid, instructions, playerPos))
	})
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
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
	// (can move if there exists a free space before a wall is encountered)
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

func transformGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for y, row := range grid {
		newRow := make([]rune, len(row)*2)
		for x, cell := range row {
			switch cell {
			case '#':
				newRow[x*2], newRow[x*2+1] = '#', '#'
			case 'O':
				newRow[x*2], newRow[x*2+1] = '[', ']'
			case '.':
				newRow[x*2], newRow[x*2+1] = '.', '.'
			case '@':
				newRow[x*2], newRow[x*2+1] = '@', '.'
			}
		}
		newGrid[y] = newRow
	}
	return newGrid
}

func tryMovePart2(grid [][]rune, playerPos *Coordinate, dy, dx int) {
	height := len(grid)
	width := len(grid[0])
	x, y := playerPos.x, playerPos.y

	nextX, nextY := x+dx, y+dy

	if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height || grid[nextY][nextX] == '#' {
		return
	}

	if grid[nextY][nextX] == '[' {
		if !canPushWideBoxesCone(grid, nextY, nextX, dy, dx) {
			return
		}
		moveWideBoxesCone(grid, nextY, nextX, dy, dx)
	} else if grid[nextY][nextX] == ']' {
		if !canPushWideBoxesCone(grid, nextY, nextX-1, dy, dx) {
			return
		}
		moveWideBoxesCone(grid, nextY, nextX-1, dy, dx)
	}

	grid[y][x] = '.'
	grid[nextY][nextX] = '@'
	playerPos.x, playerPos.y = nextX, nextY
}

func canPushWideBoxesCone(grid [][]rune, startY, startX, dy, dx int) bool {
	width := len(grid[0])

	if dy == 0 {
		y, x := startY, startX
		for x >= 0 && x+1 < width {
			if grid[y][x] == '#' {
				return false // Blocked by wall
			}
			if grid[y][x] == '.' {
				return true // Found free space
			}
			x += dx
		}
		return false
	}

	canPush := canPushVerticalCone(grid, startY, startX, dy)
	return canPush
}

func canPushVerticalCone(grid [][]rune, startY, startX, dy int) bool {
    height := len(grid)
    width := len(grid[0])

    visited := make(map[Coordinate]bool)

    var check func(y, x int) bool
    check = func(y, x int) bool {
        if y < 0 || y >= height || x < 0 || x+1 >= width {
            return false // Out of bounds
        }
        if grid[y][x] == '#' || grid[y][x+1] == '#' {
            return false // Wall blocks movement
        }
        if grid[y][x] != '[' || grid[y][x+1] != ']' {
            return true // Not a box, space is valid
        }

        coord := Coordinate{x: x, y: y}
        if visited[coord] {
            return true
        }
        visited[coord] = true

        // Check the space in the direction of movement
        nextY := y + dy
        if nextY < 0 || nextY >= height {
            return false
        }

        // If the next space contains another box, validate it recursively
        if grid[nextY][x] == '[' || grid[nextY][x] == ']' || grid[nextY][x+1] == '[' || grid[nextY][x+1] == ']' {
            if !check(nextY, x) {
                return false
            }
        } else if !(grid[nextY][x] == '.' && grid[nextY][x+1] == '.') {
            // Space is not free
            return false
        }

        // Recursively validate the cone for all offsets at the nextY level
        for dxOffset := -1; dxOffset <= 1; dxOffset++ {
            nextX := x + dxOffset
            if nextX >= 0 && nextX+1 < width {
                if grid[nextY][x] == '[' || grid[nextY][x] == ']' || grid[nextY][x+1] == '[' || grid[nextY][x+1] == ']' {
                    if grid[nextY][nextX] != '#' && grid[nextY][nextX+1] != '#' && !check(nextY, nextX) {
                        return false
                    }
                }
            }
        }

        return true
    }

    return check(startY, startX)
}

func moveWideBoxesCone(grid [][]rune, startY, startX, dy, dx int) {
	// Horizontal check is simple as part 1, find the first free space then shift everything over
	if dy == 0 {
		width := len(grid[0])
		y, x := startY, startX

		freeX := -1
		for x >= 0 && x < width {
			if grid[y][x] == '.' {
				freeX = x
				break
			}
			x += dx
		}

		if freeX == -1 {
			return
		}

		targetStart := freeX
		playerStart := startX
		direction := dx

		if direction > 0 { // Moving right
			for i := targetStart; i > playerStart; i-- {
				grid[y][i] = grid[y][i-1]
			}
		} else { // Moving left
			for i := targetStart; i < playerStart; i++ {
				grid[y][i] = grid[y][i+1]
			}
		}

	}

	moveVerticalCone(grid, startY, startX, dy, dx)
}

// Recursively move all boxes that are influenced by pushing a source box
func moveVerticalCone(grid [][]rune, startY, startX, dy, dx int) {
	height := len(grid)
	width := len(grid[0])
	visited := make(map[Coordinate]bool)

	positions := []Coordinate{}
	var collectPositions func(y, x int)
	collectPositions = func(y, x int) {
		// Base case: Out of bounds or not a valid box
		if y < 0 || y >= height || x < 0 || x+1 >= width || grid[y][x] != '[' || grid[y][x+1] != ']' {
			return
		}

		coord := Coordinate{x: x, y: y}
		if visited[coord] {
			return
		}
		visited[coord] = true

		positions = append(positions, coord)

		// Collect all positions influenced by the vertical cone
		for dxOffset := -1; dxOffset <= 1; dxOffset++ {
			collectPositions(y+dy, x+dxOffset)
		}
	}

	// Collect all affected positions starting from the initial box
	collectPositions(startY, startX)

	// First pass: Clear all current positions
	for _, pos := range positions {
		grid[pos.y][pos.x] = '.'
		grid[pos.y][pos.x+1] = '.'
	}

	// Second pass: Move all boxes to their new positions
	for _, pos := range positions {
		newY, newX := pos.y+dy, pos.x+dx
		if newY >= 0 && newY < height && newX >= 0 && newX+1 < width {
			grid[newY][newX] = '['
			grid[newY][newX+1] = ']'
		}
	}
}

func updateGridPart2(grid [][]rune, instruction rune, playerPos *Coordinate) [][]rune {
	switch instruction {
	case '>':
		tryMovePart2(grid, playerPos, 0, 1)
	case '<':
		tryMovePart2(grid, playerPos, 0, -1)
	case '^':
		tryMovePart2(grid, playerPos, -1, 0)
	case 'v':
		tryMovePart2(grid, playerPos, 1, 0)
	}
	return grid
}

func solvePart2(grid [][]rune, instructions string, playerPos Coordinate) int {
	for _, instruction := range instructions {
		grid = updateGridPart2(grid, instruction, &playerPos)
	}

	ret := 0
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == '[' && grid[y][x+1] == ']' {
				ret += 100*y + x
			}
		}
	}

	return ret
}

