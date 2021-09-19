### 开发文档
> 一些编译指令
> 编译带参数验证.proto文件

```shell
protoc --proto_path=C:/Users/harder/go/pkg/mod --proto_path=C:/Users/harder/Desktop/github.com/abingzo/com.youyu.api/lib/internal/proto_file --govalidators_out=. --go_out=plugins=grpc:. rpc_service.proto
```

> Macos
```shell
protoc --proto_path=/Users/harder/go/pkg/mod --proto_path=/Users/harder/Desktop/Git-Repo/github.com/abingzo/com.youyu.api/lib/internal/proto_file --govalidators_out=. --go_out=plugins=grpc:. rpc_service.proto
```
