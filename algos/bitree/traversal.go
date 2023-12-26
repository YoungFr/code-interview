package main

// .______   .______       _______   ______   .______       _______   _______ .______
// |   _  \  |   _  \     |   ____| /  __  \  |   _  \     |       \ |   ____||   _  \
// |  |_)  | |  |_)  |    |  |__   |  |  |  | |  |_)  |    |  .--.  ||  |__   |  |_)  |
// |   ___/  |      /     |   __|  |  |  |  | |      /     |  |  |  ||   __|  |      /
// |  |      |  |\  \----.|  |____ |  `--'  | |  |\  \----.|  '--'  ||  |____ |  |\  \----.
// | _|      | _| `._____||_______| \______/  | _| `._____||_______/ |_______|| _| `._____|

// LC 144 - 二叉树先序遍历的迭代式算法
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
func PreorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0)
	ans := make([]int, 0)
	// 1. 将头结点压入栈中
	if root != nil {
		stk = append(stk, root)
	}
	for len(stk) > 0 {
		// 2. 弹出栈顶节点并打印值
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		ans = append(ans, top.Val)
		// 3. 将栈顶节点的右孩子（不为空）压入栈中
		if top.Right != nil {
			stk = append(stk, top.Right)
		}
		// 4. 将栈顶节点的左孩子（不为空）压入栈中
		if top.Left != nil {
			stk = append(stk, top.Left)
		}
	}
	return ans
}

//  __  .__   __.   ______   .______       _______   _______ .______
// |  | |  \ |  |  /  __  \  |   _  \     |       \ |   ____||   _  \
// |  | |   \|  | |  |  |  | |  |_)  |    |  .--.  ||  |__   |  |_)  |
// |  | |  . `  | |  |  |  | |      /     |  |  |  ||   __|  |      /
// |  | |  |\   | |  `--'  | |  |\  \----.|  '--'  ||  |____ |  |\  \----.
// |__| |__| \__|  \______/  | _| `._____||_______/ |_______|| _| `._____|

// LC 94 - 二叉树中序遍历的迭代式算法
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
func InorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0)
	ans := make([]int, 0)
	for len(stk) > 0 || root != nil {
		if root != nil {
			// 1. 压入根节点。然后从根节点开始，不断压入当前节点的左孩子
			stk = append(stk, root)
			root = root.Left
		} else {
			// 2. 当左孩子为空时，弹出栈顶元素并打印值
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans = append(ans, top.Val)
			// 3. 然后将栈顶元素的右孩子设为当前节点，重复步骤 1 的操作
			root = top.Right
		}
	}
	return ans
}

// .______     ______        _______.___________.  ______   .______       _______   _______ .______
// |   _  \   /  __  \      /       |           | /  __  \  |   _  \     |       \ |   ____||   _  \
// |  |_)  | |  |  |  |    |   (----`---|  |----`|  |  |  | |  |_)  |    |  .--.  ||  |__   |  |_)  |
// |   ___/  |  |  |  |     \   \       |  |     |  |  |  | |      /     |  |  |  ||   __|  |      /
// |  |      |  `--'  | .----)   |      |  |     |  `--'  | |  |\  \----.|  '--'  ||  |____ |  |\  \----.
// | _|       \______/  |_______/       |__|      \______/  | _| `._____||_______/ |_______|| _| `._____|

// LC 145 - 二叉树后序遍历的迭代式算法
// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
func PostorderTraversal(root *TreeNode) []int {
	s1 := make([]*TreeNode, 0)
	s2 := make([]*TreeNode, 0)
	// 1. 将根节点压入 s1 栈
	if root != nil {
		s1 = append(s1, root)
	}
	for len(s1) > 0 {
		// 2. 弹出 s1 栈的栈顶节点并压入 s2 栈
		top := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, top)
		// 3. 将栈顶节点的左孩子（不为空）压入 s1 栈
		if top.Left != nil {
			s1 = append(s1, top.Left)
		}
		// 4. 将栈顶节点的右孩子（不为空）压入 s1 栈
		if top.Right != nil {
			s1 = append(s1, top.Right)
		}
	}
	ans := make([]int, 0)
	// 5. 依次弹出 s2 中的元素就是后序遍历的结果
	for len(s2) > 0 {
		ans = append(ans, s2[len(s2)-1].Val)
		s2 = s2[:len(s2)-1]
	}
	return ans
}
