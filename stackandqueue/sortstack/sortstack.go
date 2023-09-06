package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 描述
// 一个栈中元素的类型为整型，现在想将该栈从顶到底按从大到小的顺序排序，
// 只许申请一个栈。除此之外，可以申请新的变量，但不能申请额外的数据结构。
//
// 输入描述
// 第一行输入一个 N 表示栈中元素的个数。
// 第二行输入 N 个整数表示栈顶到栈底的各个元素。
//
// 输出描述
// 输出一行表示排序后的栈中栈顶到栈底的各个元素。
//
// 备注
// 1 <= N <= 10000
// -1000000 <= an <= 1000000
func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 默认的缓冲区大小是 64 * 1024 字节
	// 最大的可能输入大小是 10000 个 "-1000000 " 所以要显式设置缓冲区大小
	// 否则读取第二行时会产生 bufio.Scanner: token too long 错误
	sc.Buffer(make([]byte, len("-1000000 ")*10000), len("-1000000 ")*10000)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// 只访问最后一个元素保证这是一个栈
	stk := make([]int, 0, n)

	sc.Scan()
	nums := strings.Split(sc.Text(), " ")
	for i := len(nums) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(nums[i])
		stk = append(stk, num)
	}

	// 辅助栈
	aux := make([]int, 0, n)
	for len(stk) > 0 {
		// 弹出要排序的栈的栈顶元素
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		// 弹出辅助栈中所有比 top 小的元素并压回要排序的栈
		for len(aux) > 0 && aux[len(aux)-1] < top {
			stk = append(stk, aux[len(aux)-1])
			aux = aux[:len(aux)-1]
		}
		aux = append(aux, top)
	}
	// 辅助栈现在从顶到底已经按从小到大的顺序排序
	for len(aux) > 0 {
		stk = append(stk, aux[len(aux)-1])
		aux = aux[:len(aux)-1]
	}

	var ans, sep string
	for len(stk) > 0 {
		ans += sep + strconv.Itoa(stk[len(stk)-1])
		stk = stk[:len(stk)-1]
		sep = " "
	}
	fmt.Println(ans)
}
