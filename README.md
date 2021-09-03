# wee-gin-shop

## 项目介绍
`wee-gin-shop`项目是一套电商系统，包括前台商城系统及后台管理系统，基于时下流行的云原生语言go以及g,grpc开发。项目整体采用微服务架构，对内采取grpc通信。通过此电商项目，学习了解微服务中间的各种组件，掌握从0到1成熟的微服务落地方案。readme会根据进度后续逐渐完善，形成一份快速微服务入门文档。
### 技术栈
| 技术                 | 说明                | 地址                                       
| -------------------- | -------------------------------------- | ---------------------------------------------- |
|gin|go语言web框架| https://github.com/gin-gonic/gin
|grpc|高性能RPC框架|https://grpc.io/docs/languages/go/quickstart/
|gorm|ORM框架|https://gorm.io/|
|mysql|关系型数据库|
|rocketmq|分布式消息队列|https://github.com/apache/rocketmq
|elaticsearch|全文搜索引擎
|kibana|可视化日志工具
|consul|注册中心|https://www.consul.io/
|nacos|分布式配置中心|https://nacos.io/
|jaeger|链路追踪|https://www.jaegertracing.io/
|sentinel|限流|https://sentinelguard.io/
|oss|对象存储|
|kong|网关
|jenkins|自动化部署|https://github.com/jenkinsci/jenkins

### 目录
- [Grpc快速入门](#Grpc快速入门)

### Grpc快速入门
#### Grpc安装
1. 安装插件
```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
2. 安装protoc   
https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3 下载压缩包  
3. 解压安装
4. 将安装包下/bin目录加入环境变量并重启ide或则终端
5. 运行以下命令，标识安装成功
```bash
protoc --version
libprotoc 3.17.3
```
#### Grpc Hello world
1. 定义proto文件  
示例代码项目目录example/grpc-example/proto
2. 生成代码，进入示例代码目录,运行以下命令,生成对应的hello.pb.go和hello_grpc.pb.go文件 
```bash
protoc --go_out=. --go-grpc_out=. .\hello.proto
```
3. 在确保路径，proto文件没有错误后还是不能生成对应的pb文件请优先检查protoc、protoc-gen-go、protoc-gen-go-grpc版本
4. 参照example/grpc-example/server example/grpc-example/client的示例代码实现grpc的服务端、客户端
