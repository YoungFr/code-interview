package main

//  ___    ____           ___    __    __    ______   .___________.    __    ___     ___
// |__ \  |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_   ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____| |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 206 - 反转链表
// https://leetcode.cn/problems/reverse-linked-list/description/

func ReverseList(head *ListNode) *ListNode {
	// // 迭代 PASS
	// // 遍历所有节点并插入到哑节点后边
	// dummy := &ListNode {
	//     Val:  0,
	//     Next: nil,
	// }
	// curr := head
	// for curr != nil {
	//     // 保存当前节点
	//     t := curr
	//     curr = curr.Next
	//     // 将 t 指针指向的节点插到哑节点后边
	//     t.Next = dummy.Next
	//     dummy.Next = t
	// }
	// return dummy.Next

	// // 递归一 PASS
	// // 如果头节点为空或者没有下一个节点直接返回
	// if head == nil || head.Next == nil {
	//     return head
	// }
	// // 将以 head.Next 为头节点的链表反转
	// ans := reverseList(head.Next)
	// // 将指针 rear 指向反转后的链表的最后一个节点
	// rear := ans
	// for rear.Next != nil {
	//     rear = rear.Next
	// }
	// // 然后把 head 节点连到反转后的链表后边
	// rear.Next = head
	// rear.Next.Next = nil
	// return ans

	// 递归二 PASS
	// 上边的递归一版本能够通过测试但是时间非常慢
	// 因为我们每次都要遍历整个链表来找到反转后的链表的最后一个节点
	// 事实上递归没有改变 head.Next 的指向
	// 而 head.Next 指向的恰好是反转后的链表的最后一个节点
	// 所以只要把上边的 rear 换成 head.Next 就可以通过
	if head == nil || head.Next == nil {
		return head
	}
	// head ->  1 <- 2 <- ... <- n
	ans := ReverseList(head.Next)
	// head <-> 1 <- 2 <- ... <- n
	head.Next.Next = head
	// head <-  1 <- 2 <- ... <- n
	head.Next.Next.Next = nil // 这行等价于 head.Next = nil
	return ans
}
