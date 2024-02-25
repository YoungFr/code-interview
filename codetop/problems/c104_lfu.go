package problems

import "container/list"

// LC 104 - LFU 缓存
// https://leetcode.cn/problems/lfu-cache/description/

// 缓存项
type entry struct {
	key   int
	value int
	freq  int // 使用频率
}

type LFUCache struct {
	// 键为频率值为一个双向链表
	// 存放所有频率相同的缓存项
	freq map[int]*list.List

	// 键为索引项中的 key 字段
	// 值为索引项在双链表中的内存地址
	keys map[int]*list.Element

	// 当前最少使用的频率
	minFreq int

	// 缓存容量
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freq:     make(map[int]*list.List),
		keys:     make(map[int]*list.Element),
		minFreq:  0,
		capacity: capacity,
	}
}

func (c *LFUCache) Get(key int) int {
	if ele, ok := c.keys[key]; ok {
		// 1. 获取旧有频率和键对应的值
		oldFreq := ele.Value.(entry).freq
		value := ele.Value.(entry).value

		// 2. 从旧有频率对应的链表中删除
		c.freq[oldFreq].Remove(ele)
		// 3. 如果删除后链表长度为 0 则删除旧有频率键
		if c.freq[oldFreq].Len() == 0 {
			delete(c.freq, oldFreq)
			// 3.1 如果被删除的旧有频率等于当前最少使用的频率
			//     则将当前最少使用的频率加一
			if oldFreq == c.minFreq {
				c.minFreq++
			}
		}

		// 4. 插入到 newFreq(oldFreq+1) 对应的链表头部并更新 keys 信息
		newFreq := oldFreq + 1
		if c.freq[newFreq] == nil {
			c.freq[newFreq] = list.New()
		}
		c.keys[key] = c.freq[newFreq].PushFront(entry{key, value, newFreq})

		return value
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if ele, ok := c.keys[key]; !ok {
		// 1. 插入前缓存已满需要删除 minFreq 对应的链表的尾部元素
		if len(c.keys) == c.capacity {
			b := c.freq[c.minFreq].Back()
			delete(c.keys, b.Value.(entry).key)
			c.freq[c.minFreq].Remove(b)
			if c.freq[c.minFreq].Len() == 0 {
				delete(c.freq, c.minFreq)
			}
		}

		// 2. 新插入的项将其插入 freq 中 1 对应的链表的头部并更新 keys 和 minFreq 信息
		if c.freq[1] == nil {
			c.freq[1] = list.New()
		}
		c.keys[key] = c.freq[1].PushFront(entry{key, value, 1})
		c.minFreq = 1
	} else {
		// 3. 键 key 存在时和 Get 操作唯一的区别是
		//    要以新值 value 插入 newFreq(oldFreq+1) 对应的链表
		oldFreq := ele.Value.(entry).freq
		c.freq[oldFreq].Remove(ele)
		if c.freq[oldFreq].Len() == 0 {
			delete(c.freq, oldFreq)
			if oldFreq == c.minFreq {
				c.minFreq++
			}
		}
		newFreq := oldFreq + 1
		if c.freq[newFreq] == nil {
			c.freq[newFreq] = list.New()
		}
		c.keys[key] = c.freq[newFreq].PushFront(entry{key, value, newFreq})
	}
}
