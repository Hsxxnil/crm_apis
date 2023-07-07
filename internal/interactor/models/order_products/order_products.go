package order_products

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 訂單ID
	OrderID string `json:"order_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 產品ID
	ProductID string `json:"product_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 訂單產品數量
	Quantity int `json:"quantity,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 訂單產品單價
	UnitPrice float64 `json:"unit_price,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 訂單產品報價
	QuotePrice float64 `json:"quote_price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 訂單產品小計
	SubTotal float64 `json:"sub_total,omitempty" swaggerignore:"true"`
	// 訂單產品描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 訂單產品ID
	OrderProductID string `json:"order_product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 訂單ID
	OrderID *string `json:"order_id,omitempty" form:"order_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty" form:"product_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 訂單產品數量
	Quantity *int `json:"quantity,omitempty" form:"quantity"`
	// 訂單產品單價
	UnitPrice *float64 `json:"unit_price,omitempty" form:"unit_price"`
	// 訂單產品報價
	QuotePrice *float64 `json:"quote_price,omitempty" form:"quote_price"`
	// 訂單產品小計
	SubTotal *float64 `json:"sub_total,omitempty" form:"sub_total"`
	// 訂單產品描述
	Description *string `json:"description,omitempty" form:"description"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	// 多筆
	OrderProducts []*struct {
		// 訂單產品ID
		OrderProductID string `json:"order_product_id,omitempty"`
		// 訂單ID
		OrderID string `json:"order_id,omitempty"`
		// 產品ID
		ProductID string `json:"product_id,omitempty"`
		// 產品名稱
		ProductName string `json:"product_name,omitempty"`
		// 訂單產品數量
		Quantity int `json:"quantity,omitempty"`
		// 訂單產品單價
		UnitPrice float64 `json:"unit_price,omitempty"`
		// 訂單產品報價
		QuotePrice float64 `json:"quote_price,omitempty"`
		// 訂單產品小計
		SubTotal float64 `json:"sub_total,omitempty"`
		// 訂單產品描述
		Description string `json:"description,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"order_products"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 訂單產品ID
	OrderProductID string `json:"order_product_id,omitempty"`
	// 訂單ID
	OrderID string `json:"order_id,omitempty"`
	// 產品ID
	ProductID string `json:"product_id,omitempty"`
	// 產品名稱
	ProductName string `json:"product_name,omitempty"`
	// 訂單產品數量
	Quantity int `json:"quantity,omitempty"`
	// 訂單產品單價
	UnitPrice float64 `json:"unit_price,omitempty"`
	// 訂單產品報價
	QuotePrice float64 `json:"quote_price,omitempty"`
	// 訂單產品小計
	SubTotal float64 `json:"sub_total,omitempty"`
	// 訂單產品描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 訂單產品ID
	OrderProductID string `json:"order_product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 訂單產品數量
	Quantity *int `json:"quantity,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 訂單產品單價
	UnitPrice *float64 `json:"unit_price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 訂單產品報價
	QuotePrice *float64 `json:"quote_price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 訂單產品小計
	SubTotal float64 `json:"sub_total,omitempty" swaggerignore:"true"`
	// 訂單產品描述
	Description *string `json:"description,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// OrderSingle return structure file for orders
type OrderSingle struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty"`
	// 產品名稱
	ProductName string `json:"product_name,omitempty"`
	// 產品定價
	ProductPrice float64 `json:"standard_price,omitempty"`
	// 訂單產品數量
	Quantity int `json:"quantity,omitempty"`
	// 訂單產品單價
	UnitPrice float64 `json:"unit_price,omitempty"`
	// 訂單產品報價
	QuotePrice float64 `json:"quote_price,omitempty"`
	// 訂單產品小計
	SubTotal float64 `json:"sub_total,omitempty"`
	// 訂單產品描述
	Description string `json:"description,omitempty"`
}
