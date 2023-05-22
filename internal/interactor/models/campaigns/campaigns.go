package campaigns

import (
	"time"

	"app.eirc/internal/interactor/models/sort"

	"app.eirc/internal/interactor/models/opportunity_campaigns"

	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 行銷活動名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 行銷活動狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 行銷活動是否啟用
	IsEnable bool `json:"is_enable,omitempty"`
	// 行銷活動類型
	Type string `json:"type,omitempty"`
	// 父系行銷活動ID
	ParentCampaignID string `json:"parent_campaign_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 行銷活動開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 行銷活動結束日期
	EndDate time.Time `json:"end_date,omitempty"`
	// 行銷活動描述
	Description string `json:"description,omitempty"`
	// 行銷活動已傳送數量
	Sent int `json:"sent,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預算成本
	BudgetCost float64 `json:"budget_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預期回應(%)
	ExpectedResponses float64 `json:"expected_responses,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動實際成本
	ActualCost float64 `json:"actual_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預期收入
	ExpectedIncome float64 `json:"expected_income,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 行銷活動名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 行銷活動狀態
	Status *string `json:"status,omitempty" form:"status"`
	// 行銷活動是否啟用
	IsEnable *bool `json:"is_enable,omitempty" form:"is_enable"`
	// 行銷活動類型
	Type *string `json:"type,omitempty" form:"type"`
	// 父系行銷活動ID
	ParentCampaignID *string `json:"parent_campaign_id,omitempty" form:"parent_campaign_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 行銷活動開始日期
	StartDate *time.Time `json:"start_date,omitempty" form:"start_date"`
	// 行銷活動結束日期
	EndDate *time.Time `json:"end_date,omitempty" form:"end_date"`
	// 行銷活動描述
	Description *string `json:"description,omitempty" form:"description"`
	// 行銷活動已傳送數量
	Sent *int `json:"sent,omitempty" form:"sent"`
	// 行銷活動預算成本
	BudgetCost *float64 `json:"budget_cost,omitempty" form:"budget_cost"`
	// 行銷活動預期回應(%)
	ExpectedResponses *float64 `json:"expected_responses,omitempty" form:"expected_responses"`
	// 行銷活動實際成本
	ActualCost *float64 `json:"actual_cost,omitempty" form:"actual_cost"`
	// 行銷活動預期收入
	ExpectedIncome *float64 `json:"expected_income,omitempty" form:"expected_income"`
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

// Filter struct is used to store the search field
type Filter struct {
	// 行銷活動名稱
	FilterName *string `json:"name,omitempty"`
	// 父系行銷活動名稱
	FilterParentCampaignName *string `json:"parent_campaign_name,omitempty"`
	// 行銷活動類型
	FilterType *string `json:"type,omitempty"`
	// 業務員名稱
	FilterSalespersonName *string `json:"salesperson_name,omitempty"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Campaigns []*struct {
		// 行銷活動ID
		CampaignID string `json:"campaign_id,omitempty"`
		// 行銷活動名稱
		Name string `json:"name,omitempty"`
		// 行銷活動狀態
		Status string `json:"status,omitempty"`
		// 行銷活動是否啟用
		IsEnable bool `json:"is_enable"`
		// 行銷活動類型
		Type string `json:"type,omitempty"`
		// 父系行銷活動ID
		ParentCampaignID string `json:"parent_campaign_id,omitempty"`
		// 父系行銷活動名稱
		ParentCampaignName string `json:"parent_campaign_name,omitempty"`
		// 行銷活動開始日期
		StartDate time.Time `json:"start_date,omitempty"`
		// 行銷活動結束日期
		EndDate time.Time `json:"end_date,omitempty"`
		// 行銷活動描述
		Description string `json:"description,omitempty"`
		// 行銷活動已傳送數量
		Sent int `json:"sent,omitempty"`
		// 行銷活動預算成本
		BudgetCost float64 `json:"budget_cost,omitempty"`
		// 行銷活動預期回應(%)
		ExpectedResponses float64 `json:"expected_responses,omitempty"`
		// 行銷活動實際成本
		ActualCost float64 `json:"actual_cost,omitempty"`
		// 行銷活動預期收入
		ExpectedIncome float64 `json:"expected_income,omitempty"`
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
	} `json:"campaigns"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty"`
	// 行銷活動名稱
	Name string `json:"name,omitempty"`
	// 行銷活動狀態
	Status string `json:"status,omitempty"`
	// 行銷活動是否啟用
	IsEnable bool `json:"is_enable"`
	// 行銷活動類型
	Type string `json:"type,omitempty"`
	// 父系行銷活動ID
	ParentCampaignID string `json:"parent_campaign_id,omitempty"`
	// 父系行銷活動名稱
	ParentCampaignName string `json:"parent_campaign_name,omitempty"`
	// 行銷活動開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 行銷活動結束日期
	EndDate time.Time `json:"end_date,omitempty"`
	// 行銷活動描述
	Description string `json:"description,omitempty"`
	// 行銷活動已傳送數量
	Sent int `json:"sent,omitempty"`
	// 行銷活動預算成本
	BudgetCost float64 `json:"budget_cost,omitempty"`
	// 行銷活動預期回應(%)
	ExpectedResponses float64 `json:"expected_responses,omitempty"`
	// 行銷活動實際成本
	ActualCost float64 `json:"actual_cost,omitempty"`
	// 行銷活動預期收入
	ExpectedIncome float64 `json:"expected_income,omitempty"`
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

// SingleOpportunities return structure file containing opportunities
type SingleOpportunities struct {
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty"`
	// 行銷活動名稱
	Name string `json:"name,omitempty"`
	// 行銷活動狀態
	Status string `json:"status,omitempty"`
	// 行銷活動是否啟用
	IsEnable bool `json:"is_enable"`
	// 行銷活動類型
	Type string `json:"type,omitempty"`
	// 父系行銷活動ID
	ParentCampaignID string `json:"parent_campaign_id,omitempty"`
	// 父系行銷活動名稱
	ParentCampaignName string `json:"parent_campaign_name,omitempty"`
	// 行銷活動開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 行銷活動結束日期
	EndDate time.Time `json:"end_date,omitempty"`
	// 行銷活動描述
	Description string `json:"description,omitempty"`
	// 行銷活動已傳送數量
	Sent int `json:"sent,omitempty"`
	// 行銷活動預算成本
	BudgetCost float64 `json:"budget_cost,omitempty"`
	// 行銷活動預期回應(%)
	ExpectedResponses float64 `json:"expected_responses,omitempty"`
	// 行銷活動實際成本
	ActualCost float64 `json:"actual_cost,omitempty"`
	// 行銷活動預期收入
	ExpectedIncome float64 `json:"expected_income,omitempty"`
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
	OpportunityCampaigns []opportunity_campaigns.CampaignSingle `json:"opportunities,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 行銷活動名稱
	Name *string `json:"name,omitempty"`
	// 行銷活動狀態
	Status *string `json:"status,omitempty"`
	// 行銷活動是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 行銷活動類型
	Type *string `json:"type,omitempty"`
	// 父系行銷活動ID
	ParentCampaignID *string `json:"parent_campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 行銷活動開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 行銷活動結束日期
	EndDate *time.Time `json:"end_date,omitempty"`
	// 行銷活動描述
	Description *string `json:"description,omitempty"`
	// 行銷活動已傳送數量
	Sent *int `json:"sent,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預算成本
	BudgetCost *float64 `json:"budget_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預期回應(%)
	ExpectedResponses *float64 `json:"expected_responses,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動實際成本
	ActualCost *float64 `json:"actual_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 行銷活動預期收入
	ExpectedIncome *float64 `json:"expected_income,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
