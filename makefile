GO_CMD = go

# project info
# 项目的基本信息
# 项目版本信息
PROJECT_VERSION = 0.1.3
# 项目名称
PROJECT_Named = com.youyu.api
# 项目绝对路径
PROJECT_ABS_PATH = $(shell pwd)

# build use dir
# 编译时会用到的文件夹
DIR_APP = ./app
DIR_APP_RPC = $(DIR_APP)/rpc/server
DIR_APP_JOB = $(DIR_APP)/job
DIR_CONFIG = ./build/conf
DIR_SCRIPT = ./build/script
DIR_LOG = ./build/log
DIR_INFO = ./build/dir
DIR_BUILD = ./project
DIR_VERSION = "com.youyu.api/lib/utils/version"
DIR_BUILD_CACHE = $(DIR_BUILD)/build_cache

# build use source
# 编译时会用到的源信息
BUILD_TIME = `date +%F`
SOURCES := $(wildcard ./app/*/cmd)
RPC_SOURCES := $(wildcard $(DIR_APP_RPC)/*/cmd)

# build app info
# 需要写入的版本信息
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-s -w -X ${DIR_VERSION}.gitTag=${gitTag} -X ${DIR_VERSION}.buildDate=${buildDate} -X ${DIR_VERSION}.gitCommit=${gitCommit} -X ${DIR_VERSION}.gitTreeState=${gitTreeState} -X ${DIR_VERSION}.version=${PROJECT_VERSION} -X ${DIR_VERSION}.gitBranch=${gitBranch}"

# func
# 获取文件的父目录
define getFAD
	$(shell echo $(1) | awk '{len=split($$0,a,"/");print a[len-1] > "$(DIR_BUILD_CACHE)/tmp"}')
endef

.PHONY: linux-create
linux-create:
	# 创建编译缓存区
	mkdir -p $(DIR_BUILD_CACHE)
	# 拷贝文件
	cp -r $(DIR_LOG) $(DIR_BUILD)
	cp -r $(DIR_SCRIPT)/docker/* $(DIR_BUILD)
	cp -r $(DIR_SCRIPT)/mysql $(DIR_BUILD)/script
	cp -r $(DIR_CONFIG)/pro/ $(DIR_BUILD)/conf
	cp -r $(DIR_INFO) $(DIR_BUILD)

linux-clean:
	rm -rf $(DIR_BUILD)


build-linux: linux-clean
build-linux: linux-create
build-linux:export GOOS=linux
build-linux:export GOARCH=amd64
build-linux:export GO111MODULE=on
build-linux:export CGO_ENABLED=0
build-linux:
	# 注入版本信息并编译工具程序
	# 提取程序名
	#$(call getFAD, $(DIR_APP)/uts/cmd)
	#cat $(DIR_BUILD_CACHE)/tmp
	#tmp=$$(cat $(DIR_BUILD_CACHE)/tmp) ; cd $(DIR_APP)/uts/cmd && $(GO_CMD) build -ldflags ${ldflags} -o $$tmp ; mv $$tmp $(PROJECT_ABS_PATH)/$(DIR_BUILD)
	# 编译主服务器程序和工具程序
	@for str in $(SOURCES);\
	do \
	echo $$str | awk '{len=split($$0,a,"/");print a[len-1] > "./project/build_cache/tmp"}' ;\
	tmp=$$(cat ./project/build_cache/tmp) ;\
	echo "$$tmp......编译中";\
	cd $$str && $(GO_CMD) build -x -o $$tmp -ldflags ${ldflags};\
	pwd ;\
	mv $$tmp $(PROJECT_ABS_PATH)/$(DIR_BUILD) ;\
	cd $(PROJECT_ABS_PATH) ;\
	pwd ;\
	done
	# 循环遍历统一编译同一参数的Rpc程序
	@for str in $(RPC_SOURCES);\
	do \
	echo $$str | awk '{len=split($$0,a,"/");print a[len-1] > "./project/build_cache/tmp"}' ;\
	tmp=$$(cat ./project/build_cache/tmp) ;\
	echo "$$tmp......编译中";\
	cd $$str && $(GO_CMD) build -x -o $$tmp -ldflags ${ldflags};\
	pwd ;\
	mv $$tmp $(PROJECT_ABS_PATH)/$(DIR_BUILD) ;\
	cd $(PROJECT_ABS_PATH) ;\
	pwd ;\
	done
build-windows:
	for str in $(RPC_SOURCES);\
	do \
	echo $$str | awk '{len=split($$0,a,"/");print a[len-1] > "./project/build_cache/tmp"}' ;\
	tmp=$$(cat ./project/build_cache/tmp) ;\
	echo $$tmp;\
	cd $$str && $(GO_CMD) build -x -o $$tmp -ldflags ${ldflags} ;\
	mv $$tmp $(PROJECT_ABS_PATH)/$(DIR_BUILD) ;\
	cd $(PROJECT_ABS_PATH) ;\
	done