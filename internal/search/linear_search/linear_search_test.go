package linear_search

import (
	"fmt"
	"testing"

	"github.com/pachirode/algs-base/internal/pkg/student"
	"github.com/pachirode/algs-base/internal/pkg/utils"
)

func TestLinearSearch(t *testing.T) {
	// int
	ints := utils.GenerateSortedSlice[int](4)
	fmt.Printf("Search(ints, 1): %v\n", Search(ints, 1))
	fmt.Printf("Search(ints, 5): %v\n", Search(ints, 5))

	// string
	strs := []string{"a", "b", "c"}
	fmt.Printf("Search(strs, \"b\"): %v\n", Search(strs, "b"))
	fmt.Printf("Search(strs, \"z\"): %v\n", Search(strs, "z"))

	// float64
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Printf("Search(floats, 2.2): %v\n", Search(floats, 2.2))
	fmt.Printf("Search(floats, 4.4): %v\n", Search(floats, 4.4))

	// Student 自定义类型
	students := []student.Student{
		{Name: "张三", ID: 1001, Socre: 88.5},
		{Name: "李四", ID: 1002, Socre: 92.0},
		{Name: "王五", ID: 1003, Socre: 76.5},
	}
	fmt.Printf("Search(students, 李四): %v\n", Search(students, student.Student{Name: "李四", ID: 1002, Socre: 92.0}))
	fmt.Printf("Search(students, 赵六): %v\n", Search(students, student.Student{Name: "赵六", ID: 9999, Socre: 0}))
}

func BenchmarkLinearSearch(b *testing.B) {
	data := utils.GenerateSortedSlice[int](100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(data, 9999)
	}
}
