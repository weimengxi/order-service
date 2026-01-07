package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/my-org/order-service/internal/model"
)

// 模拟订单数据存储
var orders = []model.Order{
	{
		ID:          "ORD-20240101-001",
		UserID:      1,
		Status:      "completed",
		TotalAmount: 7999.00,
		Currency:    "CNY",
		Items: []model.OrderItem{
			{ProductID: 100, ProductName: "iPhone 15 Pro", Quantity: 1, UnitPrice: 7999.00, TotalPrice: 7999.00},
		},
		ShippingAddress: model.Address{
			Name: "张三", Phone: "13800138000", Province: "广东省", City: "深圳市", District: "南山区", Street: "科技园路1号", ZipCode: "518000",
		},
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now(),
	},
	{
		ID:          "ORD-20240101-002",
		UserID:      2,
		Status:      "pending",
		TotalAmount: 2999.00,
		Currency:    "CNY",
		Items: []model.OrderItem{
			{ProductID: 101, ProductName: "AirPods Pro", Quantity: 1, UnitPrice: 1999.00, TotalPrice: 1999.00},
			{ProductID: 102, ProductName: "Apple Watch Band", Quantity: 2, UnitPrice: 500.00, TotalPrice: 1000.00},
		},
		ShippingAddress: model.Address{
			Name: "李四", Phone: "13900139000", Province: "北京市", City: "北京市", District: "海淀区", Street: "中关村大街1号", ZipCode: "100000",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
var orderCounter = 3

// HealthCheck 健康检查
// @Summary      健康检查
// @Description  检查服务是否正常运行
// @Tags         System
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "order-service",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// GetOrders 获取订单列表
// @Summary      获取订单列表
// @Description  分页获取订单列表，支持按用户ID和状态筛选
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        page     query     int     false  "页码"  default(1)
// @Param        size     query     int     false  "每页数量"  default(10)
// @Param        user_id  query     int     false  "用户ID"
// @Param        status   query     string  false  "订单状态"  Enums(pending, paid, shipped, completed, cancelled)
// @Success      200      {object}  model.Response{data=model.OrderListResponse}
// @Failure      400      {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /orders [get]
func GetOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	userIDStr := c.Query("user_id")
	status := c.Query("status")

	// 过滤订单
	var filteredOrders []model.Order
	for _, order := range orders {
		if userIDStr != "" {
			userID, _ := strconv.ParseInt(userIDStr, 10, 64)
			if order.UserID != userID {
				continue
			}
		}
		if status != "" && order.Status != status {
			continue
		}
		filteredOrders = append(filteredOrders, order)
	}

	// 分页
	start := (page - 1) * size
	end := start + size
	if start > len(filteredOrders) {
		start = len(filteredOrders)
	}
	if end > len(filteredOrders) {
		end = len(filteredOrders)
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data: model.OrderListResponse{
			Total:  int64(len(filteredOrders)),
			Page:   page,
			Size:   size,
			Orders: filteredOrders[start:end],
		},
	})
}

// GetOrder 获取订单详情
// @Summary      获取订单详情
// @Description  根据订单ID获取订单详细信息
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "订单ID"
// @Success      200  {object}  model.Response{data=model.Order}
// @Failure      404  {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /orders/{id} [get]
func GetOrder(c *gin.Context) {
	id := c.Param("id")

	for _, order := range orders {
		if order.ID == id {
			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "success",
				Data:    order,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "Order not found",
	})
}

// CreateOrder 创建订单
// @Summary      创建订单
// @Description  创建新订单
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order  body      model.CreateOrderRequest  true  "订单信息"
// @Success      201    {object}  model.Response{data=model.Order}
// @Failure      400    {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /orders [post]
func CreateOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	// 模拟创建订单
	now := time.Now()
	orderID := fmt.Sprintf("ORD-%s-%03d", now.Format("20060102"), orderCounter)
	orderCounter++

	var items []model.OrderItem
	var totalAmount float64
	for _, item := range req.Items {
		// 模拟获取商品信息
		unitPrice := 999.00 // 模拟价格
		orderItem := model.OrderItem{
			ProductID:   item.ProductID,
			ProductName: fmt.Sprintf("Product %d", item.ProductID),
			Quantity:    item.Quantity,
			UnitPrice:   unitPrice,
			TotalPrice:  unitPrice * float64(item.Quantity),
		}
		items = append(items, orderItem)
		totalAmount += orderItem.TotalPrice
	}

	newOrder := model.Order{
		ID:          orderID,
		UserID:      req.UserID,
		Status:      "pending",
		TotalAmount: totalAmount,
		Currency:    "CNY",
		Items:       items,
		ShippingAddress: model.Address{
			Name:     req.Address.Name,
			Phone:    req.Address.Phone,
			Province: req.Address.Province,
			City:     req.Address.City,
			District: req.Address.District,
			Street:   req.Address.Street,
			ZipCode:  req.Address.ZipCode,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
	orders = append(orders, newOrder)

	c.JSON(http.StatusCreated, model.Response{
		Code:    0,
		Message: "Order created successfully",
		Data:    newOrder,
	})
}

// UpdateOrderStatus 更新订单状态
// @Summary      更新订单状态
// @Description  根据订单ID更新订单状态
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id      path      string                        true  "订单ID"
// @Param        status  body      model.UpdateOrderStatusRequest  true  "订单状态"
// @Success      200     {object}  model.Response{data=model.Order}
// @Failure      400     {object}  model.ErrorResponse
// @Failure      404     {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /orders/{id}/status [put]
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	for i, order := range orders {
		if order.ID == id {
			orders[i].Status = req.Status
			orders[i].UpdatedAt = time.Now()

			if req.Status == "paid" {
				now := time.Now()
				orders[i].PaidAt = &now
			}

			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "Order status updated successfully",
				Data:    orders[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "Order not found",
	})
}

// CancelOrder 取消订单
// @Summary      取消订单
// @Description  根据订单ID取消订单
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "订单ID"
// @Success      200  {object}  model.Response{data=model.Order}
// @Failure      400  {object}  model.ErrorResponse
// @Failure      404  {object}  model.ErrorResponse
// @Security     BearerAuth
// @Router       /orders/{id}/cancel [post]
func CancelOrder(c *gin.Context) {
	id := c.Param("id")

	for i, order := range orders {
		if order.ID == id {
			if order.Status == "completed" || order.Status == "shipped" {
				c.JSON(http.StatusBadRequest, model.ErrorResponse{
					Code:    400,
					Message: "Cannot cancel order in current status",
				})
				return
			}

			orders[i].Status = "cancelled"
			orders[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, model.Response{
				Code:    0,
				Message: "Order cancelled successfully",
				Data:    orders[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, model.ErrorResponse{
		Code:    404,
		Message: "Order not found",
	})
}

// GetOrderStatistics 获取订单统计
// @Summary      获取订单统计
// @Description  获取订单统计信息
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Response{data=model.OrderStatistics}
// @Security     BearerAuth
// @Router       /orders/statistics [get]
func GetOrderStatistics(c *gin.Context) {
	var stats model.OrderStatistics
	stats.TotalOrders = int64(len(orders))

	for _, order := range orders {
		stats.TotalAmount += order.TotalAmount
		switch order.Status {
		case "pending":
			stats.PendingOrders++
		case "completed":
			stats.CompletedOrders++
		case "cancelled":
			stats.CancelledOrders++
		}
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    stats,
	})
}
