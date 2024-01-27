package problems

import "math/rand"

// LC 912 - 排序数组
// https://leetcode.cn/problems/sort-an-array/description/

// 随机选择主元的双向切分快速排序算法 => 虽然通过了所有用例但是耗时太长
// 随机选择主元的三向切分快速排序算法 => 时间击败 94.93% 空间击败 76.51%
func SortArray(nums []int) []int {
	// QuickSort(nums, 0, len(nums)-1)
	// return nums
	ThreeWayQsort(nums, 0, len(nums)-1)
	return nums
}

// 随机选择主元的三向切分快速排序算法
func ThreeWayQsort(a []int, l, r int) {
	if l < r {
		lt, gt := threeWayPartition(a, l, r)
		ThreeWayQsort(a, l, lt-1)
		ThreeWayQsort(a, gt+1, r)
	}
}

// 随机选择主元的三向切分算法
//
// 假设第一个返回值为 lt 第二个返回值为 gt 则有:
// a[l...lt-1] 中的所有元素严格小于 pivot
// a[lt....gt] 中的所有元素全部等于 pivot
// a[gt+1...r] 中的所有元素严格大于 pivot
func threeWayPartition(a []int, l int, r int) (int, int) {
	// 随机选择主元并把它放到数组的最右边
	randIdx := rand.Int()%(r-l+1) + l
	a[randIdx], a[r] = a[r], a[randIdx]
	pivot := a[r]

	// 对除主元外的元素进行处理
	lt := l
	gt := r - 1
	for i := lt; i <= gt; i++ {
		// 如果当前元素大于主元
		// 将它和 a[gt] 交换并将 gt 减一
		// 交换后的 a[i] 可能仍然是大于主元的
		// 如果我们此时就处理下一个元素就不能做到正确切分
		// 所以要不断重复此过程直到 i 和 gt 相遇或是 a[i] 小于等于主元为止
		// 过程中数组 a[gt+1...r-1] 中的所有元素都是严格大于主元的
		if a[i] > pivot {
			for i <= gt && a[i] > pivot {
				a[i], a[gt] = a[gt], a[i]
				gt--
			}
		}
		// 经过了上面的处理
		// 当前元素只能小于等于主元
		// 如果它小于主元就和 a[lt] 交换并将 lt 加一
		// 数组 a[l...lt-1] 中的所有元素都是严格小于主元的
		if a[i] < pivot {
			a[i], a[lt] = a[lt], a[i]
			lt++
		}
	}
	// 现在 a[gt+1] 是大于主元的部分的第一个元素
	// 将它和主元 a[r] 交换
	// 从而数组 a[lt...gt+1] 中的所有元素都是等于主元的
	a[gt+1], a[r] = a[r], a[gt+1]
	return lt, gt + 1
}

// 随机选择主元的双向切分快速排序算法
func QuickSort(a []int, l, r int) {
	if l < r {
		m := partition(a, l, r)
		QuickSort(a, l, m-1)
		QuickSort(a, m+1, r)
	}
}

// 随机选择主元的双向切分算法
func partition(a []int, l, r int) int {
	k := rand.Int()%(r-l+1) + l
	a[k], a[r] = a[r], a[k]

	x := a[r]                // the pivot
	i := l - 1               // highest index into the low side
	for j := l; j < r; j++ { // process each element other than the pivot
		if a[j] < x { // does this element belong on the low side?
			i++                     // index of a new slot in the low side
			a[i], a[j] = a[j], a[i] // put this element there
		}
	}
	a[i+1], a[r] = a[r], a[i+1] // pivot goes just to the right of the low side
	return i + 1                // new index of the pivot
}
