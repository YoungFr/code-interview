package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 描述
// 有一个整型数组 arr 和一个大小为 w 的窗口从数组的最左边滑到最右边，
// 窗口每次向右边滑一个位置，求每一种窗口状态下的最大值。
// 如果数组长度为 n，窗口大小为 w，则一共产生 n-w+1 个窗口的最大值。
//
// 输入描述
// 第一行输入 n 和 w 分别代表数组长度和窗口大小。
// 第二行输入 n 个整数 xi 表示数组中的各个元素。
//
// 输出描述
// 输出一个长度为 n-w+1 的数组 res，res[i]表示每一种窗口状态下的最大值。
//
// 备注
// 1 <= w <= n <= 1000000
// -1000000 <= xi <= 1000000
func main() {
	n, w := 0, 0
	fmt.Scan(&n, &w)

	arr := make([]int, 0, n)

	r := bufio.NewReaderSize(os.Stdin, 10*1024*1024)
	line, _, _ := r.ReadLine()

	for _, s := range strings.Split(string(line), " ") {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	ans := make([]int, 0, n-w+1)
	// 双端队列 => 单调队列
	deq := make([]int, 0, n-w+1)

	for i := 0; i < n; i++ {
		if len(deq) == 0 {
			deq = append(deq, i)
		} else {
			// 队头元素过期
			if deq[0] == i-w {
				deq = deq[1:]
			}
			// 队列中的元素保持严格单调递减
			for len(deq) > 0 && arr[deq[len(deq)-1]] <= arr[i] {
				deq = deq[:len(deq)-1]
			}
			deq = append(deq, i)
		}
		if i+1 >= w {
			ans = append(ans, arr[deq[0]])
		}
	}

	var str strings.Builder
	var sep string
	for i := 0; i < len(ans); i++ {
		str.WriteString(sep)
		str.WriteString(strconv.Itoa(ans[i]))
		sep = " "
	}
	fmt.Println(str.String())
}
