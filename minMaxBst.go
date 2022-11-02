// Find the minimum and maximum elements in a Binary Search Tree
package main

import (
	"fmt"

	"github.com/souvikhaldar/go-ds/bst"
)

func main() {
	t := bst.NewIterativeBst()
	t.Insert(5)
	t.Insert(9)
	t.Insert(1)
	t.Insert(1)
	t.Insert(3)
	temp := t.Root
	for temp.Left != nil {
		temp = temp.Left
	}
	// min element is left most leaf
	fmt.Println("Min is: ", temp.Value)
	temp = t.Root
	for temp.Right != nil {
		temp = temp.Right
	}
	// max element is right most leaf
	fmt.Println("Max is: ", temp.Value)
}
