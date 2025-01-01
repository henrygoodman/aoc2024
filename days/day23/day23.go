package day23

import (
	"aoc2024/common"
	"fmt"
	"sort"
	"strings"
)

func Solve() {
	input, err := common.ReadInput(23)
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

func serialize(slice []string) string {
	return strings.Join(slice, "|")
}

func solvePart1(input []string) int {
	mappings := make(map[string][]string)
	detectedTriples := make(map[string]struct{})
	sum := 0

	for _, mapping := range input {
		parts := strings.Split(mapping, "-")
		if len(parts) != 2 {
			continue
		}
		src, dst := strings.ToLower(strings.TrimSpace(parts[0])), strings.ToLower(strings.TrimSpace(parts[1]))
		mappings[src] = append(mappings[src], dst)

		if !contains(mappings[dst], src) {
			mappings[dst] = append(mappings[dst], src)
		}
	}

	for src, neighbors := range mappings {
		if len(neighbors) < 2 {
			continue
		}

		for _, dst := range neighbors {
			if len(mappings[dst]) < 2 {
				continue
			}

			for _, dstEdge := range mappings[dst] {
				if dstEdge == src || dstEdge == dst {
					continue
				}

				// Check if dstEdge is also a neighbor of src
				if !contains(mappings[src], dstEdge) {
					continue
				}

				// Create and sort the triple to normalize
				triple := []string{src, dst, dstEdge}
				sort.Strings(triple)
				key := serialize(triple)

				if _, exists := detectedTriples[key]; exists {
					continue
				}

				detectedTriples[key] = struct{}{}

				foundT := false
				for _, machine := range triple {
					if strings.HasPrefix(machine, "t") {
						foundT = true
						break
					}
				}

				if foundT {
					sum++
				}
			}
		}
	}

	return sum
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func solvePart2(input []string) int {
	return 5
}
