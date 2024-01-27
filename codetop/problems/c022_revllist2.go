package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// LC 92 - 反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/description/

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	p := dummy

	// 移动指针 p 使其指向位置 left 节点的前驱节点
	for i := 1; i < left; i++ {
		p = p.Next
	}

	// 断开位置 left 节点和其前驱节点的连接
	tp := p.Next
	p.Next = nil

	// 现在 curr 指向了位置 left 的节点
	curr := tp

	// 指针 rear 也指向位置 left 的节点
	// 它同时也是反转完成后的最后一个节点
	rear := tp

	// 将从位置 left 到位置 right 的节点插到指针 p 后边
	for i := 1; i <= right-left+1; i++ {
		t := curr
		curr = curr.Next
		t.Next = p.Next
		p.Next = t
	}

	// 将反转后的链表的最后一个节点和 curr 指向的链表相连
	rear.Next = curr
	return dummy.Next
}
