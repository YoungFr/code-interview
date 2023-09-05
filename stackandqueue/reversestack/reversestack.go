package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 描述
// 一个栈依次压入1,2,3,4,5，那么从栈顶到栈底分别为5,4,3,2,1。
// 将这个栈转置后，从栈顶到栈底为1,2,3,4,5，也就是实现栈中元素的逆序。
// 但是只能用递归函数来实现，不能用其他数据结构。
//
// 输入描述
// 输入数据第一行一个整数 N 为栈中元素的个数。
// 接下来一行 N 个整数表示一个栈依次压入的每个元素。
//
// 输出描述
// 输出一行表示栈中元素逆序后的栈顶到栈底的每个元素
func main() {
	sc := bufio.NewScanner(os.Stdin)

	var n int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	// 只访问 stk 的最后一个元素保证这是一个栈
	stk := make([]string, 0, n)

	// 将读入的元素压入栈中
	sc.Scan()
	stk = append(stk, strings.Split(sc.Text(), " ")...)

	reverseStack(stk)

	var ans, sep string
	for len(stk) > 0 {
		ans += sep + stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		sep = " "
	}
	fmt.Println(ans)
}

func reverseStack(stk []string) {
	if len(stk) == 0 {
		return
	}
	bottom := getAndRemoveBottom(&stk)
	reverseStack(stk)
	stk = append(stk, bottom)
}

// 移除并返回栈底元素
func getAndRemoveBottom(stk *[]string) string {
	top := (*stk)[len(*stk)-1]
	*stk = (*stk)[:len(*stk)-1]
	if len(*stk) == 0 {
		return top
	} else {
		bottom := getAndRemoveBottom(stk)
		*stk = append(*stk, top)
		return bottom
	}
}
