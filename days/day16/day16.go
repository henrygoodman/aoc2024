package day16

import (
	"aoc2024/common"
	"container/heap"
	"math"
)

type State struct {
	x, y      int
	direction int
	cost      int
	path      []int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var directions = [][2]int{
	{0, 1},  // East
	{1, 0},  // South
	{0, -1}, // West
	{-1, 0}, // North
}

func Solve() {
	input, err := common.ReadInput(16)
	if err != nil {
		return
	}

	grid := common.ParseGrid(input)

	common.Time("Part 1", func() {
		println("Part 1 Answer:", solvePart1(grid))
	})

	common.Time("Part 2", func() {
		println("Part 2 Answer:", solvePart2(grid))
	})
}

func solvePart1(grid [][]rune) int {
	start, end := findStartAndEnd(grid)
	return dijkstra(grid, start, end)
}

func dijkstra(grid [][]rune, start, end [2]int) int {
	height, width := len(grid), len(grid[0])
	minCost := make([][][]int, height)
	for i := range minCost {
		minCost[i] = make([][]int, width)
		for j := range minCost[i] {
			minCost[i][j] = make([]int, 4)
			for k := range minCost[i][j] {
				minCost[i][j][k] = math.MaxInt32
			}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{x: start[0], y: start[1], direction: 0, cost: 0})
	minCost[start[0]][start[1]][0] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		if current.x == end[0] && current.y == end[1] {
			return current.cost
		}

		if current.cost > minCost[current.x][current.y][current.direction] {
			continue
		}

		for i, d := range directions {
			newX, newY := current.x+d[0], current.y+d[1]

			if newX < 0 || newX >= height || newY < 0 || newY >= width || grid[newX][newY] == '#' {
				continue
			}

			newCost := current.cost
			if i != current.direction {
				newCost += 1000
			}
			newCost++

			if newCost < minCost[newX][newY][i] {
				minCost[newX][newY][i] = newCost
				heap.Push(pq, State{x: newX, y: newY, direction: i, cost: newCost})
			}
		}
	}

	return -1
}

func solvePart2(grid [][]rune) int {
	start, end := findStartAndEnd(grid)
	return len(findAllBestPathTiles(grid, start, end))
}

// Same as part 1, just keep track of the current path (using a bool map of flattened coords since state is not important)
func findAllBestPathTiles(grid [][]rune, start, end [2]int) map[int]bool {
	height, width := len(grid), len(grid[0])
	minCost := make([][][]int, height)
	for i := range minCost {
		minCost[i] = make([][]int, width)
		for j := range minCost[i] {
			minCost[i][j] = make([]int, 4)
			for k := range minCost[i][j] {
				minCost[i][j][k] = math.MaxInt32
			}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{x: start[0], y: start[1], direction: 0, cost: 0, path: []int{start[0]*width + start[1]}})
	minCost[start[0]][start[1]][0] = 0

	bestTiles := make(map[int]bool)
	minFinalCost := math.MaxInt32

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		if current.cost > minFinalCost {
			continue
		}

		if current.x == end[0] && current.y == end[1] {
			if current.cost < minFinalCost {
				minFinalCost = current.cost
				bestTiles = make(map[int]bool)
			}
			for _, tile := range current.path {
				bestTiles[tile] = true
			}
			continue
		}

		for i, d := range directions {
			newX, newY := current.x+d[0], current.y+d[1]

			if newX < 0 || newX >= height || newY < 0 || newY >= width || grid[newX][newY] == '#' {
				continue
			}

			newCost := current.cost
			if i != current.direction {
				newCost += 1000
			}
			newCost++

			if newCost <= minCost[newX][newY][i] {
				minCost[newX][newY][i] = newCost
				newPath := append([]int(nil), current.path...)
				newPath = append(newPath, newX*width+newY)
				heap.Push(pq, State{x: newX, y: newY, direction: i, cost: newCost, path: newPath})
			}
		}
	}

	return bestTiles
}

func findStartAndEnd(grid [][]rune) ([2]int, [2]int) {
	var start, end [2]int
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				start = [2]int{i, j}
			}
			if cell == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	return start, end
}
