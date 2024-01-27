package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LC 232 - 用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/description/

// 描述
// 用两个栈实现队列，支持队列的基本操作。
//
// 输入描述
// 第一行输入一个整数 N 表示对队列进行的操作总数。
// 下面 N 行每行输入一个字符串 S 表示操作的种类。
// 如果 S 为 "add" 则后面还有一个整数 X 表示向队列尾部加入整数 X 。
// 如果 S 为 "poll" 则表示弹出队列头部操作。
// 如果 S 为 "peek" 则表示询问当前队列中头部元素是多少。
//
// 输出描述
// 对于每一个为 "peek" 的操作，输出一行表示当前队列中头部元素是多少。
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

	q := NewTwoStacksQueue()
	for i := 0; i < n; i++ {
		sc.Scan()
		line := sc.Text()
		switch {
		case strings.HasPrefix(line, "add"):
			val, _ := strconv.Atoi(strings.Split(line, " ")[1])
			q.add(val)
		case line == "poll":
			q.poll()
		case line == "peek":
			fmt.Println(q.peek())
		default:
			panic("invalid input")
		}
	}
}

func NewTwoStacksQueue() *TwoStacksQueue {
	return &TwoStacksQueue{
		pushstk: make([]int, 0),
		popstk:  make([]int, 0),
	}
}

// 只访问 pushstk 和 popstk 的最后一个元素来保证这是两个栈
type TwoStacksQueue struct {
	pushstk []int
	popstk  []int
}

func (q *TwoStacksQueue) push2pop() {
	// 当 popstk 为空时将 pushstk 中的元素全部 “倒入” popstk 中
	if len(q.popstk) == 0 {
		for len(q.pushstk) > 0 {
			q.popstk = append(q.popstk, q.pushstk[len(q.pushstk)-1])
			q.pushstk = q.pushstk[:len(q.pushstk)-1]
		}
	}
}

func (q *TwoStacksQueue) add(val int) {
	q.pushstk = append(q.pushstk, val)
	q.push2pop()
}

func (q *TwoStacksQueue) poll() int {
	if len(q.popstk) == 0 {
		panic("queue is empty.")
	}
	ans := q.popstk[len(q.popstk)-1]
	q.popstk = q.popstk[:len(q.popstk)-1]
	q.push2pop()
	return ans
}

func (q *TwoStacksQueue) peek() int {
	if len(q.popstk) == 0 {
		panic("queue is empty.")
	}
	return q.popstk[len(q.popstk)-1]
}
