package skills

import "sort"

//   ___     ___           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \   / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  \__, |  \__, |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    / /     / /      /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//   /_/     /_/      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 31 - 下一个排列
// https://leetcode.cn/problems/next-permutation/description/

func NextPermutation(nums []int) {
	// 如何做到更大：
	// 1. 将一个数字 i 和它右侧的一个大于它的数 j 交换
	// 如何做到下一个：
	// 1. 交换的位置要尽可能靠右
	// 2. 数字 j 应该是所有大于 i 的数中最小的

	n := len(nums)

	// 从后往前查找
	i := n - 1
	for i >= 1 {
		if nums[i] > nums[i-1] {
			break
		}
		i--
	}

	// 整个排列是降序 => 恢复成升序
	if i == 0 {
		sort.Ints(nums)
		return
	}

	// 要被交换的数字是 nums[i-1] 且它右侧的数是降序排列的
	// 从后往前找到第一个大于 nums[i-1] 的数
	j := n - 1
	for j >= i {
		if nums[j] > nums[i-1] {
			break
		}
		j--
	}

	nums[i-1], nums[j] = nums[j], nums[i-1]
	sort.Ints(nums[i:])
}
