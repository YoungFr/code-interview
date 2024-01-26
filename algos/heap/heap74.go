package heap

import "math/rand"

//  ______   _  _            ___    __    __    ______   .___________.    __    ___     ___
// |____  | | || |          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//     / /  | || |_        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//    / /   |__   _|      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//   / /       | |       /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  /_/        |_|      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 215 - 数组中的第 k 个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/

func FindKthLargest(nums []int, k int) int {
	// 方法一 PASS
	//
	// 排序后返回 nums[len(nums)-k] 位置的元素
	// 时间击败 32.75%
	// 空间击败 33.69%
	//
	// sort.Ints(nums)
	// return nums[len(nums)-k]

	// 方法二 PASS
	//
	// 构建大顶堆后进行 k-1 次堆顶元素的删除后再返回堆顶元素
	// 时间击败 75.50%
	// 空间击败 56.67%
	//
	// down := func(a []int, i int, size int) {
	//     root := i
	//     for {
	//         child := 2*root+1
	//         if child >= size {
	//             break
	//         }
	//         if child+1 < size && a[child+1] > a[child] {
	//             child++
	//         }
	//         if a[root] >= a[child] {
	//             break
	//         }
	//         a[root], a[child] = a[child], a[root]
	//         root = child
	//     }
	// }
	// n := len(nums)
	// // 构建大顶堆
	// for i := n/2-1; i >= 0; i-- {
	//     down(nums, i, n)
	// }
	// // 进行 k-1 次堆顶元素的删除
	// for i := 0; i < k-1; i++ {
	//     nums[0], nums[n-1] = nums[n-1], nums[0]
	//     n--
	//     down(nums, 0, n)
	// }
	// return nums[0]

	// 方法三 TIMEOUT
	//
	// 随机选择主元的双向切分快速选择算法
	// 虽然通过了所有用例但是耗时太长
	//
	// return Qselect(nums, 0, len(nums)-1, k)

	// 方法四 PASS
	//
	// 随机选择主元的三向切分快速选择算法
	// 时间击败 94.08%
	// 空间击败 52.35%
	//
	return ThreeWayQselect(nums, 0, len(nums)-1, k)
}

func Qselect(a []int, l int, r int, k int) int {
	// a[l...m-1] 范围中的元素都小于 a[m]
	// a[m+1...r] 范围中的元素都大于等于 a[m]
	// 从而 a[m] 是 a[l...r] 中第 rank = r-m+1 大的元素
	m := partition(a, l, r)
	rank := r - m + 1

	switch {
	// 假设我们要找第 k=1 大的但此时找到的是第 rank=3 大的
	// 说明答案只能在 a[m+1...r] 中且答案在这个数组中也是第 k 大的
	// 原因: a[m] 是整体数组 a[l...r] 中第 rank 大的且 a[m+1...r] 中的元素都大于等于 a[m]
	// 说明 a[m+1...r] 中的元素在整体数组中是第 1, 2, ..., (r-m = rank-1) 大的
	// 恰好它们在这个部分数组 a[m+1...r] 中也是第 1, 2, ..., rank-1 大的
	case rank > k:
		return Qselect(a, m+1, r, k)

	// 假设我们要找第 k=3 大的但此时找到的是第 rank=1 大的
	// 说明答案只能在 a[l...m-1] 中且答案在这个数组中是第 k-rank 大的
	// 原因: a[m] 是整体数组 a[l...r] 中第 rank 大的且 a[l...m-1] 中的元素都小于 a[m]
	// 说明 a[l...m-1] 中的元素在整体数组中是第 rank+1, rank+2, ..., (rank+m-l = r-l+1) 大的
	// 而它们在这个部分数组中又是第 1, 2, ..., m-l(一共 m-l 个元素) 大的
	// 所以有对应关系: 整体中第 rank+m 大 <=> 部分中第 m 大
	// 我们要找整体第 k 大 => 即 rank+m=k => 对应部分中的第 k-rank 大的元素
	case rank < k:
		return Qselect(a, l, m-1, k-rank)

	default:
		return a[m]
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

func ThreeWayQselect(nums []int, l int, r int, k int) int {
	lt, gt := threeWayPartition(nums, l, r)
	// 找到了第 a 大和第 b 大 (a <= b) 的元素
	a := r - gt + 1
	b := r - lt + 1
	switch {
	case k < a:
		return ThreeWayQselect(nums, gt+1, r, k)
	case k > b:
		return ThreeWayQselect(nums, l, lt-1, k-b)
	default:
		return nums[lt]
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
