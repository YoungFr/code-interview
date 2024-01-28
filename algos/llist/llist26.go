package llist

//  ___      __           ___    __    __    ______   .___________.    __    ___     ___
// |__ \    / /          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//    ) |  / /_         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//   / /  | '_ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  / /_  | (_) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____|  \___/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 142 - 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/description/

func DetectCycle(head *ListNode) *ListNode {
	// 参照 llist25.go 中第一种方法的注释
	// 使用哈希表记录访问过的节点
	// 如果遇到一个先前访问过的节点
	// 说明链表中有环且这个节点就是环的入口
	vis := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if vis[p] {
			return p
		}
		vis[p] = true
		p = p.Next
	}
	return nil

	// 快慢双指针解法
	// TODO
}
