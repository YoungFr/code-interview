package problems

// LC 113 - 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/

func PathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var dfs func(root *TreeNode, targetSum int, path []int)

	dfs = func(root *TreeNode, targetSum int, path []int) {
		// 先序 dfs 框架部分
		if root == nil {
			return
		}

		// 把当前访问到的节点保存到路径中
		path = append(path, root.Val)

		// 叶节点
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			ans = append(ans, append([]int(nil), path...)) // 路径的副本
			return
		}

		// 先序 dfs 框架部分
		dfs(root.Left, targetSum-root.Val, path)
		dfs(root.Right, targetSum-root.Val, path)
	}

	dfs(root, targetSum, make([]int, 0))
	return
}
