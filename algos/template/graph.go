package template

import "math"

// 并查集模板
type UF struct {
	C int   // 连通分量个数
	V []int // 节点数组
}

func NewUF(n int) UF {
	V := make([]int, n)
	for i := 0; i < n; i++ {
		V[i] = i
	}
	return UF{C: n, V: V}
}

func (u *UF) Find(x int) int {
	if u.V[x] != x {
		u.V[x] = u.Find(u.V[x])
	}
	return u.V[x]
}

func (u *UF) Union(x int, y int) {
	rx, ry := u.Find(x), u.Find(y)
	if rx != ry {
		u.V[rx] = ry
		u.C--
	}
}

// LC 743 - 网络延迟时间
// https://leetcode.cn/problems/network-delay-time/description/

const inf = math.MaxInt >> 1

// n                   节点个数
// base                第一个节点从几开始编号
// weightedEdges       图中所有的加权边
// weightedEdges[i][0] 起点
// weightedEdges[i][1] 终点
// weightedEdges[i][2] 权值
// directed            是否为有向图
// start               起点
func Dijkstra(n int, base int, weightedEdges [][]int, directed bool, start int) []int {
	// 1. 如果存在从 x 到 y 的边则 graph[x][y] = weight 否则为无穷大
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				graph[i][j] = inf
			}
		}
	}
	for i := 0; i < len(weightedEdges); i++ {
		e := weightedEdges[i]
		x, y, w := e[0], e[1], e[2]
		x -= base
		y -= base
		graph[x][y] = w
		if !directed {
			graph[y][x] = w
		}
	}

	// 2. 初始化起点到自己的距离为零；到其他节点的距离为无穷大
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[start-base] = 0

	used := make([]bool, n)
	for i := 0; i < n; i++ {
		// 3. 在线性时间复杂度内从没有访问过的节点中找到距离最小的节点
		x := 0
		for i := 0; i < n; i++ {
			if !used[i] {
				x = i
				break
			}
		}
		for y, u := range used {
			if !u && dist[y] < dist[x] {
				x = y
			}
		}
		// 4. 使用该节点进行松弛操作
		for y, w := range graph[x] {
			dist[y] = min(dist[y], dist[x]+w)
		}
		// 5. 将该节点标为已访问
		used[x] = true
	}
	return dist
}
