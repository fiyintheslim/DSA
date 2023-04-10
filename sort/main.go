package main

import (
	"fmt"
	"DSA/sort_algo/bubble"
)

func main() {
	fmt.Println("Ready and rearing to go")
	slc := []int{2, 3, 4, 2, 1, 3, 5, 7, 9, 4}
	sort.Bubble(slc)
}
