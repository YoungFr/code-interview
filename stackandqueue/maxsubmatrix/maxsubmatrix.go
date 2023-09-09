package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 描述
// 给定一个整型矩阵 map，其中的值只有 0 和 1 两种，
// 求其中全是 1 的所有矩形区域中，最大的矩形区域里 1 的数量。
//
// 输入描述
// 第一行输入两个整数 n 和 m，代表 n*m 的矩阵。
// 接下来输入一个 n*m 的矩阵。
//
// 输出描述
// 输出其中全是 1 的所有矩形区域中，最大的矩形区域里 1 的数量。
//
// 备注
// 1 <= n,m <= 2000
func main() {
	m, n := 0, 0
	fmt.Scan(&m, &n)

	// 一个 m 行 n 列的矩阵
	matrix := make([][]int, 0, m)

	r := bufio.NewReaderSize(os.Stdin, 10*1024*1024)
	for i := 0; i < m; i++ {
		row := make([]int, 0, n)
		line, _, _ := r.ReadLine()
		for _, s := range strings.Split(string(line), " ") {
			if s == "1" {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		matrix = append(matrix, row)
	}

	// 求以矩阵的每一行为底时最大的的全是 1 的矩阵面积
	heights := matrix[0]
	ans := maxRect(heights)

	for i := 1; i < m; i++ {
		// 计算以当前行为底时柱状图的高度
		// 柱状图的高度是当前行往上连续 1 的数量
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, maxRect(heights))
	}

	fmt.Println(ans)
}

// LC 84. Largest Rectangle in Histogram
// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
func maxRect(heights []int) int {
	n := len(heights)
	l_less := make([]int, n)
	r_less := make([]int, n)
	stk := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			l_less[i] = -1
		} else {
			l_less[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	for len(stk) > 0 {
		stk = stk[:len(stk)-1]
	}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			r_less[i] = n
		} else {
			r_less[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	area := 0
	for i := 0; i < n; i++ {
		area = max(area, heights[i]*(r_less[i]-l_less[i]-1))
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
