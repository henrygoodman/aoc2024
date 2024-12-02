package common

type Pair struct {
    First  int
    Second int
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
