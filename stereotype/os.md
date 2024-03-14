# 网络系统

### 零拷贝

- 传统 I/O ：CPU 指令 -> 磁盘控制器 -> 中断 -> CPU 拷贝数据到内存（期间不能执行其他操作）

- DMA 方式：DMA 控制器负责拷贝数据并发中断信号给 CPU

- 利用 `read/write` 的文件传输方式：4 次上下文切换、4 次数据拷贝

  ```
     read         read        write          write
  文件 -> 内核缓冲区 -> 用户缓冲区 -> socket 缓冲区 -> 网卡
       1            2           3             4
  ```

- 零拷贝

  - `mmap + write` ：利用 `mmap` 直接把内核缓冲区的数据映射到用户空间，减少了第 1 次拷贝

    4 次上下文切换、3 次数据拷贝

  - `sendfile` ：直接把内核缓冲区的数据拷贝到 socket 缓冲区，减少了第 2 次数据拷贝

    2 次上下文切换、3 次数据拷贝

  - `sendfile` + SG-DMA 技术：利用 `ethtool -k eth0 | grep scatter-gather` 命令查看网卡是否

    支持 scatter-gather 特性，将描述符和数据长度传到 socket 缓冲区后网卡的 SG-DMA 控制器直接将

    内核缓存中的数据拷贝到网卡的缓冲区，减少了第 2、3 次数据拷贝

    2 次上下文切换、2 次数据拷贝

- 大文件的传输：零拷贝技术是缓冲 I/O 使用了内核缓冲区（PageCache）。如果传输的是大文件，PageCache 会被

  大文件长期占据，所以针对大文件要使用 <font color=red>**异步 I/O + 直接 I/O**</font> 来绕过内核缓冲区。

### I/O 多路复用

- 基本模型

  ```
  socket()      socket()
                bind()
                listen()
  connect() --> accept() --> 服务器内核维护半连接队列和全连接队列
  write()   --> read()
  read()    --> write()
  ```

- 多进程模型：在 `accept` 函数返回后通过 `fork` 创建子进程，子进程会复制父进程的文件描述符，可以直接

  和客户端通信。

- 多线程模型：在 `accept` 函数返回后通过 `pthread_create` 创建线程，将已连接 socket 的文件描述符传递给

  线程函数在线程里和客户端进行通信。优化：全局的已连接 socket 队列 + 线程池。

- I/O 多路复用 —— <font color=red>**只使用一个进程来维护多个 socket **</font>

