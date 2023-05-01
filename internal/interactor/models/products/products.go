package products

import (
	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
	"github.com/shopspring/decimal"
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
	Price decimal.Decimal `json:"price,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
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
	Price *decimal.Decimal `json:"price,omitempty" form:"price"`
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
		// 產品是否啟用
		IsEnable bool `json:"is_enable"`
		// 產品描述
		Description string `json:"description,omitempty"`
		// 產品價格
		Price decimal.Decimal `json:"price,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// create_users data
		CreatedByUsers *users.Base `json:"created_by_users,omitempty" swaggerignore:"true"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// update_users data
		UpdatedByUsers *users.Base `json:"updated_by_users,omitempty" swaggerignore:"true"`
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
	Price decimal.Decimal `json:"price,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// create_users data
	CreatedByUsers *users.Base `json:"created_by_users,omitempty" swaggerignore:"true"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// update_users data
	UpdatedByUsers *users.Base `json:"updated_by_users,omitempty" swaggerignore:"true"`
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
	Price *decimal.Decimal `json:"price,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

type OrderProductSingle struct {
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
	Price decimal.Decimal `json:"price,omitempty"`
}
