# Go
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GORELEASEARG=-ldflags '-s -w'
GODEBUGARG=-ldflags '-s -w'
BINARY_MAIN_NAME=business
BINARY_RPC_NAME=data_rpc
BINARY_CENT_NAME=cent_rpc
BUILD=linux
APP_VERSION=0.1.2

# 环境变量
export GIN_MODE=release
export BUPS_MODE=release
export GOPROXY="https://goproxy.cn,direct"
export GO111MODULE=on
export CGO_ENABLED=0
export UPX_OPTIONS=off

# 文件与目录路径
build_path=./bulid_release
config_version=pro
ifeq ($(BUPS_MODE),debug)
build_path=./build_debug
config_version=dev
else
build_path=./build_release
config_version=pro
endif

sources := $(wildcard ./app/*)
# path
Main_APP_CMD=./cmd/main_server
DataRpc_APP_CMD=./cmd/rpc_cmd
CentRpc_APP_CMD=./cmd/rpc_cent_cmd
project_config_path=./conf/$(config_version)
project_script_path=./script
project_docker_script_path=$(project_script_path)/docker
project_mysql_script_path=$(project_script_path)/mysql

# 项目配置文件的所在
config_path=$(build_path)/conf
build_path_cache=$(build_path)/cache
build_path_cache_rsa=$(build_path_cache)/rsa
build_path_dir=$(build_path)/dir
build_path_log=$(build_path)/log

# file
config_path_file=app.conf.toml
business_conf_path_file=business.conf.toml
build_path_log_data=data_rpc.log
build_path_log_gin=gin.log
build_path_log_cent=cent_rpc.log
build_source_file=main.go
# docker file
build_docker_business=business.dockerfile
build_docker_dataRpc=data_rpc.dockerfile
build_docker_centRpc=cent_rpc.dockerfile

mod:
	@$(GOMOD) download -json

.PHONY:create
create:
	# 创建目录
	mkdir -p $(build_path_cache)
	mkdir -p $(build_path_cache_rsa)
	mkdir -p $(config_path)
	mkdir -p $(build_path_dir)
	mkdir -p $(build_path_log)
	mkdir -p $(build_path)/script/mysql
	# 创建文件
	touch $(build_path_log)/$(build_path_log_data)
	touch $(build_path_log)/$(build_path_log_gin)
	touch $(build_path_log)/$(build_path_log_cent)
	touch $(build_path_dir)/$(build_path_dir_info)
	# 拷贝项目的配置文件
	cp $(project_config_path)/* $(config_path)
	# 拷贝项目运行所需要的脚本
	cp $(project_docker_script_path)/* $(build_path)
	cp -r $(project_mysql_script_path) $(build_path)/script/mysql

clean:
	echo 正在清理编译目录:$(build_path)
	rm -rf $(build_path)

build-darwin:
	@echo $(BUPS_MODE)
	@echo $(GOOS)
	cd $(Main_APP_CMD) && $(GOBUILD) -o $(BINARY_MAIN_NAME) -ldflags '-s -w'
	# 移动编译之后的文件
	mv $(Main_APP_CMD)/$(BINARY_MAIN_NAME) $(build_path)

build-windows:
	echo $(Main_APP_CMD)

# 编译linux版本的局部变量
build-linux:export GOOS=linux
build-linux:export GOARCH=amd64
build-linux:
	# 编译主程序文件
	# 切换编译目录
	cd $(Main_APP_CMD) && $(GOBUILD) -o $(BINARY_MAIN_NAME) -ldflags '-s -w'
	# 移动编译之后的文件
	mv $(Main_APP_CMD)/$(BINARY_MAIN_NAME) $(build_path)

	#编译Rpc程序
	cd $(DataRpc_APP_CMD) && $(GOBUILD) -o $(BINARY_RPC_NAME) -ldflags '-s -w'
	# 移动编译之后的文件
	mv $(DataRpc_APP_CMD)/$(BINARY_RPC_NAME) $(build_path)

	# 编译配置和日志中心
	cd $(CentRpc_APP_CMD) && $(GOBUILD) -o $(BINARY_CENT_NAME) -ldflags '-s -w'
	# 移动编译之后的文件
	mv $(CentRpc_APP_CMD)/$(BINARY_CENT_NAME) $(build_path)


# v0.1之后废弃改用docker-compose管理容器
build-docker:
	# 编译business docker镜像
	docker build -f $(build_docker_business) -t $(BINARY_MAIN_NAME):v$(APP_VERSION) .
	# 编译dataRpc docker 镜像
	docker build -f $(build_docker_dataRpc) -t $(BINARY_RPC_NAME):v$(APP_VERSION) .
	# 编译centRpc docker 镜像
	docker build -f $(build_docker_centRpc) -t $(BINARY_CENT_NAME):v$(APP_VERSION) .

build-linux-sub:
	@echo $(GOOS)