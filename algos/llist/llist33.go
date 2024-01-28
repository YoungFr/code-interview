package llist

//  ____    ____           ___    __    __    ______   .___________.    __    ___     ___
// |___ \  |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) |   __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 148 - 排序链表
// https://leetcode.cn/problems/sort-list/description/

func SortList(head *ListNode) *ListNode {
	// 分析使用的方法
	// 要做到在 O(NlogN) 复杂度下排序
	// 方法有快速排序、归并排序和堆排序
	// 快速排序和堆排序都需要利用数组随机访问的性质
	// 所以只剩归并排序
	// 归并排序又分为自顶向下和自底向上方法

	// 自顶向下方法 PASS
	// 时间复杂度 O(NlogN)
	// 空间复杂度 O(logN)

	// if head == nil || head.Next == nil {
	// 	return head
	// }

	// // 快慢指针找到链表的中间节点
	// slow := head
	// fast := head
	// for fast.Next != nil && fast.Next.Next != nil {
	// 	slow = slow.Next
	// 	fast = fast.Next.Next
	// }

	// // 链表的头节点作为左半段的头节点
	// lhead := head
	// // 中间节点作为左半段链表的尾节点
	// lrear := slow
	// // 中间节点的后继节点作为右半段的头节点
	// rhead := slow.Next
	// // 断开左右两段链表的连接
	// lrear.Next = nil

	// // 在 llist27.go 中实现的合并两个有序链表的函数
	// return MergeTwoLists(SortList(lhead), SortList(rhead))

	// 自底向上方法 PASS
	// 时间复杂度 O(NlogN)
	// 空间复杂度 O(1)

	// 为链表添加一个哑节点
	dummy := &ListNode{}
	dummy.Next = head

	// 首先需要计算链表的长度
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	// 自底向上归并
	// 每轮循环都将链表拆分为若干个长度为 size 的子链表
	// 最后一个子链表的长度可以比 size 小
	// 然后每两个链表进行合并
	// 不断加倍 size 并重复前两步直到 size 大于等于链表长度

	for size := 1; size < length; size = size + size {
		// 指针 curr 指向头节点
		curr := dummy.Next
		// 尾指针指向哑节点
		// 每轮归并后的链表都会重新连到哑节点的后边
		rear := dummy

		// 每轮循环的目标都是找到两个长为 size 的子链表
		// 然后将它们合并后连到 rear 后边
		for curr != nil {
			// 1. 第一个长为 size 的链表的头节点
			lhead := curr

			// 2. 将指针 curr 后移 size-1 次后
			//    在 lhead 和 curr 之间就有 size 个节点
			//    在移动过程中要注意 curr 不为空
			//    因为第一段链表的长度不足 size 也是有可能的
			//    比如当 length = 5, size = 2 时第二次进入循环的情况
			for i := 0; i < size-1 && curr != nil; i++ {
				curr = curr.Next
			}

			// 3. 第一个长为 size 的链表的尾节点
			lrear := curr

			// 4. 第二个子链表可能为空 => 先假设其为空
			var rhead *ListNode

			// 5. 如果指针 curr 不为空
			//    那么第二个长为 size 的链表的头节点是 curr 的后继节点
			if curr != nil {
				curr = curr.Next
				rhead = curr
			}

			// 6. 同第 2 步
			for i := 0; i < size-1 && curr != nil; i++ {
				curr = curr.Next
			}
			rrear := curr

			// 7. 先移动 curr 指针
			//    因为我们后续要断开找到的两段链表和后边的连接
			if curr != nil {
				curr = curr.Next
			}
			if lrear != nil {
				lrear.Next = nil
			}
			if rrear != nil {
				rrear.Next = nil
			}

			// 当 length = 5, size = 2 的链表第一次进入循环
			// 并执行到这里时链表的形式如下
			// dummy -> O --> O -> nil O --> O -> nil O -> nil
			//   ^      ^     ^        ^     ^        ^
			//   |      |     |        |     |        |
			//  rear  lhead lrear    rhead rrear     curr

			// 合并后连到 rear 后边
			rear.Next = MergeTwoLists(lhead, rhead)

			// 再将指针 rear 指向最后一个节点
			for rear.Next != nil {
				rear = rear.Next
			}
		}
	}

	return dummy.Next
}
