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
// a[l...lt-1] 中的元素严格小于 pivot
// a[lt...gt]  中的元素全部等于 pivot
// a[gt+1...r] 中的元素严格大于 pivot
func threeWayPartition(a []int, l int, r int) (int, int) {
	randIdx := rand.Int()%(r-l+1) + l
	a[randIdx], a[r] = a[r], a[randIdx]
	pivot := a[r]
	lt, gt := l, r-1
	for i := lt; i <= gt; i++ {
		if a[i] > pivot {
			for i <= gt && a[i] > pivot {
				a[i], a[gt] = a[gt], a[i]
				gt--
			}
		}
		if a[i] < pivot {
			a[i], a[lt] = a[lt], a[i]
			lt++
		}
	}
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
