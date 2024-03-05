package problems

// LC 662 - 二叉树的最大宽度
// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/

func WidthOfBinaryTree(root *TreeNode) int {
	// 完全二叉树
	// 根节点 i
	// 左孩子 2 * i + 1
	// 有孩子 2 * i + 2

	// 节点数目是 [1, 3000]

	q := make([]item, 0)
	q = append(q, item{root, 0})

	ans := 1

	for len(q) > 0 {
		// 先计算本层的最大宽度
		ans = max(ans, q[len(q)-1].i-q[0].i+1)

		// 层序遍历
		for sz := len(q); sz > 0; sz-- {
			t := q[0]
			q = q[1:]
			if t.n.Left != nil {
				q = append(q, item{t.n.Left, 2*t.i + 1})
			}
			if t.n.Right != nil {
				q = append(q, item{t.n.Right, 2*t.i + 2})
			}
		}
	}

	return ans
}

type item struct {
	n *TreeNode
	i int
}
