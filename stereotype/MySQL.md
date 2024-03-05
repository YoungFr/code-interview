# 字段类型

- **数值、字符串、日期时间**概览

  ![mysql_data_types](./assets/mysql_data_types.png)

- 整数类型的 `UNSIGNED` 属性：不允许负数的场景，例如 ID 属性

- 定长字符串 `CHAR` 类型：存储时在右边添加空格、检索时去掉

  变长字符串 `VARCHAR` 类型：存储时用 1 或 2 个额外字节记录字符串的长度

- 定点数 `DECIMAL` 类型：精确的小数值

  浮点数 `FLOAT/DOUBLE` 类型：近似的小数值

- 不推荐使用的长文本数据 `TEXT` 和二进制大对象 `BLOB` 类型：

  不能有默认值、检索效率低、不能使用内存临时表、不能直接创建索引、消耗带宽

- 无时区信息的 `DATATIME` 和有时区信息的 `TIMESTAMP` 类型

- `NULL` 和 `''` 的区别：<font color=red>为什么 MySQL 不建议使用 `NULL`作为列默认值？</font>

  `NULL` 表示不确定的值，两个 `NULL` 不一定相等

  `''` 长度为 0 不占用空间而 `NULL` 占用空间

  `NULL` 影响聚合函数的结果

  `NULL` 要使用 `IS NULL` 或 `IS NOT NULL` 判断而 `''` 可以使用比较运算符

# 存储引擎

- 命令 `SHOW ENGINES` 查看所有存储引擎，<font color=red>只有 InnoDB 支持事务</font>

- MySQL 的存储引擎使用<font color=red>插件式架构</font>

  可以为不同的表设置不同的存储引擎 => <font color=red>存储引擎是基于表的而不是基于数据库的</font>

- MySQL 默认的存储引擎：MyISAM <- 5.5 -> InnoDB

- 二者对比：

  No <- 是否支持**行级锁** -> Yes

  No <- 是否支持**事务** -> Yes

  No <- 是否支持**外键** -> Yes

  No <- 是否支持异常崩溃后的**安全恢复** -> Yes

  No <- 是否支持**MVCC** -> Yes

  索引和数据文件分离 <- **索引实现** -> 数据文件本身就是索引文件

- 更为详细的对比如下

  ![comparison_of_common_mysql_storage_engines](./assets/comparison_of_common_mysql_storage_engines.png)

# 索引

- **索引就相当于数据的目录**

# 事务

- 逻辑上的一组操作要么都执行，要么都不执行

  没有特指**分布式事务**时就是指**数据库事务**

  ```mysql
  # 开启一个事务
  START TRANSACTION;
  # 多条 SQL 语句
  SQL1,SQL2...
  ## 提交事务
  COMMIT;
  ```

- 事务具有<font color=red>**ACID**</font>属性

  

  ```
                      手段                   =>                目的
  
  A(最小执行单位) + I(并发事务之间独立) + D(提交后是持久改变) => C(执行前后数据一致)
  
    undolog  /  锁（悲观）+ MVCC（乐观）  /  redolog     <= 对应的 MySQL 实现方法
  ```

  

  DDIA：前三者是**数据库**的属性而一致性是**应用程序**的属性

- 并发事务带来的问题

  **脏读**（Dirty Read） -> 一个事务读取了另一个事务**未提交**的数据

  **不可重复读**（Unrepeatable Read） -> 一个**事务内对同一数据的两次读取**结果不同 => **数据不一致**

  **幻读**（Phantom Read） -> 事务内**读取一条数据时不存在但插入时又存在** => **数量不一致**

- SQL 标准定义的四个事务隔离级别

  **读未提交**（READ-UNCOMMITED）:sob:  :sob:  :sob:

  **读已提交**（READ-COMMITED）:smile:  :sob:  :sob: <font color=red>MVCC</font>

  **可重复读**（REPEATABLE-READ） —— InnoDB 引擎默认支持的隔离级别 :smile:  :smile:  :sob: <font color=red>MVCC</font>

  **可串行化**（SERIALIZABLE）:smile:  :smile:  :smile: <font color=red>锁</font>

# 锁



# 日志

