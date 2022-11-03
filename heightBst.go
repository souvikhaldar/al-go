// Find the height a of Binary Search Tree.

// The height of a binary search tree is the maximum number of edges from root node to the leaf node

package main

import (
	"fmt"

	"github.com/souvikhaldar/go-ds/bst"
	"github.com/souvikhaldar/go-ds/queue"
	"github.com/souvikhaldar/go-ds/stack"
)

func main() {
	b := bst.NewIterativeBst()
	b.Insert(4)
	b.Insert(2)
	b.Insert(1)
	b.Insert(0)
	b.Insert(5)
	b.Insert(6)
	b.Insert(7)
	b.Insert(8)
	b.Insert(9)
	temp := b.Root
	if temp == nil {
		fmt.Println(bst.ErrEmptyTree)
		return
	}

	heights := make([]int, 0)
	var edge int
	q := queue.NewQueue()
	s := stack.NewStack()
	visited := make(map[int]struct{})

	q.Enqueue(temp)

	for !q.IsEmpty() {
		t, err := q.Dequeue()
		if err != nil {
			break
		}
		temp, ok := t.(*bst.Node)
		if !ok {
			break
		}
		_, ok = visited[temp.Value]
		if !ok {
			if temp.Value != b.Root.Value {
				visited[temp.Value] = struct{}{}
				edge++
			}
		} else {
			t, err := q.Dequeue()
			if err != nil {
				break
			}
			temp, ok = t.(*bst.Node)
			if !ok {
				break
			}

		}

		for temp != nil || !s.IsEmpty() {
			if _, ok := visited[temp.Value]; ok {
				t, err := s.Pop()
				if err != nil {
					break
				}
				temp, ok = t.(*bst.Node)
				if !ok {
					break
				}
			}
			if bst.IsALeafNode(temp) {
				heights = append(heights, edge)
				t, err := q.Dequeue()
				if err != nil {
					break
				}
				temp, ok = t.(*bst.Node)
				if !ok {
					break
				}
				// traverse a new path to leaf
				edge = 0

			} else if bst.IsNodeWithOneChild(temp) {
				edge++

				// go either left sub-tree or right
				if temp.Left != nil {
					temp = temp.Left
				} else {
					temp = temp.Right
				}

			} else {
				edge++
				// go the left subtree and hold the right
				temp = temp.Left
			}
		}
	}

	max := 0
	for _, i := range heights {
		if i > max {
			max = i
		}
	}

	fmt.Println("Height of the bst is: ", max)
}
