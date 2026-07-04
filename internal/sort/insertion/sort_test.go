package insertion

import (
	"testing"

	"github.com/pachirode/algs-base/internal/pkg/utils"
)

func TestSort(t *testing.T) {
	// int：用逆序输入验证排序结果是否升序排列
	ints := utils.GenerateSortedSlice[int](20)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	result := Sort(ints)
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("ints: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, result[i], i-1, result[i-1])
		}
	}

	// string：用逆序输入验证排序结果是否升序排列
	strs := []string{"c", "b", "a"}
	resultStr := Sort(strs)
	for i := 1; i < len(resultStr); i++ {
		if resultStr[i] < resultStr[i-1] {
			t.Errorf("strs: 排序结果未按升序排列, result[%d]=%q < result[%d]=%q", i, resultStr[i], i-1, resultStr[i-1])
		}
	}

	// float64：用逆序输入验证排序结果是否升序排列
	floats := utils.GenerateSortedSlice[float64](10)
	for i, j := 0, len(floats)-1; i < j; i, j = i+1, j-1 {
		floats[i], floats[j] = floats[j], floats[i]
	}
	resultFloat := Sort(floats)
	for i := 1; i < len(resultFloat); i++ {
		if resultFloat[i] < resultFloat[i-1] {
			t.Errorf("floats: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, resultFloat[i], i-1, resultFloat[i-1])
		}
	}
}

func TestSortOptimize(t *testing.T) {
	// int：用逆序输入验证排序结果是否升序排列
	ints := utils.GenerateSortedSlice[int](20)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	result := SortOptimize(ints)
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("ints: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, result[i], i-1, result[i-1])
		}
	}

	// string：用逆序输入验证排序结果是否升序排列
	strs := []string{"c", "b", "a"}
	resultStr := SortOptimize(strs)
	for i := 1; i < len(resultStr); i++ {
		if resultStr[i] < resultStr[i-1] {
			t.Errorf("strs: 排序结果未按升序排列, result[%d]=%q < result[%d]=%q", i, resultStr[i], i-1, resultStr[i-1])
		}
	}

	// float64：用逆序输入验证排序结果是否升序排列
	floats := utils.GenerateSortedSlice[float64](10)
	for i, j := 0, len(floats)-1; i < j; i, j = i+1, j-1 {
		floats[i], floats[j] = floats[j], floats[i]
	}
	resultFloat := SortOptimize(floats)
	for i := 1; i < len(resultFloat); i++ {
		if resultFloat[i] < resultFloat[i-1] {
			t.Errorf("floats: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, resultFloat[i], i-1, resultFloat[i-1])
		}
	}
}

// 基准测试辅助函数：生成逆序输入（最坏情况），然后调用排序函数。
func benchSort(fn func([]int) []int, n int) func(*testing.B) {
	return func(b *testing.B) {
		data := make([]int, n)
		for i := 0; i < n; i++ {
			data[i] = n - i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			input := make([]int, n)
			copy(input, data)
			fn(input)
		}
	}
}

// Sort 基础版本
func BenchmarkSort_10(b *testing.B)   { benchSort(Sort[int], 10)(b) }
func BenchmarkSort_100(b *testing.B)  { benchSort(Sort[int], 100)(b) }
func BenchmarkSort_1000(b *testing.B) { benchSort(Sort[int], 1000)(b) }

// SortOptimize 优化版本
func BenchmarkSortOptimize_10(b *testing.B)   { benchSort(SortOptimize[int], 10)(b) }
func BenchmarkSortOptimize_100(b *testing.B)  { benchSort(SortOptimize[int], 100)(b) }
func BenchmarkSortOptimize_1000(b *testing.B) { benchSort(SortOptimize[int], 1000)(b) }
