package bitree

import "math"

//  _  _     ____           ___    __    __    ______   .___________.    __    ___     ___
// | || |   |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | || |_    __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |__   _|  |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    | |    ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//    |_|   |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 98 - 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/

func IsValidBST(root *TreeNode) bool {
	// 使用迭代式中序遍历算法判断一个二叉树是否是有效的二叉搜索树
	// 有效的二叉搜索树的中序遍历结果是一个单调递增的序列
	stk := make([]*TreeNode, 0)
	var (
		last = math.MinInt64
		curr = 0
	)
	for len(stk) > 0 || root != nil {
		if root != nil {
			// 当前节点不为空时压入当前节点
			// 并将它的左孩子设置为当前节点
			stk = append(stk, root)
			root = root.Left
		} else {
			// 当前节点为空时栈顶节点就是当前遍历到的值
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			curr = top.Val
			// 当前的值必须大于上一个值
			if curr <= last {
				return false
			}
			// 保存当前遍历到的值
			last = curr
			// 将栈顶节点的右孩子设置为当前节点
			root = top.Right
		}
	}
	return true
}
