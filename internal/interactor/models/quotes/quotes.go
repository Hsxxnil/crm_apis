package quotes

import (
	"time"

	"crm/internal/interactor/models/sort"

	"crm/internal/interactor/models/quote_products"

	"crm/internal/interactor/models/page"
	"crm/internal/interactor/models/section"
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
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 報價到期日期
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description string `json:"description,omitempty"`
	// 報價稅額
	Tax float64 `json:"tax,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 報價運費及其他費用
	ShippingAndHandling float64 `json:"shipping_and_handling,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
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
	// 報價是否為最終版
	IsFinal *bool `json:"is_final,omitempty" form:"is_final"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" form:"opportunity_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價到期日期
	ExpirationDate *time.Time `json:"expiration_date,omitempty" form:"expiration_date"`
	// 報價描述
	Description *string `json:"description,omitempty" form:"description"`
	// 報價稅額
	Tax *float64 `json:"tax,omitempty" form:"tax"`
	// 報價運費及其他費用
	ShippingAndHandling *float64 `json:"shipping_and_handling,omitempty" form:"shipping_and_handling"`
	// 報價號碼
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
	// 報價名稱
	FilterName string `json:"name,omitempty"`
	// 商機名稱
	FilterOpportunityName string `json:"opportunity_name,omitempty"`
	// 報價狀態
	FilterStatus string `json:"status,omitempty"`
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
		// 報價是否為最終版
		IsFinal bool `json:"is_final"`
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 商機名稱
		OpportunityName string `json:"opportunity_name,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 報價到期日期
		ExpirationDate time.Time `json:"expiration_date,omitempty"`
		// 報價描述
		Description string `json:"description,omitempty"`
		// 報價稅額
		Tax float64 `json:"tax,omitempty"`
		// 報價運費及其他費用
		ShippingAndHandling float64 `json:"shipping_and_handling,omitempty"`
		// 報價小計
		SubTotal float64 `json:"sub_total"`
		// 報價總價
		TotalPrice float64 `json:"total_price"`
		// 報價折扣
		Discount float64 `json:"discount"`
		// 報價總計
		GrandTotal float64 `json:"grand_total"`
		// 報價號碼
		Code string `json:"code,omitempty"`
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
	// 報價是否為最終版
	IsFinal bool `json:"is_final"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 商機名稱
	OpportunityName string `json:"opportunity_name,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 報價到期日期
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description string `json:"description,omitempty"`
	// 報價稅額
	Tax float64 `json:"tax,omitempty"`
	// 報價運費及其他費用
	ShippingAndHandling float64 `json:"shipping_and_handling,omitempty"`
	// 報價小計
	SubTotal float64 `json:"sub_total"`
	// 報價總價
	TotalPrice float64 `json:"total_price"`
	// 報價折扣
	Discount float64 `json:"discount"`
	// 報價總計
	GrandTotal float64 `json:"grand_total"`
	// 報價號碼
	Code string `json:"code,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// SingleProducts return structure file containing products
type SingleProducts struct {
	// 報價ID
	QuoteID string `json:"quote_id,omitempty"`
	// 報價名稱
	Name string `json:"name,omitempty"`
	// 報價狀態
	Status string `json:"status,omitempty"`
	// 報價與商機是否同步化
	IsSyncing bool `json:"is_syncing"`
	// 報價是否為最終版
	IsFinal bool `json:"is_final"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 商機名稱
	OpportunityName string `json:"opportunity_name,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 報價到期日期
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description string `json:"description,omitempty"`
	// 報價稅額
	Tax float64 `json:"tax,omitempty"`
	// 報價運費及其他費用
	ShippingAndHandling float64 `json:"shipping_and_handling,omitempty"`
	// 報價小計
	SubTotal float64 `json:"sub_total"`
	// 報價總價
	TotalPrice float64 `json:"total_price"`
	// 報價折扣
	Discount float64 `json:"discount"`
	// 報價總計
	GrandTotal float64 `json:"grand_total"`
	// 報價號碼
	Code string `json:"code,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
	// quote_products data
	QuoteProducts []quote_products.QuoteSingle `json:"products,omitempty"`
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
	// 報價是否為最終版
	IsFinal *bool `json:"is_final,omitempty"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 報價到期日期
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description *string `json:"description,omitempty"`
	// 報價稅額
	Tax *float64 `json:"tax,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 報價運費及其他費用
	ShippingAndHandling *float64 `json:"shipping_and_handling,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
