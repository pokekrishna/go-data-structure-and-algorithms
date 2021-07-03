package merge

import (
	"fmt"
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
)

const debug = true

func Sort(l []int){
	ll := len(l)
	merge(l, 0, ll)

}

func merge(l []int, low, high int){
	debugLn("received", l, low, high)


	if high - low < 2  {
		debugLn("ret:", l[low:high])
		return // do nothing
	}

	merge(l, low, (low+high)/2)
	merge(l, (low+high)/2, high)

	fmt.Printf("low %d, (low+high)/2 %d, high %d\n", low, (low+high)/2, high)
	for i, j := low, (low+high)/2; i < (low+high)/2 && j < high; {
		debugLn("l[i]", l[i])
		debugLn("l[j]", l[j])
		if l[i] < l [j] {
			debugLn("swapping", l[i], "and", l[j])
			selection.SwapValues(l, i, j)
			i++
		} else if l [j] < l[i]{

		}
	}
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}