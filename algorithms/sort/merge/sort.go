package merge

import (
	"fmt"
)

const debug = false

// Sort calls mergeSort() on the input list.
func Sort(l []int){
	mergeSort(l, 0, len(l))
}

// mergeSort breaks down the list into sublist by breaking at half the length
// recursively, until length is equal to 1 - returns in that case. When both the
// left and right sublist returns, merge the two list and make changes in the
// original list.
func mergeSort(l []int, low, high int){
	debugLn("received", l, low, high)

	// BASE CASE: when length of input list (or sublist) is 1 or smaller
	if high - low < 2  {
		return // do nothing
	}

	// split the input list into sublist - left and right and call mergeSort on them
	mergeSort(l, low, (low+high)/2)
	mergeSort(l, (low+high)/2, high)

	//merge the two sublists onto 'l'
	merge(l, low, (low+high)/2, high)
}

// merge takes in two individually sorted sublist (Left and Right) by their start and end
// markers 'low', 'mid', 'high', and merges them onto 'l' and making sure that
// the list l[low:high] is sorted.
//
// merge should finish in O(N)
func merge(l []int, low, mid, high int){
	debugLn(fmt.Sprintf("low %d, mid %d, high %d", low, mid, high))

	// Create copies of the sub lists so that changes can be made in l
	leftLen := mid-low
	left := make([]int, leftLen)
	copy(left, l[low:mid])
	rightLen := high-mid
	right := make([]int, rightLen)
	copy(right, l[mid:high])

	// sort and merge onto 'l'. use j and k for progression markers in left and right
	// respectively
	for i, j, k := low, 0, 0; i < high; i ++ {
		debugLn("Deciding for index:", i)
		if j == leftLen {
			// means, left list is completely traversed, just copy right elements from now
			l[i] = right[k]
			k++
		} else if k == rightLen {
			// means, right list is completely traversed, just copy left elements from now
			l[i] = left[j]
			j++
		} else if left[j] <= right[k]{
			// means, left element is smaller or equal, copy it to l
			l[i] = left[j]
			j++
		} else if right[k] < left[j] {
			// means, right element is smaller, copy it to l
			l[i] = right[k]
			k++
		}

	}
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}