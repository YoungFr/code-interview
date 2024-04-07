[toc]

# 环境准备

在 Windows 上安装 VMWare 注意事项：安装完成后使用 `ncpa.cpl` 命令确保 VMnet1 和 VMnet8 虚拟网卡已正常启用。

# 基本命令

### 列出文件

`ls`

每个用户在 Linux 中的专属目录称为 **HOME 目录**，默认路径为 `/home/username` 。

Linux 的命令行在执行时需要一个**工作目录**，打开终端时会默认将工作目录设置为 HOME 目录。

```
-a 列出包含隐藏文件的全部文件(以 . 开头的文件称为隐藏文件)
-l 以竖向列表的形式展示内容并显示更多信息
-h 和 -l 搭配使用以更人性化的形式显示文件的大小单位
```

### 目录切换

`cd` —— 无任何选项。带参数执行时切换到指定的目录，不带参数执行时切换到 HOME 目录。

`pwd` —— 无任何选项和参数。输出当前所在的工作目录。

### 特殊路径

```
.  表示当前目录
.. 表示上一级目录
~  表示 HOME 目录
```

### 创建目录

`mkdir` —— 创建新的目录，参数是要创建的目录的路径。

```
-p 自动创建不存在的父目录以创建连续的多层级的目录
```

### 文件操作

`touch` —— 指定要创建的文件路径来创建文件

`cat` —— 查看文件内容

`more` —— 以翻页的形式 (用空格翻页) 查看文件内容

`cp` —— 复制文件/文件夹

```
-r 递归地复制 -> 在复制文件夹时使用
```

`mv` —— 移动或重命名文件/文件夹，如果目标路径不存在则为重命名

`rm` —— 删除 (一个或多个) 文件/文件夹

```
-r 递归地删除 -> 在删除文件夹时使用
-f 强制删除 -> 由 root 用户使用
```

### 文件查找

`which` —— 命令的本质是一个二进制可执行程序，使用 `which command-name` 可以查看命令的程序文件放在哪里

`find` —— 查找

- 按文件名字查找

  ```
  find starting-point -name "file-name-pattern"
   
  Examples:
  find / -name "test"
  find . -name "*test"
  ```

- 按文件大小查找

  ```
  find starting-point -size +|-n[kMG]
  
  Examples:
  find / -size -4k
  find . -size +1M
  ```

`grep`

```
grep [-n] "pattern" file-path
-n 表示是否打印匹配行的行号

Examples:
grep "include" test.txt
grep -n "a" ../test.txt
```

`wc`

```
wc [-c -m -l -w] file-path

不加任何选项的输出:
a b c file-name -> 行数 单词数 字节数 文件名

-c 字节数
-m 字符数
-l 行数
-w 单词数
```

### 管道操作

管道符 `|` 的作用是将左边命令的输出作为右边命令的输入。

```shell
# 查找文件中包含 apple 的行
cat a.txt | grep -n "apple"
# 查找当前文件夹中所有文件的数量
ls -al | wc -l
```

### 其他操作

`echo` ——  命令 `echo "contents"` 用于输出指定内容：

```shell
echo "Hello World!"
```

反引号的使用 —— 被反引号包围的内容会作为命令执行：

```shell
echo `pwd`
```

重定向符 `>` 用于将左侧命令的结果**覆盖**写入到右侧文件中而 `>>` 用于**追加**写入：

```shell
ls > test.txt
echo "a new line" >> test.txt
```

`tail`

```shell
tail [-f -num] file-path
用于查看文件尾部内容及跟踪文件的最新更改
-f   持续跟踪
-num 查看尾部多少行(默认十行)
```

### 文本编辑

`vi/vim`

三种工作模式：

- 命令模式：所有按键都理解为命令，不能自由进行文本编辑

  |   快捷键   |                作用                |
  | :--------: | :--------------------------------: |
  |    `i`     |   在**当前光标位置**进入输入模式   |
  |    `a`     | 在**当前光标位置之后**进入输入模式 |
  |    `I`     |   在当前行的**开头**进入输入模式   |
  |    `A`     |   在当前行的**结尾**进入输入模式   |
  |    `0`     |       移动光标到当前行的开头       |
  |    `$`     |       移动光标到当前行的结尾       |
  |    `/`     |            进入搜索模式            |
  |    `n`     |            向下继续搜索            |
  |    `N`     |            向上继续搜索            |
  |  `(n)dd`   | 删除当前光标（向下 n 行）行的内容  |
  |  `(n)yy`   | 复制当前光标（向下 n 行）行的内容  |
  |    `p`     |           粘贴复制的内容           |
  |    `u`     |              撤销修改              |
  | `CTRL + R` |            反向撤销修改            |

- 输入模式：自由编辑

  只需要记住按 `ESC` 键从输入模式退出到命令模式。

- 底线模式：用于文件的保存和退出

  在命令模式按 `:` 进入底线模式。

  |   命令    |    作用    |
  | :-------: | :--------: |
  |   `:wq`   | 保存并退出 |
  |   `:q`    |   仅退出   |
  |   `:q!`   |  强制退出  |
  |   `:w`    |   仅保存   |
  | `:set nu` |  显示行号  |

# 用户和组

### 用户切换

超级管理员 `root` 拥有最大的系统操作权限。

