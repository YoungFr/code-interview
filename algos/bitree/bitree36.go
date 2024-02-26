package bitree

//  ____      __           ___    __    __    ______   .___________.    __    ___     ___
// |___ \    / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) |  / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <  | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

//  __  .__   __.   ______   .______       _______   _______ .______
// |  | |  \ |  |  /  __  \  |   _  \     |       \ |   ____||   _  \
// |  | |   \|  | |  |  |  | |  |_)  |    |  .--.  ||  |__   |  |_)  |
// |  | |  . `  | |  |  |  | |      /     |  |  |  ||   __|  |      /
// |  | |  |\   | |  `--'  | |  |\  \----.|  '--'  ||  |____ |  |\  \----.
// |__| |__| \__|  \______/  | _| `._____||_______/ |_______|| _| `._____|

// LC 94 - 二叉树中序遍历的迭代式算法
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

func InorderTraversal1(root *TreeNode) []int {
	stk := make([]*TreeNode, 0)
	ans := make([]int, 0)
	for len(stk) > 0 || root != nil {
		if root != nil {
			// 1. 压入根节点。然后从根节点开始，不断压入当前节点的左孩子
			stk = append(stk, root)
			root = root.Left
		} else {
			// 2. 当左孩子为空时，弹出栈顶元素并打印值
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans = append(ans, top.Val)
			// 3. 然后将栈顶元素的右孩子设为当前节点，重复步骤 1 的操作
			root = top.Right
		}
	}
	return ans
}

// LC 94 - 二叉树中序遍历的递归式算法
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

func InorderTraversal2(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
