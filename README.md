# com.youyu.api

> @author : `abingzo`

> 游鱼旅游网站的后台，基于GO语言构建，在本文档中，您将获得关于该项目的知识

> 本程序只对Linux兼容，测试服务器系统为`Centos7`

---

### 代码规范

- `common`包的每个子包必须有包注释，每个函数/方法应该都有功能注释和返回的错误类型注释
- `cmd`文件夹中只能存放启动一个服务器相关的程序
- 只有`cmd`文件夹下的初始化操作才允许`panic`操作，业务逻辑不允许`panic`
- `rpc/server/*`的具体服务实现文件，实现要在注释中明确其返回的错误类型

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

### 错误处理的约定

> 使用`pkg/errors`处理错误，配合包级别变量，包级别变量定义逻辑错误，不考虑与第三方库的错误类型兼容
>
> 包级别错误变量存放于`*/modelName/errors.go`，比如`app/rpc/model/errors.go`中存储的是数据库`Dao`的错误变量
>
> 以`Model-Article`的添加文章的方法为例
>
> `AddArticle()`接受一个文章数据`Struct`的参数输入,第一步的动作是查询文章`id`是否存在，如果存在则会返回逻辑错误，如果`gorm`本身有错误，则返回本身的一个已存储堆栈信息的错误
>
> `方法`本身返回的错误应该明示，如下

```go
// AddArticle 获取文章的数据
// pkg/errors处理错误
// article_id 已存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *Article) AddArticle(article *Article) (*Article, error) {
	// 查找文章是否存在
	result := DB.Where("article_id = ?", article.Id).First(self)
	if result.RowsAffected > 0 {
		return nil, errors.WithStack(ArticleIdAlreadyExists)
	}
	result = DB.Create(article)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	result = DB.Create(&ArticleStatistics{
		Id:         article.Id,
		Fabulous:   0,
		Hot:        0,
		CommentNum: 0,
	})
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return article, nil
}
```

> `Rpc`调用者应该处理的错误
>
> 根据方法明示的错误处理不同的逻辑,`default`则处理非逻辑错误

```go
articleModel, err := md.AddArticle(&model.Article{
		Id:         hex.EncodeToString(hash.Sum(nil)),
		Abstract:   article.ArticleAbstract,
		Content:    article.ArticleContent,
		Title:      article.ArticleTitle,
		Tag:        utils.TagListToSplitStrings(article.ArticleTag),
		Uid:        article.Uid,
		CreateTime: t,
		UpdateTime: t,
	})
	switch errors.Cause(err) {
	case model.ArticleIdAlreadyExists:
		// article id 存在则重试
		goto aeLoop
	case nil:
		return &rpc.Article{
			ArticleId: articleModel.Id,
		}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "add article failed"))
		return &rpc.Article{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
```

> 包级别错误的定义
>
> 应该使用`errors`而不是`pkg/errors`，因为后者会在原始错误中存储错误堆栈信息
>
> 以下以`app/rpc/model/errors.go`为例

```go
var (
	// CreateSameExistence 写入的数据不能跟数据表中已存在的内容一样
	CreateSameExistence = errors.New("the same content exists when it is created")
	// UserNameAlreadyExists 用户名已经存在
	UserNameAlreadyExists = errors.New("the user name already exists")
	// UserDoesNotExist 用户不存在
	UserDoesNotExist = errors.New("the user does not exist")
	// UserPasswordORUserNameErr 用户密码或用户名错误
	UserPasswordORUserNameErr = errors.New("user password or user name error")
	// ArticleIdAlreadyExists 文章id已存在"
	ArticleIdAlreadyExists = errors.New("the article id already exists")
	// ArticleIdNotExists 文章id不存在
	ArticleIdNotExists = errors.New("the article id not exists")
	// AdvertisementIdNotExists 广告id不存在
	AdvertisementIdNotExists = errors.New("the advertisement id does not exist")
	// TagNameAlreadyExists 标签名已经存在
	TagNameAlreadyExists = errors.New("the tag name already exists")
	// TagNameNotExists 标签名不存在
	TagNameNotExists = errors.New("the tag name not exists")
	// TagIdNotExists 标签id不存在
	TagIdNotExists = errors.New("the tag id does not exist")
)
```

---

### 本项目的其他文档

- [项目配置解析](./doc/PROJECT_CONFIG.md)
- [项目的部署文档](./doc/DEPLOY.md)
- [项目的结构解析](./doc/PROJECT_STRUCT.md)
- [项目的版本发行日志](./RELEASE_NODE.md)

