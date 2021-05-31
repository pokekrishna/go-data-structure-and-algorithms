package selection

import "fmt"

// Sort is a typical implementation of selection sort algo.
// Time complexity cannot be more than O(n*n)
var operationCounter int
func Sort(sourceList []int) () {
	sort(sourceList, 0)
	operationCounter = 0
}

func sort(sourceList []int, startIndex int) error{



	if startIndex == len(sourceList) - 1{
		operationCounter++
		fmt.Printf("Interation #%v: \n", operationCounter)
		return nil
	}
	smallestIndex := startIndex
	for index := startIndex+1; index < len(sourceList); index ++{
		operationCounter++
		fmt.Printf("Interation #%v: \n", operationCounter)
		if (sourceList)[index] < (sourceList)[smallestIndex]{
			// found a smaller number
			// update the smallest index
			smallestIndex = index
		}
	}

	// so far, the smallest index has been found in the remaining list
	// now do, swap if smallest index is not the start index itself
	if smallestIndex != startIndex{
		SwapValues(sourceList, startIndex, smallestIndex)
	}
	fmt.Println(sourceList)
	return sort(sourceList, startIndex+1)

}

func SwapValues(sourceList []int, index1 int, index2 int){
	if len(sourceList) == 0 {
		return
	}
	if index1 == index2 {
		return
	}
	(sourceList)[index2] += (sourceList)[index1]
	(sourceList)[index1] = (sourceList)[index2] - (sourceList)[index1]
	(sourceList)[index2] = (sourceList)[index2] - (sourceList)[index1]
}