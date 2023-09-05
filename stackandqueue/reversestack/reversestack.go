package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
