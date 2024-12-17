package day11

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
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

var stoneIndex = make(map[string]int)
var nextIndex int
var stoneIndexMutex sync.Mutex

func processStone(stone string, iterations int, cache []int, maxIterations int, cacheMutex *sync.Mutex) int {
	if iterations == 0 {
		return 1
	}

	stoneIndexMutex.Lock()
	idx, exists := stoneIndex[stone]
	if !exists {
		idx = nextIndex
		stoneIndex[stone] = idx
		nextIndex++
	}
	stoneIndexMutex.Unlock()

	flatIndex := idx*(maxIterations+1) + iterations

	cacheMutex.Lock()
	if cache[flatIndex] != -1 {
		result := cache[flatIndex]
		cacheMutex.Unlock()
		return result
	}
	cacheMutex.Unlock()

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
		totalCount += processStone(next, iterations-1, cache, maxIterations, cacheMutex)
	}

	cacheMutex.Lock()
	cache[flatIndex] = totalCount
	cacheMutex.Unlock()

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

	var wg sync.WaitGroup
	var mu sync.Mutex
	var cacheMutex sync.Mutex
	total := 0

	for _, stone := range stones {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			count := processStone(s, 25, cache, 25, &cacheMutex)
			mu.Lock()
			total += count
			mu.Unlock()
		}(stone)
	}

	wg.Wait()
	return total
}

func solvePart2(input []string) int {
	if len(input) != 1 {
		fmt.Println("Input must be 1 line")
		return -1
	}

	stones := strings.Fields(input[0])
	cacheSize := 5000 * 76
	cache := make([]int, cacheSize)
	for i := range cache {
		cache[i] = -1
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var cacheMutex sync.Mutex
	total := 0

	for _, stone := range stones {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			count := processStone(s, 75, cache, 75, &cacheMutex)
			mu.Lock()
			total += count
			mu.Unlock()
		}(stone)
	}

	wg.Wait()
	return total
}
