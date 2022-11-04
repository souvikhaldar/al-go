package main

import "fmt"

func SelectionSort(inp []int) []int {
	for i := 0; i < len(inp); i++ {
		inter := inp[i:]
		pos, min := minimum(inter)
		inp[i], inp[pos+i] = min, inp[i]
	}
	return inp
}
func minimum(inp []int) (pos, val int) {
	min := 999999
	for p, i := range inp {
		if i < min {
			min = i
			pos = p
		}
	}
	return pos, min
}

func main() {
	inp := []int{4, 6, 2, 8, 1, 9, 5, 4, 4, 76, 123, 87, 9, -4, -2, 34}
	fmt.Println("Input: ", inp)
	fmt.Println("After sorting: ", SelectionSort(inp))
}
