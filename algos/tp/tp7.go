package tp

//   ___    ______          ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | | | |     / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | | | |    / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | |_| |   / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 42 - 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/

func Trap1(height []int) int {
	n := len(height)
	// 前缀最大值
	lmax := make([]int, n)
	lmax[0] = height[0]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}
	// 后缀最大值
	rmax := make([]int, n)
	rmax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += (min(lmax[i], rmax[i]) - height[i])
	}
	return ans
}

func Trap2(height []int) int {
	n := len(height)
	ans := 0
	l, r := 0, n-1
	pmax, smax := 0, 0
	for l <= r {
		pmax = max(pmax, height[l])
		smax = max(smax, height[r])
		if pmax < smax {
			// height[l] 位置的前缀最大值已知为 pmax
			// height[l] 位置的后缀最大值不可能小于 smax 了
			// => min(pre_max[height[l]], suf_max[height[l]]) = pmax
			ans += pmax - height[l]
			l++
		} else {
			// 同理
			ans += smax - height[r]
			r--
		}
	}
	return ans
}
