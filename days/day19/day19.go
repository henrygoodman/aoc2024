package day19

import (
	"aoc2024/common"
	"fmt"
	"strings"
	"sync"
)

// Solve is the entry point. Reads input and times the solutions.
func Solve() {
	input, err := common.ReadInput(19)
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

// solve is a backtracking algorithm that tries building the design with valid towels.
// Part 1: returns 1 if possible, 0 if not.
// Part 2: returns the count of all possible ways.
func solve(design string, towels []string, returnCount bool) int {
	cache := make(map[string]int)

	var backtrack func(string) int
	backtrack = func(design string) int {
		if val, exists := cache[design]; exists {
			return val
		}

		if len(design) == 0 {
			return 1
		}

		if !returnCount {
			// Part 1: only check if it is possible
			for _, t := range towels {
				if strings.HasPrefix(design, t) && backtrack(design[len(t):]) == 1 {
					cache[design] = 1
					return 1
				}
			}
			cache[design] = 0
			return 0
		}

		// Part 2: Count all possible ways
		count := 0
		for _, t := range towels {
			if strings.HasPrefix(design, t) {
				count += backtrack(design[len(t):])
			}
		}
		cache[design] = count
		return count
	}

	return backtrack(design)
}

// solvePart1 spawns a goroutine for each design and sums the results.
func solvePart1(input []string) int {
	towels := strings.Split(input[0], ",")
	for i := range towels {
		towels[i] = strings.TrimSpace(towels[i])
	}

	designs := input[2:]
	results := make(chan int, len(designs))
	var wg sync.WaitGroup

	for _, d := range designs {
		d := d // capture loop variable
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- solve(d, towels, false)
		}()
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	return sum
}

// solvePart2 is the same idea, but calls solve(..., true) to count all possible ways.
func solvePart2(input []string) int {
	towels := strings.Split(input[0], ",")
	for i := range towels {
		towels[i] = strings.TrimSpace(towels[i])
	}

	designs := input[2:]
	results := make(chan int, len(designs))
	var wg sync.WaitGroup

	for _, d := range designs {
		d := d
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- solve(d, towels, true)
		}()
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	return sum
}
