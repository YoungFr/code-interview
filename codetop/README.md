# CodeTop

来自 [codetop](https://codetop.cc/home) 的按照出现频率高低排列的面试题目。

1. [无重复字符的最长子串](../algos/slidingwin/sw8.go) 滑动窗口

1. [反转链表](../algos/llist/llist23.go) 迭代反转 or 递归反转

1. [LRU 缓存机制](../algos/llist/llist35.go) 哈希表 + 双链表 + 每次 `Get` 和 `Put` 操作都将缓存项移到队头

1. [数组中的第 k 个最大元素](../algos/heap/heap74.go) 堆的删除（下沉）操作 + 原地建堆 or 三向切分快速选择

1. [k 个一组反转链表](../algos/llist/llist31.go) 分组反转链表

1. [三数之和](../algos/tp/tp3.go) 排序 + 双指针 + 元素去重

1. [最大子数组和](../algos/array/array13.go) 动态规划 - 详见 [README](../algos/README.md) 第 13 题

1. [手撕快速排序](./problems/c008_qsort.go) 切分函数的实现 + 主元的选取

1. [合并 2 个有序链表](../algos/llist/llist27.go) 迭代合并 or 递归合并

1. [两数之和](../algos/hash/hash1.go) 构建键为元素、值为下标的哈希表

1. [最长回文子串](../algos/mddp/mddp93.go) 中心扩展法

1. [二叉树的层序遍历](./problems/c038_bitreerightview.go) 队列的经典应用

1. [搜索旋转排序数组](../algos/bs/bs66.go) 先用二分查找找到最小值再确定在哪个区间上使用二分查找搜索目标值

1. [岛屿数量](../algos/graph/graph51.go) 网格图的深度优先搜索

1. [有效的括号](../algos/stkandq/stk61.go) 利用栈进行括号匹配

1. [环形链表](../algos/llist/llist25.go) Floyd 判圈算法

1. [买卖股票的最佳时机](../algos/greedy/greedy77.go) 用一个变量维护前 `i-1` 天的最小价格作为买入价格，然后枚举卖出价格来获取最大利润

1. 二叉树的最近公共祖先

1. [合并两个有序数组](./problems/c019_merge.go) 从后往前遍历确定应该使用哪个数组中的数字

1. [全排列](../algos/backtrack/bt55.go) 排列型回溯

1. [二叉树的锯齿形层次遍历](./problems/c021_zigzaglevel.go) 层次遍历中添加下一层节点时使用栈来实现锯齿形遍历

1. [反转链表 II](./problems/c022_revllist2.go) 需要部分反转的链表的头节点的 `Next` 域在反转后指向剩余未反转的部分

1. [螺旋矩阵](../algos/matrix/mat19.go) 按圈层从外向内打印矩阵

1. [相交链表](../algos/llist/llist22.go) 双指针获取两个链表的长度差

1. [合并 k 个升序链表](../algos/llist/llist34.go) 堆的插入（上浮）操作和删除（下沉）操作

1. [字符串相加](./problems/c026_addstrings.go) 模仿竖式加法并对较短的数进行补零

1. 最长上升子序列

1. [重排链表](./problems/c028_reorderllist.go) 交替合并两个链表

1. [接雨水](../algos/tp/tp7.go) 前后缀分解 or 双指针

1. [环形链表 II](../algos/llist/llist26.go) 哈希表 + 快慢指针

1. [删除链表的倒数第 n 个节点](../algos/llist/llist29.go) 没啥意义的题

1. 二叉树中的最大路径和

1. [合并区间](../algos/array/array14.go) 将区间按照左端点从小到大排序后所有可以合并的区间是连续的

1. 编辑距离

1. [二叉树的中序遍历](../algos/bitree/bitree36.go) 不断压入当前节点的左孩子，如果为空了就打印栈顶元素然后转向右边

1. 最长公共子序列

1. [二分查找](../algos/bs/bs64.go) 找到第一个大于等于 `target` 的值的下标并判断 `k<len(a) && a[k]==target` 是否成立

1. [二叉树的右视图](./problems/c038_bitreerightview.go) 层序遍历记录每层最右边节点的值

1. [用栈实现队列](../algos/stkandq/twostacksqueue/twostacksqueue.go) 在 `popstk` 为空时要将 `pushstk` 中的元素全部 “倒入” `popstk` 中

1. [复原 IP 地址](./problems/c040_restoreip.go)
   
   子集型回溯 => 每两个数字之间的 `'.'` 都有选或不选两种可能 => 选取所有元素都在 `[0,255]` 且长度为 4 的结果

   如果去掉筛选长度为 4 的结果这一步就是 [分割回文串](../algos/backtrack/bt61.go) 一题，子集型回溯的例题是 [子集](../algos/backtrack/bt56.go) 一题

1. [删除排序链表中的重复元素 II](./problems/c041_deldups2.go)

   统计和当前节点值相等的节点的个数
   
   如果大于 1 则将当前节点的前驱节点的 `Next` 域指向下一个与当前节点值不等的节点

1. 寻找两个正序数组的中位数

1. [下一个排列](../algos/skills/skills99.go) 将一个尽可能靠右的数和它右侧比它大的数中最小的那个交换

1. [排序链表](../algos/llist/llist33.go) 对链表使用自顶向下和自底向上归并排序算法进行排序

1. [x 的平方根](./problems/c045_sqrtx.go) 二分查找第一个满足自身的平方和大于等于 x 的元素

1. 字符串转换整数

1. [爬楼梯](../algos/dp/dp81.go) 斐波那契数列

1. 括号生成

1. 两数相加

1. 滑动窗口最大值

1. 比较版本号

1. 缺失的第一个正数

1. [链表中倒数第 k 个节点](./problems/c053_kthnode.go) 快慢指针找链表的倒数第 k 个节点

1. [最小覆盖子串](../algos/substr/substr12.go) 使用 `check` 函数的滑动窗口

1. [子集](../algos/backtrack/bt56.go) 子集型回溯的例题 => 每个元素都有选或不选两种可能

1. [零钱兑换](../algos/dp/dp85.go) 动态规划 - 详见 [README](../algos/README.md) 第 85 题

1. 从前序与中序遍历序列构造二叉树

1. [最小栈](../algos/stkandq/minstack/minstack.go) 使用辅助栈保存当前的最小值

1. 最长有效括号

1. 翻转字符串里的单词

1. 字符串相乘

1. 对称二叉树

1. 平衡二叉树

1. [二叉树的前序遍历](./problems/c064_preorder.go) 先访问根节点然后将其所有子节点从右向左逆序压栈

1. [二叉树的最大深度](../algos/bitree/bitree40.go) 左右子树最大深度的较大值加一

1. [求根到叶子节点数字之和](./problems/c066_sumnums.go) 遍历二叉树计算根到叶子节点路径表示的数字的值

1. [二叉树的直径](../algos/bitree/bitree40.go) 遍历每个节点并计算以每个节点作为根节点时的最长路径并更新最大值

1. [验证二叉搜索树](../algos/bitree/bitree43.go) 使用迭代式算法判断中序遍历结果是否为单调递增的

1. [旋转图像](../algos/matrix/mat20.go) 根据矩阵各种对称的转移式进行推导

1. [组合总和](../algos/backtrack/bt58.go)

1. [路径总和 II](./problems/c071_pathsum2.go) 
   
   在先序遍历时保存路径
   
   如果访问到了叶节点且目标和与叶节点的值相等则返回一条从根到叶节点的路径

1. [字符串解码](../algos/stkandq/stk63.go) 数字栈 + 字符串栈

1. 用 `Rand7()` 实现 `Rand10()`

1. 最小路径和

1. [在排序数组中查找元素的第一个和最后一个位置](../algos/bs/bs65.go) `bsearch(nums, target) & bsearch(nums, target+1) - 1`

1. 最大正方形

1. 搜索二维矩阵 II

1. [回文链表](../algos/llist/llist24.go) 快慢指针寻找中间节点 + 反转一半链表

1. [路径总和](./problems/c079_pathsum1.go)

   当前节点的值为 `val` 问是否存在从当前节点到叶子节点的路径和等于 `target` ？ =>

   是否存在从当前节点的左右孩子节点到叶子节点的路径和等于 `target-val` ？

1. 寻找峰值

1. 最长公共前缀

1. [多数元素](../algos/skills/skills97.go) Boyer-Moore 投票算法

1. 最长重复子数组

1. 翻转二叉树

1. 不同路径

1. 最长连续序列

1. [二叉树最大宽度](./problems/c087_widthofbt.go) 将二叉树视为完全二叉树并在层序遍历时保存每个节点的编号

1. 买卖股票的最佳时机 II

1. 乘积最大子数组

1. 岛屿的最大面积

1. 删除排序链表中的重复元素

1. 基本计算器 II

1. [打家劫舍](../algos/dp/dp83.go) 动态规划 - 详见 [README](../algos/README.md) 第 83 题

1. 最大数

1. [单词拆分](../algos/dp/dp86.go) 动态规划 - 详见 [README](../algos/README.md) 第 86 题

1. [两两交换链表中的节点](../algos/llist/llist31.go) 第 5 题 k 个一组反转链表中 k = 2 时的情况

1. [手撕堆排序](./problems/c097_heapsort.go) 原地构建大顶堆 + 进行 n-1 次堆顶元素的删除

1. 二叉树的序列化与反序列化

1. 移动零

1. 和为 k 的子数组

1. 长度最小的子数组

1. [寻找旋转排序数组中的最小值](../algos/bs/bs67.go)

   数组满足前半部分都大于 `nums[-1]` 而后半部分都小于等于 `nums[-1]` 且

   答案是第一个满足小于等于 `nums[-1]` 的元素

1. 验证 IP 地址

1. [LFU 缓存](./problems/c104_lfu.go)

   双哈希表
   
   `freq` —— 键为频率 + 值为所有使用频率等于键的索引项按使用时间先后排序的双向链表

   `keys` —— 键为索引项中的 key 字段 + 值为索引项在链表中的内存地址

1. 复制带随机指针的链表

1. 基本计算器

1. 每日温度

1. [用两个栈实现队列](../algos/stkandq/twostacksqueue/twostacksqueue.go) 在 `popstk` 为空时要将 `pushstk` 中的元素全部 “倒入” `popstk` 中

1. 只出现一次的数字

1. 全排列 II

1. [课程表](../algos/graph/graph53.go) 对 DAG 进行拓扑排序的 Kahn 算法

1. 对角线遍历

1. 移掉 k 位数字

1. 盛最多水的容器

1. 二叉搜索树与双向链表

1. 二叉树的完全性校验

1. 检测循环依赖

1. [手撕归并排序](./problems/c118_mergesort.go) 函数 `merge(a, lo, mi, hi)` 函数的作用是归并 `a[lo...mi]` 和 `a[mi+1...hi]` 两个有序子数组

1. 买卖股票的最佳时机 III

1. 排序奇升偶降链表

1. 单词搜索

1. 跳跃游戏

1. 二叉搜索树的第 k 大节点

1. `Pow(x, n)`

1. 数组中的逆序对

1. 旋转链表

1. 组合总和 II

1. 删除二叉搜索树中的节点

1. 删除排序数组中的重复项

1. 整数反转

1. 零钱兑换 II

1. [搜索二维矩阵](../algos/bs/bs64.go) 将二维矩阵的行列下标对应到一维数组的下标后进行二分查找

1. 螺旋矩阵 II

1. 最小的 k 个数

1. 二叉树的后序遍历

1. 青蛙跳台阶问题

1. 连续子数组的最大和

1. [荷兰国旗问题](../algos/skills/skills98.go) 以 1 为主元进行三向切分

1. 字典序的第 k 小数字

1. 树的子结构
