package array

//  __     __           ___    __    __    ______   .___________.    __    ___     ___
// /_ |   / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | |  / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | | | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_|  \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 238 - 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/description/

func ProductExceptSelf(nums []int) []int {
	// 保证数组中任意元素的全部前缀元素和后缀的乘积都在 32 位整数范围内
	n := len(nums)

	// 前缀积
	// 元素 prep[i] 表示 nums[i] 之前（不包括 nums[i] 本身）的所有元素的积
	prep := make([]int, n)
	prep[0] = 1
	for i := 1; i < n; i++ {
		prep[i] = prep[i-1] * nums[i-1]
	}

	// 后缀积
	// 元素 prep[i] 表示 nums[i] 之后（不包括 nums[i] 本身）的所有元素的积
	sufp := make([]int, n)
	sufp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		sufp[i] = sufp[i+1] * nums[i+1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = prep[i] * sufp[i]
	}
	return ans
}
