package day02

import (
	"aoc2024/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(2)
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
	maxIncrement := 3
	numSafeRows := 0

	for _, line := range input {
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue
		}

		if isSafe(parts, maxIncrement) {
			numSafeRows++
			continue
		}
	}
	return numSafeRows
}

func solvePart2(input []string) int {
	maxIncrement := 3
	numSafeRows := 0

	for _, line := range input {
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue
		}

		if isSafe(parts, maxIncrement) {
			numSafeRows++
			continue
		}

		safe := false
		for i := 0; i < len(parts); i++ {
			adjusted := removeIndex(parts, i)

			if isSafe(adjusted, maxIncrement) {
				safe = true
				break
			}
		}

		if safe {
			numSafeRows++
		}
	}

	return numSafeRows
}

func isSafe(parts []string, maxIncrement int) bool {
	if len(parts) < 2 {
		return true
	}

	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])
	initialSign := common.Sign(first - second)
	if initialSign == 0 || int(math.Abs(float64(first-second))) > maxIncrement {
		return false
	}

	for _, window := range common.Window(parts[1:], 2) {
		a, _ := strconv.Atoi(window[0])
		b, _ := strconv.Atoi(window[1])
		currentSign := common.Sign(a - b)

		if currentSign != initialSign || currentSign == 0 || int(math.Abs(float64(a-b))) > maxIncrement {
			return false
		}
	}

	return true
}

func removeIndex(parts []string, index int) []string {
    return append(parts[:index], parts[index+1:]...)
}
