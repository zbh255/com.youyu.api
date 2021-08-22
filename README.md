# com.youyu.api

> @author : `abingzo`

> 游鱼旅游网站的后台，基于GO语言构建，在本文档中，您将获得关于该项目的知识

> 本程序只对Linux兼容，测试服务器系统为`Centos7`

---

### 代码规范

- `common`包的每个子包必须有包注释，每个函数/方法应该都有功能注释和返回的错误类型注释

- `cmd`文件夹中只能存放启动一个服务器相关的程序
- 只有`cmd`文件夹下的初始化操作才允许`panic`操作，业务逻辑不允许`panic`

---

### 一些约定

> 注: *为任何可能的文件夹

>  `com.youyu.api/app/*`放置的是项目要编译的主程序和工具程序，必须要有以个`cmd`文件夹，里面存放`main.go`文件，编译程序的名字约定为`cmd`父文件夹的名字,比如下方代码块的目录结构,按照约定，编译的程序名为`business`和`uts`
>
> `com.youyu.api/app/rpc/server/*`则不同上，因为是`Rpc service`服务，所以`cmd`目录放在`server/*`下,按照约定，编译出来的程序名字也是`cmd`的父目录

```shell
|-- app
|   |-- business
|   |   |-- cmd
|   |   |   `-- main.go
|   |   `-- controller
|   |       |-- advertisement.go
|   |       |-- article.go
|   |       |-- base.go
|   |       `-- connect.go
|   |-- rpc
|   |   |-- client
|   |   |   |-- client.go
|   |   |   `-- io.go
|   |   |-- model
|   |   |   `-- database_table_bind.go
|   |   |-- proto_files
|   |   |   |-- rpc_cent.pb.go
|   |   |   `-- rpc_service.pb.go
|   |   `-- server
|   |       |-- cent_rpc
|   |       |   |-- CentApi.go
|   |       |   `-- cmd
|   |       |       `-- main.go
|   |       `-- data_rpc
|   |           |-- cmd
|   |           |   `-- main.go
|   |           `-- MysqlApi.go
|   `-- uts
|       |-- cmd
|       |   `-- main.go
|       `-- uts.go
```

---

### 本项目的其他文档

- [项目配置解析](./doc/PROJECT_CONFIG.md)
- [项目的部署文档](./doc/DEPLOY.md)
- [项目的结构解析](./doc/PROJECT_STRUCT)
- [项目的版本发行日志](./RELEASE_NODE)

