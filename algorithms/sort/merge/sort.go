package merge

import (
	"fmt"
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

	debugLn(fmt.Sprintf("low %d, (low+high)/2 %d, high %d\n", low, (low+high)/2, high))
	for i, j := low, (low+high)/2; i < j && j < high; i ++{
		debugLn("l[i]", l[i])
		debugLn("l[j]", l[j])
		if l[i] < l [j] {
			debugLn("l[i] is already smaller than l[j]. doing nothing. i++")
		} else if l [j] < l[i]{
			debugLn("inserting l[j] at i, do shifting")
			key := l[j]
			for k := j; k > i; k --{
				l[k]=l[k-1]
			}
			l[i] = key
			j++
			debugLn("after shifting", l)
		}
	}
}

func debugLn(v ...interface{}){
	if debug{
		fmt.Println(v...)
	}
}