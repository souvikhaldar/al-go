package main

import "fmt"

func BubbleSort(inp []int) []int {
	n := len(inp)
	m := n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if inp[j] > inp[j+1] {
				inp[j], inp[j+1] = inp[j+1], inp[j]
			}
		}
		m--
	}
	return inp
}

func main() {
	inp := []int{4, 3, 6, 7, 7, 1, 9, 3, 76, 89, 12, 4, -1, 8}
	fmt.Println("Input: ", inp)
	fmt.Println("After sorting: ", BubbleSort(inp))
}
