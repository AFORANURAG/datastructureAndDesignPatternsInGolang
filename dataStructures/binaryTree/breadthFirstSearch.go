package main

func LevelOrder(root *Node) [][]int {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	levelValues := make([][]int, 0)
	for len(queue) > 0 {
		curr := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			tempNode := queue[0]
			queue = queue[1:]
			if tempNode != nil {
				curr = append(curr, tempNode.Value)
				if tempNode.Left != nil {
					queue = append(queue, tempNode.Left)
				}
				if tempNode.Right != nil {
					queue = append(queue, tempNode.Right)
				}
			}
		}
		levelValues = append(levelValues, curr)

	}
	return levelValues
}
