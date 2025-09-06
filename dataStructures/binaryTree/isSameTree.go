package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	q1p := make([]*TreeNode, 0)
	q2q := make([]*TreeNode, 0)
	q1p = append(q1p, p)
	q2q = append(q2q, q)
	if p.Value != q.Value {
		return false
	}

	for len(q1p) > 0 && len(q2q) > 0 {
		l1 := len(q1p)
		l2 := len(q2q)
		c1 := make([]int, 0)
		c2 := make([]int, 0)
		for i := 0; i < l1; i++ {
			tempNodeP := q1p[0]
			q1p = q1p[1:]
			if tempNodeP != nil {
				c1 = append(c1, tempNodeP.Value)
				if tempNodeP.Left != nil {
					q1p = append(q1p, tempNodeP.Left)
				}
				if tempNodeP.Right != nil {
					q1p = append(q1p, tempNodeP.Right)
				}

			}
		}
		for i := 0; i < l2; i++ {
			tempNodeQ := q2q[0]
			q2q = q2q[1:]
			if tempNodeQ != nil {
				c2 = append(c2, tempNodeQ.Value)
				if tempNodeQ.Left != nil {
					q2q = append(q2q, tempNodeQ.Left)
				}
				if tempNodeQ.Right != nil {
					q2q = append(q2q, tempNodeQ.Right)
				}
			}
		}
		if len(c1) != len(c2) {
			return false
		}
		for i := 0; i < len(c2); i++ {
			if c1[i] != c2[i] {
				return false
			}
		}

	}
	return true

}
