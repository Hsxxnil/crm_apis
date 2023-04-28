package contracts

import (
	"time"

	"app.eirc/internal/interactor/models/accounts"

	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 契約狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 契約開始日期
	StartDate time.Time `json:"start_date,omitempty" binding:"required" validate:"required"`
	// 契約有效期限(月)
	Term int `json:"term,omitempty" binding:"required" validate:"required"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 契約描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 契約ID
	ContractID string `json:"contract_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 契約狀態
	Status *string `json:"status,omitempty" form:"status"`
	// 契約開始日期
	StartDate *time.Time `json:"start_date,omitempty" form:"start_date"`
	// 契約有效期限(月)
	Term *int `json:"term,omitempty" form:"term"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約描述
	Description *string `json:"description,omitempty" form:"description"`
	// 契約號碼
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
	Contracts []*struct {
		// 契約ID
		ContractID string `json:"contract_id,omitempty"`
		// 契約狀態
		Status string `json:"status,omitempty"`
		// 契約開始日期
		StartDate time.Time `json:"start_date,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		Accounts accounts.AccountName `json:"accounts,omitempty"`
		// 契約有效期限(月)
		Term int `json:"term,omitempty"`
		// 契約描述
		Description string `json:"description,omitempty"`
		// 契約號碼
		Code int `json:"code,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"contracts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 契約ID
	ContractID string `json:"contract_id,omitempty"`
	// 契約狀態
	Status string `json:"status,omitempty"`
	// 契約開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 契約有效期限(月)
	Term int `json:"term,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	Accounts accounts.AccountName `json:"accounts,omitempty"`
	// 契約描述
	Description string `json:"description,omitempty"`
	// 契約號碼
	Code int `json:"code,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 契約ID
	ContractID string `json:"contract_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 契約狀態
	Status *string `json:"status,omitempty"`
	// 契約開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 契約有效期限(月)
	Term *int `json:"term,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約描述
	Description *string `json:"description,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

type ContractCode struct {
	// 契約號碼
	Code *int `json:"code,omitempty"`
}
