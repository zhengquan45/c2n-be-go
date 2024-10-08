# 基础镜像
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器
COPY . .

# 下载依赖并编译项目
RUN go mod tidy
RUN go build -o myapp .


# 设置执行命令
CMD ["/app/myapp"]
