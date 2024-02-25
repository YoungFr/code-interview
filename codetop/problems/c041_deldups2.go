package problems

// LC 82 - 删除排序链表中的重复元素 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/

func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	p := head

	for p != nil {
		// 统计和当前节点值相等的节点的个数
		equal := 0
		v := p.Val
		for p != nil && p.Val == v {
			equal++
			p = p.Next
		}
		if equal > 1 {
			// 删除所有节点
			prev.Next = p
		} else {
			// 保留当前节点
			prev = prev.Next
		}
	}

	return dummy.Next
}
