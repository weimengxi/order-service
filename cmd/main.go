package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"order-service/internal/config"
	"order-service/internal/handler"
	"order-service/internal/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Order Service API
// @version         3.0.0
// @description     订单管理服务 - 提供订单创建、查询、状态管理等功能
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.company.com/support
// @contact.email  support@company.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /api/v3

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// 加载配置
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = "configs/config.local.yaml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置 Gin 模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 健康检查
	r.GET("/health", handler.HealthCheck)

	// API 路由
	api := r.Group(cfg.Server.BasePath)
	router.SetupRoutes(api)

	// Swagger 文档
	if cfg.Swagger.Enabled {
		// 配置 Swagger UI 使用自定义 JSON 路径
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.URL("/docs/swagger/swagger.json"),
		))
		r.Static("/docs/swagger", "./docs/swagger")
		log.Printf("Swagger UI available at http://%s/swagger/index.html", cfg.Swagger.Host)
	}

	// 启动服务
	addr := ":" + cfg.Server.Port
	log.Printf("Order Service starting on %s", addr)
	log.Printf("Environment: %s", cfg.Server.Environment)
	log.Printf("API Base Path: %s", cfg.Server.BasePath)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
