package binary

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int {
		1,2, 9, 20, 31, 45, 63, 70, 100,
	}

	iterations := []int{}
	for _,item := range list{
		counter = 1
		fmt.Printf("Finding %d\n---\n", item)
		fmt.Println("index:", BinarySearch(item, list))
		iterations = append(iterations, counter)
		fmt.Printf("---\n\n")
	}
	fmt.Println("list:", list)
	fmt.Println("Iterations: ",iterations)

	counter =1
	fmt.Printf("Finding %d\n---\n", -100)
	fmt.Println("index:", BinarySearch(-100, list))

	counter =1
	fmt.Printf("Finding %d\n---\n", 1000)
	fmt.Println("index:", BinarySearch(1000, list))
}
