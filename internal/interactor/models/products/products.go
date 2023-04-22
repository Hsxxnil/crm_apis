package products

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 產品名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 產品識別碼
	Code string `json:"code,omitempty"`
	// 是否啟用
	IsEnable bool `json:"is_enable,omitempty"`
	// 產品描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 產品名稱
	Name *string `json:"name,omitempty" from:"name"`
	// 產品識別碼
	Code *string `json:"code,omitempty" from:"code"`
	// 是否啟用
	IsEnable *bool `json:"is_enable,omitempty" from:"is_enable"`
	// 產品描述
	Description *string `json:"description,omitempty" from:"description"`
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
	Products []*struct {
		// 產品ID
		ProductID string `json:"product_id,omitempty"`
		// 產品名稱
		Name string `json:"name,omitempty"`
		// 產品識別碼
		Code string `json:"code,omitempty"`
		// 是否啟用
		IsEnable bool `json:"is_enable"`
		// 產品描述
		Description string `json:"description,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"products"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty"`
	// 產品名稱
	Name string `json:"name,omitempty"`
	// 產品識別碼
	Code string `json:"code,omitempty"`
	// 是否啟用
	IsEnable bool `json:"is_enable"`
	// 產品描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 產品名稱
	Name *string `json:"name,omitempty"`
	// 產品識別碼
	Code *string `json:"code,omitempty"`
	// 是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 產品描述
	Description string `json:"description,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
