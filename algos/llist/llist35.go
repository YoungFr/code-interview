package llist

import "container/list"

// ____    _____          ___    __    __    ______   .___________.    __    ___     ___
// |___ \  | ____|        /  /   |  |  |  |  /  __  \  |           |   /_ |  / _ \   / _ \
//   __) | | |__         /  /    |  |__|  | |  |  |  | `---|  |----`    | | | | | | | | | |
//  |__ <  |___ \       /  /     |   __   | |  |  |  |     |  |         | | | | | | | | | |
//  ___) |  ___) |     /  /      |  |  |  | |  `--'  |     |  |         | | | |_| | | |_| |
// |____/  |____/     /__/       |__|  |__|  \______/      |__|         |_|  \___/   \___/

// LC 146 - LRU 缓存
// https://leetcode.cn/problems/lru-cache/description/

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	maxEntries int
	ll         *list.List
	cache      map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxEntries: capacity,
		ll:         list.New(),
		cache:      make(map[int]*list.Element),
	}
}

// 缓存项
type entry struct {
	key   int
	value int
}

func (l *LRUCache) Put(key int, value int) {
	// 1. 键 key 在缓存中存在
	//    把缓存项移到队头并修改键对应的值
	if ee, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}

	// 2. 键 key 在缓存中不存在
	//    新建一个缓存项插入队头
	//    并将键和缓存项加入哈希表
	ele := l.ll.PushFront(&entry{key, value})
	l.cache[key] = ele

	// 3. 插入操作导致容量超过 capacity 则删除最久未使用的缓存项
	if l.ll.Len() > l.maxEntries {
		// 如果队尾元素不为空
		// 则从链表中删除该缓存项、从哈希表中删除键
		ele := l.ll.Back()
		l.ll.Remove(ele)
		delete(l.cache, ele.Value.(*entry).key)
	}
}

func (l *LRUCache) Get(key int) int {
	// 1. 键 key 在缓存中存在
	//    把缓存项移到队头并返回键对应的值
	if ele, hit := l.cache[key]; hit {
		l.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}

	// 2. 键 key 在缓存中不存在直接返回 -1 即可
	return -1
}
