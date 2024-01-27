package skills

//   ___     ___           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \   / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  \__, |  > _ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    / /  | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//   /_/    \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 75 - 颜色分类
// https://leetcode.cn/problems/sort-colors/description/

// 随机选择主元的三向切分算法
//
// 假设第一个返回值为 lt 第二个返回值为 gt 则有:
// a[l...lt-1] 中的所有元素严格小于 pivot
// a[lt....gt] 中的所有元素全部等于 pivot
// a[gt+1...r] 中的所有元素严格大于 pivot
// func threeWayPartition(a []int, l int, r int) (int, int) {
// 	// 随机选择主元并把它放到数组的最右边
// 	randIdx := rand.Int()%(r-l+1) + l
// 	a[randIdx], a[r] = a[r], a[randIdx]
// 	pivot := a[r]

// 	// 对除主元外的元素进行处理
// 	lt := l
// 	gt := r - 1
// 	for i := lt; i <= gt; i++ {
// 		// 如果当前元素大于主元
// 		// 将它和 a[gt] 交换并将 gt 减一
// 		// 交换后的 a[i] 可能仍然是大于主元的
// 		// 如果我们此时就处理下一个元素就不能做到正确切分
// 		// 所以要不断重复此过程直到 i 和 gt 相遇或是 a[i] 小于等于主元为止
// 		// 过程中数组 a[gt+1...r-1] 中的所有元素都是严格大于主元的
// 		if a[i] > pivot {
// 			for i <= gt && a[i] > pivot {
// 				a[i], a[gt] = a[gt], a[i]
// 				gt--
// 			}
// 		}
// 		// 经过了上面的处理
// 		// 当前元素只能小于等于主元
// 		// 如果它小于主元就和 a[lt] 交换并将 lt 加一
// 		// 数组 a[l...lt-1] 中的所有元素都是严格小于主元的
// 		if a[i] < pivot {
// 			a[i], a[lt] = a[lt], a[i]
// 			lt++
// 		}
// 	}
// 	// 现在 a[gt+1] 是大于主元的部分的第一个元素
// 	// 将它和主元 a[r] 交换
// 	// 从而数组 a[lt...gt+1] 中的所有元素都是等于主元的
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
