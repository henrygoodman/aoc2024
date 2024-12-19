package day12

import (
	"aoc2024/common"
	"fmt"
	"sort"
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

func floodFillSides(grid [][]rune, visited [][]bool, x, y int) (area, sides int) {
	rows := len(grid)
	cols := len(grid[0])
	letter := grid[x][y]

	directions := [][2]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	stack := [][2]int{{x, y}}
	visited[x][y] = true

	boundary := [][2]int{}

	for len(stack) > 0 {
		cx, cy := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		area++

		for d := 0; d < len(directions); d++ {
			nx, ny := cx+directions[d][0], cy+directions[d][1]

			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != letter {
				boundary = append(boundary, [2]int{cx, cy})
				continue
			}

			if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	boundary = removeDuplicates(boundary)
	sides = countRegionSidesByScan(boundary)

	return area, sides
}

// We only ever care about the set of boundary points to perform the scanning
func removeDuplicates(boundary [][2]int) [][2]int {
	uniquePoints := make(map[[2]int]struct{})
	result := make([][2]int, 0)

	for _, point := range boundary {
		if _, exists := uniquePoints[point]; !exists {
			uniquePoints[point] = struct{}{}
			result = append(result, point)
		}
	}

	return result
}

func countRegionSidesByScan(boundary [][2]int) int {
	if len(boundary) == 0 {
		return 0
	}

	sides := 0

	// Sort boundary points by rows for vertical scan
	sort.Slice(boundary, func(i, j int) bool {
		if boundary[i][0] == boundary[j][0] {
			return boundary[i][1] < boundary[j][1]
		}
		return boundary[i][0] < boundary[j][0]
	})

	// Perform vertical scan to count distinct spans
	currentRow := boundary[0][0]
	startCol := boundary[0][1]

	for i := 1; i < len(boundary); i++ {
		if boundary[i][0] != currentRow {
			// End of the current row
			if startCol != boundary[i-1][1] {
				sides++ // Add one side for the span
			}
			currentRow = boundary[i][0]
			startCol = boundary[i][1]
		} else if boundary[i][1] != boundary[i-1][1]+1 {
			// End of a contiguous span
			if startCol != boundary[i-1][1] {
				sides++
			}
			startCol = boundary[i][1]
		}
	}
	// Count the last span in the last row
	if startCol != boundary[len(boundary)-1][1] {
		sides++
	}

	// Sort boundary points by columns for horizontal scan
	sort.Slice(boundary, func(i, j int) bool {
		if boundary[i][1] == boundary[j][1] {
			return boundary[i][0] < boundary[j][0]
		}
		return boundary[i][1] < boundary[j][1]
	})

	// Perform horizontal scan to count distinct spans
	currentColumn := boundary[0][1]
	startRow := boundary[0][0]
	for i := 1; i < len(boundary); i++ {
		if boundary[i][1] != currentColumn {
			// End of the current column
			if startRow != boundary[i-1][0] {
				sides++
			}
			currentColumn = boundary[i][1]
			startRow = boundary[i][0]
		} else if boundary[i][0] != boundary[i-1][0]+1 {
			// End of a contiguous span
			if startRow != boundary[i-1][0] {
				sides++
			}
			startRow = boundary[i][0]
		}
	}
	// Count the last span in the last column
	if startRow != boundary[len(boundary)-1][0] {
		sides++
	}


	return sides
}

// This allows us to eliminate the annoyance of '1-wide' cells, that normally can contribute to multiple corners/spans
// In 3x3 repr, corners ONLY ever have 1 set of adjacent edges (i.e. no N/U/O configurations)
func expandGrid(grid [][]rune) [][]rune {
	rows := len(grid)
	cols := len(grid[0])

	expanded := make([][]rune, rows*3)
	for i := range expanded {
		expanded[i] = make([]rune, cols*3)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			letter := grid[i][j]
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					expanded[i*3+x][j*3+y] = letter
				}
			}
		}
	}

	return expanded
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
	expandedGrid := expandGrid(grid)
	rows := len(expandedGrid)
	cols := len(expandedGrid[0])


	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	total := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				area, sides := floodFillSides(expandedGrid, visited, i, j)
				total += area/9 * sides
			}
		}
	}

	return total
}
