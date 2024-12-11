package day06

import (
	"aoc2024/common"
	"fmt"
	"sync"
)

func Solve() {

	
	input, err := common.ReadInput(6)
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

var directions = []Coordinate{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}


func isValidPosition(grid [][]rune, position Coordinate) bool {
	return position.y >= 0 && position.y < len(grid) && position.x >= 0 && position.x < len(grid[0])
}

func getNextDirection(curr Coordinate) Coordinate {
	for i, dir := range directions {
		if dir == curr {
			return directions[(i+1)%len(directions)]
		}
	}
	return curr
}

func countTiles(grid [][]rune, start Coordinate) int {
	currDirection := directions[3]
	curr := start

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	tilesVisited := 0

	for {
		if isValidPosition(grid, curr) && !visited[curr.y][curr.x] {
			visited[curr.y][curr.x] = true
			tilesVisited++
		}

		next := Coordinate{curr.y + currDirection.y, curr.x + currDirection.x}

		if !isValidPosition(grid, next) {
			break
		}
		if grid[next.y][next.x] == '#' {
			currDirection = getNextDirection(currDirection)
		} else {
			curr = next
		}
	}

	return tilesVisited
}

func solvePart1(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '^' {
				return countTiles(grid, Coordinate{i, j})
			}
		}
	}
	return -1
}

func solvePart2(grid [][]rune) int {
    start := findStart(grid)
    path := generatePath(grid, start)

    loopCount := 0
    results := make(chan bool, len(path))
    var wg sync.WaitGroup

    for _, candidate := range path {
        wg.Add(1)
        go func(candidate Coordinate) {
            defer wg.Done()
            gridCopy := copyGrid(grid)
            gridCopy[candidate.y][candidate.x] = '#'
            results <- createsLoop(gridCopy)
        }(candidate)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for result := range results {
        if result {
            loopCount++
        }
    }

    return loopCount
}

func findStart(grid [][]rune) Coordinate {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				return Coordinate{i, j}
			}
		}
	}
	return Coordinate{-1, -1}
}

func generatePath(grid [][]rune, start Coordinate) []Coordinate {
	currDirection := directions[3]
	curr := start

	rows := len(grid)
	cols := len(grid[0])
	visited := make([]bool, rows*cols)
	var path []Coordinate

	getIndex := func(coord Coordinate) int {
		return coord.y*cols + coord.x
	}

	for {
		index := getIndex(curr)
		if !visited[index] {
			visited[index] = true
			path = append(path, curr)
		}

		next := Coordinate{curr.y + currDirection.y, curr.x + currDirection.x}

		if !isValidPosition(grid, next) {
			break
		}

		if grid[next.y][next.x] == '#' {
			currDirection = getNextDirection(currDirection)
		} else {
			curr = next
		}

		if curr == start && currDirection == directions[3] {
			break
		}
	}

	return path
}

func createsLoop(grid [][]rune) bool {
	start := findStart(grid)
	curr := start
	currDirection := directions[3]

	rows := len(grid)
	cols := len(grid[0])
	visitCount := make([]int, rows*cols)

	getIndex := func(coord Coordinate) int {
		return coord.y*cols + coord.x
	}

	for {
		if !isValidPosition(grid, curr) {
			return false
		}

		index := getIndex(curr)

		visitCount[index]++
		if visitCount[index] > 4 {
			return true
		}

		next := Coordinate{curr.y + currDirection.y, curr.x + currDirection.x}

		if !isValidPosition(grid, next) {
			return false
		}

		if grid[next.y][next.x] == '#' {
			currDirection = getNextDirection(currDirection)
		} else {
			curr = next
		}
	}
}


func copyGrid(grid [][]rune) [][]rune {
	copy := make([][]rune, len(grid))
	for i := range grid {
		copy[i] = append([]rune{}, grid[i]...)
	}
	return copy
}
