package main

import "fmt"

func main () {
	var word string
	fmt.Println("PLease enter word to be searched")
	fmt.Scanln(&word)
	array := [5]string{"Hello", "there", "how", "are", "you"} 
	for i:=0; i < len(array); i++ {
		if array[i] == word {
			fmt.Println("The word", word, "was found in the array at index", i)
			return
		}
		if i == len(array) - 1 {
			fmt.Println("Word doesn't exist in the loop")
			return
		}
	}
	fmt.Println("Editting golang with vim", array)
}
