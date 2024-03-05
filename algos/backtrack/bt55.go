package backtrack

import (
	"slices"
	"sort"
)

//  _____   _____          ___    __    __    ______   .___________.    __    ___     ___
// | ____| | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | |__   | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |___ \  |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 46 - 全排列
// https://leetcode.cn/problems/permutations/description/

// 利用下一个更大排列的解法
func Permute2(nums []int) (ans [][]int) {
	sort.Ints(nums)

	dst := make([]int, len(nums))
	copy(dst, nums)
	slices.Reverse(dst)

	ans = append(ans, append([]int(nil), nums...))
	for !slices.Equal(nums, dst) {
		next(nums)
		ans = append(ans, append([]int(nil), nums...))
	}

	return
}

func next(nums []int) {
	n := len(nums)
	i := n - 1
	for i >= 1 {
		if nums[i] > nums[i-1] {
			break
		}
		i--
	}
	if i == 0 {
		slices.Reverse(nums)
		return
	}
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
