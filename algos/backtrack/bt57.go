package backtrack

//  _____   ______          ___    __    __    ______   .___________.    __    ___     ___
// | ____| |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | |__       / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |___ \     / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |   / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 17 - 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

var mapping = []string{"",
	"", "abc", "def", // 1 2 3
	"ghi", "jkl", "mno", // 4 5 6
	"pqrs", "tuv", "wxyz", // 7 8 9
}

// 回溯例题
func LetterCombinations1(digits string) (ans []string) {
	n := len(digits)

	if n == 0 {
		return ans
	}

	// 枚举第 i 个字符
	// path 是当前组成的字符串
	var dfs func(i int, path string)

	dfs = func(i int, path string) {
		// 已经枚举了 n 个字符 -> 返回
		if i == n {
			ans = append(ans, path)
			return
		}

		for _, c := range mapping[int(digits[i]-'0')] {
			// 当前操作：枚举第 i 个字符
			path += string(c)

			// 下一个子问题：枚举第 i+1 个字符
			dfs(i+1, path)

			// 撤销当前选择
			path = path[:len(path)-1]
		}
	}

	dfs(0, "")
	return ans
}

// 因为最终结果的长度一定是 n -> 通过 DFS 形成决策树
func LetterCombinations2(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return ans
	}

	// 全局的路径
	path := make([]byte, n)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		// 依次放置数字对应的字母
		for _, c := range mapping[int(digits[i]-'0')] {
			path[i] = byte(c)
			dfs(i + 1)
		}
	}

	dfs(0)
	return ans
}
