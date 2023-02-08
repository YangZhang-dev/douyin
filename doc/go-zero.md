## 1.goctl基础

### 1.1 api

api: goctl api go -api *.api -dir ../ --style=goZero

doc: goctl api dpc --dir ./

### 1.2 rpc

goctl rpc protoc user.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero

sql2pb -go_package ./pb -host 192.168.72.128 -package pb -password root -port 3306 -schema douyin -service_name user
-table user -user root > user.proto

### 1.3 protoc

protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. userModel.proto

### 1.4 model

goctl model mysql datasource -url ="root:root:@tcp(192.168.72.128:3306)/douyin" -table = "follow" -dir="./"
--style=goZero -c --home= ""

goctl model mysql ddl -src="./*.sql" -dir="" -c --style=goZero -c --home= ""

## 2. go-zero 实践

简易版抖音项目

[接口文档]((https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707520))

大概有用户服务，视频服务，用户聊天，关注用户，点赞视频，评论视频

#### 2.1 数据库

根据微服务架构来建好数据库。

建如下数据库：

1. douyin_user
2. douyin_video

#### 2.2 生成model文件

首先在commom下创建三个全局的文件夹，一个用于httpresult，一个是全局的key，一个是存放错误码

```shell
douyin\deploy> goctl model mysql ddl -src="./sql/*.sql" -dir="../model" -c --style=goZero -c --home .\goctl\1.4.3\
```

goctl model mysql ddl -src="./sql/*.sql" -dir="../model" -c --style=goZero -c --home .\goctl\1.4.3\

运行 go mod tidy 下载依赖

#### 2.3 目录结构

创建app目录

​ app下创建user和video

​ user下创建cmd和model

​ cmd下创建api和rpc

​ video同理

#### 2.4 编写api文件

在api下创建desc文件夹，用于存放api文件

对于api中的实例信息，建立文件夹存放，在desc下建立user文件夹存放user实例

可以搭配validator进行参数校验

#### 2.5 生成api代码

```shell
douyin\app\user\cmd\api\desc> goctl api go -api userApi.api -dir ../  --style=goZero --home ../../../../../deploy/goctl/1.4.3/
```

在etc/userApi.yaml中配置信息

```yaml
Name: userApi
Host: 0.0.0.0
Port: 8888

Log:
  ServiceName: userApi
  Encoding: plain
  Level: info
```

#### 2.6 编写proto文件

使用sql2pb工具生成pb文件

```shell
douyin\app\user\cmd\rpc\pb> sql2pb -go_package  ./pb -host 192.168.72.128 -package pb -password root -port 3306 -schema douyin_user -service_name user -table user -user root > userRpc.proto
```

根据自己的需求进行修改

#### 2.7 生成rpc代码

```shell
douyin\app\user\cmd\rpc\pb> goctl rpc protoc videoRpc.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero --home ../../../../../deploy/goctl/1.4.3/
```

在etc/userRpc.yaml中配置信息

```yaml
Name: userRpc
ListenOn: 0.0.0.0:8080
Mode: dev
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: userRpc

Log:
  ServiceName: userRpc
  Encoding: plain
  Level: info
```

#### 2.8 rpc配置数据库和缓存连接

在etc/userApi.yaml中配置信息

```yaml
# DB,Cache
DB:
  DataSource: root:root@tcp(127.0.0.1:3306)/douyin_user?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
Cache:
  - Host: 127.0.0.1:6379
    Pass:
```

在config/config.go配置

```go
type Config struct {
zrpc.RpcServerConf
DB struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
}
Cache   cache.CacheConf // redis缓存
}

```

在svc/serviceContext.go中配置信息

```
type ServiceContext struct {
   Config    config.Config
   UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
   return &ServiceContext{
      Config:    c,
      UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
   }
}
```

#### 2.9 api配置连接rpc

在etc/user.yaml中配置信息

```shell
UserRpcConf:
  Etcd:
    Hosts:
      - 0.0.0.0:2379
    Key: userRpc
```

在internal/config中配置

```go
type Config struct {
rest.RestConf
UserRpcConf zrpc.RpcClientConf
}
```

在svc/serviceContext.go中配置

```go
type ServiceContext struct {
Config                config.Config
OptionalJWTMiddleware rest.Middleware
UserRpc               userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
return &ServiceContext{
Config:                c,
OptionalJWTMiddleware: middleware.NewOptionalJWTMiddleware().Handle,
UserRpc:               userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
}
}
```

sudo du -sh ./* --exclude proc
