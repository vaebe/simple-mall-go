# Golang版本；Alpine镜像的体积较小。
FROM golang:1.20.6-alpine3.17 as builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace

# 拷贝源代码到镜像中
COPY . .
COPY config.yaml .

# 打包
RUN go build -buildmode=pie -o /workspace/simple-mall-go

# 运行时镜像。
# Alpine兼顾了镜像大小和运维性。
FROM alpine:3.14

# 先创建文件目录在创建文件
RUN mkdir -p /app/tmp
RUN touch /app/tmp/logs /app/tmp/error

# 复制构建产物。
COPY --from=builder /workspace/simple-mall-go /app/simple-mall-go
COPY --from=builder /workspace/config.yaml /app/config.yaml

# 设置工作目录为/app
WORKDIR /app

EXPOSE 51015

# 指定默认的启动命令。
CMD ["./simple-mall-go"]