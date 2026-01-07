package router

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/handler"
)

// SetupRoutes 配置路由
func SetupRoutes(r *gin.RouterGroup) {
	// 订单相关路由
	orders := r.Group("/orders")
	{
		orders.GET("", handler.GetOrders)
		orders.GET("/statistics", handler.GetOrderStatistics)
		orders.GET("/:id", handler.GetOrder)
		orders.POST("", handler.CreateOrder)
		orders.PUT("/:id/status", handler.UpdateOrderStatus)
		orders.POST("/:id/cancel", handler.CancelOrder)
	}
}
