package leads

import (
	"app.eirc/internal/entity/postgresql/db/accounts"
	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 線索狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 線索描述
	Description string `json:"description,omitempty" binding:"required" validate:"required"`
	// 線索來源
	Source string `json:"source,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 線索分級
	Rating string `json:"rating,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 線索狀態
	Status *string `json:"status,omitempty" form:"status"`
	// 線索描述
	Description *string `json:"description,omitempty" form:"description"`
	// 線索來源
	Source *string `json:"source,omitempty" form:"source"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索分級
	Rating *string `json:"rating,omitempty" form:"rating"`
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
	Leads []*struct {
		// 線索ID
		LeadID string `json:"lead_id,omitempty"`
		// 線索狀態
		Status string `json:"status,omitempty"`
		// 線索描述
		Description string `json:"description,omitempty"`
		// 線索來源
		Source string `json:"source,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		AccountName string `json:"account_name,omitempty"`
		// accounts data
		Accounts *accounts.Base `json:"accounts,omitempty" swaggerignore:"true"`
		// 線索分級
		Rating string `json:"rating,omitempty"`
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
	} `json:"leads"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 線索ID
	LeadID string `json:"lead_id,omitempty"`
	// 線索狀態
	Status string `json:"status,omitempty"`
	// 線索描述
	Description string `json:"description,omitempty"`
	// 線索來源
	Source string `json:"source,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// accounts data
	Accounts *accounts.Base `json:"accounts,omitempty" swaggerignore:"true"`
	// 線索分級
	Rating string `json:"rating,omitempty"`
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
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 線索狀態
	Status *string `json:"status,omitempty"`
	// 線索描述
	Description *string `json:"description,omitempty"`
	// 線索來源
	Source *string `json:"source,omitempty"`
	// 線索分級
	Rating *string `json:"rating,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
