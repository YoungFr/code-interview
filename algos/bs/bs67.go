package bs

//    __    ______          ___    __    __    ______   .___________.    __    ___     ___
//   / /   |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_       / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \     / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |   / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 153 - 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/

func FindMin(nums []int) int {
	// 根据题意
	// 数组在旋转 kn 次后保持不变
	// 其他情况下则会形成两个升序的子数组
	// 比如
	// 2 3 4 5 6 7 | 1
	// 3 4 5 6 7 | 1 2
	// 4 5 6 7 | 1 2 3
	// ...
	// 且满足左边子数组的最小值大于右边子数组的最大值
	// 这就是解决本题的关键

	// 右边子数组的最大值是已知的 => nums[len(nums)-1]
	// 从而数组满足二段性
	// 即前半部分都大于 nums[len(nums)-1] 且
	// 后半部分都小于等于 nums[len(nums)-1]
	// 答案恰好是第一个满足小于等于 nums[len(nums)-1] 的元素

	lo := 0
	hi := len(nums)
	for lo < hi {
		mi := int(uint(lo+hi) >> 1)
		if nums[mi] > nums[len(nums)-1] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return nums[lo]
}