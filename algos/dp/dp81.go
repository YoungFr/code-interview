package dp

//   ___    __          ___    __    __    ______   .___________.    __    ___     ___
//  / _ \  /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | (_) |  | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  > _ <   | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 70 - 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description/

func ClimbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 2; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
