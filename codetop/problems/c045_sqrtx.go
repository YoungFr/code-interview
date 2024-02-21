package problems

// LC 69 - x 的平方根
// https://leetcode.cn/problems/sqrtx/description/

func MySqrt(x int) int {
	lo := 0
	hi := x
	// 找到第一个满足自身的平方和大于等于 x 的元素
	for lo < hi {
		mi := int(uint(lo+hi) >> 1)
		if mi*mi < x {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	if lo*lo == x {
		return lo
	}
	return lo - 1
}
