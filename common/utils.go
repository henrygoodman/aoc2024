package common

import "fmt"

type Pair struct {
    First, Second int
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func Reduce[T any, U any](slice []T, reducer func(U, T) U, initial U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func Sign(x int) int {
    if x > 0 {
        return 1
    } else if x < 0 {
        return -1
    }
    return 0
}

func Window[T any](data []T, size int) [][]T {
	if size <= 0 || size > len(data) {
		return nil
	}

	result := make([][]T, 0, len(data)-size+1)
	for i := 0; i <= len(data)-size; i++ {
		windowCopy := make([]T, size)
		copy(windowCopy, data[i:i+size])
		result = append(result, windowCopy)
	}

	return result
}

func CircularWindow[T any](data []T, size int) [][]T {
    if size <= 0 || size > len(data) {
        return nil
    }

    result := make([][]T, 0, len(data))
    for i := 0; i < len(data); i++ {
        window := make([]T, size)
        for j := 0; j < size; j++ {
            window[j] = data[(i+j)%len(data)]
        }
        result = append(result, window)
    }

    return result
}

func AllPairs[T any](items []T) [][2]T {
    var pairs [][2]T
    for i := 0; i < len(items); i++ {
        for j := i + 1; j < len(items); j++ {
            pairs = append(pairs, [2]T{items[i], items[j]})
        }
    }
    return pairs
}

func Zip(a, b []int) []Pair {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	result := make([]Pair, n)
	for i := 0; i < n; i++ {
		result[i] = Pair{First: a[i], Second: b[i]}
	}
	return result
}

func ParseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return grid
}

func DisplayGrid(grid [][]rune) {
    for _, row := range grid {
        for _, cell := range row {
            fmt.Print(string(cell), " ")
        }
        fmt.Println()
    }
}
