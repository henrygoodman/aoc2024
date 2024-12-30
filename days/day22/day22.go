package day22

import (
	"aoc2024/common"
	"fmt"
	"strconv"
)

func Solve() {
	input, err := common.ReadInput(22)
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

func getNextPrice(price int) int {
	res := ((price << 6) ^ price) & 0xFFFFFF
	res = ((res >> 5) ^ res) & 0xFFFFFF
	res = ((res << 11) ^ res) & 0xFFFFFF
	return res
}

func solvePart1(input []string) int {
	res := 0

	for _, price := range input {
		p, _ := strconv.Atoi(price)
		for i := 0; i < 2000; i++ {
			p = getNextPrice(p)
		}
		res += p
	}
	return res
}

// Can solve in 1 pass with a map for each sequence. For each price, store the price the first
// time the sequence is encountered. Then, we can get the max value from the map at the end.

func solvePart2(input []string) int {
	const maxIndex = 19 * 19 * 19 * 19 // Maximum number of possible sequences (-9 to 9)
	sequenceScores := make([]int, maxIndex)
	seenSequences := make(map[int]map[int]bool)
	maxScore := 0

	computeIndex := func(sequence [4]int) int {
		base := 19
		offset := 9 // Offset to handle negative values
		return (sequence[0]+offset)*base*base*base +
			(sequence[1]+offset)*base*base +
			(sequence[2]+offset)*base +
			(sequence[3]+offset)
	}

	for _, price := range input {
		p, _ := strconv.Atoi(price)
		originalP := p
		priceChangeHistory := make([]int, 4)

		if seenSequences[originalP] == nil {
			seenSequences[originalP] = make(map[int]bool)
		}

		for i := 0; i < 2000; i++ {
			nextP := getNextPrice(p)
			priceChange := nextP%10 - p%10

			copy(priceChangeHistory, priceChangeHistory[1:])
			priceChangeHistory[3] = priceChange

			if i >= 3 {
				sequence := [4]int{priceChangeHistory[0], priceChangeHistory[1], priceChangeHistory[2], priceChangeHistory[3]}
				index := computeIndex(sequence)

				if !seenSequences[originalP][index] {
					sequenceScores[index] += nextP % 10
					seenSequences[originalP][index] = true
				}
			}

			p = nextP
		}
	}

	for _, sum := range sequenceScores {
		if sum > maxScore {
			maxScore = sum
		}
	}

	return maxScore
}