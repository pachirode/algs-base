package insertion

import (
	"cmp"
)

func Sort[T cmp.Ordered](data []T) []T {
	for i := 1; i < len(data); i++ {
		for j := i - 1; j > -1; j-- {
			if data[j+1] > data[j] {
				break
			}
			data[j+1], data[j] = data[j], data[j+1]
		}
	}
	return data
}

func SortOptimize[T cmp.Ordered](data []T) []T {
	for i := 1; i < len(data); i++ {
		cur := data[i]
		j := i
		for j > 0 && data[j-1] > cur {
			data[j] = data[j-1]
			j--
		}
		data[j] = cur
	}
	return data
}
