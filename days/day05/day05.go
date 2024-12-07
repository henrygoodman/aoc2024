package day05

import (
	"aoc2024/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func Solve() {
	input, err := common.ReadInput(5)
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

type PriorPost struct {
	PriorSet map[int]struct{}
	PostSet  map[int]struct{}
}

func parseUpdateLine(line string) ([]int, error) {
	numStrs := strings.Split(line, ",")
	var nums []int
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			return nil, fmt.Errorf("invalid number in update: %s", numStr)
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func buildMapping(input []string) map[int]PriorPost {
	m := make(map[int]PriorPost)
	for _, line := range input {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				continue
			}

			valStr := strings.TrimSpace(parts[0])
			mappingStr := strings.TrimSpace(parts[1])

			val, err1 := strconv.Atoi(valStr)
			mapping, err2 := strconv.Atoi(mappingStr)
			if err1 != nil || err2 != nil {
				continue
			}

			entry := m[val]
			if entry.PostSet == nil {
				entry.PostSet = make(map[int]struct{})
			}
			entry.PostSet[mapping] = struct{}{}
			m[val] = entry

			mappingEntry := m[mapping]
			if mappingEntry.PriorSet == nil {
				mappingEntry.PriorSet = make(map[int]struct{})
			}
			mappingEntry.PriorSet[val] = struct{}{}
			m[mapping] = mappingEntry
		}
	}
	return m
}

func solvePart1(input []string) int {
	m := buildMapping(input)
	var sum int64
	const workers = 16

	lines := make(chan string, len(input)) 
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for line := range lines {
			if line == "" || strings.Contains(line, "|") {
				continue
			}

			nums, err := parseUpdateLine(line)
			if err != nil || len(nums)%2 == 0 {
				continue
			}

			isValid := true
			for i, curr := range nums {
				left := nums[:i]
				right := nums[i+1:]

				entry, exists := m[curr]
				if !exists {
					isValid = false
					break
				}

				for _, num := range left {
					if _, exists := entry.PriorSet[num]; !exists {
						isValid = false
						break
					}
				}

				for _, num := range right {
					if _, exists := entry.PostSet[num]; !exists {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}

			if isValid {
				midIdx := len(nums) / 2
				atomic.AddInt64(&sum, int64(nums[midIdx]))
			}
		}
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker()
	}

	for _, line := range input {
		lines <- line
	}
	close(lines)

	wg.Wait()

	return int(sum)
}

func solvePart2(input []string) int {
	m := buildMapping(input)
	var sum int64
	const workers = 16

	lines := make(chan string, len(input))
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for line := range lines {
			if line == "" || strings.Contains(line, "|") {
				continue
			}

			nums, err := parseUpdateLine(line)
			if err != nil || len(nums)%2 == 0 {
				continue
			}

			sorted := make([]int, len(nums))
			copy(sorted, nums)
			sort.Slice(sorted, func(i, j int) bool {
				a, b := sorted[i], sorted[j]
				if _, exists := m[b].PriorSet[a]; exists {
					return true
				}
				if _, exists := m[b].PostSet[a]; exists {
					return false
				}
				return false
			})

			isSorted := true
			for i := range nums {
				if nums[i] != sorted[i] {
					isSorted = false
					break
				}
			}

			if !isSorted {
				midIdx := len(sorted) / 2
				atomic.AddInt64(&sum, int64(sorted[midIdx]))
			}
		}
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker()
	}

	for _, line := range input {
		lines <- line
	}
	close(lines)

	wg.Wait()

	return int(sum)
}
