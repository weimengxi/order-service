# User Service

> 用户管理微服务 - 提供用户账户相关的 API 服务

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![API Version](https://img.shields.io/badge/API-v3.0.0-blue)](./docs/swagger.json)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 目录

- [功能特性](#功能特性)
- [技术栈](#技术栈)
- [快速开始](#快速开始)
- [API 文档](#api-文档)
- [项目结构](#项目结构)
- [配置说明](#配置说明)
- [开发指南](#开发指南)
- [部署指南](#部署指南)
- [版本管理](#版本管理)
- [贡献指南](#贡献指南)

---

## 🚀 功能特性

- ✅ 用户注册、登录、注销
- ✅ 用户信息 CRUD 操作
- ✅ JWT 认证和授权
- ✅ 密码加密存储
- ✅ 用户权限管理
- ✅ RESTful API 设计
- ✅ 自动生成 Swagger 文档

---

## 🛠 技术栈

- **语言**: Go 1.21+
- **框架**: Gin Web Framework
- **数据库**: PostgreSQL 14+
- **缓存**: Redis 7+
- **文档**: Swaggo (Swagger/OpenAPI)
- **日志**: Zap
- **配置**: Viper
- **容器**: Docker

---

## ⚡ 快速开始

### 前置要求

**本地开发：**
- Go 1.21 或更高版本
- PostgreSQL 14+ （可选，可使用 Docker）
- Redis 7+ （可选，可使用 Docker）
- Make （可选，用于快捷命令）

**Docker 开发：**
- Docker 20.10+
- Docker Compose 2.0+

### 方式一：本地直接运行（推荐开发调试）

```bash
# 1. 克隆仓库
git clone https://github.com/your-org/user-service.git
cd user-service

# 2. 下载依赖
go mod download

# 3. 安装 Swagger 文档生成工具
go install github.com/swaggo/swag/cmd/swag@latest

# 4. 启动依赖服务（使用 Docker）
docker-compose up -d postgres redis

# 5. 配置环境
cp configs/config.local.yaml configs/config.yaml

# 6. 运行数据库迁移
make migrate-up

# 7. 生成 API 文档
make docs

# 8. 启动服务
make run-local
# 或者
CONFIG_FILE=configs/config.local.yaml go run cmd/main.go
```

服务启动后访问：
- 应用服务: http://localhost:8001
- API 文档: http://localhost:8001/swagger/index.html
- 健康检查: http://localhost:8001/health

### 方式二：Docker 开发环境（推荐团队协作）

```bash
# 1. 克隆仓库
git clone https://github.com/your-org/user-service.git
cd user-service

# 2. 启动完整开发环境（包含所有依赖）
docker-compose -f docker-compose.dev.yml up -d

# 3. 查看服务状态
docker-compose -f docker-compose.dev.yml ps

# 4. 查看日志（支持热重载）
docker-compose -f docker-compose.dev.yml logs -f user-service

# 5. 代码修改会自动重新编译和重启服务
```

服务启动后访问：
- 应用服务: http://localhost:8001
- API 文档: http://localhost:8001/swagger/index.html
- 文档聚合中心: http://localhost:9000

```bash
# 停止服务
docker-compose -f docker-compose.dev.yml down

# 清理数据（谨慎使用）
docker-compose -f docker-compose.dev.yml down -v
```

### 方式三：远程开发/测试环境

```bash
# 远程开发环境（自动部署）
# 推送到 main 分支后自动部署
git push origin main

# 访问远程开发环境
open https://dev-docs.company.com/docs/user-service

# 测试环境（release 分支）
git checkout release/v3
git push origin release/v3

# 访问测试环境
open https://test-docs.company.com/docs/user-service
```

---

## 📚 API 文档

### 文档访问地址

**本地开发环境：**
- **直接运行**: http://localhost:8001/swagger/index.html
- **Docker 运行**: http://localhost:8001/swagger/index.html
- **文档聚合**: http://localhost:9000/docs/user-service
- **OpenAPI JSON**: http://localhost:8001/docs/swagger.json
- **OpenAPI YAML**: http://localhost:8001/docs/swagger.yaml

**远程环境：**
- **开发环境**: https://dev-docs.company.com/docs/user-service
- **测试环境**: https://test-docs.company.com/docs/user-service
- **生产环境**: https://api-docs.company.com/docs/user-service
- **私有云**: 请联系运维团队获取访问地址

### 本地生成文档

```bash
# 生成 Swagger 文档
make docs

# 或者直接使用 swag
swag init -g cmd/main.go -o docs

# 验证文档格式
make docs-validate
```

### 文档说明

本服务使用 [Swaggo](https://github.com/swaggo/swag) 自动生成 API 文档：
- 文档与代码同步更新
- 支持在线测试 API (Try it out)
- 提供完整的请求/响应示例
- 包含认证和参数验证说明

### API 概览

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 用户注册 | POST | /api/v3/users/register | 创建新用户 |
| 用户登录 | POST | /api/v3/users/login | 用户认证 |
| 获取用户信息 | GET | /api/v3/users/:id | 获取用户详情 |
| 更新用户信息 | PATCH | /api/v3/users/:id | 更新用户资料 |
| 删除用户 | DELETE | /api/v3/users/:id | 删除用户账户 |
| 用户列表 | GET | /api/v3/users | 分页查询用户列表 |

完整 API 文档请访问 Swagger UI。

---

## 📁 项目结构

```
user-service/
├── cmd/
│   └── main.go                 # 应用入口
├── internal/                   # 私有应用代码
│   ├── handler/                # HTTP 处理器
│   │   └── user.go
│   ├── service/                # 业务逻辑层
│   │   └── user.go
│   ├── repository/             # 数据访问层
│   │   └── user.go
│   ├── model/                  # 数据模型
│   │   └── user.go
│   ├── middleware/             # 中间件
│   │   ├── auth.go
│   │   └── cors.go
│   └── router/                 # 路由配置
│       └── router.go
├── pkg/                        # 可复用的公共库
│   ├── logger/                 # 日志封装
│   ├── database/               # 数据库连接
│   ├── redis/                  # Redis 客户端
│   └── utils/                  # 工具函数
├── configs/                    # 配置文件
│   ├── config.local.yaml       # 本地开发配置
│   ├── config.dev.yaml         # 远程开发环境配置
│   ├── config.test.yaml        # 测试环境配置
│   ├── config.prod.yaml        # 生产环境配置
│   └── config.example.yaml     # 配置示例模板
├── docs/                       #
│   └── swagger/                # Swagger 文档 (自动生成)
│       ├── docs.go                 # (不提交到 Git)
│       ├──swagger.json            # (提交到 Git)
│       └── swagger.yaml            # (提交到 Git)
├── scripts/                    # 脚本文件
│   ├── migrations/             # 数据库迁移
│   │   ├── 000001_init.up.sql
│   │   └── 000001_init.down.sql
│   └── seed.sql                # 测试数据
├── deployments/                # 部署配置
│   ├── docker-compose.yml      # 基础配置
│   ├── docker-compose.dev.yml  # 开发环境
│   ├── docker-compose.test.yml # 测试环境
│   └── kubernetes/             # K8s 部署清单
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── configmap.yaml
│       └── ingress.yaml
├── tests/                      # 测试文件
│   ├── integration/
│   │   └── user_test.go
│   └── unit/
│       └── handler_test.go
├── .github/
│   └── workflows/
│       ├── ci.yml              # 持续集成
│       └── deploy.yml          # 自动部署
├── .gitignore
├── .swagger-version            # 当前 API 版本标识
├── Dockerfile                  # 生产环境镜像
├── Dockerfile.dev              # 开发环境镜像（支持热重载）
├── Makefile                    # 快捷命令
├── go.mod
├── go.sum
├── README.md
└── CHANGELOG.md                # 变更日志
```

---

## ⚙️ 配置说明

### 配置文件

配置文件位于 `configs/` 目录，支持多环境配置：

#### 本地开发配置 (config.local.yaml)

```yaml
server:
  port: 8001
  base_path: "/api/v3"
  environment: "local"
  read_timeout: 10s
  write_timeout: 10s

swagger:
  enabled: true
  title: "User Service API"
  version: "3.0.0"
  description: "用户管理服务 - 本地开发"
  host: "localhost:8001"
  schemes: ["http"]

database:
  driver: "postgres"
  host: "localhost"              # 本地数据库
  port: 5432
  username: "postgres"
  password: "postgres"
  database: "users_dev"
  max_open_conns: 10
  max_idle_conns: 5

redis:
  host: "localhost"              # 本地 Redis
  port: 6379
  password: ""
  db: 0
  pool_size: 10

jwt:
  secret: "local-dev-secret-key"
  expire_hours: 24

logging:
  level: "debug"                 # 本地开发详细日志
  format: "console"              # 控制台友好格式
  output: "stdout"
```

#### 远程开发环境配置 (config.dev.yaml)

```yaml
server:
  port: 8001
  base_path: "/api/v3"
  environment: "development"

swagger:
  enabled: true
  title: "User Service API"
  version: "3.0.0"
  description: "用户管理服务 - 开发环境"
  host: "dev.api.company.com"
  schemes: ["https"]

database:
  host: "postgres"               # Docker Compose 服务名
  port: 5432
  username: "postgres"
  password: "${DB_PASSWORD}"     # 从环境变量读取
  database: "users"
  max_open_conns: 50
  max_idle_conns: 10

redis:
  host: "redis"                  # Docker Compose 服务名
  port: 6379
  password: "${REDIS_PASSWORD}"
  db: 0

jwt:
  secret: "${JWT_SECRET}"
  expire_hours: 24

logging:
  level: "info"
  format: "json"
  output: "stdout"
```

#### 生产环境配置 (config.prod.yaml)

```yaml
server:
  port: 8001
  base_path: "/api/v3"
  environment: "production"

swagger:
  enabled: false                 # 生产环境关闭或限制访问
  title: "User Service API"
  version: "3.0.0"

database:
  host: "postgres.internal.svc.cluster.local"  # K8s 内部 DNS
  port: 5432
  username: "${DB_USERNAME}"
  password: "${DB_PASSWORD}"
  database: "users"
  max_open_conns: 100
  max_idle_conns: 20

redis:
  host: "redis.internal.svc.cluster.local"
  port: 6379
  password: "${REDIS_PASSWORD}"
  db: 0

jwt:
  secret: "${JWT_SECRET}"
  expire_hours: 24

logging:
  level: "warn"                  # 生产环境只记录警告和错误
  format: "json"
  output: "stdout"
```

### 环境变量

支持通过环境变量覆盖配置：

```bash
# 本地开发
export CONFIG_FILE=configs/config.local.yaml
export SERVER_PORT=8001
export DB_HOST=localhost
export DB_PASSWORD=postgres
export REDIS_HOST=localhost
export JWT_SECRET=local-secret

# Docker 开发（在 docker-compose.dev.yml 中配置）
CONFIG_FILE=/app/configs/config.dev.yaml
DB_HOST=postgres
REDIS_HOST=redis

# 生产环境（在 K8s ConfigMap/Secret 中配置）
CONFIG_FILE=/config/config.prod.yaml
DB_HOST=postgres.internal.svc.cluster.local
DB_PASSWORD=<from-secret>
REDIS_HOST=redis.internal.svc.cluster.local
JWT_SECRET=<from-secret>
```

### 配置优先级

```
环境变量 > 配置文件 > 默认值
```

---

## 💻 开发指南

### 开发环境设置

```bash
# 1. 安装开发工具
make install-tools

# 2. 启动依赖服务（数据库、Redis）
docker-compose up -d postgres redis

# 3. 运行数据库迁移
make migrate-up

# 4. 生成 Swagger 文档
make docs

# 5. 启动开发服务器（带热重载）
make dev
```

### 添加新的 API 接口

#### 1. 定义数据模型

```go
// internal/model/user.go
package model

type User struct {
    ID        int64     `json:"id" example:"1"`
    Username  string    `json:"username" example:"john_doe"`
    Email     string    `json:"email" example:"john@example.com"`
    CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
    Username string `json:"username" binding:"required,min=3,max=32" example:"john_doe"`
    Email    string `json:"email" binding:"required,email" example:"john@example.com"`
    Password string `json:"password" binding:"required,min=8" example:"SecurePass123!"`
}
```

#### 2. 实现业务逻辑

```go
// internal/service/user.go
package service

type UserService struct {
    repo repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
    // 业务逻辑实现
    return s.repo.Create(ctx, user)
}
```

#### 3. 添加 HTTP 处理器（含 Swagger 注解）

```go
// internal/handler/user.go
package handler

// CreateUser 创建用户
// @Summary      创建新用户
// @Description  注册一个新的用户账户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.CreateUserRequest  true  "用户信息"
// @Success      201   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req model.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.CreateUser(c.Request.Context(), &req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, user)
}
```

#### 4. 注册路由

```go
// internal/router/router.go
package router

func SetupRouter(handler *handler.UserHandler) *gin.Engine {
    r := gin.Default()

    v3 := r.Group("/api/v3")
    {
        users := v3.Group("/users")
        {
            users.POST("", handler.CreateUser)
            users.GET("/:id", handler.GetUser)
            // ... 其他路由
        }
    }

    return r
}
```

#### 5. 生成并测试文档

```bash
# 生成文档
make docs

# 启动服务
make run

# 访问 Swagger UI 测试
open http://localhost:8001/swagger/index.html
```

### 编码规范

- 遵循 [Effective Go](https://golang.org/doc/effective_go.html)
- 使用 `golangci-lint` 进行代码检查
- 所有公开 API 必须添加 Swagger 注解
- 单元测试覆盖率不低于 80%
- Git commit 遵循 [Conventional Commits](https://www.conventionalcommits.org/)

### 运行测试

```bash
# 运行所有测试
make test

# 运行单元测试
make test-unit

# 运行集成测试
make test-integration

# 查看测试覆盖率
make test-coverage

# 生成覆盖率报告
make coverage-html
```

### Makefile 命令

```bash
# 帮助信息
make help              # 查看所有可用命令

# 开发相关
make install-tools     # 安装开发工具（swag, golangci-lint 等）
make deps              # 下载依赖
make docs              # 生成 API 文档
make run-local         # 本地运行（使用 config.local.yaml）
make run-dev           # 开发模式运行（使用 config.dev.yaml）
make dev               # 开发模式运行（带热重载，需要 air）
make build             # 编译二进制文件

# 测试相关
make test              # 运行所有测试
make test-unit         # 运行单元测试
make test-integration  # 运行集成测试
make test-coverage     # 查看测试覆盖率
make coverage-html     # 生成 HTML 覆盖率报告

# 代码质量
make lint              # 代码检查
make fmt               # 格式化代码
make vet               # Go vet 检查

# 数据库相关
make migrate-up        # 执行数据库迁移
make migrate-down      # 回滚数据库迁移
make migrate-create    # 创建新的迁移文件
make seed              # 填充测试数据

# Docker 相关
make docker-build      # 构建 Docker 镜像
make docker-run        # 运行 Docker 容器
make docker-stop       # 停止 Docker 容器
make docker-logs       # 查看 Docker 日志

# Docker Compose
make compose-up        # 启动开发环境 (docker-compose.dev.yml)
make compose-down      # 停止开发环境
make compose-logs      # 查看日志
make compose-ps        # 查看服务状态
make compose-restart   # 重启服务

# 文档相关
make docs-validate     # 验证 Swagger 文档
make docs-open         # 在浏览器中打开文档

# 清理
make clean             # 清理构建文件
make clean-all         # 清理所有生成文件（包括文档）
```

### Makefile 示例

```makefile
# Makefile
.PHONY: help run-local run-dev docs test docker-build compose-up

# 变量定义
APP_NAME := user-service
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)

help: ## 显示帮助信息
  @echo "Available commands:"
  @grep -E '^[a-zA-Z_-]+:.*?## .*$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $1, $2}'

install-tools: ## 安装开发工具
  @echo "Installing development tools..."
  go install github.com/swaggo/swag/cmd/swag@latest
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  go install github.com/cosmtrek/air@latest
  @echo "✅ Tools installed"

deps: ## 下载依赖
  @echo "Downloading dependencies..."
  go mod download
  go mod verify
  @echo "✅ Dependencies downloaded"

docs: ## 生成 API 文档
  @echo "Generating Swagger documentation..."
  swag init -g cmd/main.go -o docs --parseDependency --parseInternal
  @echo "✅ Documentation generated"

docs-validate: ## 验证文档
  @echo "Validating Swagger documentation..."
  @command -v swagger-cli >/dev/null 2>&1 || (echo "Installing swagger-cli..." && npm install -g @apidevtools/swagger-cli)
  swagger-cli validate docs/swagger.json
  @echo "✅ Documentation is valid"

docs-open: docs ## 打开文档
  @echo "Opening Swagger UI..."
  @(open http://localhost:8001/swagger/index.html || xdg-open http://localhost:8001/swagger/index.html || start http://localhost:8001/swagger/index.html) 2>/dev/null

run-local: docs ## 本地运行
  @echo "Starting service (local config)..."
  CONFIG_FILE=configs/config.local.yaml go run cmd/main.go

run-dev: docs ## 开发环境运行
  @echo "Starting service (dev config)..."
  CONFIG_FILE=configs/config.dev.yaml go run cmd/main.go

dev: ## 开发模式（热重载）
  @echo "Starting service with hot reload..."
  air -c .air.toml

build: ## 编译
  @echo "Building $(APP_NAME)..."
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) cmd/main.go
  @echo "✅ Build complete: bin/$(APP_NAME)"

test: ## 运行测试
  @echo "Running tests..."
  go test -v -race -coverprofile=coverage.out ./...
  @echo "✅ Tests passed"

test-coverage: test ## 测试覆盖率
  @echo "Test coverage:"
  go tool cover -func=coverage.out | tail -1

coverage-html: test ## 生成 HTML 覆盖率报告
  go tool cover -html=coverage.out -o coverage.html
  @echo "✅ Coverage report: coverage.html"

lint: ## 代码检查
  @echo "Running linter..."
  golangci-lint run ./...
  @echo "✅ Lint passed"

fmt: ## 格式化代码
  @echo "Formatting code..."
  go fmt ./...
  gofmt -s -w .
  @echo "✅ Code formatted"

migrate-up: ## 执行迁移
  @echo "Running migrations..."
  migrate -path scripts/migrations -database "postgres://postgres:postgres@localhost:5432/users_dev?sslmode=disable" up
  @echo "✅ Migrations applied"

migrate-down: ## 回滚迁移
  @echo "Rolling back migrations..."
  migrate -path scripts/migrations -database "postgres://postgres:postgres@localhost:5432/users_dev?sslmode=disable" down 1

docker-build: ## 构建 Docker 镜像
  @echo "Building Docker image..."
  docker build -t $(APP_NAME):$(VERSION) -t $(APP_NAME):latest .
  @echo "✅ Image built: $(APP_NAME):$(VERSION)"

docker-build-dev: ## 构建开发镜像
  @echo "Building development Docker image..."
  docker build -f Dockerfile.dev -t $(APP_NAME):dev .
  @echo "✅ Dev image built"

docker-run: ## 运行 Docker 容器
  docker run -d --name $(APP_NAME) -p 8001:8001 \
    -e CONFIG_FILE=/app/configs/config.dev.yaml \
    $(APP_NAME):latest

compose-up: ## 启动 Docker Compose
  @echo "Starting Docker Compose environment..."
  docker-compose -f docker-compose.dev.yml up -d
  @echo "✅ Environment started"
  @echo "Access documentation: http://localhost:9000"

compose-down: ## 停止 Docker Compose
  @echo "Stopping Docker Compose environment..."
  docker-compose -f docker-compose.dev.yml down
  @echo "✅ Environment stopped"

compose-logs: ## 查看日志
  docker-compose -f docker-compose.dev.yml logs -f

compose-ps: ## 查看服务状态
  docker-compose -f docker-compose.dev.yml ps

compose-restart: ## 重启服务
  docker-compose -f docker-compose.dev.yml restart $(APP_NAME)

clean: ## 清理构建文件
  @echo "Cleaning..."
  rm -rf bin/
  rm -f coverage.out coverage.html
  @echo "✅ Cleaned"

clean-all: clean ## 清理所有
  @echo "Cleaning all generated files..."
  rm -rf docs/docs.go
  @echo "✅ All cleaned"
```

---

## 🚢 部署指南

### 本地 Docker 部署

```bash
# 1. 构建镜像
make docker-build

# 2. 运行容器
docker run -d \
  --name user-service \
  -p 8001:8001 \
  -e CONFIG_FILE=/app/configs/config.dev.yaml \
  -e DB_HOST=host.docker.internal \
  -e REDIS_HOST=host.docker.internal \
  user-service:latest

# 3. 查看日志
docker logs -f user-service

# 4. 停止服务
docker stop user-service
docker rm user-service
```

### Docker Compose 部署

#### 开发环境

```bash
# 启动所有服务（包括依赖）
docker-compose -f docker-compose.dev.yml up -d

# 查看服务状态
docker-compose -f docker-compose.dev.yml ps

# 查看日志
docker-compose -f docker-compose.dev.yml logs -f user-service

# 重启某个服务
docker-compose -f docker-compose.dev.yml restart user-service

# 停止所有服务
docker-compose -f docker-compose.dev.yml down

# 停止并删除数据卷（谨慎使用）
docker-compose -f docker-compose.dev.yml down -v
```

#### 测试环境

```bash
# 使用测试环境配置
export VERSION=v3.2.1
export REGISTRY=ghcr.io/your-org

docker-compose -f docker-compose.test.yml up -d
```

### Kubernetes 部署

#### 前置准备

```bash
# 1. 构建并推送镜像
docker build -t ghcr.io/your-org/user-service:v3.2.1 .
docker push ghcr.io/your-org/user-service:v3.2.1

# 2. 创建命名空间
kubectl create namespace production

# 3. 创建 Secret（存储敏感信息）
kubectl create secret generic user-service-secrets \
  --from-literal=db-password=<password> \
  --from-literal=redis-password=<password> \
  --from-literal=jwt-secret=<secret> \
  -n production
```

#### 部署服务

```bash
# 应用所有 K8s 配置
kubectl apply -f deployments/kubernetes/

# 或者逐个应用
kubectl apply -f deployments/kubernetes/configmap.yaml
kubectl apply -f deployments/kubernetes/deployment.yaml
kubectl apply -f deployments/kubernetes/service.yaml
kubectl apply -f deployments/kubernetes/ingress.yaml

# 查看部署状态
kubectl get pods -n production -l app=user-service
kubectl get svc -n production user-service
kubectl get ingress -n production

# 查看日志
kubectl logs -f deployment/user-service -n production

# 查看服务详情
kubectl describe deployment user-service -n production
```

#### 更新部署

```bash
# 方式1: 更新镜像
kubectl set image deployment/user-service \
  user-service=ghcr.io/your-org/user-service:v3.2.2 \
  -n production

# 方式2: 重新应用配置
kubectl apply -f deployments/kubernetes/deployment.yaml

# 查看滚动更新状态
kubectl rollout status deployment/user-service -n production

# 查看更新历史
kubectl rollout history deployment/user-service -n production

# 回滚到上一个版本
kubectl rollout undo deployment/user-service -n production
```

#### 扩缩容

```bash
# 手动扩容
kubectl scale deployment user-service --replicas=5 -n production

# 查看副本数
kubectl get deployment user-service -n production

# 配置自动扩缩容（HPA）
kubectl autoscale deployment user-service \
  --cpu-percent=70 \
  --min=3 \
  --max=10 \
  -n production
```

#### 故障排查

```bash
# 查看 Pod 状态
kubectl get pods -n production -l app=user-service

# 查看 Pod 日志
kubectl logs <pod-name> -n production

# 查看最近的日志
kubectl logs <pod-name> -n production --tail=100

# 进入容器
kubectl exec -it <pod-name> -n production -- /bin/sh

# 查看事件
kubectl get events -n production --sort-by='.lastTimestamp'

# 查看资源使用
kubectl top pods -n production -l app=user-service
```

### 配置管理

#### 本地环境
```bash
CONFIG_FILE=configs/config.local.yaml ./bin/user-service
```

#### Docker 环境
```bash
docker run -e CONFIG_FILE=/app/configs/config.dev.yaml user-service
```

#### Kubernetes 环境
通过 ConfigMap 和 Secret 管理配置：

```yaml
# ConfigMap 用于非敏感配置
apiVersion: v1
kind: ConfigMap
metadata:
  name: user-service-config
data:
  config.prod.yaml: |
    server:
      port: 8001
      base_path: "/api/v3"
    # ...

# Secret 用于敏感信息
apiVersion: v1
kind: Secret
metadata:
  name: user-service-secrets
type: Opaque
data:
  db-password: <base64-encoded>
  jwt-secret: <base64-encoded>
```

### 健康检查

服务提供以下健康检查端点：

```bash
# 存活探针（Liveness Probe）
curl http://localhost:8001/health

# 就绪探针（Readiness Probe）
curl http://localhost:8001/ready

# 响应示例
{
  "status": "healthy",
  "version": "v3.2.1",
  "timestamp": "2026-01-07T10:00:00Z",
  "checks": {
    "database": "ok",
    "redis": "ok"
  }
}
```

### 监控和日志

```bash
# 查看应用日志（JSON 格式）
docker logs user-service | jq .

# K8s 日志
kubectl logs -f deployment/user-service -n production

# 导出日志到文件
kubectl logs deployment/user-service -n production > user-service.log

# 查看 Prometheus 指标
curl http://localhost:8001/metrics
```

---

## 🏷️ 版本管理

### 当前版本

- **API 版本**: v3.0.0
- **服务版本**: v3.2.1
- **Git 分支**: `release/v3`

### 版本分支

```
main                    # 最新稳定代码
├── develop             # 开发分支
├── release/v3          # v3 稳定版（公有云生产环境）
├── release/v2          # v2 稳定版（部分私有云客户）
└── release/v1          # v1 维护版（老客户，仅 bugfix）
```

### 版本说明

| 版本 | 状态 | 部署环境 | 支持截止 | 说明 |
|------|------|---------|---------|------|
| v3.x | **Current** | 公有云、新客户 | - | 当前推荐版本 |
| v2.x | **Stable** | 部分私有云 | 2026-06-30 | 稳定支持 |
| v1.x | **Deprecated** | 老客户 | 2025-12-31 | 仅维护，建议升级 |

### API 变更日志

详见 [CHANGELOG.md](./CHANGELOG.md)

**v3.0.0 主要变更**
- 🎉 新增：用户导出功能
- ✨ 改进：优化查询性能
- 🔒 安全：增强密码策略
- 💥 破坏性变更：响应格式统一调整

**迁移指南**
- [从 v1 迁移到 v2](./docs/migration/v1-to-v2.md)
- [从 v2 迁移到 v3](./docs/migration/v2-to-v3.md)

---

## 🤝 贡献指南

欢迎贡献！请遵循以下流程：

### 开发流程

1. **Fork 仓库并克隆**
   ```bash
   git clone https://github.com/your-username/user-service.git
   ```

2. **切换到对应版本分支**
   ```bash
   git checkout release/v3
   ```

3. **创建功能分支**
   ```bash
   git checkout -b feature/add-export-api
   ```

4. **开发并提交**
   ```bash
   git add .
   git commit -m "feat: add user export API"
   ```

5. **推送并创建 Pull Request**
   ```bash
   git push origin feature/add-export-api
   ```

### Commit 规范

遵循 [Conventional Commits](https://www.conventionalcommits.org/)：

```
<type>(<scope>): <subject>

<body>

<footer>
```

**类型 (type)**
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建/工具链更新

**示例**
```
feat(user): add user export API

- Add export endpoint
- Support CSV and JSON formats
- Add pagination support

Closes #123
```

### Code Review 标准

- ✅ 代码遵循项目规范
- ✅ 包含必要的单元测试
- ✅ 更新了相关文档
- ✅ Swagger 注解完整
- ✅ 通过 CI 检查

---

## 📞 联系方式

- **负责团队**: User Service Team
- **技术负责人**: @tech-lead
- **邮箱**: user-team@company.com
- **Slack**: #user-service
- **Issue Tracker**: https://github.com/your-org/user-service/issues

---

## 📄 许可证

本项目采用 [MIT License](LICENSE)

---

## 🙏 致谢

感谢所有为本项目做出贡献的开发者！

---

**最后更新**: 2026-01-07
**维护者**: @your-name