package common

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(day int) ([]string, error) {
	fileType := "txt"

	filename := fmt.Sprintf("inputs/day%02d.%s", day, fileType)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
