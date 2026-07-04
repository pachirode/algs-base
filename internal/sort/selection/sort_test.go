package selection

import (
	"testing"

	"github.com/pachirode/algs-base/internal/pkg/utils"
)

func TestSortNoInplace(t *testing.T) {
	// int：用逆序输入验证排序结果是否升序排列
	ints := utils.GenerateSortedSlice[int](20)
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	result := SortNoInplace(ints)
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("ints: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, result[i], i-1, result[i-1])
		}
	}

	// string：用逆序输入验证排序结果是否升序排列
	strs := []string{"c", "b", "a"}
	resultStr := SortNoInplace(strs)
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
	resultFloat := SortNoInplace(floats)
	for i := 1; i < len(resultFloat); i++ {
		if resultFloat[i] < resultFloat[i-1] {
			t.Errorf("floats: 排序结果未按升序排列, result[%d]=%v < result[%d]=%v", i, resultFloat[i], i-1, resultFloat[i-1])
		}
	}
}

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

// benchmarkSortNoInplace 是 benchmark 的公共逻辑：生成逆序输入然后排序。
func benchmarkSortNoInplace[T utils.Numeric](b *testing.B, n int) {
	data := utils.GenerateSortedSlice[T](n)
	// 逆序，构造最坏情况
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortNoInplace(data)
	}
}

func BenchmarkSortNoInplace_Int_10(b *testing.B)   { benchmarkSortNoInplace[int](b, 10) }
func BenchmarkSortNoInplace_Int_100(b *testing.B)  { benchmarkSortNoInplace[int](b, 100) }
func BenchmarkSortNoInplace_Int_1000(b *testing.B) { benchmarkSortNoInplace[int](b, 1000) }
