package sort

import "testing"

func TestMergesort(t *testing.T) {
	tesecase := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{-1, 0, -2, -11, 5, 10},
			expected: []int{-11, -2, -1, 0, 5, 10},
		},
		{
			input:    []int{0, 0, -1, 1},
			expected: []int{-1, 0, 0, 1},
		},
	}

	for _, c := range tesecase {
		if after := Mergesort(c.input); !equal(after, c.expected) {
			t.Fatalf("Wrong")
		}
	}
}

func BenchmarkMergesort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quicksort([]int{1, -1, 0})
	}
}
