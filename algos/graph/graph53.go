package graph

//  _____   ____           ___    __    __    ______   .___________.    __    ___     ___
// | ____| |___ \         /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
// | |__     __) |       /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
// |___ \   |__ <       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 207 - 课程表
// https://leetcode.cn/problems/course-schedule/description/

func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 任何一个有向无环图（Directed Acyclic Graph, DAG）都有拓扑排序序列
	// 拓扑排序的 Kahn 算法
	// 1. 找到入度为 0 的节点放入结果列表
	// 2. 删除该节点以及和该节点相连的所有边
	// 3. 重复步骤 1 和 2 直到找不到入度为 0 的节点
	//    此时如果结果列表中的元素个数和节点总数相等则能完成拓扑排序

	// 邻接表
	// 数组元素 adj[i] = [a1, a2, ..., an] 表示
	// 图中存在 i->a1, i->a2, ..., i->an 共 n 条边

	// Example => 图中有 0->1, 0->2, 1->2 三条有向边
	//   0 -> [1, 2]
	//   1 -> [2]
	//   2 -> []
	adj := make([][]int, numCourses)
	// 记录每个节点的入度
	indeg := make([]int, numCourses)

	// 元素 prerequisites[i] = [ai, bi] 表示要学习 ai 课程必须先学习 bi 课程
	// 所以有向边是 bi->ai 也就是 prerequisites[i][1]->prerequisites[i][0] 边
	for _, prerequisite := range prerequisites {
		adj[prerequisite[1]] = append(adj[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}

	// 变量 n 记录最终放入结果列表中元素的个数
	// 这里我们只需要判断 n 最终是否和节点总数相等即可，输出最终拓扑排序结果的题目参见
	// 第 210 题 https://leetcode.cn/problems/course-schedule-ii/description/
	n := 0
	q := make([]int, 0)

	// 遍历入度表，找到所有初始入度为 0 的节点放入队列
	for v, d := range indeg {
		if d == 0 {
			q = append(q, v)
			n++
		}
	}

	// 如果还有入度为 0 的节点则循环
	for len(q) > 0 {
		for sz := len(q); sz > 0; sz-- {
			v := q[0] // 取出队头元素
			q = q[1:]

			// 删除和节点 v 相连的所有边的实现方法
			//   1. 将所有节点 v 指向的节点的入度减一
			//   2. 将入度被减为 0 的节点放入队列
			for _, u := range adj[v] {
				indeg[u]--
				if indeg[u] == 0 {
					q = append(q, u)
					n++
				}
			}
		}
	}

	return n == numCourses
}
