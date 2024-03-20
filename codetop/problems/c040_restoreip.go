package problems

import (
	"strconv"
	"strings"
)

// LC 93 - 复原 IP 地址
// https://leetcode.cn/problems/restore-ip-addresses/description/

func RestoreIpAddresses(s string) (ans []string) {
	m := make(map[string]bool)
	for i := 0; i < 256; i++ {
		m[strconv.Itoa(i)] = true
	}
	n := len(s)

	// s: 0 . 1 . 1 . 2 . 3 . 2 . 5 . 5 .
	// i: 0   1   2   3   4   5   6   7
	// i:   0   1   2   3   4   5   6   7
	// 选 4 个 '.' 且最后一个 (编号为 n-1 的索引) 必须选取
	// 注意到 '.' 的索引和它前边的数字的索引是相同的

	// start: 子串的起始索引
	// chose: 下一次要选择/不选择的 '.' 的索引
	var dfs func(start, chose int, a []string)
	dfs = func(start, chose int, a []string) {
		// 剪枝
		if len(a) > 4 {
			return
		}

		// 下边的步骤与
		// LC 61 - 切分回文串
		// https://leetcode.cn/problems/palindrome-partitioning/description/
		// 完全相同
		// 每次选择的 '.' 如果能让 s[start...chose] 子串形成 0-255 之间的值
		// 就将它添加到结果中
		// 唯一的区别是我们只需要能将整个字符串切分成 4 段的那些结果
		if chose == n {
			if len(a) == 4 {
				ans = append(ans, strings.Join(a, "."))
			}
			return
		}

		// 选择了 chose 索引
		if t := s[start : chose+1]; m[t] {
			a = append(a, t)
			dfs(chose+1, chose+1, a)
			a = a[:len(a)-1]
		}

		// 未选择 chose 索引
		if chose < n-1 {
			dfs(start, chose+1, a)
		}
	}

	dfs(0, 0, make([]string, 0))
	return
}
