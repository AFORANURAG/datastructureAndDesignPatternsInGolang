package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func NewTree(value int) *TreeNode {
	if value == 0 {
		return nil
	}
	root := &TreeNode{Left: nil, Right: nil, Value: value}
	root.Left = NewTree(value - 1)
	root.Right = NewTree(value - 1)
	return root
}

// root,left and right
func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Println(root.Value)
	InOrderTraversal(root.Right)
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
	fmt.Println(root.Value)
}

//	3
//
// 2  2
// push root
// print
// pop
// check if queue has something if no print new line
// else
// push left if exists
// push right if exists

// [3]
// print and remove
// nothing in queue
// new line
// [2,3,]
// 2

func LevelOrderTraversal(root *TreeNode) {
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)
	// we need an identifier to tell us hey this level is completed,
	// please move to next level
	// this identifier will be reached only after all the levels in the current level has been
	// completed
	// in other words
	// it will be the last TreeNode in a level(just an identifier)
	// so i was thinking about to add an identifier
	// so i was right, i was close, i can doit

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		if temp == nil {
			fmt.Println("\n", "")
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			fmt.Printf(strconv.Itoa(temp.Value) + " ")

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}

	}
}

func main() {
	fmt.Println("<------------Function starts here-------------->")
	// tree := NewTree(3)
	// levelOrderValues := LevelOrder(tree)
	// fmt.Println("levelOrderValues are", levelOrderValues)
	// create a bst
	n := &TreeNode{Value: 10}
	n.Left = &TreeNode{Value: 5}
	n.Left.Left = &TreeNode{Value: 2}
	n.Left.Right = &TreeNode{Value: 7}
	n.Right = &TreeNode{Value: 15}
	n.Right.Left = &TreeNode{Value: 12}
	n.Right.Right = &TreeNode{Value: 17}
	// predecessor := InOrderPredecessor(n, 15)
	InOrderTraversal(n)
	// fmt.Println("predecessor is", predecessor.Value)

}
