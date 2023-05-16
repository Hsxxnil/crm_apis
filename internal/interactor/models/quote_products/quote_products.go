package quote_products

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 報價ID
	QuoteID string `json:"quote_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 產品ID
	ProductID string `json:"product_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 報價產品數量
	Quantity int `json:"quantity,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 報價產品單價
	UnitPrice float64 `json:"unit_price,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 報價產品小計
	SubTotal float64 `json:"sub_total,omitempty" swaggerignore:"true"`
	// 報價產品折扣
	Discount float64 `json:"discount,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// CreateList struct is used to create multiple data
type CreateList struct {
	QuoteProducts []*Create `json:"quote_products"`
}

// Field is structure file for search
type Field struct {
	// 報價產品ID
	QuoteProductID string `json:"quote_product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 報價ID
	QuoteID *string `json:"quote_id,omitempty" form:"quote_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty" form:"product_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價產品數量
	Quantity *int `json:"quantity,omitempty" form:"quantity"`
	// 報價產品單價
	UnitPrice *float64 `json:"unit_price,omitempty" form:"unit_price"`
	// 報價產品小計
	SubTotal *float64 `json:"sub_total,omitempty" form:"sub_total"`
	// 報價產品折扣
	Discount *float64 `json:"discount,omitempty" form:"discount"`
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
	QuoteProducts []*struct {
		// 報價產品ID
		QuoteProductID string `json:"quote_product_id,omitempty"`
		// 報價ID
		QuoteID string `json:"quote_id,omitempty"`
		// 產品ID
		ProductID string `json:"product_id,omitempty"`
		// 產品名稱
		ProductName string `json:"product_name,omitempty"`
		// 報價產品數量
		Quantity int `json:"quantity,omitempty"`
		// 報價產品單價
		UnitPrice float64 `json:"unit_price,omitempty"`
		// 報價產品小計
		SubTotal float64 `json:"sub_total,omitempty"`
		// 報價產品折扣
		Discount float64 `json:"discount,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"quote_products"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 報價產品ID
	QuoteProductID string `json:"quote_product_id,omitempty"`
	// 報價ID
	QuoteID string `json:"quote_id,omitempty"`
	// 產品ID
	ProductID string `json:"product_id,omitempty"`
	// 產品名稱
	ProductName string `json:"product_name,omitempty"`
	// 報價產品數量
	Quantity int `json:"quantity,omitempty"`
	// 報價產品單價
	UnitPrice float64 `json:"unit_price,omitempty"`
	// 報價產品小計
	SubTotal float64 `json:"sub_total,omitempty"`
	// 報價產品折扣
	Discount float64 `json:"discount,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 報價產品ID
	QuoteProductID string `json:"quote_product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價產品數量
	Quantity *int `json:"quantity,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 報價產品單價
	UnitPrice *float64 `json:"unit_price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 報價產品小計
	SubTotal float64 `json:"sub_total,omitempty"`
	// 報價產品折扣
	Discount *float64 `json:"discount,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// QuoteSingle return structure file for quotes
type QuoteSingle struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty"`
	// 產品名稱
	ProductName string `json:"product_name,omitempty"`
	// 產品定價
	ProductPrice float64 `json:"standard_price,omitempty"`
	// 報價產品數量
	Quantity int `json:"quantity,omitempty"`
	// 報價產品單價
	UnitPrice float64 `json:"unit_price,omitempty"`
	// 報價產品小計
	SubTotal float64 `json:"sub_total,omitempty" swaggerignore:"true"`
	// 報價產品折扣
	Discount float64 `json:"discount,omitempty"`
}
