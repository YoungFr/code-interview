package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LC 155 - 最小栈
// https://leetcode.cn/problems/min-stack/description/

func NewMinStack() *MinStack {
	return &MinStack{
		valStack: make([]int, 0),
		minStack: make([]int, 0),
	}
}

type MinStack struct {
	valStack []int
	minStack []int
}

func (ms *MinStack) push(val int) {
	// 新元素压入 valStack 栈
	ms.valStack = append(ms.valStack, val)
	// 如果新元素小于等于 minStack 的栈顶元素
	// 则把新元素也压入 minStack 栈
	if len(ms.minStack) == 0 {
		ms.minStack = append(ms.minStack, val)
	} else if val <= ms.getmin() {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) pop() int {
	if len(ms.valStack) == 0 {
		panic("stack is empty.")
	}
	// 弹出 valStack 的栈顶元素
	ans := ms.valStack[len(ms.valStack)-1]
	ms.valStack = ms.valStack[:len(ms.valStack)-1]
	// 如果弹出的元素等于 minStack 的栈顶元素
	// 则把 minStack 的栈顶元素也弹出
	if ans == ms.getmin() {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
	return ans
}

func (ms *MinStack) getmin() int {
	if len(ms.minStack) == 0 {
		panic("stack is empty.")
	}
	return ms.minStack[len(ms.minStack)-1]
}

// 描述
// 实现一个特殊功能的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。
//
// 输入描述
// 第一行输入一个整数 N 表示对栈进行的操作总数。
// 下面N行每行输入一个字符串 S 表示操作的种类。
// 如果 S 为 "push" 则后面还有一个整数 X 表示向栈里压入整数 X。
// 如果 S 为 "pop" 则表示弹出栈顶操作。
// 如果 S 为 "getMin" 则表示询问当前栈中的最小元素是多少。
//
// 输出描述
// 对于每个 getMin 操作，输出一行表示当前栈中的最小元素是多少。
//
// 备注
// 1 <= N <= 1000000
// -1000000 <= X <= 1000000
// 数据保证没有不合法的操作
func main() {
	sc := bufio.NewScanner(os.Stdin)

	var n int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	ms := NewMinStack()
	for i := 0; i < n; i++ {
		sc.Scan()
		line := sc.Text()
		if strings.HasPrefix(line, "push") {
			val, _ := strconv.Atoi(strings.Split(line, " ")[1])
			ms.push(val)
		}
		if line == "pop" {
			ms.pop()
		}
		if line == "getMin" {
			fmt.Println(ms.getmin())
		}
	}
}
