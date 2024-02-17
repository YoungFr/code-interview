LC [Hot 100](https://leetcode.cn/studyplan/top-100-liked/) and More...

# 双指针

6. [三数之和](./tp/tp3.go) 排序 + 双指针 + 元素去重

# 滑动窗口

8. [无重复字符的最长子串](./slidingwin/sw8.go) 滑动窗口

# 矩阵

20. [旋转图像](./matrix/mat20.go) 根据矩阵各种对称的转移式进行推导

# 链表

22. [相交链表](./llist/llist22.go) 双指针获取两个链表的长度差

23. [反转链表](./llist/llist23.go) 迭代反转 + 递归反转

24. [回文链表](./llist/llist24.go) 快慢指针寻找中间节点 + 反转一半链表

25. [判断链表中是否有环](./llist/llist25.go) Floyd 判圈算法

26. [寻找链表中环的入口](./llist/llist26.go) 哈希表 + 快慢指针

27. [合并 2 个有序链表](./llist/llist27.go) 迭代合并 + 递归合并

28. 两数相加

29. [删除链表的倒数第 n 个节点](./llist/llist29.go) 没啥意义的题

30. [2 个一组翻转链表](./llist/llist30.go) 分组反转链表

31. [k 个一组翻转链表](./llist/llist31.go) 分组反转链表

32. 随机链表的复制

33. [排序链表](./llist/llist33.go) 对链表使用自顶向下和自底向上归并排序算法进行排序

34. [合并 k 个升序链表](./llist/llist34.go) 堆的插入（上浮）操作和删除（下沉）操作

35. [LRU 缓存](./llist/llist35.go) 哈希表 + 双链表 + 每次 `Get` 和 `Put` 操作都将缓存项移到队头

# 图论

51. [岛屿数量](./graph/graph51.go) 网格图的深度优先搜索

# 二分查找

63. [搜索插入位置](./bs/bs63.go) 二分查找找到数组中一个大于等于 target 的元素的下标

64. 搜索二维矩阵

65. [在排序数组中查找元素的第一个和最后一个位置](./bs/bs65.go)  

    利用 `bsearch(nums, target)` 和 `bsearch(nums, target+1) - 1` 寻找元素的第一个和最后一个位置

# 堆

74. [数组中的第 k 个最大元素](./heap/heap74.go) 堆的删除（下沉）操作 + 原地建堆 + 三向切分快速选择

75. [前 k 个高频元素](./heap/heap75.go) 建立大顶堆后进行 k 次删除获取前 k 大的值

76. [数据流的中位数](./heap/heap76.go) 大顶堆保存数组前半部分最大值 + 小顶堆保存数组后半部分最小值

# 动态规划

85. [零钱兑换](./dp/dp85.go) 数组元素 $f[i]$ 表示组成金额 $i$ 所需的最少硬币数

    $$
    f[i] = 
    \begin{cases}
    -1 \rightarrow {组成小于零的金额是不可能的} \\
    0 \space\space\space \rightarrow {组成的金额为零则不需要硬币} \\
    min(f[i - c_j])+1 \rightarrow \\ 
    \quad {枚举组成金额\space i \space的最后一枚硬币的面值\space c_j \space的大小} \\
    \quad {如果所有\space f[i-c_j] \space都等于\space -1 \space则无法组成大小为\space i \space的金额} \\
    \quad {否则结果是所有大于等于零的\space f[i-c_j] \space中的最小值再加一}
    \end{cases}
    $$

86. 单词拆分

# 技巧

97. [多数元素](./skills/skills97.go) Boyer-Moore 投票算法

98. [荷兰国旗问题](./skills/skills98.go) 以 1 为主元进行三向切分
