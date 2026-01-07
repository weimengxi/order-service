# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod go.sum* ./
RUN go mod download

# 安装 swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 复制源代码
COPY . .

# 生成 Swagger 文档
RUN swag init -g cmd/main.go -o docs --parseDependency --parseInternal

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Runtime stage
FROM alpine:3.18

WORKDIR /app

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata curl

# 设置时区
ENV TZ=Asia/Shanghai

# 复制二进制文件
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/configs ./configs

# 暴露端口
EXPOSE 8002

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8002/health || exit 1

# 运行应用
CMD ["./main"]
