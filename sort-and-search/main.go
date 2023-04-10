package main

import (
	"fmt"
	"DSA/sort-and-search/sort"
	"DSA/sort-and-search/search"
)

func main (){
	slice := []int{23, 2, 3, 45, 62, 3, 4, 5, 6, 7, 8, 99, 9, 0, 1}
	sorted := sort.Bubble(slice)
	fmt.Println("Not sorted", slice)
	fmt.Println("Sorted", sorted)
	search.Binary(slice, 99)
}
