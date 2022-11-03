package main

import "fmt"

func SelectionSort(input []int) []int {
	n := len(input)
	result := make([]int, n)
	for i := range input {
		var pos int
		result[i], pos = min(input)
		input[pos] = 99999999
	}
	return result
}

// min returns the min value in the inp array and its position
func min(inp []int) (int, int) {
	min := 9999999
	var posMin int
	for pos, i := range inp {
		if i < min {
			posMin = pos
			min = i
		}
	}
	return min, posMin
}

func main() {
	input := []int{2, 7, 1, 5, 6, 6, 9, -1, 43, 75, 17}
	fmt.Println("Input: ", input)
	fmt.Println("After sorting: ", SelectionSort(input))
}
