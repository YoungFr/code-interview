package problems

func HeapSort(a []int) {
	n := len(a)
	// 构建大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		down(a, i, n)
	}
	// 进行 n-1 次堆顶元素的删除
	for n > 1 {
		a[0], a[n-1] = a[n-1], a[0]
		n--
		down(a, 0, n)
	}
}

func down(a []int, i int, size int) {
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
