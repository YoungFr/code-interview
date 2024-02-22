package bs

//    __      __           ___    __    __    ______   .___________.    __    ___     ___
//   / /     / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_    / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \  | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 33 - 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func Search(nums []int, target int) int {
	// 153 题的进阶
	// 数组在旋转后会形成两个升序的子数组
	// 先用二分查找找到数组中最小值的下标
	// 根据 target 和 nums[-1] 的大小关系
	// 确定其在（如果存在）哪个子数组中
	// 然后再使用一次二分查找

	// 数组的最小值是第一个小于等于 nums[-1] 的元素
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

	// 现在 nums[0...lo-1] 和 nums[lo...len(nums)-1] 是两个升序子数组
	if target > nums[len(nums)-1] {
		lo, hi = 0, lo-1
	} else {
		// lo = lo
		hi = len(nums) - 1
	}

	// nums[lo...hi] 是一个升序数组
	for lo <= hi {
		mi := int(uint(lo+hi) >> 1)
		switch {
		case target < nums[mi]:
			hi = mi - 1
		case target > nums[mi]:
			lo = mi + 1
		case target == nums[mi]:
			return mi
		}
	}
	return -1
}
