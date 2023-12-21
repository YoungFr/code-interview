package systemdesign

// Specific implementations of systems change
// over time, but the underlying concepts do not.

// 一、程序的结构与表示
//
// 1. 信息是位 + 上下文
//
// - 计算机系统中的所有信息都是用 [比特] (a bunch of bits) 来表示的，
//   区分不同数据对象的唯一方式是我们看待它们时的 [上下文] (context) 。
// - 像 hello.c 这样只由 ASCII 字符构成的文件称为 [文本文件] (text file) ，而其他
//   所有文件都称为 [二进制文件] (binary file) 。
//
// 2. 程序被其他程序翻译成不同的格式
//
//  -text             -text             -text           -binary         -binary
// hello.c => cpp => hello.i => cc1 => hello.s => as => hello.o => ld => hello
//                                                               ^
//                                                              /
//                                                         printf.o
// - 预处理器根据 '#' 开头的 [指令] (directive) 修改原始的 C 程序。
// - 编译器将预处理后的 C 程序编译为汇编语言程序。
// - 汇编器将汇编程序翻译为机器语言指令，并把它们打包成一种叫
//   做 [可重定位目标程序] (relocatable object program) 的格式。
// - 链接器将已经预编译好的目标文件 printf.o 合并到 hello.o 中得到
//   一个 [可执行目标文件] (executable object file) 并存储在磁盘中。
//
// 3. 了解编译系统大有益处
//
// - 优化程序性能
// - 理解链接错误
// - 避免安全漏洞

// 二、程序的运行
//
// 4. 处理器读取并解释内存中的指令
//
// 4.1 系统的硬件组成 => ../resources/hardware_organization.png
// - [总线] (bus) 用于传输被称为 [字] (word) 的定长字节块，一个字包含的
//   字节数称为系统的 [字长] (word size) 。
// - I/O 设备通过 [控制器] (controller) 或 [适配器] (adapter) 与
//   I/O 总线相连，它们的区别主要在于封装方式 (packaging) 不同。
//   控制器是设备本身或主板上的芯片组而适配器是一块插在主板上的卡。
// - [主存] (main memory) 是一个临时存储设备，用于在处理器执行程序时存放程序和要处理的数据。
//   物理上来看主存是一组 [动态随机存取存储器] (dynamic random access memory) 芯片。
//   逻辑上来看主存是一个线性的字节数组。
// - [中央处理单元] (central processing unit) 简称为 [处理器] (processor) ，它
//   的核心是一个 [程序计数器] (program counter, PC) 。程序计数器在任何时候都包含
//   主存中某条指令的地址。
// - 从系统通电开始到系统断电，处理器会不断地从 PC 指向的内存处读取指令，解释指令并执行
//   相应的操作，然后更新 PC 的值。这些操作则是围绕主存、[寄存器文件] (register file) 和
//   [算术/逻辑单元] (arithmetic/logic unit, ALU) 来进行的。寄存器文件是由一些单字长的
//   寄存器组成的存储设备，而 ALU 则用来计算数据和地址的值。
// - 处理器的指令执行模型由 [指令集架构] (instruction set architecture, ISA) 决定，它
//   描述了每条机器指令执行的效果。表面上看起来处理器是指令集架构的简单实现，但实际上现代处理
//   器使用了非常复杂的机制来加速程序的执行。处理器的 [微体系结构] (microarchitecture) 用
//   于描述这种具体的实现。
//
// 4.2 运行 hello 程序
// - 读取输入 => ../resources/reading_input.png
// - 使用 DMA 方式加载可执行文件到主存 => ../resources/loading_executable.png
// - 执行程序中的机器指令 => ../resources/writing_output.png
//
// 5. 高速缓存至关重要
//
// - 根据程序的 [局部性] (locality) 原理，现代计算机系统使用 [高速缓存存储器] (cache memory) 来
//   暂存处理器近期可能要用到的信息。L1、L2 和 L3 高速缓存使用 [静态随机存取存储器] (static random
//   access memory) 技术来实现。
//
// 6. 存储设备形成层次结构
//
// - 存储器层次结构的主要思想是上一层的存储器作为低一层存储器的高速缓存。
//                                              /\
//           ↑                                 /  \
//           |                             L0 /Regs\
//           |                               /______\
//           |                              /        \
//  smaller  |                          L1 / L1(SRAM) \
//  faster   |                            /____________\
//  costlier |                           /              \
//           |                       L2 /    L2(SRAM)    \
//           |                         /__________________\
//                                    /                    \
//                                L3 /       L3(SRAM)       \
//                                  /________________________\
//           |                     /                          \
//           |                 L4 /      Main Memory(DRAM)     \
//  larger   |                   /______________________________\
//  slower   |                  /                                \
//  cheaper  |              L5 /  Local Secondary Storage(Disks)  \
//           |                /____________________________________\
//           |               /                                      \
//           |           L6 /        Remote Secondary Storage        \
//           ↓             /   (Distributed File System, Web Server)  \
//                        /____________________________________________\

