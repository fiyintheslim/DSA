package stack

import (
	"fmt"
)

type Node struct {
	value string
	Prev *Node
}

type Stack struct {
	Length int
	Head *Node
}


func (stk *Stack) Push (input string) {
	current := &Node{value: input}

	if stk.Head == nil {
		stk.Head = current
	} else {
		current.Prev = stk.Head
		stk.Head = current
	}
	stk.Length++
}

func (stk *Stack) Pop () {
	if stk.Head == nil {
		fmt.Println("Nothing stack to pop")
	} else {
		stk.Head = stk.Head.Prev
		stk.Length--
	}

}

