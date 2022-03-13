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
## 四、httpserver部署到k8s集群中
### 4.1 第一部分
#### 4.1.1 要求
- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离
#### 4.1.2 实现
详细请看httpserver-deploy.yaml文件
### 4.2 第二部分
#### 4.2.1 要求
- 在第一部分的基础上更加完备的部署spec
- Service
- Ingress
- 如何保证整个应用的高可用
- 如何通过证书保证httpServer的通讯安全
#### 4.2.2 实现
- Service的实现请看httpserver-svc.yaml
- Ingress以及安全通讯的实现参考如下
  
````
# 1、Ubuntu上安装helm
sudo snap install helm --classic
# 2、通过helm方式安装ingress-nginx
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress

#3、查看安装情况
kubectl get pod -n ingress
kubectl get svc -n ingress

#4、无证书的实现ingres，请参考httpserver-ingress.yaml

#5、证书通过cert-manager实现
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/cert-manager.crds.yaml

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.7.1 \

#6、签发CA证书的配置
参考httpserver-issuer.yaml文件

#7、配置证书
kubectl get cert
kubectl get CertificateRequest

#8、通讯安全的Ingress
请参考httpserver-ingress-ca.yaml文件

````
## 五、httpserver添加监控
### 5.1 要求
- 为 HTTPServer 添加 0-2 秒的随机延时；
- 为 HTTPServer 项目添加延时 Metric；
- 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
- 从 Promethus 界面中查询延时指标数据；
- （可选）创建一个 Grafana Dashboard 展现延时分配情况。
### 5.2 实现
- 原main.go中添加prometheus相关代码
- 重新制作Docker镜像，命名为httpserver:v1.0.1-metrics,步骤参考"第三章"
- 在K8S集群中安装loki，同时安装Promethus，安装步骤参考如下：
````
1、helm repo add grafana https://grafana.github.io/helm-charts
2、helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false
````
- 在K8S中把loki和Promethus的对应ClusterIP转为NodePort方式对外暴露端口
- 编写deploy.yaml,详情查看httpserver-metrics-deploy.yaml
- 在K8S中按照httpserver-metrics-deploy.yaml进行部署
- 在Promethus界面中验证指标数据展现的情况
- 可以添加对应告警规则，利用AlertManager发送短信/邮件提醒作用
