# Simple_bank



一个基于Golang的简易银行,根据视频教程[Backend Master Class](https://bit.ly/backendmaster)写成,在某些地方有所改动

主要功能:创建用户时发送验证邮箱,用户的账号之间转账,用户登录的权限验证

### 数据库设计

![image-20230717150548337](https://cdn.jsdelivr.net/gh/LyntNy4n/md_image@main/img/image-20230717150548337.png)

蓝色线表示有外键约束

使用PostgreSQL数据库

- users 存储用户的各种信息,每个用户可以创建多个账号(account),一个账号对应一个币种
- verify_emails 存储创建用户时验证邮件的各种信息
- sessions 存储登录用户时的各种信息,比如刷新令牌(refresh token)
- accounts 存储用户下的账号信息
- entries 存储账号的收支记录
- transfers 存储不同账号之间的转账记录

### 数据库操作

使用`sqlc`,在sqlc.yaml文件中设置好连接的数据库即可将SQL查询生成go代码文件

支持事务操作

附带各种操作的测试

使用`gomigrate`进行数据库迁移

### 配置加载

使用`viper`,在.env文件中获取环境变量,更改配置不需要重新编译

### API

#### RESTful API

使用`Gin`搭建web服务器

使用`gomock`创建假数据库,使用`httptest`创建假web服务器,容易测试API

也可以使用grpc网关代替`Gin`

#### GRPC

定义proto文件后,使用`protoc`生成对应go代码,搭建grpc服务器

使用`grpc gateway`把http请求转化为grpc请求,做到同时服务http和grpc请求

使用`swagger UI`构建文档

使用`zerolog`对不同的gprc请求打印对应的结构化日志(json格式)

#### 鉴权

使用`JWT`/`Pasteo`获取短期的`访问令牌(access token)`

数据库内部存储`会话(session)`,会话中包含一个长时间的刷新令牌,以获得长期登录权限

### 异步+Redis

使用`asynq`与`Redis`来解决任务处理持久化问题

#### 发送验证邮件

在创建用户的同时发送验证邮件

使用`Redis`进行该任务处理

使用事务,避免数据库创建用户成功但Redis处理失败,导致下一次无法创建用户(重复)的问题

### 运行
无论使用哪种方法运行,都需要新建一个`private.env`文件设置自己的邮件地址和密码:
```
EMAIL_SENDER_ADDRESS=youremail
EMAIL_SENDER_PASSWORD=yourpassword
```
注意:邮箱密码一般都不会是账号密码,而是需要生成一个应用专用密码

可以在配置好各种环境后,运行`make server`命令来启动程序

本项目也配置好了docker-compose,其二进制文件与postgres一起组成容器
只需要运行`docker-compose up`即可
