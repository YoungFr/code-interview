package problems

// LC 143 - 重排链表
// https://leetcode.cn/problems/reorder-list/

func ReorderList(head *ListNode) {
	// 节点数 0, 1, 2 => 直接返回
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 找到链表的中间节点
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开左右两端链表的连接
	lRear := slow
	rHead := slow.Next
	lRear.Next = nil

	// 反转链表
	var reverseList func(head *ListNode) *ListNode
	reverseList = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		ans := reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
		return ans
	}

	l := head
	// 右半段链表反转后的头节点
	r := reverseList(rHead)

	// 将 l 和 r 两个链表中的节点交替合并
	var tl, tr *ListNode
	for l != nil && r != nil {
		// O -> O -> O -> nil
		// ^    ^
		// l    tl
		//
		// Q -> Q -> Q -> nil
		// ^    ^
		// r    tr
		tl = l.Next
		tr = r.Next

		// O    O -> O -> nil
		// |    ^
		// |   l/tl
		// ↓
		// Q -> Q -> Q -> nil
		// ^    ^
		// r    tr
		l.Next = r
		l = tl

		// O -> Q -> O -> O -> nil
		//           ^
		//          l/tl
		//
		//           Q -> Q -> nil
		//           ^
		//          r/tr
		r.Next = l
		r = tr
	}
}
