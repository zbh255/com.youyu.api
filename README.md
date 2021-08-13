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

> 编译`docker`镜像

---

### conf

> 本程序的配置使用的是Toml格式,关于Tmol的语法

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
