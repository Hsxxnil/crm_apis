package orders

import (
	"time"

	"app.eirc/internal/interactor/models/sort"

	"app.eirc/internal/interactor/models/order_products"

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
	Code *string `json:"code,omitempty" form:"code"`
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
	// 訂單號碼
	FilterCode *string `json:"code,omitempty"`
	// TODO 帳戶名稱
	// TODO 契約號碼
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
		// 契約ID
		ContractID string `json:"contract_id,omitempty"`
		// 契約號碼
		ContractCode string `json:"contract_code,omitempty"`
		// 訂單描述
		Description string `json:"description,omitempty"`
		// 訂單號碼
		Code string `json:"code,omitempty"`
		// 啟用者
		ActivatedBy string `json:"activated_by,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
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
	// 契約號碼
	ContractCode string `json:"contract_code,omitempty"`
	// 契約ID
	ContractID string `json:"contract_id,omitempty"`
	// 訂單描述
	Description string `json:"description,omitempty"`
	// 訂單號碼
	Code string `json:"code,omitempty"`
	// 啟用者
	ActivatedBy string `json:"activated_by,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// SingleProducts return structure file containing products
type SingleProducts struct {
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
	// 契約號碼
	ContractCode string `json:"contract_code,omitempty"`
	// 契約ID
	ContractID string `json:"contract_id,omitempty"`
	// 訂單描述
	Description string `json:"description,omitempty"`
	// 訂單號碼
	Code string `json:"code,omitempty"`
	// 啟用者
	ActivatedBy string `json:"activated_by,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
	// order_products data
	OrderProducts []order_products.OrderSingle `json:"products,omitempty"`
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
	// 啟用者
	ActivatedBy *string `json:"activated_by,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
}
