package hash

//   ___    __          ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | | | |  | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | | | |  | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | |_| |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 1 - 两数之和
// https://leetcode.cn/problems/two-sum/description/

func TwoSum(nums []int, target int) []int {
	var ans = make([]int, 0, 2)
	// 键为数组元素
	// 值为对应下标
	var has = make(map[int]int)

	for i, v := range nums {
		if j, ok := has[target-v]; ok {
			ans = append(ans, i, j)
			break
		} else {
			has[v] = i
		}
	}

	return ans
}
