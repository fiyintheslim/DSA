package main

import (
	"fmt"
	"DSA/stack/stack_impl"
)

func main () {
	var input string
	stk := stack.Stack{}
	for {
		fmt.Println("Please enter value to be added to the stack")
		fmt.Println("exit!: To close  traverse!: Look through the queue destack!: Take off stack")
		fmt.Scanln(&input)
		if input == "exit!" {
			break
		} else if input == "traverse!" {
			fmt.Println("Traverse operation")
			continue
		} else if input == "destack!" {
			fmt.Println("Destack operation")
			continue
		}
		stk.Push(input)
		fmt.Println("", stk)
	}
}
