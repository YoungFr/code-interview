package llist

//  ___    _____          ___    __    __    ______   .___________.    __    ___     ___
// |__ \  | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) | | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /  |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_   ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____| |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 141 - 环形链表
// https://leetcode.cn/problems/linked-list-cycle/description/

func HasCycle(head *ListNode) bool {
	// 使用哈希表记录访问过的节点 PASS
	// 指针从头节点开始不断往后移动
	// 如果遇到一个先前访问过的节点
	// 说明链表中有环且这个节点就是环的入口
	// 这段代码只要修改一下返回值就可以用来
	// 解决 LC 142 - 寻找链表中开始入环的第一个节点
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
