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
		p,_ := strconv.Atoi(price)
		for range 2000 {
			nextP := getNextPrice(p)
			p = nextP
		}
		res += p
	}
	return res
}

func solvePart2(input []string) int {
    return 5
}
