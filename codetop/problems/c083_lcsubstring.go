package problems

// LC 718 - 最长重复子数组/最长公共子串
// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/

func FindLength(nums1 []int, nums2 []int) (ans int) {
	// 相当于求两个字符串的最长公共子串
	m := len(nums1)
	n := len(nums2)

	// dp[i][j]
	// 表示 nums1 的前 i 个数字和 nums2 的前 j 个数字
	// 形成的最长公共子数组的长度
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 只有相等时才能转移
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}
