package main

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

//  ___    _____          ___    __    __    ______   .___________.    __    ___     ___
// |__ \  | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) | | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /  |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_   ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____| |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 141 - 环形链表
// https://leetcode.cn/problems/linked-list-cycle/description/

func HasCycle(head *ListNode) bool {
	// // 使用哈希表记录访问过的节点 PASS
	// // 指针从头节点开始不断往后移动
	// // 如果遇到一个先前访问过的节点
	// // 说明链表中有环且这个节点就是环的入口
	// // 这段代码只要修改一下返回值就可以用来
	// // 解决 LC 142 - 寻找链表中开始入环的第一个节点
	// vis := make(map[*ListNode]bool)
	// p := head
	// for p != nil {
	//     if vis[p] {
	//         return true
	//     }
	//     vis[p] = true
	//     p = p.Next
	// }
	// return false

	// Floyd 判圈算法 PASS
	// 如果链表中有环
	// 一旦指针进入环内就会一直在环内移动
	// 定义一快一慢的两个指针
	// 在两个指针进入环中以后它们一定会在某个节点相遇

	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != slow {
		// 快指针要移动两步
		// 所以要保证 fast 和 fast.Next 不为空
		// 因为快指针在慢指针前边
		// 所以也可以保证 slow 不为空
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
