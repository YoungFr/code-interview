package stkandq

import (
	"strings"
	"unicode"
)

//    __    ____           ___    __    __    ______   .___________.    __    ___     ___
//   / /   |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_     __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 394 - 字符串解码
// https://leetcode.cn/problems/decode-string/description/

func DecodeString(s string) string {
	// 保存编码中的数字
	num := 0

	// 数字栈
	nums := make([]int, 0)

	// 字符串栈
	strs := make([]string, 0)

	for _, r := range s {
		if unicode.IsDigit(r) {
			// 计算重复次数
			num = 10*num + int(r-'0')
		} else if r == ']' {
			// 如果数字在处理字符前不是零则压栈
			if num != 0 {
				nums = append(nums, num)
				num = 0
			}

			// 这一步很关键
			// 需要不断弹出字符串栈顶的元素并拼接直到遇到左括号
			// 拼接成的字符串才是我们要重复的字符串
			t := ""
			for strs[len(strs)-1] != "[" {
				t = strs[len(strs)-1] + t
				strs = strs[:len(strs)-1]
			}

			// 弹出左括号
			strs = strs[:len(strs)-1]

			// 弹出数字栈的栈顶元素
			k := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			// 重复 k 次后再压栈
			strs = append(strs, strings.Repeat(t, k))
		} else {
			// 如果数字在处理字符前不是零则压栈
			if num != 0 {
				nums = append(nums, num)
				num = 0
			}

			// 除右括号以外的其他字符直接压栈
			strs = append(strs, string(r))
		}
	}

	// 把栈中剩余的元素按从底到顶的顺序拼接起来才是最终的答案
	return strings.Join(strs, "")
}
