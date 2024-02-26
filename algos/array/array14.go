package array

import "slices"

//  __   ____           ___    __    __    ______   .___________.    __    ___     ___
// /_ | |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | |  |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_| |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 56 - 合并区间
// https://leetcode.cn/problems/merge-intervals/description/

func Merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	// 把第一个区间加入答案列表
	ans = append(ans, intervals[0])

	// 遍历剩余区间
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > ans[len(ans)-1][1] {
			// 如果当前区间与最后一个区间没有交集
			// 则另起一个区间
			ans = append(ans, intervals[i])
		} else {
			// 否则更新最后一个区间的右端点进行合并
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}
	}

	return ans
}
