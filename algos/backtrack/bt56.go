package backtrack

//  _____     __           ___    __    __    ______   .___________.    __    ___     ___
// | ____|   / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | |__    / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |___ \  | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 78 - 子集
// https://leetcode.cn/problems/subsets/description/

func Subsets(nums []int) (ans [][]int) {
	n := len(nums)

	// 子集型回溯的典型例题
	// 每个元素都有选或不选两种可能

	// 参数 i 表示当前在对 nums 中的第 i 个元素做出选择
	// 参数 set 保存做完选择后集合中包含的元素
	var dfs func(i int, set []int)

	dfs = func(i int, set []int) {
		// 做出了 n 次选择
		if i == n {
			// 这里要把当前 set 中的元素全部拷贝到一个新的切片中
			// 否则后边 set 中的元素会被更改破坏
			ans = append(ans, append([]int(nil), set...))
			return
		}

		// 选择了第 i 个元素
		set = append(set, nums[i])
		dfs(i+1, set)
		set = set[:len(set)-1]

		// 不选择第 i 个元素
		dfs(i+1, set)
	}

	dfs(0, make([]int, 0))
	return
}
