package tp

import "sort"

//   ___      __           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \    / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | | | |  / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | | | | | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | |_| | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 15 - 三数之和
// https://leetcode.cn/problems/3sum/description/

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	n := len(nums)
	for i := 0; i <= n-3; i++ {
		// 对第一个元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// i < j < k
		j := i + 1
		k := n - 1
		// 必须是 3 个数 => 变量 j 和 k 不能相等
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum < 0:
				j++
			case sum > 0:
				k--
			default:
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				// 去重
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			}
		}
	}

	return ans
}

// LC 167 - 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/

// LC 2824 - 统计和小于目标的下标对数目
// https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/description/

// LC 16 - 最接近的三数之和
// https://leetcode.cn/problems/3sum-closest/description/

// LC 18 - 四数之和
// https://leetcode.cn/problems/4sum/description/

// LC 611 - 有效三角形的个数
// https://leetcode.cn/problems/valid-triangle-number/description/
