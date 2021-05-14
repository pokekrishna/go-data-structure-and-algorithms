package singlylinkedlist_test

import (
	sll "github.com/pokekrishna/dsa/data-structures/linear/linkedlist/singlylinkedlist"
	"testing"
)

func TestAppend(t *testing.T){
	var head *sll.Node
	var tail *sll.Node
	for i := 1; i <=10; i++ {
		if head == nil {
			head = &sll.Node{Item: i}
			tail = head
		} else {
			n := &sll.Node{Item: i}

			if err := sll.Append(tail, n); err != nil{
				t.Error("Cannot append", err)
			} else {
				tail = n
			}

		}
	}

	if head.Item != 1{
		t.Error("should receive 1 at head, but didnt ")
	}

	got := head.Next().Next().Item
	if got != 3 {
		t.Error("should receive 3 at 3rd node, but got", got)
	}
}