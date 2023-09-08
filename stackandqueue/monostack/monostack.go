package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 描述
// 给定一个不含有重复值的数组 arr，找到每一个 i 位置左边和右边
// 离 i 位置最近且值比 arr[i] 小的位置。返回所有位置相应的信息。
//
// 输入描述
// 第一行输入一个数字 n 表示数组 arr 的长度。
// 第二行输入 n 个数字表示数组中的元素。
//
// 输出描述
// 输出 n 行，每行两个数字 L 和 R；如果不存在，则值为 -1，下标从 0 开始。
//
// 备注
// 1 <= n <= 1000000
// -1000000 <= arr[i] <= 1000000
func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, 0, n)

	r := bufio.NewReaderSize(os.Stdin, 10*1024*1024)
	nums, _, _ := r.ReadLine()
	for _, num := range strings.Split(string(nums), " ") {
		a, _ := strconv.Atoi(num)
		arr = append(arr, a)
	}

	ans := make([][2]int, n)
	stk := make([]int, 0, n)

	//          单调栈
	//
	//            @
	//            @     @     @
	//            @  @  @  @  @
	//            @  @  @  @  @  @
	// -----------------------------
	// |   i   |  0  1  2  3  4  5 |
	// -----------------------------
	// | l-ans | -1 -1  1 -1  3 -1 |
	// | r-ans |  1  5  3  5  5 -1 |
	// -----------------------------
	//
	// 上、下一个严格小于的元素的位置 => 从顶到底的位置所表示的值的大小严格递减
	// 上、下一个小于等于的元素的位置 => 从顶到底的位置所表示的值的大小单调递减
	// 上、下一个严格大于的元素的位置 => 从顶到底的位置所表示的值的大小严格递增
	// 上、下一个大于等于的元素的位置 => 从顶到底的位置所表示的值的大小单调递增

	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && arr[i] <= arr[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			ans[i][1] = -1
		} else {
			ans[i][1] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	// 清空栈
	for len(stk) > 0 {
		stk = stk[:len(stk)-1]
	}

	for i := 0; i < n; i++ {
		for len(stk) > 0 && arr[i] <= arr[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			ans[i][0] = -1
		} else {
			ans[i][0] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	var str strings.Builder
	var sep string
	for i := 0; i < n; i++ {
		str.WriteString(sep)
		str.WriteString(strconv.Itoa(ans[i][0]))
		str.WriteString(" ")
		str.WriteString(strconv.Itoa(ans[i][1]))
		sep = "\n"
	}
	fmt.Println(str.String())
}
