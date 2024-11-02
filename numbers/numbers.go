package numbers_util

import "sort"

func UniqueInts(ints []int) []int {
	seen := make(map[int]struct{})
	result := []int{}

	for _, v := range ints {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func DescendingInts(ints []int) []int {
	result := make([]int, len(ints))
	copy(result, ints)

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}
