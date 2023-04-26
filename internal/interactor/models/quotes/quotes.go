package quotes

import (
	"time"

	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
	"github.com/shopspring/decimal"
)

// Create struct is used to create achieves
type Create struct {
	// 報價名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 報價狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 報價與商機是否同步化
	IsSyncing bool `json:"is_syncing,omitempty"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 報價到期日期
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description string `json:"description,omitempty"`
	// 報價稅額
	Tax decimal.Decimal `json:"tax,omitempty"`
	// 報價運輸和處理費
	ShippingAndHandling decimal.Decimal `json:"shipping_and_handling,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 產品ID
	QuoteID string `json:"quote_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 報價名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 報價狀態
	Status *string `json:"status,omitempty" form:"status"`
	// 報價與商機是否同步化
	IsSyncing *bool `json:"is_syncing,omitempty" form:"is_syncing"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" form:"opportunity_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價到期日期
	ExpirationDate *time.Time `json:"expiration_date,omitempty" form:"expiration_date"`
	// 報價描述
	Description *string `json:"description,omitempty" form:"description"`
	// 報價稅額
	Tax *decimal.Decimal `json:"tax,omitempty" form:"tax"`
	// 報價運輸和處理費
	ShippingAndHandling *decimal.Decimal `json:"shipping_and_handling,omitempty" form:"shipping_and_handling"`
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
	Quotes []*struct {
		// 報價ID
		QuoteID string `json:"quote_id,omitempty"`
		// 報價名稱
		Name string `json:"name,omitempty"`
		// 報價狀態
		Status string `json:"status,omitempty"`
		// 報價與商機是否同步化
		IsSyncing bool `json:"is_syncing"`
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 報價到期日期
		ExpirationDate time.Time `json:"expiration_date,omitempty"`
		// 報價描述
		Description string `json:"description,omitempty"`
		// 報價稅額
		Tax decimal.Decimal `json:"tax,omitempty"`
		// 報價運輸和處理費
		ShippingAndHandling decimal.Decimal `json:"shipping_and_handling,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"quotes"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 報價ID
	QuoteID string `json:"quote_id,omitempty"`
	// 報價名稱
	Name string `json:"name,omitempty"`
	// 報價狀態
	Status string `json:"status,omitempty"`
	// 報價與商機是否同步化
	IsSyncing bool `json:"is_syncing"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 報價到期日期
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description string `json:"description,omitempty"`
	// 報價稅額
	Tax decimal.Decimal `json:"tax,omitempty"`
	// 報價運輸和處理費
	ShippingAndHandling decimal.Decimal `json:"shipping_and_handling,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 產品ID
	QuoteID string `json:"quote_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 報價名稱
	Name *string `json:"name,omitempty"`
	// 報價狀態
	Status *string `json:"status,omitempty"`
	// 報價與商機是否同步化
	IsSyncing *bool `json:"is_syncing,omitempty"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價到期日期
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description *string `json:"description,omitempty"`
	// 報價稅額
	Tax *decimal.Decimal `json:"tax,omitempty"`
	// 報價運輸和處理費
	ShippingAndHandling *decimal.Decimal `json:"shipping_and_handling,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
