# 项目配置

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
    grpc_poll_max_cap_size= 300 # 最大连接数
    grpc_poll_max_idle-size = 200 # 最大空闲连接数
    grpc_poll_max_idle_timeout = 15 # 连接存活时间,以秒为单位
```

---