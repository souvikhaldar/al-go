package sorting

import (
	"errors"
	"fmt"
)

// Reference- https://youtu.be/TzeBrDU-JaY
func MergeSort(input []int) ([]int, error) {
	if len(input) == 0 {
		err := errors.New("Input is empty")
		return nil, err
	}
	divide(input)
	return input, nil
}

func divide(input []int) {
	// a list with one element is already sorted
	iLen := len(input)
	if iLen == 1 {
		return
	}
	// divide the list into two sub lists
	mid := int(iLen / 2)
	left := input[:mid]
	right := input[mid:]
	fmt.Println("Left: ", left)
	fmt.Println("Right: ", right)
	// now divide each sublists iteratively
	divide(left)
	divide(right)
	// merge the two divided sublists in sorted manner
	input = make([]int, len(left), len(right))
	merge(left, right, input)
	fmt.Println("Left and right merged: ", input)
}

func merge(left, right []int, input []int) {
	iLen := len(left) + len(right)
	var lPos, rPos int // positions of left and right sublist

	// traversing entire length of sorted final list
	for iPos := 0; iPos < iLen; iPos++ {
		// reached the end of left sublist already
		if lPos == len(left) {
			// copy all elements of remaining right sublist
			// to the remaining part of the final sorted list
			input = append(input[:iPos], right[rPos:]...)
			// now we can return since we are done traversing
			// both the sublists, one was already traversed and
			// other we just copied over in above step
			return
		}

		// reached the end of right sublist already
		if rPos == len(right) {
			// copy all elements remaining in the left sublist
			// to the remaining part of the final sorted list
			input = append(input[:iPos], left[lPos:]...)
			return
		}

		if left[lPos] <= right[rPos] {
			input[iPos] = left[lPos]
			lPos++
		} else {
			input[iPos] = right[rPos]
			rPos++
		}
	}

	return
}
