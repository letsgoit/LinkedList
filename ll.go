package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (linkedlist *linkedlist) append(newnode *node) {
	if linkedlist.head == nil {
		linkedlist.head = newnode
		newnode.next = nil
	} else {
		currentnode := linkedlist.head
		for currentnode.next != nil {
			currentnode = currentnode.next
		}
		currentnode.next = newnode
	}
}

func (li *linkedlist) printlist() {
	nn := li.head
	for nn != nil {
		fmt.Printf("value %d at position %v", nn.data, nn.next)
		fmt.Println()
		nn = nn.next
	}

}
func main() {
	l := &linkedlist{}
	l.append(&node{1, nil})
	l.append(&node{2, nil})
	l.append(&node{3, nil})
	l.append(&node{4, nil})
	l.printlist()
}
