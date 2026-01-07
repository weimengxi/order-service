package model

import "time"

// Order 订单模型
type Order struct {
	ID          string     `json:"id" example:"ORD-20240101-001"`
	UserID      int64      `json:"user_id" example:"1"`
	Status      string     `json:"status" example:"pending" enums:"pending,paid,shipped,completed,cancelled"`
	TotalAmount float64    `json:"total_amount" example:"199.99"`
	Currency    string     `json:"currency" example:"CNY"`
	Items       []OrderItem `json:"items"`
	ShippingAddress Address `json:"shipping_address"`
	CreatedAt   time.Time  `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time  `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	PaidAt      *time.Time `json:"paid_at,omitempty" example:"2024-01-01T01:00:00Z"`
}

// OrderItem 订单项
type OrderItem struct {
	ProductID   int64   `json:"product_id" example:"100"`
	ProductName string  `json:"product_name" example:"iPhone 15 Pro"`
	Quantity    int     `json:"quantity" example:"1"`
	UnitPrice   float64 `json:"unit_price" example:"7999.00"`
	TotalPrice  float64 `json:"total_price" example:"7999.00"`
}

// Address 地址
type Address struct {
	Name     string `json:"name" example:"张三"`
	Phone    string `json:"phone" example:"13800138000"`
	Province string `json:"province" example:"广东省"`
	City     string `json:"city" example:"深圳市"`
	District string `json:"district" example:"南山区"`
	Street   string `json:"street" example:"科技园路1号"`
	ZipCode  string `json:"zip_code" example:"518000"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	UserID  int64              `json:"user_id" binding:"required" example:"1"`
	Items   []CreateOrderItem  `json:"items" binding:"required,min=1"`
	Address CreateAddressRequest `json:"address" binding:"required"`
}

// CreateOrderItem 创建订单项请求
type CreateOrderItem struct {
	ProductID int64 `json:"product_id" binding:"required" example:"100"`
	Quantity  int   `json:"quantity" binding:"required,min=1" example:"1"`
}

// CreateAddressRequest 创建地址请求
type CreateAddressRequest struct {
	Name     string `json:"name" binding:"required" example:"张三"`
	Phone    string `json:"phone" binding:"required" example:"13800138000"`
	Province string `json:"province" binding:"required" example:"广东省"`
	City     string `json:"city" binding:"required" example:"深圳市"`
	District string `json:"district" binding:"required" example:"南山区"`
	Street   string `json:"street" binding:"required" example:"科技园路1号"`
	ZipCode  string `json:"zip_code" example:"518000"`
}

// UpdateOrderStatusRequest 更新订单状态请求
type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped completed cancelled" example:"paid"`
}

// OrderListResponse 订单列表响应
type OrderListResponse struct {
	Total  int64   `json:"total" example:"100"`
	Page   int     `json:"page" example:"1"`
	Size   int     `json:"size" example:"10"`
	Orders []Order `json:"orders"`
}

// Response 通用响应
type Response struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
	Details string `json:"details,omitempty" example:"Invalid parameters"`
}

// OrderStatistics 订单统计
type OrderStatistics struct {
	TotalOrders     int64   `json:"total_orders" example:"1000"`
	TotalAmount     float64 `json:"total_amount" example:"999999.99"`
	PendingOrders   int64   `json:"pending_orders" example:"50"`
	CompletedOrders int64   `json:"completed_orders" example:"900"`
	CancelledOrders int64   `json:"cancelled_orders" example:"50"`
}
