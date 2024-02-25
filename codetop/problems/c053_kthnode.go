package problems

// LCR 140 - 训练计划 II
// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/

func TrainingPlan2(head *ListNode, cnt int) *ListNode {
	p := head
	for i := 0; i < cnt; i++ {
		p = p.Next
	}
	ans := head
	for p != nil {
		ans = ans.Next
		p = p.Next
	}
	return ans
}
