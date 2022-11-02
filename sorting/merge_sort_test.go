package sorting

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	inputs := [][]int{
		{
			3, 2, 8, 5, 1, 3, 8, 4,
		},
		{
			10, 56, 12, 87, 45, 12, 89, 93,
		},
	}
	outputs := [][]int{
		{
			1, 2, 3, 3, 4, 5, 8, 8,
		},
		{
			10, 12, 12, 45, 56, 87, 89, 93,
		},
	}
	sorted := make([][]int, len(inputs))
	for i := range inputs {
		fmt.Println("Input to sorting: ", inputs[i])
		st, err := MergeSort(inputs[i])
		if err != nil {
			t.Fatal("Error in sorting: ", err)
		}
		sorted[i] = st
		fmt.Println("After sorting: ", sorted[i])
	}
	for input := range inputs {
		for i := range inputs[input] {
			if inputs[input][i] != outputs[input][i] {
				t.Fatalf("Got %d, Expected: %d ", inputs[input][i], outputs[input][i])
			}
		}
	}
}
