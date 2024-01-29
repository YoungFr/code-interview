package skills

//   ___    ______          ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) |     / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  \__, |    / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    / /    / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//   /_/    /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// 169 - 多数元素
// https://leetcode.cn/problems/majority-element/description/

func MajorityElement(nums []int) int {
	candidate := -1
	votes := 0

	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			votes++
		} else {
			votes--
		}
	}

	return candidate
}
