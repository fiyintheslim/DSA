package main

import (
	"fmt"
	"DSA/linked_list/implementation"
)

func main () {
	var input string
	ll := linkedList.LinkedList{}

	for {
		fmt.Println("Enter value to be added to the linked list:")
		fmt.Scanln(&input)
		fmt.Println()
		if input == "exit"{
			break
		} else if input == "traverse" {
			ll.Traverse()
			continue
		}
		fmt.Println("You entered ", input)
		ll.AddNode(input)
		ll.Length()
	}
}
