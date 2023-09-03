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
	} else if val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) pop() int {
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
	return ms.minStack[len(ms.minStack)-1]
}
