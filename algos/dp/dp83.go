package dp

//   ___    ____           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  > _ <   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 198 - 打家劫舍
// https://leetcode.cn/problems/house-robber/description/

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	f := make([]int, n)
	f[0] = nums[0]
	f[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		f[i] = max(nums[i]+f[i-2], f[i-1])
	}
	return f[n-1]
}
