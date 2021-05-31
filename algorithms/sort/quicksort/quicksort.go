package quicksort

import "github.com/pokekrishna/dsa/algorithms/sort/selection"

func Sort(l []int) []int{

}

func sort(l []int){
	// p pivot index - choosing at half the length of the slice
	ll := len(l)
	p := ll / 2



	// pivot element, so that the l can reused as 'left partition'
	// and pivot the element can be swapped forward in the index
	// to make room.
	pe := l[p]

	// rpart is the right partition of the list. this will contain
	// only the numbers which are greater than pivot element, pe.
	//
	// it is of fixed size because slice's append operation creates
	// a new slice of double the length needed and also all the
	// elements are copied everytime append reaches the limit
	// of the slice
	rPart := make([]int, ll-1, ll-1)
	var elementsInRPart int = 0
	dynamicLL := ll
	for i := 0; i<dynamicLL; i++ {
		// skip the processing on the element at pivot
		if i == p{
			continue
		}

		// right partition
		if l[i] > pe {
			rPart = append(rPart, l[i])
			elementsInRPart ++
			selection.SwapValues(l, i, ll-elementsInRPart)
			i--
			dynamicLL--
		}
	}
}