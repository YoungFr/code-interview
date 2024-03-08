package backtrack

func CombinationSum1(candidates []int, target int) (ans [][]int) {
	n := len(candidates)

	var dfs func(i int, target int, a []int)
	dfs = func(i int, target int, a []int) {
		if i == n {
			if target == 0 {
				ans = append(ans, append([]int(nil), a...))
			}
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), a...))
			return
		}

		if target < 0 {
			return
		}

		// 选择了 candidates 的第 i 个元素
		a = append(a, candidates[i])
		// 第 i 个元素可以被重复选取
		dfs(i, target-candidates[i], a)
		a = a[:len(a)-1]

		// 未选择 candidates 的第 i 个元素
		dfs(i+1, target, a)
	}

	dfs(0, target, make([]int, 0))
	return
}
