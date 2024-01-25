package problems

import "math/rand"

// LC 912 - 排序数组
// https://leetcode.cn/problems/sort-an-array/description/

func SortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(a []int, l, r int) {
	if l < r {
		m := randomPartition(a, l, r)
		quickSort(a, l, m-1)
		quickSort(a, m+1, r)
	}
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]                // the pivot
	i := l - 1               // highest index into the low side
	for j := l; j < r; j++ { // process each element other than the pivot
		if a[j] <= x { // does this element belong on the low side?
			i++                     // index of a new slot in the low side
			a[i], a[j] = a[j], a[i] // put this element there
		}
	}
	a[i+1], a[r] = a[r], a[i+1] // pivot goes just to the right of the low side
	return i + 1                // new index of the pivot
}
