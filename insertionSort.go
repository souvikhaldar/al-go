package main

import "fmt"

func InsertionSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	sortedPos := 0
	for i := sortedPos + 1; i < len(input); i++ {
		temp := input[i]
		for j := sortedPos; j >= 0; j-- {
			if temp < input[j] {
				input[j+1] = input[j]
				input[j] = temp
			}
		}
		sortedPos++
	}
	return input
}

func main() {
	inp := []int{-3, -6, 4, 2, -9}
	fmt.Println("Input: ", inp)
	fmt.Println("After sorting: ", InsertionSort(inp))
}
