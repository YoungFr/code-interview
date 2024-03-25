package slidingwin

//   ___     ___           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \   / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | | | | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | | | |  > _ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | |_| | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 3 - 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

func LengthOfLongestSubstring(s string) int {
	ans := 0
	l := 0
	r := 0

	unique := make(map[byte]int)
	for r < len(s) {
		unique[s[r]]++
		for unique[s[r]] > 1 {
			unique[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}

// LC 209 - 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/

// LC 713 - 乘积小于 k 的子数组个数
// https://leetcode.cn/problems/subarray-product-less-than-k/description/

// LC 2958 - 最多 k 个重复元素的最长子数组
// https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/description/

// LC 2730 - 找到最长的半重复子字符串
// https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/description/

// LC 1004 - 最大连续1的个数 III
// https://leetcode.cn/problems/max-consecutive-ones-iii/description/

// LC 2962 - 统计最大元素出现至少 k 次的子数组
// https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/description/

// LC 2302 - 统计得分小于 K 的子数组数目
// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/description/

// LC 1658 - 将 x 减到 0 的最小操作数
// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
