package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp != nil {
				level = append(level, temp.Val)
				if temp.Left != nil {
					queue = append(queue, temp.Left)
				}
				if temp.Right != nil {
					queue = append(queue, temp.Right)

				}
			}
		}
		if len(level) > 0 {
			result = append(result, level)
		}

	}
	return result
}
