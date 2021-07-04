package merge

import (
	"fmt"
)

const debug = false

// Sort calls merge() on the input list.
func Sort(l []int){
	merge(l, 0, len(l))
}

// merge breaks down the list into sublist by breaking at half the length
// recursively, until length is equal to 1 - returns in that case. When both the
// left and right sublist returns, sort the two list and make changes in the
// original list.
func merge(l []int, low, high int){
	debugLn("received", l, low, high)

	// BASE CASE: when length of input list (or sublist) is 1 or smaller
	if high - low < 2  {
		return // do nothing
	}

	// split the input list into sublist - left and right and call merge on them
	merge(l, low, (low+high)/2)
	merge(l, (low+high)/2, high)

	debugLn(fmt.Sprintf("low %d, (low+high)/2 %d, high %d", low, (low+high)/2, high))

	// After above merge call has finished, sort the two input lists but make changes
	// directly on the master list
	//
	// Loop with two variables i and j. 'i' will traverse the combined list (left and
	// right) so that we know we have traversed all the elements.'j' will traverse
	// only the right sublist and making sure that the j only increases when an
	// element at index j is replaced with another index. This way we keep track that
	// all the elements on the right sublist have been read and placed in their
	// appropriate positions.
	//
	// Elements from the right sublist, when changes its position (index), causes the
	// entire list to shift. This action is similar to insertion sort.
	for i, j := low, (low+high)/2; i < j && j < high; i ++{
		if l [j] < l[i]{
			debugLn("insert l[j] at i, but before that do shifting")
			key := l[j]
			for k := j; k > i; k --{
				l[k]=l[k-1] // shifting and making room for 'key' to come at index i
			}
			l[i] = key
			j++
			debugLn("after shifting", l)
		} else if l[i] < l [j] {
			debugLn("l[i] is already smaller than l[j]. doing nothing. i++")
		}
	}
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}