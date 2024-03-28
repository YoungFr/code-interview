package array

//  __   ______          ___    __    __    ______   .___________.    __    ___     ___
// /_ | |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | |     / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | |    / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | |   / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_|  /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 41 - 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive/description/

func FirstMissingPositive(nums []int) int {
	// 首先答案一定在 [1, n+1] 中
	// 最坏情况是数组中的元素为 1, 2, 3, ..., n => 答案为 n+1
	// 其他情况答案一定在 [1, n] 中

	// a[0] = 1
	// a[1] = 2
	// ...
	// a[n-1] = n

	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		// 如果 x 在 [1, n] 中
		// 根据上边的推导 x 应该放在 x-1 位置上
		// 所以交换 nums[i] 和 nums[x-1] 元素
		for x >= 1 && x <= n && nums[x-1] != x {
			nums[x-1], nums[i] = nums[i], nums[x-1]
			x = nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
