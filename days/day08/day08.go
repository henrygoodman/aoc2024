package day08

import (
	"aoc2024/common"
	"fmt"
)

func Solve() {
	input, err := common.ReadInput(8)
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

type Coordinate struct {
	y, x int
}

func isValidPosition(grid [][]rune, position Coordinate) bool {
	return position.y >= 0 && position.y < len(grid) && position.x >= 0 && position.x < len(grid[0])
}

func findNumberOfValidAntinodes(first Coordinate, second Coordinate, grid [][]rune) []Coordinate {
    deltaY := second.y - first.y
    deltaX := second.x - first.x

    antinode1 := Coordinate{first.y - deltaY, first.x - deltaX}
    antinode2 := Coordinate{second.y + deltaY, second.x + deltaX}

    var validAntinodes []Coordinate
    if isValidPosition(grid, antinode1) {
        validAntinodes = append(validAntinodes, antinode1)
    }
    if isValidPosition(grid, antinode2) {
        validAntinodes = append(validAntinodes, antinode2)
    }

    return validAntinodes
}

func solvePart1(grid [][]rune) int {
    m := make(map[rune][]Coordinate)
    rows, cols := len(grid), len(grid[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            char := grid[i][j]
            if char != '.' {
                m[char] = append(m[char], Coordinate{i, j})
            }
        }
    }

    uniquePositions := make(map[Coordinate]struct{})

    for _, positions := range m {
        for _, window := range common.AllPairs(positions) {
            currentAntinodes := findNumberOfValidAntinodes(window[0], window[1], grid)
            for _, antinode := range currentAntinodes {
                uniquePositions[antinode] = struct{}{}
            }
        }
    }

    return len(uniquePositions)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return abs(a)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func findCollinearPoints(first Coordinate, second Coordinate, grid [][]rune) []Coordinate {
    deltaY := second.y - first.y
    deltaX := second.x - first.x

    gcdValue := gcd(deltaY, deltaX)
    stepY := deltaY / gcdValue
    stepX := deltaX / gcdValue

    var points []Coordinate

    // Step in both directions from the first point
    for i := 0; ; i++ {
        y := first.y + i*stepY
        x := first.x + i*stepX
        if !isValidPosition(grid, Coordinate{y, x}) {
            break
        }
        points = append(points, Coordinate{y, x})
    }

    // Step backward from the first point
    for i := -1; ; i-- {
        y := first.y + i*stepY
        x := first.x + i*stepX
        if !isValidPosition(grid, Coordinate{y, x}) {
            break
        }
        points = append(points, Coordinate{y, x})
    }

    return points
}

func solvePart2(grid [][]rune) int {
    m := make(map[rune][]Coordinate)
    rows, cols := len(grid), len(grid[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            char := grid[i][j]
            if char != '.' {
                m[char] = append(m[char], Coordinate{i, j})
            }
        }
    }

    uniquePositions := make(map[Coordinate]struct{})

    for _, positions := range m {
        for _, pair := range common.AllPairs(positions) {
            collinearPoints := findCollinearPoints(pair[0], pair[1], grid)
            for _, point := range collinearPoints {
                uniquePositions[point] = struct{}{}
            }
        }
    }

    return len(uniquePositions)
}