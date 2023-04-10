package BS

import (
	"fmt"
	"math"
)

func BinarySearch (arr []int, val int) (bool, int) {
	lo := 0
	hi := len(arr) - 1

	for {
		notRounded := float64((lo + hi)/2)
		mid := int(math.Floor(notRounded))
		fmt.Println("low:", lo, "Mid:", mid, "High:", hi)
		if arr[mid] == val {
			return true, mid 
		} else if val < arr[mid] {
			hi = mid
			continue
		} else if val > arr[mid] {
			lo = mid + 1
			continue
		}

		if hi < lo {
			return false, -1
		}
	}
}

