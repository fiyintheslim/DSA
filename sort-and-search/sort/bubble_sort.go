package sort


func Bubble(slc []int) []int {
	for i:=0; i < len(slc); i++ {
		for j:=0; j < (len(slc) - 1 - i); j++{
			next := slc[j + 1]
			if slc[j] > next {
				slc [j + 1] = slc[j]
				slc[j] = next
			}
		}
	}
	return slc
}
