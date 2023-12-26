package main

//  ___    _  _            ___    __    __    ______   .___________.    __    ___     ___
// |__ \  | || |          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) | | || |_        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /  |__   _|      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_     | |       /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____|    |_|      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 234 - 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description/

func IsPalindrome(head *ListNode) bool {
	// 这道题是非常重要的一道题
	// 因为它涉及寻找链表的中间节点、反转链表、判断回文等操作
	// 这些代码都需要能在面试时现场写出来

	// 只有一个节点时是回文链表
	if head.Next == nil {
		return true
	}

	// 使用快慢指针寻找链表的中间节点
	endOfFirstHalf := func(head *ListNode) *ListNode {
		slow := head
		fast := head
		// 对于长度是奇数的链表
		// 指针 slow 最终指向最中间的节点
		// 对于长度是偶数的链表
		// 指针 slow 最终指向的是前半部分的最后一个节点
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	// LC 206 - 反转链表
	reverseList := func(head *ListNode) *ListNode {
		dummy := &ListNode{}
		curr := head
		for curr != nil {
			t := curr
			curr = curr.Next
			t.Next = dummy.Next
			dummy.Next = t
		}
		return dummy.Next
	}

	// 找到前半部分链表的尾节点
	firstHalfEnd := endOfFirstHalf(head)
	// 反转后半部分链表
	secondHalfStart := reverseList(firstHalfEnd.Next)
	firstHalfEnd.Next = nil

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	// 根据前边 endOfFirstHalf 函数返回的结果我们可以知道
	// 链表长度是偶数 => 前后两条链表的长度相等
	// 链表长度是奇数 => 前半部分链表的长度比后半部分的长度多一
	// 所以将 p2 != nil 作为循环条件
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
