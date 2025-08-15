package main

import (
	"fmt"
)

// it should be generic
type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (n *Node[T]) BuildTree(root *Node[T], data T, zero T) *Node[T] {
	// take input from user

	// Since T is a generic type, we can't directly compare with nil
	// Instead, we'll return a new node with the scanned value

	if data == zero {
		return &Node[T]{Value: data, Left: nil, Right: nil}
	}
	node := &Node[T]{Value: data, Left: nil, Right: nil}

	fmt.Println("Enter the value of the left node")
	var value T
	fmt.Scanln(&value)
	node.Left = n.BuildTree(node.Left, value, zero)

	fmt.Println("Enter the value of the right node")
	fmt.Scanln(&value)
	node.Right = n.BuildTree(node.Right, value, zero)

	return node
}

func (n *Node[T]) Delete(value T) *Node[T] {
	return nil
}

func (n *Node[T]) LevelOrderTraversal(root *Node[T]) {
	// take a queue
	queue := []*Node[T]{}
	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Println("\n", "")
			// if queue is has, append nil
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			fmt.Printf("%v ", node.Value)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

}

func (n *Node[T]) InOrderTraversal(root *Node[T]) {
	// l n r
	if root == nil {
		return
	}
	n.InOrderTraversal(root.Left)
	fmt.Println(root.Value)
	n.InOrderTraversal(root.Right)

}

func (n *Node[T]) PreOrderTraversal(root *Node[T]) {
	// n l r
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	n.PreOrderTraversal(root.Left)
	n.PreOrderTraversal(root.Right)

}

func (n *Node[T]) PostOrderTraversal(root *Node[T]) {
	// l r n
	if root == nil {
		return
	}
	n.PostOrderTraversal(root.Left)
	n.PostOrderTraversal(root.Right)
	fmt.Println(root.Value)

}

func main() {
	tree := &Node[int]{}
	root := tree.BuildTree(nil, 10, -1)
	tree.LevelOrderTraversal(root)
	fmt.Println("Hello world")
}
