package bs

//    __    _  _            ___    __    __    ______   .___________.    __    ___     ___
//   / /   | || |          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//  / /_   | || |_        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// | '_ \  |__   _|      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
// | (_) |    | |       /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
//  \___/     |_|      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 74 - 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/description/

func SearchMatrix(matrix [][]int, target int) bool {
	// 根据题目描述
	// 如果按照从上到下、从左到右的顺序遍历整个矩阵
	// 得到的结果将是一个有序数组
	// 所以可以使用二分查找
	// 问题关键是将二维矩阵的行列下标对应到一维数组的下标
	//
	// matrix[i][j] => a[i * n + j] (其中 n 是矩阵的列数)
	// a[k] => matrix[k / n][k % n] (其中 n 是矩阵的列数)

	m := len(matrix)
	n := len(matrix[0])

	lo := 0
	hi := m * n

	for lo < hi {
		mi := int(uint(lo+hi) >> 1)
		if matrix[mi/n][mi%n] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo < m*n && matrix[lo/n][lo%n] == target
}
