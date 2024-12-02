package days

import (
	"aoc2024/common"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day01() {
	input, err := common.ReadInput(1)
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

func solvePart1(input []string) int {
	// Approach:
	// - parse into L/R list
	// - sort each L/R list
	// - zip result and compute abs, sum all diffs

	// Parse
	var left, right []int
	for _, line := range input {
		parts := strings.Fields(line)

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing numbers in line: %q\n", line)
			continue
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	// Sort
	sort.Ints(left)
	sort.Ints(right)

	// Zip and compute diffs
	cumulativeSum := 0
	for i := 0; i < len(left); i++ {
		absDiff := int(math.Abs(float64(left[i] - right[i])))
		cumulativeSum += absDiff
	}

	return cumulativeSum
}

func solvePart2(input []string) int {
	// Approach
	// - parse into L/R list (no need to sort)
	// - add R list to freq counter (hashmap) during parsing
	// - return mapping using L as key * map[L]
	
	freq := make(map[int]int)

	// Parse
	var left []int
	for _, line := range input {
		parts := strings.Fields(line)

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing numbers in line: %q\n", line)
			continue
		}

		left = append(left, leftNum)
		freq[rightNum]++
	}

	result := 0
	for _, l := range left {
		result += l * freq[l]
	}

	return result
}
