package main

import "fmt"

func insertionSortArray(list [3]int) ([3]int, int) {
	var i, j int
	for i = 1; i < len(list); i++ {
		j = i
		for ; j > 1 && list[j] < list[j-1]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}
	return list, len(list)
}

func insertionSortSlice(list []int) ([]int, int) {
	var i, j int
	for i = 1; i < len(list); i++ {
		j = i
		for ; j > 1 && list[j] < list[j-1]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}
	return list, len(list)
}

func main() {
	// this is array
	fmt.Println("ARRAY")
	inputArray := [3]int{1, 12, 2}
	fmt.Println("Original                     - ", inputArray, len(inputArray))
	outputArray, outputArrayLength := insertionSortArray(inputArray)
	fmt.Println("Original after sort function - ", inputArray, len(inputArray))
	fmt.Println("Modified                     - ", outputArray, outputArrayLength)

	// this is slice
	fmt.Println("SLICE")
	inputSlice := []int{1, 12, 2}
	fmt.Println("Original                     - ", inputSlice, len(inputSlice))
	outputSlice, outputSliceLength := insertionSortSlice(inputSlice)
	fmt.Println("Original after sort function - ", inputSlice, len(inputSlice), "WTF??")
	fmt.Println("Modified                     - ", outputSlice, outputSliceLength)
}
