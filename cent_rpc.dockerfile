FROM alpine

# 复制需要的文件
COPY ./build_release/conf /app/conf
COPY ./build_release/log /app/log
COPY ./build_release/cent_rpc /app

# 容器的工作目录
WORKDIR ./app

# 对外服务的端口
EXPOSE 5001

# 启动容器时打开的命令
CMD ["./cent_rpc"]