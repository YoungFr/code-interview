package skills

//   ___     ___           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \   / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  \__, |  > _ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    / /  | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//   /_/    \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 75 - 颜色分类
// https://leetcode.cn/problems/sort-colors/description/

// 三向切分算法
// a[l...lt-1] 中的元素严格小于 pivot
// a[lt...gt]  中的元素全部等于 pivot
// a[gt+1...r] 中的元素严格大于 pivot
//
// func threeWayPartition(a []int, l int, r int) (int, int) {
// 	pivot := a[r]
// 	lt, gt := l, r-1
// 	for i := lt; i <= gt; i++ {
// 		if a[i] > pivot {
// 			for i <= gt && a[i] > pivot {
// 				a[i], a[gt] = a[gt], a[i]
// 				gt--
// 			}
// 		}
// 		if a[i] < pivot {
// 			a[i], a[lt] = a[lt], a[i]
// 			lt++
// 		}
// 	}
// 	a[gt+1], a[r] = a[r], a[gt+1]
// 	return lt, gt + 1
// }

// 以 1 为主元运行三向切分算法
func SortColors(nums []int) {
	lt := 0
	gt := len(nums) - 1
	for i := lt; i <= gt; i++ {
		if nums[i] > 1 {
			for i <= gt && nums[i] > 1 {
				nums[i], nums[gt] = nums[gt], nums[i]
				gt--
			}
		}
		if nums[i] < 1 {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
		}
	}
}
