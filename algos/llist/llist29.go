package llist

//  ___     ___           ___    __    __    ______   .___________.    __    ___     ___
// |__ \   / _ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) | | (_) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /   \__, |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_     / /      /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____|   /_/      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 19 - 删除链表的倒数第 N 个节点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 两趟遍历 PASS
	// 删除倒数第 n 个节点 <=> 删除正数第 sz-n+1 个节点
	// 先计算节点个数
	// sz := 0
	// p := head
	// for p != nil {
	//     sz++
	//     p = p.Next
	// }
	// // 删除某个节点时需要指向该节点的前驱节点的指针
	// // 由于头节点没有前驱节点
	// // 所以添加一个哑节点指向头节点
	// // 作为头节点的前驱节点
	// dummy := &ListNode{}
	// dummy.Next = head
	// // 将指针 prev 向后移动 sz-n 个位置
	// // 使其指向要删除的节点的前驱节点
	// prev := dummy
	// for i := 0; i < sz-n; i++ {
	//     prev = prev.Next
	// }
	// // 删除某个节点即
	// // 将该节点的前驱节点的指针域指向该节点的后继节点
	// prev.Next = prev.Next.Next
	// // 哑节点的后继节点是链表的头节点
	// return dummy.Next

	// 一趟遍历 PASS
	// 先将一个指针指向第 n 个节点
	// 那么它再移动 sz-n 次就指向第 sz 个节点
	// 此时让第二个指针指向第 1 个节点
	// 然后两个指针同时移动
	// 当第一个指针指向第 sz 个节点时第二个指针就指向了第 sz-n+1 个节点

	// 为了方便编写代码
	// 将第 0 个节点也就是哑节点作为两个指针移动的起始节点
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	// 从第 0 个节点开始移动 n 次后
	// 指针 fast 指向第 n 个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	prev := dummy
	// 将指针 fast 移动到指向最后一个节点时
	// 指针 prev 指向第 sz-n 个节点也就是要删除的节点的前驱节点
	for fast.Next != nil {
		prev = prev.Next
		fast = fast.Next
	}
	// 删除某个节点即
	// 将该节点的前驱节点的指针域指向该节点的后继节点
	prev.Next = prev.Next.Next
	return dummy.Next
}
