package binary

// TODO: see if this program can have lower complexity
// by using slice operator provided by Go

import (
	"fmt"
)

var counter int = 1

func BinarySearch(item int, list [] int) (index int){
	index = splitList(item, 0, len(list) - 1, list)
	return
}

func splitList(item int, startIndex int, endIndex int, list [] int) (halvedIndex int){
	fmt.Printf("Iteration #%d: ", counter)

	// break out scenario
	if startIndex > endIndex || endIndex == 0 {
		return -1
	}

	halvedIndex = (startIndex+endIndex)/2
	counter ++
	if list[halvedIndex] == item{
		fmt.Println("Item Found at halvedIndex")
		return
	} else if list[halvedIndex] > item{
		// item is on the left of the current slice
		fmt.Printf("Item must be on the left of list. Now, Searching in index %d through %d\n",
			startIndex, halvedIndex)
		return splitList(item, startIndex, halvedIndex, list)
	} else {
		fmt.Printf("Item must be on the right of list. Now, Searching in index %d through %d\n",
			halvedIndex+1, endIndex)
		return splitList(item, halvedIndex+1, endIndex, list)
	}
}

