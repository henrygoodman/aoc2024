package day19

import (
	"aoc2024/common"
	"fmt"
	"strings"
)

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

// Backtracking algorithm, try building with each valid towel
// backtrack until we reach the length of the design
// What makes a valid towel? design starts with towel

func solve(design string, towels []string, returnCount bool) int {
	cache := make(map[string]int)

	var backtrack func(design string) int
	backtrack = func(design string) int {
		if val, exists := cache[design]; exists {
			return val
		}

		if len(design) == 0 {
			return 1
		}

		if !returnCount {
			// Part1: only check if possible
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


func solvePart1(input []string) int {
	towels := strings.Split(input[0], ",")
	for i := range towels {
		towels[i] = strings.TrimSpace(towels[i])
	}

	designs := input[2:]
	sum := 0
	for _, d := range designs {
		sum += solve(d, towels, false)
	}
	return sum
}

func solvePart2(input []string) int {
	towels := strings.Split(input[0], ",")
	for i := range towels {
		towels[i] = strings.TrimSpace(towels[i])
	}

	designs := input[2:]
	sum := 0
	for _, d := range designs {
		sum += solve(d, towels, true)
	}
	return sum
}



