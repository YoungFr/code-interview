package substr

//  __   ___           ___    __    __    ______   .___________.    __    ___     ___
// /_ | |__ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | |    ) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | |   / /       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | |  / /_      /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_| |____|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 76 - 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/description/

func MinWindow(s string, t string) string {
	// 哈希表 dst 记录字符串 t 中每个字符出现的次数
	dst := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		dst[t[i]]++
	}
	// 哈希表 win 记录滑动窗口中每个字符出现的次数
	win := make(map[byte]int)
	// 如果对于 dst 中的每个字符 c 都
	// 满足 win[c] >= dst[c] 则是一个合法的窗口
	check := func() bool {
		for c := range dst {
			if win[c] < dst[c] {
				return false
			}
		}
		return true
	}
	var (
		// lans 设为 -1 来标记其是否被更新
		lans = -1
		rans = len(s)
		l    = 0
		r    = 0
	)
	for r < len(s) {
		win[s[r]]++
		for l <= r && check() {
			win[s[l]]-- // 尝试移动左指针
			if check() {
				l++ // 可行 => 缩小左边界
			} else {
				win[s[l]]++ // 不可行 => 还原
				// 还原后窗口是可行的 => 更新答案
				if r-l < rans-lans {
					lans, rans = l, r
				}
				// 这里必须 break 否则会陷入死循环
				break
			}
		}
		r++
	}
	if lans == -1 {
		return ""
	}
	return s[lans : rans+1]
}
