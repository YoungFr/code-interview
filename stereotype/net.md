- 网络模型概览

  应用层工作在<font color=red>用户态</font>，传输层及以下工作在<font color=red>内核态</font>。

  传输层根据报文中的<font color=red>端口号</font>决定将报文发送给哪个应用。

  <font color=red>网络层</font>负责将数据从一个设备传输到另一个设备；设备的地址通过 IP 地址标识；IP 协议的两个功能：<font color=red>寻址</font>和<font color=red>路由</font>。

  IP 地址分成<font color=red>网络号</font>和<font color=red>主机号</font>两部分。

  将<font color=red>子网掩码</font>与 IP 地址相与得到网络号，将子网掩码取反后再与 IP 地址相与得到主机号。

- 键入网址到显示网页的过程

  1. 解析 URL 确定 Web 服务器名和要请求的资源的路径

  2. 生成 HTTP 请求信息

  3. 使用 DNS 查询域名对应的 IP 地址

     涉及到的服务器：<font color=red>本地</font> DNS 服务器（DNS 解析器）、<font color=red>根</font>域名服务器、<font color=red>顶级</font>域名服务器、<font color=red>权威</font>域名服务器

     DNS 缓存：浏览器缓存 => OS 缓存 => hosts 文件 => 本地 DNS 服务器

  4. 可靠传输 —— 三次握手建立 TCP 连接

  5. 远程定位 —— 使用 IP 发送网络包

  6. 两点传输 —— MAC 和交换机

  7. 离开子网 —— 路由器

- Linux 如何收发网络包

  TODO

- 什么是 HTTP 协议？ —— 超文本传输协议 HTTP 是一个在计算机世界里专门在<font color=red>两点之间</font>（服务器 <--> 客户端、服务器 <-> 服务器）<font color=red>传输</font>文字、图片、音频、视频等<font color=red>超文本</font>数据的约定和规范。

- 状态码

  `1XX` —— 表示中间状态的提示信息

  `2XX`  —— 成功处理客户端请求

  `200 OK` -> 正常；`202 No Content` -> 响应头没有 body 数据；`206 Partial Content` -> 分块下载或断点续传

  `3XX` —— 重定向

  

  