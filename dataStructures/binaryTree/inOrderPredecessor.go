package main

// Insert into BST
func insert(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: key}
	}
	if key < root.Value {
		root.Left = insert(root.Left, key)
	} else {
		root.Right = insert(root.Right, key)
	}
	return root
}

// Find Predecessor and Successor
func findPreSuc(root *TreeNode, key int) (pre, suc *TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Value == key {
		// Case 1: Left subtree → predecessor
		if root.Left != nil {
			temp := root.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			pre = temp
		}

		// Case 2: Right subtree → successor
		if root.Right != nil {
			temp := root.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			suc = temp
		}
		return pre, suc
	}

	// Case 3: search down the tree
	if key < root.Value {
		// root could be successor
		suc = root
		p, s := findPreSuc(root.Left, key)
		if p != nil {
			pre = p
		}
		if s != nil {
			suc = s
		}
	} else {
		// root could be predecessor
		pre = root
		p, s := findPreSuc(root.Right, key)
		if p != nil {
			pre = p
		}
		if s != nil {
			suc = s
		}
	}
	return pre, suc
}
