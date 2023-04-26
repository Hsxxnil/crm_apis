package campaigns

import (
	"time"

	"app.eirc/internal/interactor/models/special"
	"github.com/shopspring/decimal"
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
	StartDate time.Time `gorm:"column:start_date;type:TIMESTAMP;" json:"start_date"`
	// 行銷活動結束日期
	EndDate time.Time `gorm:"column:end_date;type:TIMESTAMP;" json:"end_date"`
	// 行銷活動描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 行銷活動已傳送數量
	Sent int `gorm:"column:sent;type:int;" json:"sent"`
	// 行銷活動預算成本
	BudgetCost decimal.Decimal `gorm:"column:budget_cost;type:numeric;" json:"budget_cost"`
	// 行銷活動預期回應(%)
	ExpectedResponses int `gorm:"column:expected_responses;type:int;" json:"expected_responses"`
	// 行銷活動實際成本
	ActualCost decimal.Decimal `gorm:"column:actual_cost;type:numeric;" json:"actual_cost"`
	// 行銷活動預期收入
	ExpectedIncome decimal.Decimal `gorm:"column:expected_income;type:numeric;" json:"expected_income"`
	special.UseTable
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
	BudgetCost *decimal.Decimal `json:"budget_cost,omitempty"`
	// 行銷活動預期回應(%)
	ExpectedResponses *int `json:"expected_responses,omitempty"`
	// 行銷活動實際成本
	ActualCost *decimal.Decimal `json:"actual_cost,omitempty"`
	// 行銷活動預期收入
	ExpectedIncome *decimal.Decimal `json:"expected_income,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "campaigns"
}
