package llist

//  ____    __          ___    __    __    ______   .___________.    __    ___     ___
// |___ \  /_ |        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) |  | |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <   | |      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  | |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/   |_|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 25 - k 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 前言
	// 本题是 LC 24 - 两两交换链表中的节点的通用版本
	// 只要将下边代码中的 k 都改成 2 就可以通过第 24 题

	// 当 k 为 1 时或者
	// 节点数目为 0 或 1 时直接返回头节点
	if k == 1 || head == nil || head.Next == nil {
		return head
	}

	// LC 206 - 反转链表
	// 使用递归的方式反转链表并返回反转后链表的头结点
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

	// 增加一个哑节点并将尾指针指向它
	// 稍后翻转后的链表将连在尾指针 rear 后边
	dummy := &ListNode{}
	rear := dummy

	// 翻转的思路概括
	// 1. 找到 k 个节点
	// 2. 将这 k 个节点和后边断开连接
	// 3. 翻转 k 个节点并连到尾指针后边
	// 4. 尾指针后移到最后一个节点
	// 5. 再将未处理的节点连到尾指针后边

	// 1. 指针 p 初始指向头结点
	p := head
outer:
	for {
		// 2. 用 h1 保存当前 p 的值
		h1 := p
		// 3. 然后将指针 p 后移 k-1 个位置
		//    此时指针 h1 到指针 p 之间就包含了 k 个节点
		//    在这个过程中要保证 p 和 p.Next 不为空
		//    如果这个条件不满足
		//    就说明链表末尾未处理的节点数已经不足 k 个
		//    将 h1 连在 rear 后边退出循环即可
		for i := 0; i < k-1; i++ {
			if p != nil && p.Next != nil {
				p = p.Next
			} else {
				rear.Next = h1
				break outer
			}
		}
		// 4. 现在指针 p 指向了链表中未处理的节点中的第 k 个
		//    用 h1r 保存 p 的值后将 p 后移
		//    这是为了将前 k 个节点和后边断开连接
		//    下边 3 行的顺序是固定的
		h1r := p
		p = p.Next
		h1r.Next = nil

		// 第一次进入循环
		// 执行到这里时链表的样式如下
		//
		//               +------- k ------+
		//     O -> nil  O -> O -> ... -> O -> nil  O -> O -> ... O -> nil
		//     ^         ^                ^         ^
		//     |         |                |         |
		// dummy/rear head/h1            h1r        p

		// 5. 将 h1 作为头节点的链表翻转并连到 rear 后边
		rear.Next = reverseList(h1)
		// 6. 指针 rear 后移 k 个位置
		//    指向以哑节点为头节点的链表的最后一个节点
		for i := 0; i < k; i++ {
			rear = rear.Next
		}
		// 7. 再将 p 指向的链表连到 rear 后边
		rear.Next = p
	}
	return dummy.Next
}
