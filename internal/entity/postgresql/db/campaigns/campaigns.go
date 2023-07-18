package campaigns

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/sort"

	"app.eirc/internal/entity/postgresql/db/opportunity_campaigns"

	model "app.eirc/internal/interactor/models/campaigns"
	"app.eirc/internal/interactor/models/special"
)

// Table struct is campaigns database table struct
type Table struct {
	// 行銷活動ID
	CampaignID string `gorm:"<-:create;column:campaign_id;type:uuid;not null;primaryKey;" json:"campaign_id"`
	// 行銷活動名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 行銷活動狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 行銷活動是否啟用
	IsEnable bool `gorm:"column:is_enable;type:bool;not null;" json:"is_enable"`
	// 行銷活動類型
	Type string `gorm:"column:type;type:text;" json:"type"`
	// 父系行銷活動ID
	ParentCampaignID string `gorm:"column:parent_campaign_id;type:uuid;not null;" json:"parent_campaign_id"`
	// 行銷活動開始日期
	StartDate time.Time `gorm:"column:start_date;type:timestamp;" json:"start_date"`
	// 行銷活動結束日期
	EndDate time.Time `gorm:"column:end_date;type:timestamp;" json:"end_date"`
	// 行銷活動描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 行銷活動已傳送數量
	Sent int `gorm:"column:sent;type:int;" json:"sent"`
	// 行銷活動預算成本
	BudgetCost float64 `gorm:"column:budget_cost;type:numeric;" json:"budget_cost"`
	// 行銷活動預期回應(%)
	ExpectedResponses float64 `gorm:"column:expected_responses;type:numeric;" json:"expected_responses"`
	// 行銷活動實際成本
	ActualCost float64 `gorm:"column:actual_cost;type:numeric;" json:"actual_cost"`
	// 行銷活動預期收入
	ExpectedIncome float64 `gorm:"column:expected_income;type:numeric;" json:"expected_income"`
	// 業務員ID
	SalespersonID string `gorm:"column:salesperson_id;type:uuid;not null;" json:"salesperson_id"`
	// salespeople  data
	Salespeople users.Table `gorm:"foreignKey:SalespersonID;references:UserID" json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.UseTable
	// opportunity_campaigns data
	OpportunityCampaigns []opportunity_campaigns.Table `gorm:"foreignKey:CampaignID;" json:"opportunities,omitempty"`
}

// Base struct is corresponding to campaigns table structure file
type Base struct {
	// 行銷活動ID
	CampaignID *string `json:"campaign_id,omitempty"`
	// 行銷活動名稱
	Name *string `json:"name,omitempty"`
	// 行銷活動狀態
	Status *string `json:"status,omitempty"`
	// 行銷活動是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 行銷活動類型
	Type *string `json:"type,omitempty"`
	// 父系行銷活動ID
	ParentCampaignID *string `json:"parent_campaign_id,omitempty"`
	// 行銷活動開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 行銷活動結束日期
	EndDate *time.Time `json:"end_date,omitempty"`
	// 行銷活動描述
	Description *string `json:"description,omitempty"`
	// 行銷活動已傳送數量
	Sent *int `json:"sent,omitempty"`
	// 行銷活動預算成本
	BudgetCost *float64 `json:"budget_cost,omitempty"`
	// 行銷活動預期回應(%)
	ExpectedResponses *float64 `json:"expected_responses,omitempty"`
	// 行銷活動實際成本
	ActualCost *float64 `json:"actual_cost,omitempty"`
	// 行銷活動預期收入
	ExpectedIncome *float64 `json:"expected_income,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty"`
	// salespeople  data
	Salespeople users.Base `json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// opportunity_campaigns data
	OpportunityCampaigns []opportunity_campaigns.Base `json:"opportunities,omitempty"`
	special.UseBase
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "campaigns"
}
