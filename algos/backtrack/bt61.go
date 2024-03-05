package backtrack

//    __    __          ___    __    __    ______   .___________.    __    ___     ___
//   / /   /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_    | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \   | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 131 - 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/

func Partition(s string) (ans [][]string) {
	n := len(s)

	// 判断一个字符串是否是回文字符串
	check := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	// 子集型回溯
	// 假设每两个字符之间和最后一个字符后边都有一个逗号
	// a , a , b , c ,
	// 每个逗号都有选或不选两种可能
	// 注意 s[n-1] 后边的逗号必须选择
	// 因为所有切分成的子串组合起来必须等于原来的字符串

	// 此题的思路与
	// https://leetcode.cn/problems/restore-ip-addresses/description/
	// 完全相同
	// 代码略作修改就可以通过 [LC 93 - 复原 IP 地址] 题

	var dfs func(start int, i int, a []string)
	dfs = func(start, i int, a []string) {
		if i == n {
			ans = append(ans, append([]string(nil), a...))
			return
		}

		// 是否选择 s[i] 和 s[i+1] 之间的逗号

		// 选择
		if t := s[start : i+1]; check(t) {
			a = append(a, t)
			dfs(i+1, i+1, a)
			a = a[:len(a)-1]
		}

		// 不选
		if i < n-1 {
			dfs(start, i+1, a)
		}
	}

	dfs(0, 0, make([]string, 0))
	return
}
