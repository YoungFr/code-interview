# 面经

Q: Linux 进程的<font color=red>内存布局</font>

<details><summary>A</summary>代码段（.text）、数据段（已初始化（.data）、未初始化（.bss））、堆、栈</details>

Q: 内存<font color=red>溢出</font>与内存<font color=red>泄漏</font>的解决办法

<details><summary>溢出</summary>检查算法和数据结构、释放资源、增加大小、内存池</details>

<details><summary>泄漏</summary>内存分析、智能指针、单元测试、日志监控、性能优化</details>

Q: <font color=red>extern "C"</font> 语法？

<details><summary>A</summary>在 C++ 中按 C 语言方式进行名称修饰和链接。C++ 支持重载 -> 需要保存函数名及参数信息；C 只保存函数名</details>

Q: 内存管理的<font color=red>伙伴系统</font>

<details><summary>A</summary>用于解决外部碎片问题。将内存块按照 2 的幂的大小组织成链表</details>

Q: <font color=red>中断</font>过程

<details><summary>A</summary>中断发生 -> 保存上下文 -> 中断向量表 -> 中断服务程序 -> 软件处理 -> 恢复上下文执行</details>

Q: <font color=red>死循环/阻塞协程</font>如何<font color=red>调度</font>？

<details><summary>A</summary>使用 sysmon 线程对进行系统调用的（阻塞的）和长时间运行的协程进行抢占式调度</details>

Q: 什么是<font color=red>协程泄漏</font>？

<details><summary>A</summary>由于不正确的管理或是过度使用导致协程资源无法被有效地释放从而占用大量内存空间的现象</details>

```go
// 协程泄漏的例子
func main() {
	errChan := make(chan error)
	for i := 0; i < 5; i++ {
		go func(i int) {
			// do something...
			// error occurs
			errChan <- fmt.Errorf("err %d", i)
		}(i)
	}
	fmt.Println(runtime.NumGoroutine()) // 6
	if err := <-errChan; err != nil {
		fmt.Printf("recv err: %v\n", err)
	}
	fmt.Println(runtime.NumGoroutine()) // 5
}
```

Q: 什么是 <font color=red>WebSocket</font> 协议？

<details><summary>A</summary>服务器主动发送消息？HTTP 短轮询 -> HTTP 长轮询 -> 先使用 HTTP 进行协议升级 -> 支持全双工的 WebSocket 协议</details>

Q: 常见的<font color=red>限流模式</font>？

**单机限流**

流量计数器/固定窗口计数器：控制每个固定时间窗口内的请求次数；存在**窗口边缘流量猛增**问题。

滑动时间窗算法：平滑的基于某个时间片的流量统计。请求到来 -> 移除窗口内的过期请求 -> 统计当前窗口中的请求。

这两种模式只能决定是否限流，对于**超过阈值的流量只能丢弃而不是阻塞等待**。

漏桶算法：使用一个 FIFO 队列充当缓冲区，以固定速率处理请求。关键在于**桶的大小**和**处理速率**。

令牌桶算法：规定 1 秒内请求不超过 X 次，则每隔 1/X 秒放入一个令牌。请求必须获得令牌才能处理，桶有最大容量。

漏桶和令牌桶是思路相反的两种方法：**漏桶满时是流量过多而令牌桶满时是流量过少**。

**分布式限流**

集中统计：将所有服务的统计结果存入集中式缓存（Redis）并使用分布式锁进行并发控制。

数值化令牌桶算法：请求在网关处领取一定数值的令牌，每访问一个服务使用一定量的数值，小于 0 时重新申请。



# Redis

Q: 如何使用 Redis 实现<font color=red>分布式锁</font>？

分布式锁用于控制在分布式环境下某个资源在同一时刻只能被一个服务使用。

用带 `NX` 选项的 `SET` 命令实现分布式锁。如果 `key` 不存在则插入成功表示加锁成功，否则失败：

```
SET lock_key unique_value NX PX 10000
```

**注意点**

- 使用 `PX` 选项设置过期时间，避免客户端拿到锁后一直无法释放导致异常。

- 设置的值用于区分不同的客户端，所以**每个客户端必须使用唯一值（分布式唯一 ID 生成）**。

- 解锁的过程即使用 `DEL` 命令将键删除，但**一定要保证执行该操作的客户端就是加锁客户端**。

- 要使用 Lua 脚本保证这两个命令以原子方式执行。

  ```
  为什么 Lua 脚本具有原子性？
  所有 Lua 脚本共用一个 Lua 环境 -> 一个 Lua 脚本执行时其他 Lua 脚本无法执行
  ```

基于 Redis 实现的分布式锁的缺点在于**超时时间不好设置**和**主从异步复制导致的不可靠性**。

Redis 官方针对集群环境下的分布式锁设计了 RedLock 方案：

**使用 `2*N+1` 个Redis 主节点，客户端依次向这些节点请求加锁，如果在大于等于 `N+1` 个节点上成功即加锁成功**。

Q: <font color=red>分布式唯一 ID</font> 生成算法？

<details><summary>A</summary>要满足的三个条件：全局唯一、单调递增、信息安全；常见方案：数据库自增 ID、UUID、（类）雪花算法</details>

```
┌─┬───────────┬───────────────┬────────────┬─────────────────┐
│0│ timestamp │ datacenter ID │ machine ID │ sequence number │
└─┴───────────┴───────────────┴────────────┴─────────────────┘
 1     41             5              5             12
```

