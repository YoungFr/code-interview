# Linux System Programming

本文件夹中的文档尝试记录总结 Linux 系统编程中方方面面的知识，希望以此来加深对操作系统和计算机网络相关知识的理解，提高自己的系统编程能力。内容主要来自 [The Linux Programming Interface](https://man7.org/tlpi/index.html) 、 Advanced Programming in UNIX Environment  、 [Linux 在线帮助手册](https://man7.org/linux/man-pages/) 和其他各种在线资源。

具体的计划如下，完成时间未知。

### I/O

[I/O](./io.md)

记录和系统 I/O 相关的内容。包括文件 I/O 及其细节、文件 I/O 缓冲、各种高级 I/O 模型、终端和伪终端。

### File System

[File System](./fs.md)

记录和文件系统相关的内容。包括文件系统、文件属性、扩展属性、目录和链接。

### Process

[Process](./proc.md)

记录和进程相关的内容。包括进程的内存布局、进程凭证、进程创建、进程终止、子进程、程序执行、进程优先级与调度、守护进程和共享库。

### Thread

[Thread](./thread.md)

记录和线程相关的内容。包括线程基础、线程同步、线程安全和线程取消。

### IPC

[IPC](./ipc.md)

记录和进程间通信相关的内容。包括管道、 FIFO 、信号、消息队列、信号量、内存映射和共享内存。

### Socket

[Socket](./socket.md)

套接字（socket）也是一种 IPC 方法，但它允许通过网络连接的不同主机上的进程进行通信。因为它与网络结构和协议相关，所以需要单独一章进行介绍，具体内容包括 UNIX Domain Socket 、 Internet Domain Socket 、服务器设计与其他高级主题。

### Miscellaneous

[Miscellaneous](./misc.md)

其他和 Linux 相关的主题，比如各种工具链的使用、用户命令和 Shell 脚本。

