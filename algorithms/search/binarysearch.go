package main

// TODO: see if this program can have lower complexity
// by using slice operator provided by Go

import (
	"fmt"
	"math"
)

var counter int = 1

func BinarySearch(item int, list [] int) (index int){
	index = splitList(item, 0, len(list) - 1, list)
	return
}

func splitList(item int, startIndex int, endIndex int, list [] int) (halvedIndex int){
	fmt.Printf("Iteration #%d: ", counter)
	counter ++

	halvedIndex = divide(endIndex+1 - startIndex, 2) - 1
	if list[halvedIndex] == item{
		fmt.Println("Item Found at halvedIndex")
	} else if list[halvedIndex] > item{
		// item is on the left of the current slice
		fmt.Printf("Item must be on the left of list. Now, Searching in index %d through %d\n",
			startIndex, halvedIndex)
		return splitList(item, startIndex, halvedIndex, list)
	} else {
		fmt.Printf("Item must be on the right of list. Now, Searching in index %d through %d\n",
			halvedIndex, endIndex)
		return splitList(item, halvedIndex, endIndex, list)
	}
	return

}

func divide(numerator int, denomenator int) (answer int) {
	n := float64(numerator)
	d := float64(denomenator)

	a := math.Ceil(n/d)
	answer = int(a)
	return
}

func main(){
	list := []int {
		1, 2, 4, 5, 10, 90, 110, 190, 210,
	}
	fmt.Println(list)
	fmt.Println("index:", BinarySearch(2, list))

}