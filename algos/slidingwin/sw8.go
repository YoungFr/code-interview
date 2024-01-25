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
		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}

	return ans
}
