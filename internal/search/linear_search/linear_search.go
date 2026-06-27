package linear_search

func Search[T comparable](data []T, target T) int {
	for i, v := range data {
		if v == target {
			return i
		}
	}
	return -1
}
