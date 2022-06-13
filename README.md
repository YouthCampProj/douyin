# douyin

2022字节跳动青训营 抖音项目  
gin+gorm框架开发

## 接口实现

- [x] 基础接口  
- [x] 扩展接口-I  
- [x] 扩展接口-II  

## 测试URL

baseURL填写该地址: `https://youthcamp.114514.dev`  

## 使用方法

1. 安装ffmpeg（用于上传视频后的处理）  
MacOS: `brew install ffmepg`  
Ubuntu: `apt install ffmepg`
2. 复制`config/config.example.yaml`为`config/config.yaml`
3. 修改配置文件`config/config.yaml`  
domain配置项用于上传视频后生成的`play_url`与`cover_url` 注意将域名解析到后端所监听的IP  
mysql相关配置只需要建立数据库并分配用户权限 数据表会在首次启动时自动生成

## 技术说明

### 目录结构

```zsh
.
├── README.md
├── config
│   └── config.example.yaml #示例配置文件
├── controller  
│   ├── Comment.go  #评论
│   ├── Common.go 
│   ├── controller.go
│   ├── favorite.go #点赞
│   ├── feed.go     #视频流
│   ├── file.go     #静态资源
│   ├── publish.go  #投稿发布
│   ├── relation.go #关注关系
│   └── user.go     #用户
├── go.mod
├── go.sum
├── main.go
├── model #数据库交互实体层
│   ├── comment.go  #评论
│   ├── common.go
│   ├── favorite.go #点赞
│   ├── feed.go     #视频流
│   ├── model.go
│   ├── publish.go  #投稿发布
│   ├── relation.go #关注关系
│   ├── user.go     #用户
│   └── video.go    #视频
├── pkg
│   ├── auth    #认证组件
│   │   └── token.go
│   ├── config  #外部配置组件
│   │   └── config.go
│   ├── fileprocess #文件处理组件
│   │   └── cover.go    #封面处理
│   ├── serializer  #序列化组件
│   │   ├── codemsg.go  #错误消息
│   │   ├── comment.go  #评论
│   │   ├── favorite.go #点赞
│   │   ├── feed.go     #视频流
│   │   ├── publish.go  #投稿发布
│   │   ├── relation.go #关注关系
│   │   ├── response.go #通用响应结构
│   │   ├── token.go    #token
│   │   ├── user.go     #用户
│   │   └── video.go    #视频
│   └── snowflake   #雪花算法组件
│       └── snowflake.go
├── service
│   ├── auth.go     #认证服务
│   ├── comment.go  #评论服务
│   ├── favorite.go #点赞服务
│   ├── feed.go     #视频流服务
│   ├── publish.go  #投稿发布服务
│   ├── relation.go #关注关系服务
│   ├── service.go  #服务入口
│   └── user.go     #用户服务
├── uploads #静态资源
└── utils   #工具组件
    └── utils.go

12 directories, 47 files
```

### 架构设计

请求:  
客户端->controller->service->model->serializer

响应:  
正常响应: serializer->service->controller->客户端  
异常响应: serializer->controller->客户端

### 项目亮点

controller层在设计之初便考虑了后续通过rpc进行调用的方式，使用了预先构造请求结构的方法对函数进行调用，在后续引入grpc框架时可以做到快速替换升级。

service层主要负责实际业务逻辑的处理，我们为每个业务逻辑提供了请求-响应结构体，并且提供了对应的业务逻辑函数，后续引入grpc框架时可以做到快速替换升级。得益于gin框架的并发请求结构，实际上可以做到同时处理多个请求。

model层通过引入gorm进行数据库操作，并且使用了gorm的结构体组合模型，使得数据库操作变得简单且智能。将点赞、关注等业务分离为独立表，在后续引入redis等缓存组件时可以做到快速替换升级。

serializer组件分离至独立的包使得构造响应结构变得无比简单。  
无论是充当API网关的controller层还是实际提供服务的service层，都可以轻松调用serializer组件进行响应体构造，方便了后续重构为微服务架构的服务拆分和组装。  

## Credits

<https://github.com/gin-gonic/gin>  
<https://github.com/go-gorm/gorm>  
<https://github.com/spf13/viper>  
