package llist

//  ___    ______          ___    __    __    ______   .___________.    __    ___     ___
// |__ \  |____  |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) |     / /        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /     / /        /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_    / /        /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____|  /_/        /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 21 - 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 方法一 PASS
	//
	// 迭代合并 - 类似于合并两个有序数组所用的方法
	// dummy := &ListNode{}
	// p := dummy

	// p1 := list1
	// p2 := list2
	// for p1 != nil && p2 != nil {
	// 	if p1.Val <= p2.Val {
	// 		curr := p1
	// 		p1 = p1.Next
	// 		curr.Next = nil
	// 		p.Next = curr
	// 		p = p.Next
	// 	} else {
	// 		curr := p2
	// 		p2 = p2.Next
	// 		curr.Next = nil
	// 		p.Next = curr
	// 		p = p.Next
	// 	}
	// }
	// if p1 != nil {
	// 	p.Next = p1
	// }
	// if p2 != nil {
	// 	p.Next = p2
	// }
	// return dummy.Next

	// 方法二 PASS
	//
	// 递归合并
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}

	// 方法三 PASS
	//
	// 直接调用 LC 23 - 合并 k 个升序链表中实现的 MergeKLists 函数
	// return MergeKLists([]*ListNode{list1, list2})
}
