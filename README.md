#  第六届字节跳动青训营`minitok`低配抖音项目

基于`gorm`框架+`gin`http框架 使用`docker`进行快速部署的低配版抖音项目：

### 主要技术框架及配置：

- ORM框架 `gorm`
- http框架 `gin`
- 数据库 `mysql` `redis`
- 分布式对象存储 `minio`
- 配置解析 `viper`
- 日志管理 `zap`
- 项目部署 `docker` `docker-compose`
- 其他工具 `FFmpeg`

项目概览：
![](./imgs/onefetch.jpg)

### 快速上手启动：

首先配置好`docker`以及`docker-compose`;

然后使用命令`sudo docker-compose up`快速构建部署运行项目

