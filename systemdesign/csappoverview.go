package systemdesign

// Specific implementations of systems change
// over time, but the underlying concepts do not.

// 一、程序的结构与表示
//
// 1. 信息是位 + 上下文
// - 计算机系统中的所有信息都是用 [比特] (a bunch of bits) 来表示的，
//   区分不同数据对象的唯一方式是我们看待它们时的 [上下文] (context) 。
// - 像 hello.c 这样只由 ASCII 字符构成的文件称为 [文本文件] ，而其他
//   所有文件都称为 [二进制文件] 。
//
// 2. 程序被其他程序翻译成不同的格式
//  -text             -text             -text           -binary         -binary
// hello.c => cpp => hello.i => cc1 => hello.s => as => hello.o => ld => hello
//                                                               ^
//                                                              /
//                                                         printf.o
// - 预处理器根据 '#' 开头的 [指令] (directive) 修改原始的 C 程序。
// - 编译器将 C 程序编译为汇编语言程序。
// - 汇编器将汇编程序翻译为机器语言指令，并把它们打包成一种叫
//   做 [可重定位目标程序] (relocatable object program) 的格式。
// - 链接器将已经预编译好的目标文件 printf.o 合并到 hello.o 中得到
//   一个 [可执行目标文件] (executable object file) 并存储在磁盘中。
//
// 3. 了解编译系统大有益处
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
// - [主存] (main memory) 是一个临时存储设备，用于在处理器执行程序时
//   存放程序和要处理的数据。
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
// - 根据程序的 [局部性] (locality) 原理，现代计算机系统使用 [高速缓存存储器] (cache memory) 来
//   暂存处理器近期可能要用到的信息。L1、L2 和 L3 高速缓存使用 [静态随机存取存储器] (static random
//   access memory) 技术来实现。
//
// 6. 存储设备形成层次结构
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
