package problems

// LC 129 - 求根到叶子节点数字之和
// https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/

func SumNumbers(root *TreeNode) int {
	sum := 0

	var dfs func(root *TreeNode, val int)
	dfs = func(root *TreeNode, val int) {
		// dfs 框架部分
		if root == nil {
			return
		}

		// 从根节点到当前节点的路径所表示的数字
		v := 10*val + root.Val
		// 叶节点
		if root.Left == nil && root.Right == nil {
			sum += v
			return
		}

		// dfs 框架部分
		dfs(root.Left, v)
		dfs(root.Right, v)
	}

	dfs(root, 0)
	return sum
}
