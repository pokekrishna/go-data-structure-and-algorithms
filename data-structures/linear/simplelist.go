package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List
	l.PushBack(11)
	l.PushBack(22)
	l.PushBack(33)

	for element := l.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)
	}
}