package day11

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

var stoneIndex sync.Map // Concurrent-safe map
var globalIndex int32 // Global counter for stone indices

func getStoneIndex(stone string) int {
	// Try to load the index from the concurrent map
	idx, exists := stoneIndex.Load(stone)
	if exists {
		return idx.(int)
	}

	// Generate a new index atomically
	newIdx := int(atomic.AddInt32(&globalIndex, 1)) - 1
	stoneIndex.Store(stone, newIdx)
	return newIdx
}

func processStone(stone string, iterations int, cache []int32, inProgress []int32, maxIterations int) int {
	if iterations == 0 {
		return 1
	}

	idx := getStoneIndex(stone)
	flatIndex := idx*(maxIterations+1) + iterations

	// Check cache
	for {
		val := atomic.LoadInt32(&cache[flatIndex])
		if val != -1 {
			return int(val) // Value is already cached
		}

		// Try to mark "in progress"
		if atomic.CompareAndSwapInt32(&inProgress[flatIndex], 0, 1) {
			// This thread will compute the value
			break
		}

		// Wait for the result if another thread is computing
		for atomic.LoadInt32(&inProgress[flatIndex]) == 1 {
			// Spin-wait; could add backoff to reduce CPU usage
		}
	}

	// Compute the value
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
		intStone, _ := strconv.Atoi(stone)
		nextStones = append(nextStones, strconv.Itoa(intStone*2024))
	}

	totalCount := 0
	for _, next := range nextStones {
		totalCount += processStone(next, iterations-1, cache, inProgress, maxIterations)
	}

	// Store the result and mark "not in progress"
	atomic.StoreInt32(&cache[flatIndex], int32(totalCount))
	atomic.StoreInt32(&inProgress[flatIndex], 0)

	return totalCount
}


func solve(input []string, maxIterations int) int {
	if len(input) != 1 {
		fmt.Println("Input must be 1 line")
		return -1
	}

	stones := strings.Fields(input[0])

	// Dynamically estimate cache size (10x buffer for safety)
	cacheSize := len(stones) * 1000 * (maxIterations + 1)
	cache := make([]int32, cacheSize)      // Use int32 for atomic operations
	inProgress := make([]int32, cacheSize) // Track computation in progress

	for i := range cache {
		cache[i] = -1
	}

	var wg sync.WaitGroup
	var total int64 // Atomic total summation

	for _, stone := range stones {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			count := processStone(s, maxIterations, cache, inProgress, maxIterations)
			atomic.AddInt64(&total, int64(count))
		}(stone)
	}

	wg.Wait()
	return int(total)
}


func solvePart1(input []string) int {
	return solve(input, 25)
}

func solvePart2(input []string) int {
	return solve(input, 75)
}
