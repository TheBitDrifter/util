package iter_util

import "iter"

func Collect[T any](seq iter.Seq[T]) []T {
	collect := []T{}
	for el := range seq {
		collect = append(collect, el)
	}
	return collect
}

func Collect2[T any](seq iter.Seq2[int, T]) []T {
	collect := []T{}
	for _, el := range seq {
		collect = append(collect, el)
	}
	return collect
}
