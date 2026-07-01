package utils

import (
	"reflect"
	"testing"
)

func TestGenerateSortedSlice(t *testing.T) {
	t.Run("int positive length", func(t *testing.T) {
		got := GenerateSortedSlice[int](5)
		want := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GenerateSortedSlice[int](5) = %v, want %v", got, want)
		}
	})

	t.Run("int zero length", func(t *testing.T) {
		got := GenerateSortedSlice[int](0)
		if len(got) != 0 {
			t.Errorf("GenerateSortedSlice[int](0) len = %d, want 0", len(got))
		}
	})

	t.Run("int negative length", func(t *testing.T) {
		got := GenerateSortedSlice[int](-3)
		if len(got) != 0 {
			t.Errorf("GenerateSortedSlice[int](-3) len = %d, want 0", len(got))
		}
	})

	t.Run("float64", func(t *testing.T) {
		got := GenerateSortedSlice[float64](3)
		want := []float64{0.0, 1.0, 2.0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GenerateSortedSlice[float64](3) = %v, want %v", got, want)
		}
	})

	t.Run("int8", func(t *testing.T) {
		got := GenerateSortedSlice[int8](3)
		want := []int8{0, 1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GenerateSortedSlice[int8](3) = %v, want %v", got, want)
		}
	})

	t.Run("uint", func(t *testing.T) {
		got := GenerateSortedSlice[uint](4)
		want := []uint{0, 1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GenerateSortedSlice[uint](4) = %v, want %v", got, want)
		}
	})

	t.Run("single element", func(t *testing.T) {
		got := GenerateSortedSlice[int](1)
		want := []int{0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GenerateSortedSlice[int](1) = %v, want %v", got, want)
		}
	})
}
