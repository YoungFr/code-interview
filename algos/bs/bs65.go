package bs

//    __    _____          ___    __    __    ______   .___________.    __    ___     ___
//   / /   | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_   | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \  |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 34 - 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func SearchRange(nums []int, target int) []int {
	ans := []int{-1, -1}

	// 第一个大于等于 target 的元素的下标
	a := bsearch(nums, target)

	// 数组中的所有元素都比 target 小
	if a >= len(nums) {
		return ans
	}

	if nums[a] != target {
		// 数组中只存在第一个大于 target 而不等于 target 的元素
		return ans
	} else {
		// 数组中存在等于 target 的元素
		// 下标 a 是开始位置
		// 下标 b 是结束位置
		b := bsearch(nums, target+1) - 1
		ans[0], ans[1] = a, b
		return ans
	}
}

func bsearch(nums []int, target int) int {
	// 找到数组中第一个大于等于 target 的元素的下标
	lo := 0
	hi := len(nums)
	for lo < hi {
		mi := int(uint(lo+hi) >> 1)
		if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}
