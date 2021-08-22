# 部署文档

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