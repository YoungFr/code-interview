package llist

//  ____    _  _            ___    __    __    ______   .___________.    __    ___     ___
// |___ \  | || |          /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) | | || |_        /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <  |__   _|      /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |    | |       /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/     |_|      /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 23 - 合并 k 个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/description/

func MergeKLists(lists []*ListNode) *ListNode {
	// 朴素方法 - 两两合并后的结果再和下一个链表合并 PASS
	// 时间击败  5.03%
	// 空间击败 26.13%

	// n := len(lists)
	// // 先处理下数组长度为 0 或 1 时的特殊情况
	// if n == 0 {
	// 	// 数组中没有元素时返回 nil 值
	// 	return nil
	// }
	// if n == 1 {
	// 	// 只有一个链表时直接返回其头节点即可
	// 	return lists[0]
	// }

	// // 定义合并两个有序链表的函数
	// // LC 21 - 用递归的方式合并两个有序链表
	// var merge func(*ListNode, *ListNode) *ListNode
	// merge = func(l1, l2 *ListNode) *ListNode {
	// 	if l1 == nil {
	// 		return l2
	// 	}
	// 	if l2 == nil {
	// 		return l1
	// 	}
	// 	if l1.Val <= l2.Val {
	// 		l1.Next = merge(l1.Next, l2)
	// 		return l1
	// 	} else {
	// 		l2.Next = merge(l1, l2.Next)
	// 		return l2
	// 	}
	// }

	// ans := merge(lists[0], lists[1])
	// // 两两合并后的结果再和下一个链表合并
	// for i := 2; i < n; i++ {
	// 	ans = merge(ans, lists[i])
	// }
	// return ans

	// 使用堆进行优化 PASS
	// 时间击败 97.57%
	// 空间击败 64.80%

	// 1. 先处理下数组长度为 0 时的特殊情况
	n := len(lists)
	if n == 0 {
		return nil
	}

	// 2. 定义哑节点并将尾指针指向它
	//    后序的节点将连到尾指针后边
	dummy := &ListNode{}
	rear := dummy

	// 只要有可能
	// 我们都自己实现堆中元素的上浮和下沉方法来维护堆的性质
	// 这样做有两个好处
	// 一是不用实现 heap.Interface 接口
	// 二是可以进一步熟悉堆的操作

	// 3. 定义一个数组作为优先队列
	//    优先队列中存放的是每个链表的头节点的地址
	pq := make([]*ListNode, 0)
	// 4. 一个整型变量表示优先队列的大小
	//    在任何时刻优先队列中的元素都是从 pq[0] 到 pq[sz-1] 的所有元素
	//    变量 sz 的大小与 C++ 和 Java 中的 size 方法的返回值是相同的
	sz := 0

	// 5. 将所有不为空的链表的头节点地址放入优先队列
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			// 6. 向堆中插入元素
			//    这三步相当于 C++ 中的 push 方法和 Java 中的 add 方法

			// 6.1 堆的大小加一
			sz++
			// 6.2 将元素添加到数组末尾
			pq = append(pq, lists[i])
			// 6.3 对末尾的元素执行上浮操作
			up(pq, sz-1)
		}
	}

	// 当优先队列不为空时循环
	// 这相当于 C++ 中的 !pq.empty() 条件和 Java 中的 !pq.isEmpty() 条件
	for sz > 0 {
		// 7. 取出队头元素
		//    变量 top 保存当前所有头节点中值最小的那个的地址
		//    这相当于 C++ 中的 top 方法和 Java 中的 peek 方法
		top := pq[0]

		// 每一步操作都保证 pq[0] 不为空
		// 所以可以直接将 pq[0] 的值设置为 pq[0] 的下一个节点的地址
		pq[0] = pq[0].Next

		if pq[0] == nil {
			// 8. 当 pq[0] 为空时说明某一条链表的所有元素都处理完了
			//    需要删除队头元素
			//    这三步相当于 C++ 中的 pop 方法和 Java 中的 poll 方法

			// 8.1 将第一个和最后一个元素交换
			pq[0], pq[sz-1] = pq[sz-1], pq[0]
			// 8.2 堆的大小减一
			sz--
			// 8.3 对第一个元素执行下沉操作
			down(pq, 0, sz)
		} else {
			// 9. 当 pq[0] 不为空时说明该链表还有未处理的节点
			//    将 top.Next 设为空断开和后边节点的连接
			top.Next = nil
			// A. 改变了队头元素的值
			//    对第一个元素执行下沉操作
			down(pq, 0, sz)
		}

		// B. 将 top 指向的节点连到 rear 后边
		//    然后向后移动 rear 指针
		rear.Next = top
		rear = rear.Next
	}
	return dummy.Next
}

// 上浮
func up(a []*ListNode, i int) {
	child := i
	for {
		root := (child - 1) / 2
		// 小顶堆 => 小于等于
		// 大顶堆 => 大于等于
		if child == root || a[root].Val <= a[child].Val {
			break
		}
		a[root], a[child] = a[child], a[root]
		child = root
	}
}

// 下沉
func down(a []*ListNode, i int, size int) {
	root := i
	for {
		child := 2*root + 1
		if child >= size {
			break
		}
		// 小顶堆 => 小于
		// 大顶堆 => 大于
		if child+1 < size && a[child+1].Val < a[child].Val {
			child++
		}
		// 小顶堆 => 小于等于
		// 大顶堆 => 大于等于
		if a[root].Val <= a[child].Val {
			break
		}
		a[root], a[child] = a[child], a[root]
		root = child
	}
}
