package heap

//  ______     __           ___    __    __    ______   .___________.    __    ___     ___
// |____  |   / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//     / /   / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//    / /   | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//   / /    | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  /_/      \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 295 - 数据流的中位数
// https://leetcode.cn/problems/find-median-from-data-stream/description/

// Your MedianFinder object will be instantiated and called as such:
// obj := Constructor();
// obj.AddNum(num);
// param_2 := obj.FindMedian();

type MedianFinder struct {
	// 两个堆
	// 假设当前数据结构中的所有元素按从小到大排序后的结果是 a 数组
	//
	// 第一个堆是大顶堆
	// 保存 a[0 <= i <= (len(a)-1)/2] 中的元素
	// 第二个堆是小顶堆
	// 保存 a[(len(a)-1)/2+1 <= i <= len(a)-1] 中的元素
	//
	// 即对于长度为奇数时第一个堆比第二个大一
	// 对于长度为偶数时两个堆大小相等
	loHeap []int
	hiHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{loHeap: make([]int, 0), hiHeap: make([]int, 0)}
}

func (m *MedianFinder) AddNum(num int) {
	if len(m.loHeap) == 0 {
		m.loHeap = append(m.loHeap, num)
		return
	}
	if len(m.hiHeap) == 0 {
		m.hiHeap = append(m.hiHeap, num)
		if m.loHeap[0] > m.hiHeap[0] {
			m.loHeap[0], m.hiHeap[0] = m.hiHeap[0], m.loHeap[0]
		}
		return
	}

	lo := m.loHeap[0]
	hi := m.hiHeap[0]

	switch {
	case num < lo:
		// 在 num < lo 时新元素只能被插入左边的堆
		if len(m.loHeap) > len(m.hiHeap) {
			// 保存并删除左边的堆顶元素
			leftTop := m.loHeap[0]
			m.loHeap[0], m.loHeap[len(m.loHeap)-1] = m.loHeap[len(m.loHeap)-1], m.loHeap[0]
			m.loHeap = m.loHeap[:len(m.loHeap)-1]
			bigdown(m.loHeap, 0, len(m.loHeap))

			// 将其插入右边的堆
			m.hiHeap = append(m.hiHeap, leftTop)
			smallup(m.hiHeap, len(m.hiHeap)-1)

			// 将新元素插入左边的堆
			m.loHeap = append(m.loHeap, num)
			bigup(m.loHeap, len(m.loHeap)-1)
		} else {
			m.loHeap = append(m.loHeap, num)
			bigup(m.loHeap, len(m.loHeap)-1)
		}
	case lo <= num && num <= hi:
		// 在 lo <= num <= hi 时新元素可以插到任意一个堆
		if len(m.loHeap) > len(m.hiHeap) {
			m.hiHeap = append(m.hiHeap, num)
			smallup(m.hiHeap, len(m.hiHeap)-1)
		} else {
			m.loHeap = append(m.loHeap, num)
			bigup(m.loHeap, len(m.loHeap)-1)
		}
	case num > hi:
		// 在 num > hi 时新元素只能插入右边的堆
		if len(m.loHeap) > len(m.hiHeap) {
			m.hiHeap = append(m.hiHeap, num)
			smallup(m.hiHeap, len(m.hiHeap)-1)
		} else {
			// 保存并删除右边堆的堆顶元素
			rightTop := m.hiHeap[0]
			m.hiHeap[0], m.hiHeap[len(m.hiHeap)-1] = m.hiHeap[len(m.hiHeap)-1], m.hiHeap[0]
			m.hiHeap = m.hiHeap[:len(m.hiHeap)-1]
			smalldown(m.hiHeap, 0, len(m.hiHeap))

			// 将其插入左边堆
			m.loHeap = append(m.loHeap, rightTop)
			bigup(m.loHeap, len(m.loHeap)-1)

			// 将新元素插入右边的堆
			m.hiHeap = append(m.hiHeap, num)
			smallup(m.hiHeap, len(m.hiHeap)-1)
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.loHeap) > len(m.hiHeap) {
		return float64(m.loHeap[0])
	}
	return (float64(m.loHeap[0]) + float64(m.hiHeap[0])) / 2
}

// 大顶堆上浮
func bigup(a []int, i int) {
	child := i
	for {
		root := (child - 1) / 2
		if root == child || a[root] >= a[child] {
			break
		}
		a[root], a[child] = a[child], a[root]
		child = root
	}
}

// 大顶堆下沉
func bigdown(a []int, i int, size int) {
	root := i
	for {
		child := 2*root + 1
		if child >= size {
			break
		}
		if child+1 < size && a[child+1] > a[child] {
			child++
		}
		if a[root] >= a[child] {
			break
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}

// 小顶堆上浮
func smallup(a []int, i int) {
	child := i
	for {
		root := (child - 1) / 2
		if root == child || a[root] <= a[child] {
			break
		}
		a[root], a[child] = a[child], a[root]
		child = root
	}
}

// 小顶堆下沉
func smalldown(a []int, i int, size int) {
	root := i
	for {
		child := 2*root + 1
		if child >= size {
			break
		}
		if child+1 < size && a[child+1] < a[child] {
			child++
		}
		if a[root] <= a[child] {
			break
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}
