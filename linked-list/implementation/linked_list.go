package linkedList

import "fmt"

type Node struct {
	value string
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

func (ll *LinkedList) AddNode(val string){
	newNode := &Node{value: val}
	if ll.length == 0 {
		ll.head = newNode
		ll.tail = newNode
		
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
	ll.length++
}

func (ll *LinkedList) Length () int {
	fmt.Println("Linked list length", ll.length)
	return ll.length
}

func (ll *LinkedList) Traverse () {
	current := ll.head
	for {
		 if current == nil || current.next == nil {

			fmt.Println("No node in the linked list", current)
			break
		}
		fmt.Println("current", current.value)

		current = current.next
	}
}
