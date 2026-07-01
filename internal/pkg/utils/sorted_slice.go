package utils

// Numeric 约束所有数值类型（有符号/无符号整数 + 浮点数）。
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// GenerateSortedSlice 生成一个长度为 n 的有序切片，元素从 0 开始严格递增。
// 类型参数 T 必须为数值类型（整数或浮点）。
func GenerateSortedSlice[T Numeric](n int) []T {
	if n <= 0 {
		return []T{}
	}
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = T(i)
	}
	return res
}
