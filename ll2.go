package main

import "fmt"

// Node can store two values, 'id' and 'name'.
// Another value, 'ptr' is a pointer to another node
type Node struct {
	id   int
	name string
	ptr  *Node
}

// LinkedList struct
type LinkedList struct {
	head *Node
}

// LinkedList method to append a node to tail
func (linkedlist *LinkedList) append(newnode *Node) *LinkedList {
	if linkedlist.head == nil {
		linkedlist.head = newnode
		newnode.ptr = nil
		return linkedlist
	} // else
	// initialization; condition; increment
	for n := linkedlist.head; n != nil; n = n.ptr {
		if n.ptr == nil {
			n.ptr = newnode
			return linkedlist
		}
	}
	return linkedlist
}

// LinkedList method to print the nodes details
func (linkedlist *LinkedList) print() {
	for n := linkedlist.head; n != nil; n = n.ptr {
		fmt.Println(n.id, n.name, n.ptr)
	}
}

// LinkedList method to insert a new node before a particular node
func (linkedlist *LinkedList) insertBefore(name string, newNode *Node) bool {
	if linkedlist.head == nil {
		return false
	}

	ptrNode := linkedlist.head // points to first node
	for ptrNode.ptr != nil {   // while() loop
		if ptrNode.ptr.name == name {
			newNode.ptr = ptrNode.ptr // new node points to the next node
			ptrNode.ptr = newNode     // current node points to the new node
			return true
		}
		ptrNode = ptrNode.ptr // points to next node
	}
	return false
}

func main() {
	// create a linked list object
	linkedlist := &LinkedList{}

	// create nodes
	nodeA := &Node{1, "A", nil}
	nodeB := &Node{2, "B", nil}
	nodeC := &Node{3, "C", nil}
	nodeD := &Node{4, "D", nil}
	nodeE := &Node{5, "E", nil}

	// append the nodes to the linked list
	linkedlist.append(nodeA)
	linkedlist.append(nodeB)
	linkedlist.append(nodeC)
	linkedlist.append(nodeD)
	linkedlist.append(nodeE)

	// print the linked list
	linkedlist.print()

	// create a new node
	nodeF := &Node{6, "F", nil}
	// insert the new node before node "B"
	linkedlist.insertBefore("B", nodeF)
	// print the linked list
	linkedlist.print()
}
