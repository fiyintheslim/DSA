package main

import (
	"fmt"
	"DSA/binary-search/search"
)


func main () {
	array := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	found, index := BS.BinarySearch(array, 13)
	fmt.Println(found, index)
}

