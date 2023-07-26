package contracts

import (
	"time"

	"app.eirc/internal/interactor/models/sort"

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
	// 契約結束日期
	EndDate time.Time `json:"end_date,omitempty" swaggerignore:"true"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 契約描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
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
	// 契約結束日期
	EndDate *time.Time `json:"end_date,omitempty" form:"end_date"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" form:"opportunity_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約描述
	Description *string `json:"description,omitempty" form:"description"`
	// 契約號碼
	Code *string `json:"code,omitempty" form:"code"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" form:"salesperson_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" form:"is_deleted"`
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
	// 契約號碼
	FilterCode string `json:"code,omitempty"`
	// 帳戶名稱
	FilterAccountName string `json:"account_name,omitempty"`
	// 契約狀態
	FilterStatus string `json:"status,omitempty"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Contracts []*struct {
		// 契約ID
		ContractID string `json:"contract_id,omitempty"`
		// 契約狀態
		Status string `json:"status,omitempty"`
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 商機名稱
		OpportunityName string `json:"opportunity_name,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		AccountName string `json:"account_name,omitempty"`
		// 契約開始日期
		StartDate time.Time `json:"start_date,omitempty"`
		// 契約有效期限(月)
		Term int `json:"term,omitempty"`
		// 契約結束日期
		EndDate time.Time `json:"end_date,omitempty"`
		// 契約描述
		Description string `json:"description,omitempty"`
		// 契約號碼
		Code string `json:"code,omitempty"`
		// 業務員ID
		SalespersonID string `json:"salesperson_id,omitempty"`
		// 業務員名稱
		SalespersonName string `json:"salesperson_name,omitempty"`
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
	// 契約結束日期
	EndDate time.Time `json:"end_date,omitempty"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 商機名稱
	OpportunityName string `json:"opportunity_name,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// 契約描述
	Description string `json:"description,omitempty"`
	// 契約號碼
	Code string `json:"code,omitempty"`
	// 業務員ID
	SalespersonID string `json:"salesperson_id,omitempty"`
	// 業務員名稱
	SalespersonName string `json:"salesperson_name,omitempty"`
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
	// 契約結束日期
	EndDate time.Time `json:"end_date,omitempty" swaggerignore:"true"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約描述
	Description *string `json:"description,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 契約是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
