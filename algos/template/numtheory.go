package template

import (
	"math"
	"sort"
)

// 计算 (a ^ n) % mod 的值
func Pow(a, n, mod int) int {
	ans := 1
	a %= mod
	for n != 0 {
		if n%2 == 1 {
			ans = (ans * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return ans
}

// 判断一个数是否为质数
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	q := int(math.Sqrt(float64(n)))
	for i := 5; i <= q; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 计算区间 [M, N) 内的所有质数
func PrimeMN(m, n int) []int {
	ps := make([]int, 0)
	f := make([]bool, n)
	for p := 2; p < n; p++ {
		if !f[p] {
			if p >= m {
				ps = append(ps, p)
			}
			for j := p * p; j < n; j += p {
				f[j] = true
			}
		}
	}
	return ps
}

// 质因数分解
func PrimeDecompose(n int) []int {
	ps := make([]int, 0)
	for p := 2; p*p <= n; p++ {
		for n%p == 0 {
			ps = append(ps, p)
			n /= p
		}
	}
	if n > 1 {
		ps = append(ps, n)
	}
	return ps
}

// 计算一个数的所有不重复的质因子
func PrimeDivisors(n int) []int {
	pset := make(map[int]bool)
	for p := 2; p*p <= n; p++ {
		for n%p == 0 {
			pset[p] = true
			n /= p
		}
	}
	if n > 1 {
		pset[n] = true
	}
	ps := make([]int, 0)
	for p := range pset {
		ps = append(ps, p)
	}
	sort.Ints(ps)
	return ps
}
