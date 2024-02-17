package bs

//    __    ____           ___    __    __    ______   .___________.    __    ___     ___
//   / /   |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_     __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 35 - 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/

func SearchInsert(nums []int, target int) int {
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