Q: <font color=red>跳表</font>介绍及读写流程？



## TCP

TCP 是**面向连接的、可靠的、基于字节流**的传输层通信协议。

Q: 比较重要的<font color=red>TCP 首部</font>字段？

<details><summary>A</summary>源端口、目的端口、序列号(解决乱序)、确认应答号(解决丢包)、控制位 ACK/RST/SYN/FIN</details>

Q: TCP 有哪些保证<font color=red>可靠传输</font>的机制？

<details><summary>A</summary>连接管理、序列号、确认应答、超时重传、流量控制、拥塞控制</details>

Q: 简述<font color=red>三次握手</font>和<font color=red>四次挥手</font>的流程？

三次握手

```
    TCP A                                                 TCP B

1.  CLOSED                                                LISTEN

2.  SYN-SENT    --> <SEQ=100><CTL=SYN>                --> SYN-RECEIVED 丢失->重传相同的 SYN 报文

3.  ESTABLISHED <-- <SEQ=300><ACK=101><CTL=SYN,ACK>   <-- SYN-RECEIVED 丢失->客户端服务端都重传

4.  ESTABLISHED --> <SEQ=101><ACK=301><CTL=ACK>       --> ESTABLISHED  丢失->重传相同的 SYN+ACK 报文

5.  ESTABLISHED --> <SEQ=101><ACK=301><CTL=ACK><DATA> --> ESTABLISHED
```

四次挥手

```
    TCP A                                                TCP B

1.  ESTABLISHED                                          ESTABLISHED

    (Close)
2.  FIN-WAIT-1  --> <SEQ=100><ACK=300><CTL=FIN,ACK>  --> CLOSE-WAIT 丢失->重传相同的 FIN 报文

3.  FIN-WAIT-2  <-- <SEQ=300><ACK=101><CTL=ACK>      <-- CLOSE-WAIT 丢失->重传相同的 FIN 报文

                                                         (Close)
4.  TIME-WAIT   <-- <SEQ=300><ACK=101><CTL=FIN,ACK>  <-- LAST-ACK   丢失->重传相同的 FIN+ACK 报文

5.  TIME-WAIT   --> <SEQ=101><ACK=301><CTL=ACK>      --> CLOSED     丢失->重传相同的 FIN+ACK 报文

6.  (2 MSL)
    CLOSED
```

Q: 不能<font color=red>两次</font>和<font color=red>四次</font>握手的原因？

<details><summary>A</summary>两次：不能防止历史连接的建立和有效地同步序列号；四次：避免浪费</details>



Q: 客户端的 <font color=red>TIME_WAIT</font> 状态以及 <font color=red>2MSL</font> 的原因？

主动关闭连接的一方才有 TIME_WAIT 状态。

MSL 指报文最大生存时间，2MSL 的时间可以允许报文丢失一次。如果**在 2MSL 内收到了重传报文会重置定时器**。

**可以保证被动关闭连接的一方能被正确关闭**：2MSL 的时间可以保证最后一次的 ACK 报文被收到，如果最后一次的 ACK 报文丢失，经过 MSL 后会重传 FIN+ACK 报文，主动关闭方可以在第 2 个 MSL 内收到重传报文从而再次发送 ACK 报文。

<font color=red>TIME_WAIT 过多会占用系统资源及端口资源：客户端 -> 无法再对相同 IP+Port 的服务端发起连接；服务端：占用资源</font>。



Q: TCP 怎么处理<font color=red>丢包</font>？

<details><summary>A</summary>重传 -> 超时重传、快速重传、SACK</details>



Q: 简述 TCP <font color=red>拥塞控制</font>？

<details><summary>A</summary>TODO</details>



Q: 介绍下<font color=red>MTU</font>和<font color=red>MSS</font>的含义

MTU：一个网络包的最大长度。超过 MTU 的包要分片，**当一个 IP 分片丢失后，整个网络包的所有分片都要重传**。

MSS：**除去 IP 和 TCP 头部**后一个网络包能容纳的 **TCP 数据**的最大长度。

<font color=red>为了防止在一个分片丢失时重传整个包，当 TCP 发现数据超过 MSS 时会先分片而避免在 IP 层分片</font>。



Q: 在 Linux 上查看并处理 <font color=red>CPU 占用多的进程、进程的内存占用情况</font>？

```
查看某个进程的信息：
用 ps -ef | grep xxx 得到进程号
用 top -p <PID> 查看

查看CPU、内存占用多的进程
top => 按 p 按照 CPU 排序；按 m 按照内存排序
```

Q: 介绍<font color=red>布隆过滤器</font>的应用和原理？

<details><summary>A</summary>TODO</details>

Q: <font color=red>线程池</font>要考虑的参数、Java 线程池？

<details><summary>A</summary>TODO</details>

Q: 什么是<font color=red>可重入锁</font>及实现？

<details><summary>A</summary>指任意线程在获取锁后可以再次获取锁而不被阻塞，对不同的线程来说就是普通的互斥锁。实现：持有线程+计数器</details>

Q: <font color=red>服务器卡顿</font>分析？

<details><summary>A</summary>TODO</details>

Q: <font color=red>CAS</font>和<font color=red>ABA</font>简介？

<details><summary>A</summary>TODO</details>

Q: <font color=red>负载均衡</font>的策略与实现？

<details><summary>A</summary>TODO</details>

