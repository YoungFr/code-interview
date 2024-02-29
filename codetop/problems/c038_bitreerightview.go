package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LC 199 - 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/description/

func RightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		// 在把下一层节点添加到队列中时先将当前层最右边的节点添加到 ans 中
		ans = append(ans, q[len(q)-1].Val)

		// 二叉树的层序遍历
		for sz := len(q); sz > 0; sz-- {
			// 弹出队头元素
			h := q[0]
			q = q[1:]
			// 头节点不为空的左右孩子入队
			if h.Left != nil {
				q = append(q, h.Left)
			}
			if h.Right != nil {
				q = append(q, h.Right)
			}
		}
	}
	return ans
}
