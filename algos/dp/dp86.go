package dp

//   ___      __           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \    / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) |  / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  > _ <  | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) | | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 139 - 单词拆分
// https://leetcode.cn/problems/word-break/description/

func WordBreak(s string, wordDict []string) bool {
	// f[i] 表示子串 s[0...i] 能否使用字典中的单词拼接而成
	// f[i = 0] = has(s[0...0])
	// f[i > 0] = OR(f[0...j] && has[s[j+1...i]]) || has[s[0...i]], 0 <= j < i
	has := make(map[string]bool)
	for _, word := range wordDict {
		has[word] = true
	}

	n := len(s)
	f := make([]bool, n)
	f[0] = has[s[:1]]

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			f[i] = f[i] || (f[j] && has[s[j+1:i+1]])
		}
		f[i] = f[i] || has[s[:i+1]]
	}

	return f[n-1]
}
