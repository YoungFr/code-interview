package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 描述
// 给定数组 arr 和整数 num 返回有多少个子数组满足如下情况：
//
//	max(arr[i...j]) - min(arr[i...j]) <= num
//
// max(arr[i...j]) 表示子数组 arr[i...j] 中的最大值。
// min[arr[i...j]) 表示子数组 arr[i...j] 中的最小值。
//
// 输入描述
// 第一行输入两个数 n 和 num，其中 n 表示数组 arr 的长度。
// 第二行输入 n 个整数 xi 表示数组 arr 中的每个元素。
//
// 输出描述
// 输出给定数组中满足条件的子数组个数。
//
// 备注
// 1 <= n <= 1000000
// -1000000 <= arr[i] <= 1000000
// 0 <= num <= 2000000
func main() {
	n, num := 0, 0
	fmt.Scan(&n, &num)

	arr := make([]int, 0, n)
	r := bufio.NewReaderSize(os.Stdin, 10*1024*1024)
	line, _, _ := r.ReadLine()
	for _, s := range strings.Split(string(line), " ") {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	// 两个结论
	// 1. 如果 arr[i...j] 满足条件 => 那么 arr[i...j] 的所有子数组都满足条件
	// 2. 如果 arr[i...j] 不满足条件 => 那么所有包含 arr[i...j] 的数组都不满足条件

	ans := 0
	i, j := 0, 0
	qmin := make([]int, 0, n)
	qmax := make([]int, 0, n)

	for i < n {
		for j < n {
			// 维护最小值时队列中的元素要保证严格递增
			for len(qmin) > 0 && arr[j] <= arr[qmin[len(qmin)-1]] {
				qmin = qmin[:len(qmin)-1]
			}
			qmin = append(qmin, j)
			// 维护最大值时队列中的元素要保证严格递减
			for len(qmax) > 0 && arr[j] >= arr[qmax[len(qmax)-1]] {
				qmax = qmax[:len(qmax)-1]
			}
			qmax = append(qmax, j)
			// 根据结论 2 此时 j 没有后移的必要
			if arr[qmax[0]]-arr[qmin[0]] > num {
				break
			}
			j++
		}
		// 到达这里时要么是 arr[i...j] 不满足 max(arr[i...j]) - min(arr[i...j]) <= num 的条件
		// 要么是 j 和 n 相等
		// 总之根据结论 1 此时子数组 arr[i...i], arr[i...i+1], ..., arr[i...j-1] 都满足条件
		// 一共有 j-1-i+1 = j-i 个子数组
		ans += j - i
		// 队头元素过期
		if qmin[0] == i {
			qmin = qmin[1:]
		}
		// 队头元素过期
		if qmax[0] == i {
			qmax = qmax[1:]
		}
		i++
	}

	fmt.Println(ans)
}
