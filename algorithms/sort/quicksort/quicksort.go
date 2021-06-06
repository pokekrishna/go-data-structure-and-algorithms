package quicksort

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
)

const debug = true
// Sort takes in a slice of int and returns that same slice with the elements sorted
// using Quicksort algorithm
func Sort(l []int) []int{
	sort(l)
	return l
}

// sort is the implementation of the Quicksort algorithm with some performance optimizations.
//
// sort does not use the append() built-in function because append allocates extra
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
// Once the left and the right partitions are ready, sort is called recursively
// on the partitions making sure a new array is never created, hence there is
// no return type of the function.
func sort(l []int){
	debugLn("current list", l)
	ll := len(l)
	if ll <= 1{
		return
	}

	// p pivot index - choosing at half the length of the slice
	p := ll / 2

	// pivot element, so that the l can reused as 'left partition'
	// and pivot the element can be swapped forward in the index
	// to make room.
	pe := l[p]
	debugLn("pivot element", pe)

	var elementsInRPart int = 0

	// dynamicLL is short notation of `dynamic` list length. This is the
	// pseudo length of the list which keeps decreasing by 1 whenever a
	// swap happens in the list. The length `decreases` because once a
	// number greater than the pivot element is swapped
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
	sort(lPart)
	sort(rPart)
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}