package problems

// LC 103 - 二叉树的锯齿形层序遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func ZigzagLevelOrder(root *TreeNode) (ans [][]int) {
	var leftToRight bool = true

	// 相邻的两层的遍历顺序相反
	// 使用栈来实现这种效果
	l2r := make([]*TreeNode, 0) // left to right
	r2l := make([]*TreeNode, 0) // right to left

	// 根节点是从左往右遍历
	if root != nil {
		l2r = append(l2r, root)
	}
	for len(l2r) > 0 || len(r2l) > 0 {
		level := make([]int, 0)
		if leftToRight {
			for sz := len(l2r); sz > 0; sz-- {
				cur := l2r[len(l2r)-1]
				l2r = l2r[:len(l2r)-1]
				level = append(level, cur.Val)

				// 先左孩子后右孩子
				if cur.Left != nil {
					r2l = append(r2l, cur.Left)
				}
				if cur.Right != nil {
					r2l = append(r2l, cur.Right)
				}
			}
		} else {
			for sz := len(r2l); sz > 0; sz-- {
				cur := r2l[len(r2l)-1]
				r2l = r2l[:len(r2l)-1]
				level = append(level, cur.Val)

				// 先右孩子后左孩子
				if cur.Right != nil {
					l2r = append(l2r, cur.Right)
				}
				if cur.Left != nil {
					l2r = append(l2r, cur.Left)
				}
			}
		}
		ans = append(ans, level)
		leftToRight = !leftToRight
	}
	return
}
