# go-http-server
Create HTTP server with go language.
## 一、HTTP Server的简单实现
### 1.简单的Server接口定义
该接口包含了路由、服务启动、服务关闭功能
### 2.路由中包含了访问路径以及处理函数
### 3.服务启动
初始化服务时，可以起对应服务名称，为启动多个服务和后续日志上用得上。
指定对应地址和端口便可启动服务。
### 4.关闭服务
因简单的服务实现，固关闭服务时简单模拟了暂停几秒钟的设置，没有具体的资源处理。

## 二、HTTP Server的工程化实现
### 1.在SimpleServer的基础上进一步封装
具体封装了如下点：
- Http请求上下文
- 处理函数Handler
- 过滤器的封装
- 路由树
- Server
### 2.封装点以及思路
#### 2.1 Http请求上下文
- 创建Context结构体
- 实例化Context
- 封装Context上读数据
- 封装Context上写数据
- 封装常见StatusCode对应的写数据
#### 2.2 处理函数Handler
- 定义处理器接口
#### 2.3 过滤器
- 建立一个过滤链
- 应用到日志/健康检查等场景
#### 2.4 路由树
- 创建请求路径的路由树
- 查询请求路径的路由树
#### 2.5 Server
##### 2.5.1 服务初始化
- 定义路由树上的处理器
- 定义过滤链
- 指定服务名、处理器、过滤器
##### 2.5.2 服务启动
- 服务器地址+端口为参数
- 监听服务器地址+端口
##### 2.5.3 服务优雅退出
- 发出服务退出指令（命令）
- 拒绝新的请求
- 等待当前的所有请求处理完毕
- 释放当前服务所涉及的资源
- 关闭服务（在这阶段：超时/收到多次退出命令情况只能强制退出）
## 三、制作Docker镜像
- 1.编写Dockerfile
- 2.构建build
  docker build -t weilesi/httpserver:v1.0.0 .
- 3.拉去和启动镜像httpserver
  docker run -d --name myhttpserver -p 80:8099 weilesi/httpserver:v1.0.0
- 4.在客户端验证
  http://82.157.31.144/
  通过postman验证更多接口，例如：http://82.157.31.144/user/login
- 5.获取当前容器的pid
  docker inspect myhttpserver|grep -i pid
- 6.用nsenter命令查看ip和路由等
  nsenter -t 1258522 -n ip a
  nsenter -t 1258522 -n route
- 7.把镜像文件推到私有仓库中
  docker push weilesi/httpserver:v1.0.0



