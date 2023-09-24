package systemdesign

// 在浏览器中输入 URL 发生的事情
//
// URL => Universal Resource Locator
//
//  https://google.com/product/electric/resource
// |++++++||++++++++++|++++++++++++++++|++++++++|
//  Scheme    Domain         Path       Resource
//                    |+++++++++++++++++++++++++|
//          they together specify the resource on the server
//
// 1. 在浏览器中输入 URL 并按下 Enter 键
// 2. 域名解析 => 浏览器缓存 => 系统缓存 => 递归/迭代 DNS 查询
// 3. 建立 TCP 连接
// 4. 浏览器发送 HTTP 请求
// 5. 服务器返回 HTTP 响应
// 6. 浏览器渲染 HTTP 内容
//
// 参考：
// https://zhuanlan.zhihu.com/p/133906695
// https://github.com/alex/what-happens-when
