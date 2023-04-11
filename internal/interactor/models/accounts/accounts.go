package accounts

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 公司ID
	CompanyID string `json:"company_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	// 電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 帳號ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty" from:"company_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳號
	Account *string `json:"account,omitempty" from:"account"`
	// 中文名稱
	Name *string `json:"name,omitempty" from:"name"`
	// 密碼
	Password *string `json:"password,omitempty" from:"password"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" from:"is_deleted"`
	// 電話
	PhoneNumber *string `json:"phone_number,omitempty" from:"phone_number"`
	// 電子郵件
	Email *string `json:"email,omitempty" from:"email"`
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
	Accounts []*struct {
		// 帳號ID
		AccountID string `json:"account_id,omitempty"`
		// 公司ID
		CompanyID string `json:"company_id,omitempty"`
		// 帳號
		Account string `json:"account,omitempty"`
		// 中文名稱
		Name string `json:"name,omitempty"`
		// 密碼
		Password string `json:"password,omitempty"`
		// 是否刪除
		IsDeleted bool `json:"is_deleted,omitempty"`
		// 電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 電子郵件
		Email string `json:"email,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"accounts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 帳號ID
	AccountID string `json:"account_id,omitempty"`
	// 公司ID
	CompanyID string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 電子郵件
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
	// 帳號ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 密碼
	Password string `json:"password,omitempty"`
	// 電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 電子郵件
	Email *string `json:"email,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
