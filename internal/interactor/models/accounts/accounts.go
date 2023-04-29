package accounts

import (
	"app.eirc/internal/interactor/models/contacts"
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 帳戶名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 帳戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type string `json:"type,omitempty" binding:"required" validate:"required"`
	// 行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 父系帳戶ID
	ParentAccountID string `json:"parent_account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 帳戶類型
	Type *string `json:"type,omitempty" form:"type"`
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty" form:"industry_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty" form:"parent_account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
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
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		Name string `json:"name,omitempty"`
		// 帳戶電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 帳戶類型
		Type string `json:"type,omitempty"`
		// 行業ID
		IndustryID string `json:"industry_id,omitempty"`
		// 父系帳戶ID
		ParentAccountID string `json:"parent_account_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
		// contacts data
		Contacts []contacts.Single `json:"contacts,omitempty"`
	} `json:"accounts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	Name string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type string `json:"type,omitempty"`
	// 行業ID
	IndustryID string `json:"industry_id,omitempty"`
	// 父系帳戶ID
	ParentAccountID string `json:"parent_account_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
	// contacts data
	Contacts []contacts.Single `json:"contacts,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶名稱
	Name *string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type *string `json:"type,omitempty"`
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

type AccountName struct {
	// 帳戶名稱
	Name *string `json:"name,omitempty"`
}