普通用户一般在其 HOME 目录内不受限，但是在出了 HOME 目录的大多数地方只有读和执行权限，无修改权限。

`su/exit`

```shell
su [-] [username]

-        表示是否在切换用户后加载环境变量(建议使用)
username 要切换的用户(省略表示切换到 root 用户)

exit / CTRL + D
用于退回上一个用户
```

`sudo`

为普通用户授权使其临时以 root 用户执行命令。普通用户使用 `sudo` 需要配置：

```
切换到 root 用户后执行 visudo 命令
在文件的最后添加:
username ALL=(ALL)	NOPASSWD: ALL
```

### 用户管理

Linux 的权限管控有针对**用户**的权限控制和针对**用户组**的权限控制两个级别：

- 用户组管理

  ```shell
  # 所有命令需要 root 用户执行
  
  groupadd user-group-name # 创建
  groupdel user-group-name # 删除

- 用户管理

  ```shell
  # 所有命令需要 root 用户执行
  
  # -g 用于指定用户组/不指定时会创建同名的组/指定时用户组必须已存在
  # -d 指定用户的 HOME 路径
  useradd [-g -d] user-name # 创建用户
  # -r 是否删除用户的 HOME 目录
  userdel [-r] user-name # 删除用户
  id [username] # 查看用户所属的组
  usermod -aG user-group-name user-name # 修改用户所属的组
  ```

`getent`

```shell
# 查看系统全部用户信息
# login name:encrypted password:UID:GID:comment:home directory:login shell
getent passwd

# 查看系统全部组信息
# group name:encrypted password:GID:user list
getent group
```

### 权限控制

权限信息

```
|   0   |  1  |  2  |  3  |  4  |  5  |  6  |  7  |  8  |  9  |
| types |      owner      |      group      |   other users   |
| -/d/l | r/- | w/- | x/- | r/- | w/- | x/- | r/- | w/- | x/- |

-/d/l: 文件/文件夹/软链接
r: 文件 => 可以查看文件内容/文件夹 => 可以查看文件夹内容(ls)
w: 文件 => 可以修改文件内容/文件夹 => 可以在文件夹中进行创建、删除、改名等操作
x: 文件 => 可以将文件作为程序执行/文件夹 => 可以更改工作目录到此文件夹(cd)
```

`chmod`

```shell
# 由 (文件/文件夹的所属用户或 root 用户) 修改文件/文件夹的权限信息
# -R 表示对文件夹内的全部内容应用同样的操作
chmod [-R] permission file-path/directory-path

# 将文件权限修改为 rwxr-x--x
chmod u=rwx, g=rx, o=x hello.txt
# 将文件夹 test 及其全部内容的权限设置为 rwxr-x--x
chmod -R u=rwx, g=rx, o=x test

# 权限的 3 位数字表示法: r-4 w-2 x-1
#  0   1   2   3   4   5   6   7 
# --- --x -w- -wx r-- r-x rw- rwx
chmod 515 hello.txt # r-x--xr-x
chmod 326 hello.txt # -wx-w-rw-
```

`chown`

```shell
# 只能由 root 用户执行以修改文件/文件夹所属的用户和用户组
chown [-R] [owner][:][group] file-path/directory-path
```

# 更多操作

### 快捷按键

`CTRL + C` —— 强制停止/退出当前输入后重新输入

`CTRL + D` —— 退出登录的用户

`history` —— 历史命令

`CTRL + R` —— 输入内容去匹配历史命令。匹配 => 回车直接执行；不匹配 => 按左右键得到命令

`CTRL + A / E / <- / ->` —— 光标跳到命令开头/结尾/上一个单词/下一个单词

`CTRL + L / clear` —— 清空终端内容

### 软件安装

在 CentOS/Ubuntu 中联网安装软件的命令：

```shell
# 这些命令都要由 root 用户执行
yum/apt [-y] [install] [remove] [search] software-name # -y 表示自动确认
```

### 服务控制



### 软链接



# 软件安装

### 在 Ubuntu 上安装 MySQL 8.0 数据库

1. 切换到 root 用户

   ```shell
   sudo su -
   ```

2. 更新 apt 仓库信息

   ```shell
   apt update
   ```

3. 安装 MySQL 数据库

   ```shell
   apt install -y mysql-server
   ```

4. 启动 MySQL 服务器

   ```shell
   /etc/init.d/mysql start  # 启动
   /etc/init.d/mysql stop   # 停止
   /etc/init.d/mysql status # 查看状态
   ```

5. 登录

   ```shell
   mysql
   ```

6. 设置 root 用户的密码

   ```sql
   ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '******';
   ```

7. 退出

   ```mysql
   exit
   ```

8. 对 MySQL 进行初始化

   ```shell
   # which mysql_secure_installation
   mysql_secure_installation
   输入 root 用户的密码 -> ******
   是否启用 VALIDATE PASSWORD COMPONENT 来增强密码安全性 -> No
   是否更改 root 用户的密码 -> No
   是否移除匿名用户以增强安全性 -> Yes
   是否禁用 root 用户的远程登录 -> No
   是否移除自带的测试数据库 -> No
   是否刷新权限 -> Yes
   ```

9. 用更改后的密码重新登录

   ```shell
   mysql -u root -p
   ```

