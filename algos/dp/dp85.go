package dp

import "math"

func CoinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i < len(f); i++ {
		m := math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && f[i-coins[j]] >= 0 {
				m = min(m, f[i-coins[j]])
			}
		}
		if m == math.MaxInt {
			f[i] = -1
		} else {
			f[i] = m + 1
		}
	}
	return f[amount]
}
