# 版本发行日志

#### Version 0.1.3

- 重构目录结构
- 重构错误处理
- 优化编译使用的`makefile`
- 细化`data_rpc`的连接池控制
- 修复`gorm`连接池的获取的性能问题

#### Version 0.1.2

- 弃用部分编译选项

  > make build-docker 弃用

- 完善错误处理和日志

- 新增中心`Rpc`用于管理日志和配置

- 新增`docker-compose`用于管理容器

- `docker-compose`新增隔离网络

- 新增`grpc`连接池

#### Version 0.1

- 初始版本，提供了首页相关的测试接口和程序雏形，以及测试文件

