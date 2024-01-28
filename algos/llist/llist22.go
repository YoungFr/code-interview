package llist

//  ___    ___           ___    __    __    ______   .___________.    __    ___     ___
// |__ \  |__ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) |    ) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /    / /       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_   / /_      /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____| |____|    /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 160 - 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	// 假设两个链表 A 和 B 相交
	// 它们各自的长度分别为 la 和 lb 也就是各有 la 和 lb 个节点
	// 再假设从相交的节点（含）到最后一个节点（含）共 c 个节点
	// 那么有 la = a + c, lb = b + c
	// 不妨设 a >= b 那么链表 A 比链表 B 多 a-b 个节点
	// 如果我们能从链表 A 的第 a-b+1 个节点、链表 B 的第 1 个节点开始遍历
	// 就可以消除这种长度差 => 第一个相等的节点就是相交节点

	// 基于这种思路的代码 PASS
	// if headA == nil || headB == nil {
	//     return nil
	// }
	// var la, lb int
	// pa := headA
	// pb := headB
	// for pa != nil {
	//     la++
	//     pa = pa.Next
	// }
	// for pb != nil {
	//     lb++
	//     pb = pb.Next
	// }
	// if la >= lb {
	//     pa = headA
	//     for i := 0; i < la-lb; i++ {
	//         pa = pa.Next
	//     }
	//     pb = headB
	// } else {
	//     pb = headB
	//     for i := 0; i < lb-la; i++ {
	//         pb = pb.Next
	//     }
	//     pa = headA
	// }
	// for pa != pb {
	//     pa = pa.Next
	//     pb = pb.Next
	// }
	// return pa

	// 另一种更为巧妙的方法不需要获取链表的长度
	// 我们的核心是获取长度差 a-b 的大小
	// 假设指针 pa, pb 分别指向 A, B 链表的头节点
	// 让两个指针分别后移
	// 按照之前的假设当 pb 到达最后一个节点时 pa 还差 a-b 次移动到最后一个节点
	// 此时我们就得到了 a-b 这个差值
	// 然后让 pb 指向链表 A 的头节点
	// 两个指针继续后移当 pa 也移动到最后一个节点时 pb 正好在 A 链表上前移了 a-b 个节点
	// 此时再让 pa 指向链表 B 的头节点
	// 就达到了从链表 A 的第 a-b+1 个节点、链表 B 的第 1 个节点开始遍历的效果

	// 最终 pa 移动了 (a+c) + b 个节点
	// 最终 pb 移动了 (b+c) + a 个节点

	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
