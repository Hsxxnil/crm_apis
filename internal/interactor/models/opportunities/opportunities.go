package opportunities

import (
	"time"

	"crm/internal/interactor/models/sort"

	"crm/internal/interactor/models/opportunity_campaigns"

	"crm/internal/interactor/models/page"
	"crm/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 商機名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 商機階段
	Stage string `json:"stage,omitempty" binding:"required" validate:"required"`
	// 商機預測種類
	ForecastCategory string `json:"forecast_category,omitempty" binding:"required" validate:"required"`
	// 商機結束日期
	CloseDate time.Time `json:"close_date,omitempty" binding:"required" validate:"required"`
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 預期收入金額
	Amount float64 `json:"amount,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 商機名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 商機階段
	Stage *string `json:"stage,omitempty" form:"stage"`
	// 商機預測種類
	ForecastCategory *string `json:"forecast_category,omitempty" form:"forecast_category"`
	// 商機結束日期
	CloseDate *time.Time `json:"close_date,omitempty" form:"close_date"`
	// 線索ID
	LeadID *string `json:"lead_id,omitempty" form:"lead_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 預期收入金額
	Amount *float64 `json:"amount,omitempty" form:"amount"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" form:"salesperson_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
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

// FieldsNoPagination is the searched structure file (including filter)
type FieldsNoPagination struct {
	// 搜尋結構檔
	Field
	// 搜尋欄位
	FilterNoPagination `json:"filter"`
}

// Filter struct is used to store the search field
type Filter struct {
	// 商機名稱
	FilterName string `json:"name,omitempty"`
	// 帳戶名稱
	FilterAccountName string `json:"account_name,omitempty"`
	// 商機階段
	FilterStage []string `json:"stage,omitempty"`
	// 業務員名稱
	FilterSalespersonName string `json:"salesperson_name,omitempty"`
}

// FilterNoPagination struct is used to store the search field no pagination
type FilterNoPagination struct {
	// 商機名稱
	FilterName string `json:"name,omitempty"`
	// 商機階段
	FilterStage []string `json:"stage,omitempty"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Opportunities []*struct {
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 商機名稱
		Name string `json:"name,omitempty"`
		// 商機階段
		Stage string `json:"stage,omitempty"`
		// 商機預測種類
		ForecastCategory string `json:"forecast_category,omitempty"`
		// 商機結束日期
		CloseDate time.Time `json:"close_date,omitempty"`
		// 線索ID
		LeadID string `json:"lead_id,omitempty"`
		// 線索描述
		LeadDescription string `json:"lead_description,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		AccountName string `json:"account_name,omitempty"`
		// 預期收入金額
		Amount float64 `json:"amount,omitempty"`
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
	} `json:"opportunities"`
	// 分頁返回結構檔
	page.Total
}

// ListNoPagination is multiple return structure files without pagination
type ListNoPagination struct {
	// 多筆
	Opportunities []*struct {
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 商機名稱
		Name string `json:"name,omitempty"`
	} `json:"opportunities"`
}

// Single return structure file
type Single struct {
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 商機名稱
	Name string `json:"name,omitempty"`
	// 商機階段
	Stage string `json:"stage,omitempty"`
	// 商機預測種類
	ForecastCategory string `json:"forecast_category,omitempty"`
	// 商機結束日期
	CloseDate time.Time `json:"close_date,omitempty"`
	// 線索ID
	LeadID string `json:"lead_id,omitempty"`
	// 線索描述
	LeadDescription string `json:"lead_description,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// 預期收入金額
	Amount float64 `json:"amount,omitempty"`
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

// SingleCampaigns return structure file containing campaigns
type SingleCampaigns struct {
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 商機名稱
	Name string `json:"name,omitempty"`
	// 商機階段
	Stage string `json:"stage,omitempty"`
	// 商機預測種類
	ForecastCategory string `json:"forecast_category,omitempty"`
	// 商機結束日期
	CloseDate time.Time `json:"close_date,omitempty"`
	// 線索ID
	LeadID string `json:"lead_id,omitempty"`
	// 線索描述
	LeadDescription string `json:"lead_description,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// 預期收入金額
	Amount float64 `json:"amount,omitempty"`
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
	// opportunity_campaigns data
	OpportunityCampaigns []opportunity_campaigns.OpportunitySingle `json:"campaigns,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 商機名稱
	Name *string `json:"name,omitempty"`
	// 商機階段
	Stage *string `json:"stage,omitempty"`
	// 商機預測種類
	ForecastCategory *string `json:"forecast_category,omitempty"`
	// 商機結束日期
	CloseDate *time.Time `json:"close_date,omitempty"`
	// 預期收入金額
	Amount *float64 `json:"amount,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
