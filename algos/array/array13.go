package array

//  __   ____           ___    __    __    ______   .___________.    __    ___     ___
// /_ | |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | |  |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_| |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 53 - 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/description/

func MaxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < n; i++ {
		f[i] = max(nums[i], nums[i]+f[i-1])
		ans = max(ans, f[i])
	}

	// 状态压缩写法
	// fi := nums[0]
	// ans := fi
	// for i := 1; i < len(nums); i++ {
	//     fi = max(nums[i], nums[i] + fi)
	//     ans = max(ans, fi)
	// }

	return ans
}