// 三、程序间的交互与通信
//
// 7. 操作系统管理硬件
//
// - 操作系统是在应用程序和硬件之间插入的一层软件，所有应用程序对硬件的操作都必须通过操作系统。
//
//          +---------------------------------------+ ---+
//          |          Application Programs         |    |
//          +---------------------------------------+    | Software
//          |            Operating System           |    |
//          +---------------------------------------+ ---+
//          | Processor | Main Memory | I/O devices |    | Hardware
//          +---------------------------------------+ ---+
//
// - 操作系统通过几个基本的抽象概念来管理硬件：
//
//
//          +-------------- Processes --------------+ ---+
//          |                                       |    |
//          |           +------ Virtual Memory -----+    |
//          |           |                           |    | Operating System
//          |           |             +--- Files ---+    |
//          |           |             |             |    |
//          +---------------------------------------+ ---+
//          | Processor | Main Memory | I/O devices |    | Hardware
//          +---------------------------------------+ ---+
//
// 7.1 进程
// - 进程是操作系统对一个正在运行的程序的抽象，多个进程可以 [并发地] (concurrently) 运行。
//   并发运行的含义：一个进程的指令和另一个进程的指令是 [交错] (interleaved) 执行的。
// - 操作系统实现并发的机制称为 [上下文切换] (context switch) 。
// - 上下文：操作系统会跟踪进程运行所需的所有状态信息，包括 PC 、寄存器文件的值和主存的内容，这些
//   状态信息称为上下文。
// - 上下文切换是由 [内核] (kernel) 管理的。内核是操作系统代码中常驻主存的部分，它不是一个独立
//   的进程，而是系统用于管理进程所用的代码和数据结构的集合。
//
// 7.2 线程
// - 在现代系统中，进程实际上由多个称为 [线程] (thread) 的执行单元组成。每个线程都运行在进程的
//   上下文中，并共享同样的代码和全局数据。
//
// 7.3 虚拟内存
// - 通过 [虚拟内存] (virtual memory) 提供的抽象，每个进程看到的内存都是一致的，称为
//   [虚拟地址空间] (virtual address space) 。Linux 进程的虚拟地址空间如下：
//
//                  +--------------------------+        memory
//                  |   kernel virtual memory  |  ↑  invisible to
//                  +--------------------------+       user code
//                  |        user stack        |
//                  |  (created at run time)   |
//                  +--------------------------+
//                  |             ↓            |
//                  |                          |
//                  |             ↑            |
//                  +--------------------------+
//                  | memory-mapped region for |
//                  |     shared libraries     |  `printf` function
//                  +--------------------------+
//                  |                          |
//                  |             ↑            |
//                  +--------------------------+
//                  |       run-time heap      |
//                  |   (created by `malloc`)  |
//                  +--------------------------+ --+
//                  |      read/write data     |   |   loaded from the
//                  +--------------------------+   |       `hello`
//                  | read-only code and data  |   |   executable file
// program start -> +--------------------------+ --+
//                  |                          |
//             0 -> +--------------------------+
//
// - 程序代码和数据 (program code and data) ：代码和数据区在进程一开始运行时就被指定了大小。
//   对于所有进程，代码都是从同一固定地址开始的。其中的内容是直接按照可执行目标文件的内容初始化的。
// - 堆 (heap) ：当调用 `malloc` 和 `free` 这样的函数时，堆可以在运行时动态地扩展或收缩。
// - 共享库 (shared libraries) ：PASS
// - 栈 (stack) ：编译器使用栈来实现函数调用，栈在执行时也可以动态地扩展或收缩。
// - 内核虚拟内存 (kernel virtual memory) ：PASS
//
// 7.4 文件
// +---------------------------------------------------------------------------+
// |      "A file is a sequence of bytes, nothing more and nothing less."      |
// +---------------------------------------------------------------------------+
//
// 8. 系统之间利用网络通信
//
// - 现代系统通过网络和其他系统相连。从一个单独的系统来看，网络可以视为一个 I/O 设备 => ../resources/net.png
//
// 9. 重要主题
//
// 9.1 Amdahl 定律
// - 假设一个系统原来的运行时间是 TO
//   系统的某个部分运行时间是总时间的 a 倍 (0 < a <= 1) 即 a * TO
//   现在我们将这部分的性能提高了 k 倍从而其现在的运行时间是 a * TO / k
//   那么现在系统的总运行时间是 TN = (1 - a) * TO + (a * TO / k)
//
//                                 1
//                    TO    ---------------
//               S = ---- =             a
//                    TN     (1 - a) + ---
//                                      k
//
// - 结论
//   当我们对系统的某个部分加速时，其对系统整体性能的影响取决于该部分的重要性（a）和加速程度（k）。
//   所以，如果我们想显著地提高系统的速度，就必须提升系统中相当大部分的速度。
//
// 9.2 并发与并行
// CONCURRENCY: the general concept of a system with multiple, simultaneous activities
// PARALLELISM: the use of concurrency to make a system run faster
//
// - 线程级并发 => 多核处理器与超线程
// - 指令级并行 => 流水线与超标量操作
// - SIMD => 向量指令
//
// 9.3 虚拟机抽象
//
//  +-------------------- Virtual Machine ---------------------+
//  |                                                          |
//  |                  +-------------- Processes --------------+
//  |                  |                                       |
//  |                  +--- ISA ---+------ Virtual Memory -----+
//  |                  |           |                           |
//  |                  |           |             +--- Files ---+
//  |                  |           |             |             |
//  +------------------+-----------+-------------+-------------+
//  | Operating System | Processor | Main Memory | I/O devices |
//  +------------------+-----------+-------------+-------------+
//
