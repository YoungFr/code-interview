package greedy

//  ______   ______          ___    __    __    ______   .___________.    __    ___     ___
// |____  | |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//     / /      / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//    / /      / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//   / /      / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  /_/      /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 121 - 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

func MaxProfit(prices []int) int {
	// 从第 1 天开始枚举（不能在第 0 天卖出）卖出价格 prices[i]
	// 要想获得最大利润，我们必须知道从第 0 天
	// 到第 i-1 天之间的最低价格作为我们的买入价格
	// 这个前 i-1 天的最小买入价格可以用一个变量维护

	minPrice := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return ans
}
