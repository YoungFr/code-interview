package stkandq

//    __    __          ___    __    __    ______   .___________.    __    ___     ___
//   / /   /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_    | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \   | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 20 - 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/

func IsValid(s string) bool {
	n := len(s)
	// 长度为奇数一定不匹配
	if n%2 == 1 {
		return false
	}
	valid := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 左括号压栈
			stack = append(stack, s[i])
		} else {
			// 右括号 => 查看是否和当前的栈顶元素匹配
			if len(stack) == 0 || valid[s[i]] != stack[len(stack)-1] {
				return false
			} else {
				// 如果匹配则弹出栈顶元素
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
