package selection

import "cmp"

func SortNoInplace[T cmp.Ordered](data []T) []T {
	var tempList = make([]T, len(data))
	copy(tempList, data)

	var result = make([]T, len(data))

	for i := 0; i < len(data); i++ {
		mixIdx := 0
		for j := 1; j < len(tempList); j++ {
			if tempList[j] < tempList[mixIdx] {
				mixIdx = j
			}
		}

		result[i] = tempList[mixIdx]
		tempList[mixIdx] = tempList[len(tempList)-1-i]
	}

	return result
}

func Sort[T cmp.Ordered](data []T) []T {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

func SortOptimize[T cmp.Ordered](data []T) []T {
	for i := 0; i < len(data); i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	return data
}
