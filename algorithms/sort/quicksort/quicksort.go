package quicksort

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
)

const debug = true

func Sort(l []int) []int{
	return sort(l)
}

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