package singlylinkedlist

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func CreateDummy() {
	var prev *Node = nil
	for i := 10; i > 0; i-- {
		prev  = &Node{i, prev}
	}

	traverse(prev)
}

func traverse(head *Node) {
	for n := head; n != nil; n =  n.Next {
		fmt.Print(n.Data, "->")
	}
	fmt.Println("\\0")
}
