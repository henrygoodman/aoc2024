package day09

import (
	"aoc2024/common"
	"fmt"
)

func Solve() {
	input, err := common.ReadInput(9)
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
	if len(input) != 1 {
		fmt.Println("Expected 1 line input")
		return -1
	}

	diskMap := input[0]
	var decodedMap []int
	fileID := 0

    // Build decoded repr (list since > 10 IDs, cannot repr with digits)
	for i := 0; i < len(diskMap); i++ {
		blockSize := int(diskMap[i] - '0')

		if i%2 == 0 {
			for j := 0; j < blockSize; j++ {
				decodedMap = append(decodedMap, fileID)
			}
			fileID++
		} else {
			for j := 0; j < blockSize; j++ {
				decodedMap = append(decodedMap, -1) // -1 to represent free space
			}
		}
	}

	totalLen := len(decodedMap)

	left, right := 0, totalLen-1
	resultSum := 0

	for left < totalLen {
		if decodedMap[left] != -1 {
			resultSum += decodedMap[left] * left
			left++
		} else {
			for right > left && decodedMap[right] == -1 {
				right--
			}
			if right > left {
				resultSum += decodedMap[right] * left
				decodedMap[right] = -1
				right--
			}
			left++
		}
	}

	return resultSum
}

type freeSpan struct {
	start int
	size  int
}

func solvePart2(input []string) int {
	if len(input) != 1 {
		fmt.Println("Expected 1 line input")
		return -1
	}

	diskMap := input[0]
	var decodedMap []int
	fileID := 0

    // Build decoded repr (list since > 10 IDs, cannot repr with digits)
	for i := 0; i < len(diskMap); i++ {
		blockSize := int(diskMap[i] - '0')

		if i%2 == 0 {
			for j := 0; j < blockSize; j++ {
				decodedMap = append(decodedMap, fileID)
			}
			fileID++
		} else {
			for j := 0; j < blockSize; j++ {
				decodedMap = append(decodedMap, -1) // -1 to represent free space
			}
		}
	}

	// Precompute file positions
	filePositions := make(map[int][2]int)
	for i := 0; i < len(decodedMap); i++ {
		if decodedMap[i] != -1 {
			fileID := decodedMap[i]
			if _, exists := filePositions[fileID]; !exists {
				filePositions[fileID] = [2]int{i, i}
			} else {
				filePositions[fileID] = [2]int{filePositions[fileID][0], i}
			}
		}
	}

	freeSpans := findFreeSpans(decodedMap)

    // Starting from RHS (last fileID), try and move each to the first available free span
    // then update the freeSpans. We dont really care about updating the vacated locations, since they
    // will never be populated (we dont ever 'push' anything to the right to take its place)
	for currentFileID := fileID - 1; currentFileID >= 0; currentFileID-- {
		start, end := filePositions[currentFileID][0], filePositions[currentFileID][1]
		fileSize := end - start + 1

		for i, span := range freeSpans {
			if span.size >= fileSize && span.start < start {
				// Move the file to the free span
				for j := 0; j < fileSize; j++ {
					decodedMap[span.start+j] = currentFileID
				}
				// Mark the original file location as free space
				for j := start; j <= end; j++ {
					decodedMap[j] = -1
				}
				freeSpans[i] = freeSpan{start: span.start + fileSize, size: span.size - fileSize}
				break
			}
		}
	}

	resultSum := 0
	for i, char := range decodedMap {
		if char != -1 {
			resultSum += char * i
		}
	}

	return resultSum
}

// Find all free space spans
func findFreeSpans(diskMap []int) []freeSpan {
	var spans []freeSpan
	start := -1

	for i, val := range diskMap {
		if val == -1 {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				spans = append(spans, freeSpan{start: start, size: i - start})
				start = -1
			}
		}
	}

	if start != -1 {
		spans = append(spans, freeSpan{start: start, size: len(diskMap) - start})
	}

	return spans
}