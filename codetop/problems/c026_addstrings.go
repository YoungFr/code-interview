package problems

import "slices"

// LC 415 - 字符串相加
// https://leetcode.cn/problems/add-strings/description/
func AddStrings(num1 string, num2 string) string {
	sum := make([]byte, 0)
	i := len(num1) - 1
	j := len(num2) - 1

	// 进位
	carry := byte(0)

	for i >= 0 || j >= 0 {
		// 对位数较短的数进行补零
		d1 := byte(0)
		if i >= 0 {
			d1 = num1[i] - '0'
		}
		// 对位数较短的数进行补零
		d2 := byte(0)
		if j >= 0 {
			d2 = num2[j] - '0'
		}
		// 加和
		sum = append(sum, (d1+d2+carry)%10+'0')
		// 进位
		carry = (d1 + d2 + carry) / 10

		i--
		j--
	}
	if carry > 0 {
		sum = append(sum, carry+'0')
	}

	// 反转结果
	slices.Reverse(sum)
	return string(sum)
}
