package heap

//  ______   _____          ___    __    __    ______   .___________.    __    ___     ___
// |____  | | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//     / /  | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//    / /   |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//   / /     ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  /_/     |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 347 - 前 k 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/description/

type item struct {
	val int
	cnt int
}

func TopKFrequent(nums []int, k int) []int {
	// 统计每个元素出现的频率
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	// 将元素和其出现频率绑定后放到数组中
	h := make([]item, 0)
	for val := range cnt {
		h = append(h, item{val: val, cnt: cnt[val]})
	}

	// 建立大顶堆
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}

	// 进行 k 次堆顶元素的删除
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, h[0].val)
		h[0], h[n-1] = h[n-1], h[0]
		n--
		down(h, 0, n)
	}

	return ans
}

func down(a []item, i int, size int) {
	root := i
	for {
		child := 2*root + 1
		if child >= size {
			break
		}
		if child+1 < size && a[child+1].cnt > a[child].cnt {
			child++
		}
		if a[root].cnt >= a[child].cnt {
			break
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}
