# 基础

- 指针和引用

  **核心区别**：对象的地址 VS 变量的别名

  **定义**：* `VS` &

  **使用**：用 * 解引用 `VS` 直接使用

  **空值**：有 `VS` 无

  **可变性**：可变 `VS` 不能改变绑定

  **用途**：内存分配、数组 `VS` 参数传递、操作符重载、创建别名

- 数据类型

  16 <= short <= int <= 32 <= long <= 64 <= long long

- `const`

  **常量指针**与**指针常量**

  ```c++
  int tmp1 = 1;
  int tmp2 = 1;
  
  const int * a = &tmp1;
  a = &tmp2;
  (*a)++;    // error: increment of read-only location '* a'
  
  int * const b = &tmp1;
  b = &tmp2; // error: assignment of read-only variable 'b'
  (*b)++;
  ```

  **常量变量**：值不能修改

  **常量对象**：对象的**成员变量**不能修改

  **函数参数**：不会修改参数

  **成员函数**：不会修改对象的未被 `mutable` 修饰的非静态变量

- `static`

  **静态变量具有全局变量的生命周期，但只能作用于自己的作用域**

  **成员变量**：所有类的对象共有，必须**在类外单独定义**以便为其在全局数据区分配内存 / **成员函数**

  **全局变量**：在全局数据区分配内存且自动初始化为零，在声明它的文件外不可见 / **静态函数**：隐藏

  **局部变量**：⽣命周期延⻓到整个程序的执⾏过程，但只在声明它的函数内可⻅

- `new` 和 `malloc`

  |      对比项      |  `new`/`delete`   | `malloc`/`free`  |
  | :--------------: | :---------------: | :--------------: |
  |       本质       |      运算符       |      库函数      |
  |   内存分配大小   |     自动计算      |     手工计算     |
  |     类型安全     |        是         |        否        |
  | 分配失败的返回值 | `bad_alloc` 异常  |      `NULL`      |
  |       过程       | 调用构造/析构函数 | 只分配和释放内存 |
  |      返回值      |  具体类型的指针   |     `void *`     |

- `constexpr`

  编译期常量、用于常量表达式的函数

- `volatile`

  ⽤该关键字声明的变量表示该变量随时可能发⽣变化，与该变量有关的运算，不要进⾏编译优化

  会从内存中重新装载内容，⽽不是直接从寄存器拷⻉内容

- `extern`

  声明外部变量 —— 在函数或者⽂件外部定义的全局变量

- `std::atomic`

  `i++` 和 `a = b` 在 C++ 中都不是线程安全的 => 原因：内存 -> 寄存器 -> 内存

- 函数指针的作用：回调函数、多态、函数参数

- 类型转换

  `static_cast` —— 无运行时类型检查

  `dynamic_cast` —— 下行转换时进行类型检查，只能用于类指针、类引用和 `void` 指针
  
  `reinterpret_cast` —— 任意转换
  
  `const_cast` —— 去掉类型的 `const` 或 `volatile` 属性

# 内存管理

- C++ 程序运行时的内存分区（从低到高）：代码区 -> 常量区 -> 全局区 -> 堆区 -> 栈区

- 内存泄漏：堆内存泄漏、系统资源泄漏、未将基类的析构函数定义为虚函数

- 智能指针

  `unique_ptr` —— 提供对动态分配的单⼀对象所有权的独占管理

  `shared_ptr` —— 允许多个智能指针共享同⼀块内存资源，内部使⽤引⽤计数来跟踪对象被共享的次数

  `weak_ptr` —— 解决共享智能指针导致的循环引用问题，可以从 `shared_ptr` 创建但不增加引用计数
  
- 内存对齐：基本数据类型的变量的起始地址必须是它们大小的倍数

# 面向对象



