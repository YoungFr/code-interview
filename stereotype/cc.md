- 指针和引用

  关键区别：对象的地址 VS 变量的别名

  <font color=red>定义</font>（* VS & ）、<font color=red>使用</font>（用 * 解引用 VS 直接使用）、<font color=red>空值</font>（有 VS 无）、<font color=red>可变性</font>（可变 VS 不能改变绑定）、<font color=red>用途</font>（内存分配, 数组 VS 参数传递, 操作符重载, 创建别名）

- 数据类型

  16 <= short <= int <= 32 <= long <= 64 <= long long

- `const`

  常量指针与指针常量

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

  成员函数：

- `static`

  