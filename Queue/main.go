package main

import (
	"fmt"
	"DSA/Queue/implementation"
)


func main () {
	var input string
	q := queue.Queue{}
	for {
		fmt.Println("Enter value into queue")
		fmt.Println("exit: To exit  !peek: To peek into the first node/head  !deque: To remove from queue")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		} else if input == "!peek" {
			val := q.Peek()
			fmt.Println("Peeked value", val)
			continue
		} else if input == "!deque" {
			val := q.Deque()
			fmt.Println("Dequed node", val)
			continue
		}
		q.Enqueue(input)
		fmt.Println("Queue ready to go:", input, q)
	}
}
