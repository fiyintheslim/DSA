package sort

import "fmt"

func Bubble (arr []int) []int {
	
	for i:=0; i < len(arr); i++ {
		fmt.Println("ready to bubble sort", arr[i])
		for j:=0; j < (len(arr) - 1 - i); j++ {
			if arr[j] > arr[j + 1] {
				bigger := arr[j]
				smaller := arr[j + 1]

				arr[j] = smaller
				arr[j + 1] = bigger 
			}
		}
	}
	fmt.Println("BUbble sorted list", arr)
	return arr
}
