package insertion

func Sort(l []int){
	for i:=1; i < len(l); i++ {
		// save the item
		elem := l[i]

		// Scan through the items on the left of the 'elem' - from right to left
		// direction (index i-1 through -1). If 'elem' is less than item at j push the
		// items by one index to the right, until 'elem' is greater than item at j, to
		// make space for 'elem' to be inserted. When either scanning has reached the
		// start of the list, or 'elem' is greater than item, then INSERT the element
		// 'elem' and break out of j loop.
		for j := i-1; j >= -1 ; j-- {
			if j >= 0 && elem < l[j]{
				// shift element at j by one index
				l[j+1] = l[j]
			} else {
				l[j+1] = elem
				break
			}
		}
	}
}