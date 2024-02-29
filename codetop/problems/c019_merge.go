package problems

// LC 88 - 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/description/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	m--
	n--
	for n >= 0 { // 第二个数组中还有元素

		// 从后往前确定应该使用哪个数组中的数字
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			i--
			m--
		}

		nums1[i] = nums2[n]
		i--
		n--
	}
}
