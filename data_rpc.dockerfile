FROM alpine

# 复制需要的文件
COPY ./build_release/conf/business.conf.toml /app/conf/business.conf.toml
COPY ./build_release/data_rpc /app

# 容器的工作目录
WORKDIR ./app

# 对外服务的端口
EXPOSE 5000

# 启动容器时打开的命令
CMD ["./data_rpc"]