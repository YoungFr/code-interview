package graph

//  _____   __          ___    __    __    ______   .___________.    __    ___     ___
// | ____| /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | |__    | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |___ \   | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 200 - 岛屿数量
// https://leetcode.cn/problems/number-of-islands/description/

func NumIslands(grid [][]byte) int {
	// 思路
	// 遍历网格中的每个元素
	//   如果是水域（'0'）=> 什么也不做
	//   如果是陆地（'1'）且没有访问过 => 从这块陆地开始
	//   使用 DFS 访问所有能访问到的陆地 => 形成一个岛屿

	m := len(grid)
	n := len(grid[0])

	// 二维数组 vis 的尺寸和网格相同
	// 用于标记某块陆地是否访问过
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	// 从第 r 行 c 列的陆地开始访问所有能访问到的陆地
	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		// 如果（索引越界 or 访问到了水域 or 访问到已访问过的陆地）则返回
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '0' || vis[r][c] {
			return
		}
		// 将第 r 行 c 列的陆地标记为已访问
		vis[r][c] = true
		// 访问水平和垂直方向的陆地
		dfs(r, c-1)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r+1, c)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果是陆地（'1'）且没有访问过 => 使用 DFS 访问
			// 在这个过程中所有可访问到的陆地都会被标记为已访问
			if grid[i][j] == '1' && !vis[i][j] {
				dfs(i, j)
				ans++
			}
		}
	}

	return ans
}
