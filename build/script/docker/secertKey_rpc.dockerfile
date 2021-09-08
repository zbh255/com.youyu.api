FROM alpine

# 复制需要的文件
COPY ./conf/business.conf.toml /app/conf/business.conf.toml
COPY ./secertKey_rpc /app

# 容器的工作目录
WORKDIR ./app

# 对外服务的端口
EXPOSE 5002

# 启动容器时打开的命令
CMD ["./secertKey_rpc"]