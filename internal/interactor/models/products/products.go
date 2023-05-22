package products

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
	"app.eirc/internal/interactor/models/sort"
)

// Create struct is used to create achieves
type Create struct {
	// 產品名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 產品識別碼
	Code string `json:"code,omitempty"`
	// 產品是否啟用
	IsEnable bool `json:"is_enable,omitempty"`
	// 產品描述
	Description string `json:"description,omitempty"`
	// 產品價格
	Price float64 `json:"price,omitempty" binding:"required,gte=0" validate:"required,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 產品ID
	ProductID string `json:"product_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 產品名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 產品識別碼
	Code *string `json:"code,omitempty" form:"code"`
	// 產品是否啟用
	IsEnable *bool `json:"is_enable,omitempty" form:"is_enable"`
	// 產品描述
	Description *string `json:"description,omitempty" form:"description"`
	// 產品價格
	Price *float64 `json:"price,omitempty" form:"price"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 搜尋欄位
	Filter `json:"filter"`
	// 分頁搜尋結構檔
	page.Pagination
	// 排序欄位
	sort.Sort `json:"sort"`
}

// Filter struct is used to store the search field
type Filter struct {
	// 產品名稱
	FilterName *string `json:"name,omitempty"`
	// 產品識別碼
	FilterCode *string `json:"code,omitempty"`
	// 產品描述
	FilterDescription *string `json:"description,omitempty"`
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
		// 產品是否啟用
		IsEnable bool `json:"is_enable"`
		// 產品描述
		Description string `json:"description,omitempty"`
		// 產品價格
		Price float64 `json:"price,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
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
	// 產品是否啟用
	IsEnable bool `json:"is_enable"`
	// 產品描述
	Description string `json:"description,omitempty"`
	// 產品價格
	Price float64 `json:"price,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
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
	Code string `json:"code,omitempty"`
	// 產品是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 產品描述
	Description *string `json:"description,omitempty"`
	// 產品價格
	Price *float64 `json:"price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
