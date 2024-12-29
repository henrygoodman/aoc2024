package day20

import (
	"aoc2024/common"
	"fmt"
)

func Solve() {
	input, err := common.ReadInput(20)
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

func findStartAndEnd(grid [][]rune) (Coordinate, Coordinate) {
	var start, end Coordinate
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = Coordinate{i, j}
			}
			if cell == 'E' {
				end = Coordinate{i, j}
			}
		}
	}
	return start, end
}

// Parallel, we can solve by removing 1 wall and then finding the path (see if it saves time)
// - if we find the path originally, we only have to consider walls already on the path
// - we only have to consider walls where BOTH sides are on the path (touching path)
// Do we even have to path find?
// - can we just identify positions in the grid (3x3 block centred on a path piece) where we could
// 'remove' a wall piece and then reach a position that is now close on the path
// i.e. to find shortcuts of < 100, we need to go from step X in path to step (X+100) in path in the 3x3 area
// so maybe this can be done in 1 pass

type Coordinate struct {
	y, x int
}

func findPath(grid [][]rune, start, end Coordinate) []Coordinate {
	height := len(grid)
	width := len(grid[0])
	directions := []Coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	stack := []struct {
		coord Coordinate
		path  []Coordinate
	}{{start, []Coordinate{start}}}

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	visited[start.y][start.x] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.coord == end {
			return current.path
		}

		for _, d := range directions {
			neighbor := Coordinate{current.coord.y + d.y, current.coord.x + d.x}

			if neighbor.y < 0 || neighbor.y >= height || neighbor.x < 0 || neighbor.x >= width {
				continue
			}
			if grid[neighbor.y][neighbor.x] == '#' || visited[neighbor.y][neighbor.x] {
				continue
			}

			visited[neighbor.y][neighbor.x] = true
			stack = append(stack, struct {
				coord Coordinate
				path  []Coordinate
			}{neighbor, append(current.path, neighbor)})
		}
	}

	return nil
}


func buildPathMap(path []Coordinate) map[Coordinate]int {
	pathMap := make(map[Coordinate]int)
	for i, coord := range path {
		pathMap[coord] = i
	}
	return pathMap
}


func findShortcut(path []Coordinate, index, shortcutSize int, pathMap map[Coordinate]int) int {
	directions := []Coordinate{{0, 2}, {2, 0}, {0, -2}, {-2, 0}}
	sum := 0
	for _, d := range directions {
		searchPosition := Coordinate{path[index].y + d.y, path[index].x + d.x}
		if targetIndex, exists := pathMap[searchPosition]; exists {
			timeSave := targetIndex - index - 2
			if timeSave >= shortcutSize {
				sum++
			}
		}
	}
	return sum
}

func solvePart1(grid [][]rune) int {
	start, end := findStartAndEnd(grid)
	path := findPath(grid, start, end)
	pathMap := buildPathMap(path)

	minShortcut := 100
	sum := 0
	for i := range path[:len(path) - minShortcut] {
		sum += findShortcut(path, i, minShortcut, pathMap)
	}
	return sum
}

// Since we are allowed to leave and re-enter wall, its very easy to simply use Manhatten distance. Would need BFS if only
// allowed to enter wall once.

func solvePart2(grid [][]rune) int {
	start, end := findStartAndEnd(grid)
	path := findPath(grid, start, end)
	minShortcut := 100

	distinctShortcuts := make(map[Coordinate]map[Coordinate]bool)

	for i, startPoint := range path {
		for j := i + 1; j < len(path); j++ {
			endPoint := path[j]

			manhattanDistance := abs(startPoint.y-endPoint.y) + abs(startPoint.x-endPoint.x)

			if manhattanDistance <= 20 {
				timeSave := j - i - manhattanDistance

				if timeSave >= minShortcut {
					if _, exists := distinctShortcuts[startPoint]; !exists {
						distinctShortcuts[startPoint] = make(map[Coordinate]bool)
					}
					distinctShortcuts[startPoint][endPoint] = true
				}
			}
		}
	}

	count := 0
	for _, ends := range distinctShortcuts {
		count += len(ends)
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
