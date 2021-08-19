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

### build

```makefile
make mod
```

> `mod` 初始化需要的Go编译的包依赖

```makefile
make create
```

> 创建运行时所需要的文件

```makefile
make build-linux
```

> 编译`linux`版本的程序

```makefile
make build-docker
```

> ~~编译`docker`镜像~~ `该Target已经废弃，不建议使用`

---

### run

> 执行编译命令成功之后的文件夹名为`build_release`
>
> `build_release`的目录结构为

```shell
.
|-- business
|-- business.dockerfile
|-- cache
|   `-- rsa
|-- cent_rpc
|-- cent_rpc.dockerfile
|-- conf
|   |-- app.conf.toml
|   `-- business.conf.toml
|-- data_rpc
|-- data_rpc.dockerfile
|-- dir
|-- docker-compose.yml
|-- log
|   |-- cent_rpc.log
|   |-- data_rpc.log
|   `-- gin.log
`-- script
    `-- mysql
        `-- mysql
            `-- youyu.sql
```

> 部署使用`docker-compose`管理容器

```dockerfile
docker-compose build
```

> docker-compose run

```sh
docker-compose up -d
```

> 查看容器运行情况

```sh
docker ps
```

---

### conf

> 默认将容器内的配置目录`app/conf`映射到宿主机下的`./conf`，所以在`./conf`下修改配置即可

> 本程序的配置使用的是Toml格式,关于Tmol的语法请查看这篇官方文档的中文翻译:[知乎](https://zhuanlan.zhihu.com/p/50412485)

> 配置文件选项解析

`business.app.toml`

```toml
[cent_rpc_server] # 配置和日志中心的地址
    ip = "127.0.0.1"
    port = "5001"
```

`app.conf.toml`

```toml
[data_rpc_server]
    ip = "127.0.0.1" # 数据rpc的ip地址
    port = "5000" # 端口

# rpc_server使用
[database]
    ip_addr = "192.168.31.24"
    port = "3306"
    user_name = "root"
    user_password = "Cok774.."
[database.sync]
    db_max_idle_size = 300 # 最大空闲连接数
    db_max_open_conn_size = 400 # 最大打开连接数
    db_max_conn_life_time = 15 # 连接存活时间

[server]
    ip_addr = "127.0.0.1" # gin web_server ip
    port = "8080" # port
    protocol = "http" # 使用的协议,目前只对http进行了支持
[server.sync]
    grpc_poll_init_cap_size = 100 # 初始化连接数
    grpc_poll_max_cap_size  = 300 # 最大连接数
    grpc_poll_max_idle-size = 200 # 最大空闲连接数
    grpc_poll_max_idle_timeout = 15 # 连接存活时间,以秒为单位
```



---

### struct

> 本程序目录结构

<details>
<summary>展开查看</summary>
<pre><code>.
├─app
│  ├─business_server -- 业务服务器
│  │  └─controller -- 控制器层
│  └─rpc -- grpc调用的一些实例
│      ├─client -- grpc 客户端的获取函数
│      ├─model -- dao层
│      ├─proto_files -- protobuf文件生成的service代码存根
│      └─server -- grpc 服务器的具体实现函数
├─build_debug -- 编译缓存文件
│  ├─conf
│  └─log
├─cmd
│  ├─main_server -- gin web服务器的启动程序
│  ├─rpc_cent_cmd -- 配置中心和日志中心的启动文件
│  └─rpc_cmd -- 数据rpc的启动文件
├─common
│  ├─config -- toml配置文件的解析
│  ├─database -- 数据库连接的支持
│  ├─errors -- 错误码
│  ├─interface -- 存放公共统一的接口
│  ├─log -- zerolog的封装
│  ├─middleware -- gin的中间件服务
│  ├─path -- 项目公共的路径和文件名称
│  ├─router -- gin路由
│  └─utils -- 独立的工具包
├─conf
│  ├─dev -- 配置文件的原型
│  └─pro -- 配置文件的原型
├─internal -- grpc service文件
│  └─proto_file -- protobuf生成的go代码文件
│      └─com.youyu.api
│          └─app
│              └─rpc
│                  └─proto_files
├─script -- 程序用到的一些脚本，比如数据库初始化脚本
└─test -- 测试代码
    ├─conf
    └─log
</code></pre>
</details>
