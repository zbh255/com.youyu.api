FROM alpine

# 环境变量
ENV GIN_MODE=release

# 复制需要的文件
COPY ./conf/business.conf.toml /app/conf/business.conf.toml
COPY ./business /app
# 容器的工作目录
WORKDIR ./app

# 对外服务的端口
EXPOSE 8080
EXPOSE 80
EXPOSE 443
# 启动容器时打开的命令
CMD ["./business"]