package bitree

//  _  _      ___           ___    __    __    ______   .___________.    __    ___     ___
// | || |    / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | || |_  | | | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |__   _| | | | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    | |   | |_| |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//    |_|    \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 543 - 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/

// 解法一 PASS
// 遍历每个节点并计算以每个节点作为根节点时的最长路径并更新最大值
func DiameterOfBinaryTree1(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 二叉树的最大深度是从根节点到最远叶子节点的最长路径上的节点数
		// 那么以当前节点为根节点的最长路径上的节点数就是左右子树的最大深度再加一
		// 根据题目定义直径是指两个端点之间的边数所以还要减去一

		ans = max(ans, maxDepth(root.Left)+maxDepth(root.Right))
		dfs(root.Left)
		// 放在这里也可以 ^_^
		// ans = max(ans, maxDepth(root.Left) + maxDepth(root.Right))
		dfs(root.Right)
		// 放在这里还可以 ^_^
		// ans = max(ans, maxDepth(root.Left) + maxDepth(root.Right))
	}

	dfs(root)
	return ans
}

// 解法二 PASS
// 把计算最大深度和直径合二为一
func DiameterOfBinaryTree2(root *TreeNode) int {
	ans := 0

	var maxdepth func(root *TreeNode) int
	maxdepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 计算最大深度时会遍历每个节点一次
		// 在遍历到一个节点时
		// 利用后序遍历拿到它的左右子树的最大深度然后计算直径并更新答案

		l := maxdepth(root.Left)
		r := maxdepth(root.Right)

		ans = max(ans, l+r)

		// 二叉树的最大深度
		return max(l, r) + 1
	}

	maxdepth(root)
	return ans
}

//  ____    ______          ___    __    __    ______   .___________.    __    ___     ___
// |___ \  |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) |     / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <     / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |   / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 104 - 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
