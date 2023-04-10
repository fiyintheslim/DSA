package search

import (
	"fmt"
	"math"
)

func Binary (slc []int, val int) {
	fmt.Println("slice to be searched", slc)
	fmt.Println("Value to search for", val)
	max := len(slc) - 1
	min := 0

	for{
		mid := min + int(math.Round(float64((max - min) / 2)))
		fmt.Println("Max value", max)
		fmt.Println("Mid value", mid)
		fmt.Println("Min value", min)
		fmt.Println()
		if slc[mid] == val {
			fmt.Println("Value found at index", mid)
			break
		} else if slc[mid] > val {
			max = mid
		} else if slc[mid] < val {
			min = mid + 1
		}
		if min > max {
			fmt.Println("Value not found", mid, slc[mid])
			break
		}
	}
}
