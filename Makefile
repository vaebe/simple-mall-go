# Go 编译器和工具
GO := go

# 项目名称
APP_NAME := mall

# 目标操作系统和架构
TARGET_OS := linux
TARGET_ARCH := amd64

# 编译输出目录
BUILD_DIR := ./build

# 编译参数
BUILD_FLAGS := -v

# 构建目标
all: build

# 编译构建
build:
	@echo "开始构建..."
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) $(GO) build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)

# 清理构建
clean:
	@echo "清理构建..."
	rm -rf $(BUILD_DIR)

# 默认目标为构建
.DEFAULT_GOAL := build
