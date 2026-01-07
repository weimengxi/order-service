# Order Service

订单管理微服务，提供订单创建、查询、状态管理等功能。

## 快速开始

### Docker Compose 部署（推荐）

```bash
# 启动服务（包含 PostgreSQL 和 Redis）
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### Docker Compose 开发模式（热重载）

```bash
# 启动开发环境（使用 Air 热重载）
docker-compose -f docker-compose.dev.yml up -d

# 查看日志
docker-compose -f docker-compose.dev.yml logs -f

# 修改代码后自动重新编译并重启
```

### 本地开发

```bash
# 安装依赖
go mod download

# 方式1: 使用 Air 热重载（推荐）
go install github.com/cosmtrek/air@latest
air  # 修改代码后自动重启

# 方式2: 普通运行
make docs      # 生成 Swagger 文档
make run-local # 运行服务

# 访问 Swagger UI
open http://localhost:8002/swagger/index.html
```

### Docker 运行

```bash
# 构建镜像
docker build -t order-service:latest .

# 运行容器
docker run -p 8002:8002 order-service:latest
```

## API 文档

- **本地开发**: http://localhost:8002/swagger/index.html
- **GitHub Pages**: https://YOUR_ORG.github.io/order-service/
- **OpenAPI JSON**: http://localhost:8002/docs/swagger/swagger.json

## 目录结构

```
order-service/
├── cmd/
│   └── main.go              # 应用入口
├── internal/
│   ├── config/              # 配置管理
│   ├── handler/             # HTTP 处理器
│   ├── model/               # 数据模型
│   └── router/              # 路由配置
├── docs/
│   └── swagger/             # Swagger 生成目录
├── configs/                 # 配置文件
├── .air.toml                # Air 热重载配置
├── Dockerfile               # 生产环境 Docker
├── Dockerfile.dev           # 开发环境 Docker
├── Makefile                 # 构建脚本
└── go.mod                   # Go 模块文件
```

## 配置说明

| 环境 | 配置文件 | 说明 |
|------|---------|------|
| 本地开发 | config.local.yaml | 直连本地数据库 |
| 远程开发 | config.dev.yaml | Docker Compose 环境 |
| 生产环境 | config.prod.yaml | Kubernetes 部署 |

## 环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| CONFIG_FILE | 配置文件路径 | configs/config.local.yaml |
| GIN_MODE | Gin 运行模式 | debug |
