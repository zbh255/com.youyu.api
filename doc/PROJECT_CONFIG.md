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
# 数据库接入层rpc的地址
[data_rpc_server]
    ip = "127.0.0.1"
    port = "5000"

# 签钥和管理用户token的rpc地址
[secret_key_rpc_server]
    ip = "127.0.0.1"
    port = "5002"

# rpc_server使用
# 该栏目一下的文件为项目所使用到的组件的配置文件
# 有子配置.sync说明会使用连接池来管理了客户端连接
[database]
    ip_addr = "10.190.44.39"
    port = "3306"
    user_name = "root"
    user_password = "Cok774.."
    database_name = "youyu"
[database.sync]
    # 最大空闲连接数
    db_max_idle_size = 200
    # 空闲连接最大存活时间
    db_max_idle_life_time = 5
    # 最大打开连接数
    db_max_open_conn_size = 500
    # 连接存活时间
    db_max_conn_life_time = 10

[redis]
    ip_addr = "10.190.44.39"
    port = "6379"
    password = "cok774.."
[redis.sync]
    # 最大连接数
    max_open_conn_size = 4999
    # 最小保持的连接数
    min_open_conn_size = 300
    # 连接在池内存活的时间
    max_conn_life_time = 5
    # 建立连接超时的时间
    dial_timeout = 15
    # 当池中没有空闲连接时客户端等待空闲连接的最长时间
    pool_timeout = 10
    # 连接池的闲置连接的超时时间
    idle_timeout = 5

[server]
    ip_addr = "127.0.0.1"
    port = "8080"
    protocol = "http"
[server.sync.data_rpc]
    # 初始化连接数
    grpc_poll_init_cap_size = 100
    # 最大连接数
    grpc_poll_max_cap_size  = 500
    # 最大空闲连接数
    grpc_poll_max_idle-size = 200
    # 连接存活时间
    grpc_poll_max_idle_timeout = 10
[server.sync.secretKey_rpc]
    # 初始化连接数
    grpc_poll_init_cap_size = 300
    # 最大连接数
    grpc_poll_max_cap_size  = 4999
    # 最大空闲连接数
    grpc_poll_max_idle-size = 400
    # 连接存活时间
    grpc_poll_max_idle_timeout = 5

# 项目的配置
[project]

# 项目的权限认证相关
[project.auth]
    # 保留配置，将来可能会用到
    token_type = "jwt"
    # 保留配置，将来可能会用到
    token_signture = "HS256"
    # token默认有效时间,注意该选项仅供测试,格式: n * 分钟
    token_timeout = 1440
    # 默认签名密钥,16个字符为标准
    token_signture_key = "e126789F78901134"

# 项目的加解密相关
[project.encrypt]
    # rsa密钥长度，仅供测试用途,默认为1024
    rsa_Key_size = 1024
```

---

`err_msg.json` -> 错误信息配置

> 错误信息使用`uts`工具生成，生成命令请查看帮助(`uts -h`)

```json
{
	"-1": "应用程序不存在或已被封禁",
	"-101": "账号未登录",
	"-102": "账号被封停",
	"-105": "验证码错误",
	"-106": "账号未激活",
	"-108": "账号未激活",
	"-110": "账号未激活",
	"-111": "csrf 校验失败",
	"-112": "系统升级中",
	"-113": "账号尚未实名认证",
	"-114": "请先绑定手机",
	"-115": "请先完成实名认证",
	"-2": "Access Key错误",
	"-3": "API校验密匙错误",
	"-304": "木有改动",
	"-307": "撞车跳转",
	"-4": "调用方对该Method没有权限",
	"-400": "请求错误",
	"-401": "未认证",
	"-403": "访问权限不足",
	"-404": "啥都木有",
	"-5": "密钥已过期",
	"-500": "服务器错误",
	"-503": "过载保护,服务暂不可用",
	"-504": "服务调用超时",
	"-509": "超出限制",
	"-616": "上传文件不存在",
	"-617": "上传文件太大",
	"-625": "登录失败次数太多",
	"-626": "用户不存在",
	"-628": "密码太弱",
	"-629": "用户名或密码错误",
	"-650": "用户等级太低",
	"-652": "重复的用户",
	"-658": "Token 过期",
	"-659": "令牌签名验证失败",
	"-660": "令牌不正确",
	"-662": "密码时间戳过期",
	"0": "正确",
	"10001": "Json参数解析失败",
	"10002": "Url参数解析失败",
	"10003": "Toml参数解析失败",
	"10004": "Xml参数解析失败",
	"10005": "客户端uuid不合法",
	"10006": "编码错误",
	"10011": "sql不合法",
	"10012": "Mysql服务器错误",
	"10013": "Redis服务器错误",
	"10014": "Grpc服务器错误",
	"10201": "添加文章失败",
	"10202": "删除文章失败",
	"10203": "更新文章失败",
	"10204": "查询不到该文章",
	"10205": "添加文章热度失败",
	"10206": "添加文章点赞失败",
	"10207": "删除文章点赞失败",
	"10208": "获取文章文章的热度、点赞信息失败",
	"10301": "获取广告失败",
	"10302": "添加广告失败",
	"10303": "更新广告失败",
	"10304": "删除广告失败",
	"10401": "用户注册失败",
	"10402": "用户登录失败",
	"10403": "用户删除失败"
}
```

