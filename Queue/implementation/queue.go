package queue

import (
	"fmt"
)

type Node struct {
	value string
	next *Node
}

type Queue struct {
	Length int
	Head *Node
	Tail *Node 
}

func (q *Queue) Peek () *Node {
	return q.Head
}

func (q *Queue) Enqueue (item string) {
	newNode := &Node{value: item}
	if q.Head == nil {
		fmt.Println("Head is undefined")
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.next = newNode
		q.Tail = newNode
	}
	q.Length++

}

func (q *Queue) Deque () *Node {
	if q.Head == nil {
		fmt.Println("Nothing in the que to deque")
		return nil
	} else {
		head := q.Head
		q.Head = head.next
		q.Length--
		return head
	}

}
