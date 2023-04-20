package account_types

import (
	"app.eirc/internal/interactor/models/page"
)

// Create struct is used to create achieves
type Create struct {
	// 帳戶類型名稱
	Name string `json:",omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search
type Field struct {
	// 帳戶類型ID
	AccountTypeID string `json:"account_type_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶類型名稱
	Name *string `json:"name,omitempty" from:"name"`
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
	Industries []*struct {
		// 帳戶類型ID
		AccountTypeID string `json:"account_type_id,omitempty"`
		// 帳戶類型名稱
		Name string `json:"name,omitempty"`
	} `json:"account_types"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 帳戶類型ID
	AccountTypeID string `json:"account_type_id,omitempty"`
	// 帳戶類型名稱
	Name string `json:"name,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 帳戶類型ID
	AccountTypeID string `json:"account_type_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶類型名稱
	Name *string `json:"name,omitempty"`
}
