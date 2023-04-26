package users

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 公司ID
	CompanyID string `json:"company_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 使用者名稱
	UserName string `json:"user_name,omitempty" binding:"required" validate:"required"`
	// 使用者中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 使用者密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	// 使用者電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 使用者ID
	UserID string `json:"user_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty" form:"company_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 使用者名稱
	UserName *string `json:"user_name,omitempty" form:"user_name"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 使用者密碼
	Password *string `json:"password,omitempty" form:"password"`
	// 使用者是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" form:"is_deleted"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty" form:"email"`
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
	Users []*struct {
		// 使用者ID
		UserID string `json:"user_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 使用者名稱
		UserName string `json:"user_name,omitempty"`
		// 使用者中文名稱
		Name string `json:"name,omitempty"`
		// 使用者密碼
		Password string `json:"password,omitempty"`
		// 使用者是否刪除
		IsDeleted bool `json:"is_deleted,omitempty"`
		// 使用者電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 使用者電子郵件
		Email string `json:"email,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"users"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 使用者ID
	UserID string `json:"user_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 使用者名稱
	UserName string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name string `json:"name,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 使用者電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 使用者ID
	UserID string `json:"user_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
