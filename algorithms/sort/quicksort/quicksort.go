package quicksort

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
)

const debug = true
// Sort takes in a slice of int and returns a new slice with the elements sorted
// using Quicksort algorithm
func Sort(l []int) []int{
	return sort(l)
}

// sort is the implementation of the Quicksort algorithm with some performance optimizations.
//
// sort does not use the append() built-in function because it allocates extra
// memory for the new array when the capacity of the existing underlying array is
// hit. Also when this happens, append() ends up copying all the elements from
// the older array to the new one, which is compute heavy task.
//
// sort makes use of the slice that is passed as parameter to keep track of the
// `left partition`, the pivot element, and the `right partition` - all in the
// same slice by swapping and keeping track of the indices. This is achieved by
// positioning all the elements greater than the pivot element, on the right of
// the pivot element, and remaining onto the left of the pivot element.
//
// Once the left and the right partitions are ready, Concat combines all the
// three partitions - left , pivot element, and right, and returns
func sort(l []int) []int{
	debugLn("current list", l)
	ll := len(l)
	if ll <= 1{
		return l
	}

	// p pivot index - choosing at half the length of the slice
	p := ll / 2

	// pivot element, so that the l can reused as 'left partition'
	// and pivot the element can be swapped forward in the index
	// to make room.
	pe := l[p]
	debugLn("pivot element", pe)

	var elementsInRPart int = 0
	dynamicLL := ll
	for i := 0; i<dynamicLL; i++ {
		// right partition
		if l[i] > pe {
			elementsInRPart ++
			selection.SwapValues(l, i, ll-elementsInRPart)
			i--
			dynamicLL--
		}
	}

	lPart := l[0:ll-elementsInRPart-1]
	rPart := l[ll-elementsInRPart:ll]
	debugLn("left", lPart)
	debugLn("right", rPart)

	// process left partition, then right partition
	leftSorted := sort(lPart)
	rightSorted := sort(rPart)
	c := Concat(leftSorted, pe, rightSorted)
	debugLn("returning", c)
	return c
}

// Concat efficiently concatenates a slice, an int, and another slice without
// using the append() for the performance reasons. Using append() would result in
// creating underlying arrays twice in the worst case - once for appending the
// integer, and another for appending the slice. Also, append() would end up
// reserving more memory than needed in the worst case, and copying data from old
// array to new.
//
// These aforementioned pitfalls are resolved by making a slice of the exact size
// that is required therefore eliminating waste of memory and compute.
func Concat(leftSorted []int, pe int, rightSorted []int) []int {
	debugLn("concatenating...", leftSorted, pe, rightSorted)
	l := len(leftSorted) + len(rightSorted) + 1
	combined := make([]int, l, l)
	i := 0
	for ; i < len(leftSorted); i++{
		combined[i] = leftSorted[i]
	}
	combined[i] = pe
	i++
	for j := 0; j < len(rightSorted); j++{
		combined[i] = rightSorted[j]
		i++
	}
	debugLn("combined into ->", combined)
	return combined
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}