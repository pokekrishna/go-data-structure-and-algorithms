package quicksort

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
)

const debug = false
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
	// length of the current list
	ll := len(l)
	if ll <= 1{
		return
	}

	// pi pivot index - choosing at half the length of the slice
	pi := ll / 2
	pe := l[pi]
	debugLn("pivot element", pe)

	var lenRightPart int = 0

	// traverse the list. Move items greater than pe to the right of pivot index pi.
	// Move items less than pe to the left of pivot index pi. Make sure not traverse
	// the Right Partition because it is already traversed essentially through the
	// 'i--' statement and the loop var limit of (ll - lenRightPart).
	//
	// Whenever needed move the pivot index towards left or right, keeping the pivot
	// element same throughout the loop, in order to have the left and right parition
	// without using a separate list.
	for i := 0; i < (ll-lenRightPart); i++ {
		if l[i] > pe {
			// right partition
			lenRightPart++
			selection.SwapValues(l, i, ll-lenRightPart)
			if ll-lenRightPart == pi{
				pi = i // move pi left
 			}
			i-- // since the values have been swapped, re assess the new value at index i
		} else if i > pi {
			// Found an element less than pe but on the right of pivot pi, so swap the values
			// and update the pivot index
			selection.SwapValues(l, i, pi)
			pi = i // move pi right
		}
	}

	lPart := l[0:ll-lenRightPart-1]
	rPart := l[ll-lenRightPart :ll]
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