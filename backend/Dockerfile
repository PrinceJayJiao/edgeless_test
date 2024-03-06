# 使用golang官方镜像作为基础镜像
FROM golang:1.22

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件并下载依赖
COPY go.mod ./
RUN go mod download

# 复制应用程序代码
COPY . .

# 构建应用程序
RUN go build -o app

# 设置容器启动命令
CMD ["./app"]