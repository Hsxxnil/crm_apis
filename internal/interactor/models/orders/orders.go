package orders

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/users"

	"app.eirc/internal/entity/postgresql/db/contracts"

	"app.eirc/internal/entity/postgresql/db/accounts"

	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 訂單狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 訂單開始日期
	StartDate time.Time `json:"start_date,omitempty" binding:"required" validate:"required"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 契約ID
	ContractID string `json:"contract_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 訂單描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 訂單ID
	OrderID string `json:"order_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 訂單狀態
	Status *string `json:"status,omitempty" form:"status"`
	// 訂單開始日期
	StartDate *time.Time `json:"start_date,omitempty" form:"start_date"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約ID
	ContractID *string `json:"contract_id,omitempty" form:"contract_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 訂單描述
	Description *string `json:"description,omitempty" form:"description"`
	// 訂單號碼
	Code *int `json:"code,omitempty" form:"code"`
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
	Orders []*struct {
		// 訂單ID
		OrderID string `json:"order_id,omitempty"`
		// 訂單狀態
		Status string `json:"status,omitempty"`
		// 訂單開始日期
		StartDate time.Time `json:"start_date,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		AccountName string `json:"account_name,omitempty"`
		// accounts data
		Accounts *accounts.Base `json:"accounts,omitempty" swaggerignore:"true"`
		// 契約ID
		ContractID string `json:"contract_id,omitempty"`
		// 契約號碼
		ContractCode int `json:"contract_code,omitempty"`
		// contracts data
		Contracts *contracts.Base `json:"contracts,omitempty" swaggerignore:"true"`
		// 訂單描述
		Description string `json:"description,omitempty"`
		// 訂單號碼
		Code int `json:"code,omitempty"`
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
	} `json:"orders"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 訂單ID
	OrderID string `json:"order_id,omitempty"`
	// 訂單狀態
	Status string `json:"status,omitempty"`
	// 訂單開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// accounts data
	Accounts *accounts.Base `json:"accounts,omitempty" swaggerignore:"true"`
	// 契約號碼
	ContractCode int `json:"contract_code,omitempty"`
	// contracts data
	Contracts *contracts.Base `json:"contracts,omitempty" swaggerignore:"true"`
	// 契約ID
	ContractID string `json:"contract_id,omitempty"`
	// 訂單描述
	Description string `json:"description,omitempty"`
	// 訂單號碼
	Code int `json:"code,omitempty"`
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
	// 訂單ID
	OrderID string `json:"order_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 訂單狀態
	Status *string `json:"status,omitempty"`
	// 訂單開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約ID
	ContractID string `json:"contract_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 訂單描述
	Description *string `json:"description,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
