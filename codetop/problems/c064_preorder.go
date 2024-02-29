package problems

// LC 144 - 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/

func PreorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0)
	ans := make([]int, 0)
	// 二叉树先序遍历的迭代式算法

	// 1. 将头结点压入栈中
	if root != nil {
		stk = append(stk, root)
	}

	for len(stk) > 0 {
		// 2. 弹出栈顶节点并打印值
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		ans = append(ans, top.Val)

		// N 叉树：从右向左逆序压栈

		// 3. 将栈顶节点的右孩子（不为空）压入栈中
		if top.Right != nil {
			stk = append(stk, top.Right)
		}

		// 4. 将栈顶节点的左孩子（不为空）压入栈中
		if top.Left != nil {
			stk = append(stk, top.Left)
		}
	}

	return ans
}
