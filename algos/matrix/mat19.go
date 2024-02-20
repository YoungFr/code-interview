package matrix

//  __    ___           ___    __    __    ______   .___________.    __    ___     ___
// /_ |  / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  | | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  | |  \__, |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  | |    / /      /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  |_|   /_/      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 54 - 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/description/

func SpiralOrder(matrix [][]int) []int {
	// 按圈从外向内打印
	// 给定每圈的左上角坐标和右下角坐标打印一圈
	// 然后移动两个坐标

	// 矩阵有 m 行 n 列
	m := len(matrix)
	n := len(matrix[0])

	elemsNum := m * n

	ans := make([]int, 0, elemsNum)

	// 左上角的行坐标和列坐标
	leftTopRow := 0
	leftTopCol := 0

	// 右下角的行坐标和列坐标
	rightDownRow := m - 1
	rightDownCol := n - 1

	for {
		// 上边一行 => 从左向右
		// XXXXXX
		// OOOOOO
		// OOOOOO
		// OOOOOO
		for j := leftTopCol; j <= rightDownCol; j++ {
			ans = append(ans, matrix[leftTopRow][j])
		}
		if len(ans) == elemsNum {
			break
		}
		// 右边一列 => 从上向下
		// OOOOOO
		// OOOOOX
		// OOOOOX
		// OOOOOO
		for i := leftTopRow + 1; i <= rightDownRow-1; i++ {
			ans = append(ans, matrix[i][rightDownCol])
		}
		if len(ans) == elemsNum {
			break
		}
		// 下边一行 => 从右向左
		// OOOOOO
		// OOOOOO
		// OOOOOO
		// XXXXXX
		for j := rightDownCol; j >= leftTopCol; j-- {
			ans = append(ans, matrix[rightDownRow][j])
		}
		if len(ans) == elemsNum {
			break
		}
		// 左边一列 => 从下向上
		// OOOOOO
		// XOOOOO
		// XOOOOO
		// OOOOOO
		for i := rightDownRow - 1; i >= leftTopRow+1; i-- {
			ans = append(ans, matrix[i][leftTopCol])
		}
		if len(ans) == elemsNum {
			break
		}
		// 缩小一圈
		leftTopRow++
		leftTopCol++
		rightDownRow--
		rightDownCol--
	}

	return ans
}
