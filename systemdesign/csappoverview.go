package systemdesign

// Specific implementations of systems change over time,
// but the underlying concepts do not.

// 一、程序的结构与表示
//
// 1. 信息是位 + 上下文
// - 计算机系统中的所有信息都是用 [比特] (a bunch of bits) 来表示的，
//   区分不同数据对象的唯一方式是我们看待它们时的 [上下文] 。
// - 像 hello.c 这样只由 ASCII 字符构成的文件称为 [文本文件] ，其他
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

// 三、程序间的交互与通信
//
