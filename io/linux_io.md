# Linux I/O

# 1. File I/O

文件是 Unix 哲学的中心思想，本节聚焦于磁盘文件的 I/O 。

所有和 I/O 相关的系统调用都使用 **文件描述符 (file descriptor)** 来指代打开的文件，这些文件包括管道、FIFO、套接字、终端、设备和普通文件。<font color=red>**每个进程都有它自己的打开的文件描述符的集合**</font>。在 shell 的日常操作中，标准输入、标准输出和标准错误这三个文件描述符始终是打开的。所以通过 shell 启动的程序也会继承这三个文件描述符。

文件操作的主要系统调用是 `fd = open(pathname, flags, mode)` 、`numread = read(fd, buffer, count)` 、`numwritten = write(fd, buffer, count)` 和 `status = close(fd)` 。文件 `copy.c` 使用这 4 个系统调用实现了一个简化的 [`cp(1)`](https://man7.org/linux/man-pages/man1/cp.1.html) 命令。

```
手册页中的命令、系统调用、库函数等名字后边的小括号中的数字表示它归属于哪一个章节：
1 - 用户命令(user commands)
2 - 系统调用(system calls)
3 - 库函数(library functions)
4 - 设备(special files)
5 - 文件格式和文件系统(file formats and filesystems)
6 - 游戏(games)
7 - 杂项(overview and miscellany section)
8 - 系统管理员工具(administration and privileged commands)
```

```c
// copy.c
#include <sys/stat.h>
#include <fcntl.h>
#include "tlpi_hdr.h"

#ifndef BUF_SIZE
#define BUF_SIZE 1024
#endif

int main(int argc, char *argv[])
{
    if (argc != 3 || strcmp(argv[1], "--help") == 0)
        usageErr("%s old-file new-file\n", argv[0]);
    
    int inputFd = open(argv[1], O_RDONLY);
    if (inputFd == -1)
        errExit("opening file %s", argv[1]);

    int    openFlags = O_CREAT | O_WRONLY | O_TRUNC;
    mode_t filePerms = S_IRUSR | S_IWUSR  | 
                       S_IRGRP | S_IWGRP  |
                       S_IROTH | S_IWOTH; /* rw-rw-rw- */
    int outputFd = open(argv[2], openFlags, filePerms);
    if (outputFd == -1)
        errExit("opening file %s", argv[2]);

    /* Transfer data until we encounter end of input or an error */

    ssize_t numRead;
    char buf[BUF_SIZE];
    while ((numRead = read(inputFd, buf, BUF_SIZE)) > 0)
        if (write(outputFd, buf, numRead) != numRead)
            fatal("write() returned error or partial write occurred");
    if (numRead == -1)
        errExit("read");

    if (close(inputFd) == -1)
        errExit("close input");
    if (close(outputFd) == -1)
        errExit("close output");

    exit(EXIT_SUCCESS);
}
```

