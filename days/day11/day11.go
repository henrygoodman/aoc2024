package day11

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(11)
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


var stoneIndex = make(map[string]int)
var nextIndex int

/* Observations:
- Can we define a function f(x, i) that returns length of input after i iterations 
- 0 -> 1 -> 2024 -> 20, 24 -> 2,0,2,4
- If we multiply by 2024, its not enough to know the length, as we would need to process after?

- We see a lot of powers of 2, why?
- 0 -> 1 -> 2024 eventually breaks down to 2,0,2,4 (all powers of 2)

Do we need an iterative approach? Or can we cache by storing the number of stones for X after Y operations?
e.g. f(125, 1) = 1
f(125, 2) = 2

This requires we 'know' which numbers come from what number?
Use cache to memoize stone at particular iteration

Optimizations:
- Array for cache instead of hashmap overhead

*/
func processStone(stone string, iterations int, cache []int, maxIterations int) int {
	if iterations == 0 {
		return 1
	}

	idx, exists := stoneIndex[stone]
	if !exists {
		idx = nextIndex
		stoneIndex[stone] = idx
		nextIndex++
	}

	flatIndex := idx*(maxIterations+1) + iterations
	if cache[flatIndex] != -1 {
		return cache[flatIndex]
	}

	var nextStones []string
	if stone == "0" {
		nextStones = append(nextStones, "1")
	} else if len(stone)%2 == 0 {
		sl := len(stone)
		l, r := stone[:sl/2], strings.TrimLeft(stone[sl/2:], "0")
		if len(r) == 0 {
			r = "0"
		}
		nextStones = append(nextStones, l, r)
	} else {
		intStone, err := strconv.Atoi(stone)
		if err != nil {
			fmt.Println("Error converting stone to int:", err)
			return 0
		}
		newStone := strconv.Itoa(intStone * 2024)
		nextStones = append(nextStones, newStone)
	}

	totalCount := 0
	for _, next := range nextStones {
		totalCount += processStone(next, iterations-1, cache, maxIterations)
	}

	cache[flatIndex] = totalCount
	return totalCount
}


func solvePart1(input []string) int {
	if len(input) != 1 {
		fmt.Println("Input must be 1 line")
		return -1
	}

	stones := strings.Fields(input[0])
	cacheSize := 5000 * 26
	cache := make([]int, cacheSize)
	for i := range cache {
		cache[i] = -1
	}

	total := 0
	for _, stone := range stones {
		total += processStone(stone, 25, cache, 25)
	}

	return total
}

func solvePart2(input []string) int {
	if len(input) != 1 {
		fmt.Println("Input must be 1 line")
		return -1
	}

	stones := strings.Fields(input[0])
	cacheSize := 5000 * 75
	cache := make([]int, cacheSize)
	for i := range cache {
		cache[i] = -1
	}

	total := 0
	for _, stone := range stones {
		total += processStone(stone, 75, cache, 75)
	}

	return total
}
