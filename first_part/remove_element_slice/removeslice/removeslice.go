package removeslice

//Remove is function to remvoe element of slice base on it`s index
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
