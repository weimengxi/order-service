# API Swagger Skill - Go/Gin 项目注释规范

## 使用说明

### 方式1：AI 自动生成

直接告诉 AI 助手要为哪个函数生成 Swagger 注释：

```
给 GetUser 函数生成 Swagger 注释
```

或者指定要遵循的规范：

```
按照 Swagger Skill 规范，为这个接口补充文档注释
```

### 方式2：批量补充注释

```
检查 admin.go 中缺少 Swagger 注释的接口，按 Skill 规范补充
```

### 方式3：描述接口让 AI 生成

```
按 Swagger Skill 给 ExportData 加注释，这是导出租户数据的接口，
支持 tenant_id 和 format 参数
```

### 方式4：手动复制模板

参考下方模板，复制并修改对应字段。

---

## 快速参考

```go
// FunctionName 函数描述
// @Summary 简短摘要（一行）
// @Description 详细描述
// @Tags 分组标签
// @Accept json
// @Produce json
// @Param name path/query/body type required "描述"
// @Success 200 {object} ResponseType
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/path [method]
func (h *Handler) FunctionName(c *gin.Context) {}
```

---

## 1. 基础注释模板

### 1.1 POST 请求（创建资源）
```go
// CreateUser 创建用户
// @Summary 创建用户
// @Description 管理员创建新用户，需要提供用户名和邮箱
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body CreateUserRequest true "创建用户请求"
// @Success 200 {object} domain.User "创建成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未授权"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /api/v1/users [post]
```

### 1.2 GET 请求（获取单个资源）
```go
// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据用户ID获取用户详细信息
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "用户ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} ErrorResponse "用户不存在"
// @Router /api/v1/users/{user_id} [get]
```

### 1.3 GET 请求（分页列表）
```go
// ListUsers 列出用户
// @Summary 列出用户
// @Description 分页获取用户列表，支持关键词搜索和状态筛选
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param keyword query string false "关键词搜索"
// @Param status query string false "状态筛选" Enums(active, inactive)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} ListUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users [get]
```

### 1.4 PATCH/PUT 请求（更新资源）
```go
// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 更新用户的部分信息，只传递需要更新的字段
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "用户ID"
// @Param request body UpdateUserRequest true "更新请求"
// @Success 200 {object} domain.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/users/{user_id} [patch]
```

### 1.5 DELETE 请求（删除资源）
```go
// DeleteUser 删除用户
// @Summary 删除用户
// @Description 软删除用户，用户数据不会立即清除
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "用户ID"
// @Success 204 "删除成功"
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/users/{user_id} [delete]
```

---

## 2. 参数类型详解

### @Param 语法
```
@Param 参数名 位置 类型 是否必填 "描述"
```

### 位置类型
| 位置 | 说明 | 示例 |
|------|------|------|
| `path` | URL路径参数 | `/users/{id}` |
| `query` | URL查询参数 | `?page=1` |
| `body` | 请求体 | JSON body |
| `header` | 请求头 | `X-Token` |
| `formData` | 表单数据 | 文件上传 |

### 数据类型
| 类型 | 说明 |
|------|------|
| `string` | 字符串 |
| `int` / `integer` | 整数 |
| `bool` / `boolean` | 布尔值 |
| `number` | 浮点数 |
| `file` | 文件（formData） |
| `StructName` | 引用结构体 |

### 高级参数示例
```go
// 枚举值
// @Param status query string false "状态" Enums(pending, active, disabled)

// 默认值
// @Param page query int false "页码" default(1)

// 最小/最大值
// @Param limit query int false "数量" minimum(1) maximum(100)

// 数组参数
// @Param ids query []string false "ID列表"
```

---

## 3. 请求/响应结构体

### 3.1 定义请求结构体
```go
// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
    Name     string `json:"name" binding:"required" example:"张三"`
    Email    string `json:"email" binding:"required,email" example:"zhangsan@example.com"`
    Age      int    `json:"age" example:"25"`
    Role     string `json:"role" enums:"admin,user,guest"`
}
```

### 3.2 定义响应结构体
```go
// ListUsersResponse 用户列表响应
type ListUsersResponse struct {
    Total int64         `json:"total" example:"100"`
    Items []*domain.User `json:"items"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
    Error   string `json:"error" example:"bad_request"`
    Message string `json:"message" example:"Invalid request parameters"`
}
```

### 3.3 结构体标签
| 标签 | 说明 | 示例 |
|------|------|------|
| `example` | 示例值 | `example:"hello"` |
| `enums` | 枚举值 | `enums:"a,b,c"` |
| `default` | 默认值 | `default:"active"` |
| `format` | 格式 | `format:"date-time"` |
| `minimum` | 最小值 | `minimum:"1"` |
| `maximum` | 最大值 | `maximum:"100"` |

---

## 4. 安全认证

### 4.1 在 main.go 定义安全方案
```go
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token authentication
```

### 4.2 在接口中使用
```go
// @Security ApiKeyAuth
// @Security BearerAuth
```

---

## 5. 分组与标签

### 5.1 定义标签描述（main.go）
```go
// @tag.name Admin
// @tag.description 管理员接口

// @tag.name User
// @tag.description 用户相关接口

// @tag.name Public
// @tag.description 公开接口，无需认证
```

### 5.2 接口分组规范
| 标签 | 用途 |
|------|------|
| `Admin` | 管理后台接口 |
| `Tenant` | 租户接口 |
| `Agent` | 客户端/Agent 接口 |
| `Internal` | 内部服务调用接口 |
| `Public` | 公开接口 |

---

## 6. 生成命令

```bash
# 安装 swag
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init --dir . --generalInfo cmd/server/main.go --output ./docs/swagger

# 或使用 Makefile
make docs
```

---

## 7. 检查清单

编写 API 注释时，确保包含：

- [ ] `@Summary` - 简短描述（显示在接口列表）
- [ ] `@Description` - 详细说明（可选但推荐）
- [ ] `@Tags` - 分组标签
- [ ] `@Accept` / `@Produce` - 请求/响应格式
- [ ] `@Param` - 所有参数（path/query/body）
- [ ] `@Success` - 成功响应
- [ ] `@Failure` - 错误响应（至少 400、500）
- [ ] `@Security` - 认证方式（如需要）
- [ ] `@Router` - 路由路径和方法

---

## 8. 常见问题

### Q: 如何处理嵌套对象？
```go
type Request struct {
    User struct {
        Name string `json:"name"`
    } `json:"user"`
}
// 直接引用即可，swag 会自动解析嵌套结构
```

### Q: 如何文档化文件上传？
```go
// @Accept multipart/form-data
// @Param file formData file true "上传文件"
```

### Q: 如何添加请求示例？
在结构体字段使用 `example` 标签

### Q: 路径参数用 `{id}` 还是 `:id`？
Swagger 文档中用 `{id}`，Gin 路由用 `:id`
