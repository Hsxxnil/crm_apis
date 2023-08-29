package roles

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 公司ID
	CompanyID string `json:"company_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 角色顯示名稱
	DisplayName string `json:"display_name,omitempty" binding:"required" validate:"required"`
	// 角色名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 角色ID
	RoleID string `json:"role_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty" form:"company_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 角色顯示名稱
	DisplayName *string `json:"display_name,omitempty" form:"display_name"`
	// 角色名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 角色是否啟用
	IsEnable *bool `json:"is_enable,omitempty" form:"is_enable"`
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
	Roles []*struct {
		// 角色ID
		RoleID string `json:"role_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 角色顯示名稱
		DisplayName string `json:"display_name,omitempty"`
		// 角色名稱
		Name string `json:"name,omitempty"`
		// 角色是否啟用
		IsEnable bool `json:"is_enable,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"roles"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 角色ID
	RoleID string `json:"role_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 角色顯示名稱
	DisplayName string `json:"display_name,omitempty"`
	// 角色名稱
	Name string `json:"name,omitempty"`
	// 角色是否啟用
	IsEnable bool `json:"is_enable,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 角色ID
	RoleID string `json:"role_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 角色顯示名稱
	DisplayName *string `json:"display_name,omitempty"`
	// 角色名稱
	Name *string `json:"name,omitempty"`
	// 角色是否啟用
	IsEnable bool `json:"is_enable,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
