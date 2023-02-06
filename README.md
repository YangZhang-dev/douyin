<h1 align="center" style="font-size:50px">douyin</h1>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.18-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.7-yellowgreen"/>
<img src="https://img.shields.io/badge/go--redis-8.11.5-brightgreen"/>
<img src="https://img.shields.io/badge/gorm-1.23.5-red"/>
<img src="https://img.shields.io/badge/viper-1.12.0-orange"/>
<img src="https://img.shields.io/badge/oss-2.2.6-orange"/>
</div>

# 1. 基本介绍

douyin是一个极简版抖音的后端接口，为极简版抖音客户端提供后端服务。

### 1.1 功能介绍

- 视频：视频推送、视频投稿、发布列表
- 用户：用户注册、用户登录、用户信息
- 点赞：点赞操作、点赞列表
- 评论：评论操作、评论列表
- 关注：关注操作、关注列表、粉丝列表、好友列表
- 聊天：聊天操作、聊天记录

## 1.2 数据库设计

视频和用户是两大实体，点赞，评论，关注，聊天这些操作建立新的表。

其中将用户，关注，聊天存放于`douyin_user`库中，将视频，点赞，评论存放于`douyin_video`库中。

设计以下数据库（sql文件位于`/deploy/sql`）

![数据库设计图](.\deploy\sql\dbModel.png)

> 这里所有的外键都是逻辑外键，通过事务保证数据的正确性。

## 1.3 技术选型

本项目采用基于`go-zero`的RPC框架，包含了`go-zero`以及相关`go-zero`作者开发的一些中间件，所用到的技术栈基本是`go-zero`
项目组的自研组件。

- Go-zero
- Mysql
- Redis
- Oss

## 1.4 服务架构

| 服务名    | 用途                               | 框架                 | 协议       | 注册中心 | 监控 | 数据存储      | 日志   |
| :-------: | :--------------------------------: | :------------------: | :--------: | :------: | :------: | :-----------: | :----: |
| user-api  | 用户API微服务 | `go-zero`            | `http`     | `etcd`   | `prometheus` |               | `logx` |
| user-rpc  | 用户RPC微服务          | `go-zero` | `protobuf` | `etcd`   | `prometheus` | `mysql` `redis` | `logx` |
| video-api | 视频API微服务       | `go-zero` `mysql` `redis` `oss` | `http` | `etcd`   | `prometheus` |  | `logx` |
| video-rpc | 视频RPC微服务                      | `go-zero` |`protobuf`|`etcd`|`prometheus`|`mysql` `redis` `oss`|`logx`|

注：MapReduce，加redis,错误处理,测试ffmpeg

缓存不一致


