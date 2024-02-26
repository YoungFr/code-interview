package problems

func MergeSort(a []int, lo int, hi int) {
	if lo < hi {
		mi := int(uint(lo+hi) >> 1)
		MergeSort(a, lo, mi)
		MergeSort(a, mi+1, hi)
		merge(a, lo, mi, hi)
	}
}

// 数组 a[lo...mi] 和 a[mi+1...hi] 各自是有序的
// 将它们合并成一个大的有序数组
func merge(a []int, lo int, mi int, hi int) {
	L := make([]int, mi-lo+1)
	copy(L, a[lo:mi+1])
	R := make([]int, hi-mi)
	copy(R, a[mi+1:hi+1])
	i := 0
	j := 0
	for k := lo; k <= hi; k++ {
		if i == mi-lo+1 {
			a[k] = R[j]
			j++
		} else if j == hi-mi {
			a[k] = L[i]
			i++
		} else if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
	}
}
