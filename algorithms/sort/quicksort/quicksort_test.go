package quicksort_test

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/quicksort"
	"testing"
)

func TestQuickSort(t *testing.T){
	testcases := []struct{
		input []int
		expected []int
	}{
		{
			input: []int {1, 10, 2, 1, 9, 0},
			expected: []int{0,1,1,2,9,10},
		},
		{
			input: []int {1, 10, 2, 4, 9, 0},
			expected: []int{0,1,2,4,9,10},
		},
		{
			input: []int {},
			expected: []int{},
		},
		{
			input: []int {1},
			expected: []int{1},
		},
		{
			input: []int {10, -10},
			expected: []int{-10, 10},
		},
	}

	for i, tc := range testcases{
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got := quicksort.Sort(tc.input)
			if !assertEqual(got, tc.expected) {
				t.Errorf("got:%v, expected:%v", got, tc.expected)
			}
		})
	}
}

func assertEqual(got []int, expected []int) bool{
	if len(got) != len(expected) {
		return false
	}
	for i := range expected {
		if expected[i] != got[i] {
			return false
		}
	}
	return true
}