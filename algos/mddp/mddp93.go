package mddp

//   ___    ____           ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  \__, |  |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//    / /   ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//   /_/   |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 5 - 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/description/

func LongestPalindrome(s string) string {
	ansl := 0
	ansr := 0
	for i := 0; i < len(s); i++ {
		var a, b int
		// 以 s[i] 为回文中心
		a, b = expand(s, i, i)
		if b-a > ansr-ansl {
			ansl, ansr = a, b
		}
		// 以 | 为回文中心
		a, b = expand(s, i, i+1)
		if b-a > ansr-ansl {
			ansl, ansr = a, b
		}
	}
	return s[ansl : ansr+1]
}

// 中心扩展法
// 假设返回值为 a 和 b 则 s[a...b] 是回文串
func expand(s string, i int, j int) (int, int) {
	lo := i
	hi := j
	for lo >= 0 && hi < len(s) && s[lo] == s[hi] {
		lo--
		hi++
	}
	return lo + 1, hi - 1
}
