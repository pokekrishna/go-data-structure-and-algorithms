package singlylinkedlist

import "fmt"

type Node struct {
	Item interface{}
	next *Node
}

func Append(tail *Node, item *Node) error {
	if tail == nil {
		return fmt.Errorf("tail is nil")
	}
	if item == nil{
		return fmt.Errorf("item to append is nil")
	}
	tail.next = item
	return nil
}

func Traverse(head *Node) error {
	// TODO: complete implementation
	return nil
}

func (n *Node) Next() *Node{
	if n != nil{
		return n.next
	}
	return nil
}