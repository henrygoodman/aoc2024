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
	pairs := common.Map(input, func(line string) common.Pair {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		return common.Pair{First: leftNum, Second: rightNum}
	})

	left := common.Map(pairs, func(pair common.Pair) int { return pair.First })
	right := common.Map(pairs, func(pair common.Pair) int { return pair.Second })

	sort.Ints(left)
	sort.Ints(right)

	return common.Reduce(common.Zip(left, right), func(acc int, pair common.Pair) int {
		return acc + int(math.Abs(float64(pair.First-pair.Second)))
	}, 0)
}

func solvePart2(input []string) int {
	type Accumulator struct {
		Left []int
		Freq map[int]int
	}

	parsed := common.Reduce(input, func(acc Accumulator, line string) Accumulator {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])
		acc.Left = append(acc.Left, leftNum)
		acc.Freq[rightNum]++
		return acc
	}, Accumulator{
		Left: []int{},
		Freq: make(map[int]int),
	})

	return common.Reduce(parsed.Left, func(acc int, l int) int {
		return acc + l*parsed.Freq[l]
	}, 0)
}
