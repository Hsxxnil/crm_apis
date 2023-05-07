package opportunities

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/opportunity_campaigns"

	"app.eirc/internal/entity/postgresql/db/accounts"

	"app.eirc/internal/entity/postgresql/db/users"

	"app.eirc/internal/interactor/models/special"
	"github.com/shopspring/decimal"
)

// Table struct is opportunities database table struct
type Table struct {
	// 商機ID
	OpportunityID string `gorm:"<-:create;column:opportunity_id;type:uuid;not null;primaryKey;" json:"opportunity_id"`
	// 商機名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 商機階段
	Stage string `gorm:"column:stage;type:text;not null;" json:"stage"`
	// 商機預測種類
	ForecastCategory string `gorm:"column:forecast_category;type:text;not null;" json:"forecast_category"`
	// 商機結束日期
	CloseDate time.Time `gorm:"column:close_date;type:date;not null;" json:"close_date"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 預期收入金額
	Amount decimal.Decimal `gorm:"column:amount;type:numeric;" json:"amount"`
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
	OpportunityCampaigns []opportunity_campaigns.Table `gorm:"foreignKey:OpportunityID;" json:"campaigns,omitempty"`
}

// Base struct is corresponding to opportunities table structure file
type Base struct {
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty"`
	// 商機名稱
	Name *string `json:"name,omitempty"`
	// 商機階段
	Stage *string `json:"stage,omitempty"`
	// 商機預測種類
	ForecastCategory *string `json:"forecast_category,omitempty"`
	// 商機結束日期
	CloseDate *time.Time `json:"close_date,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 預期收入金額
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty"`
	// salespeople  data
	Salespeople users.Base `json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
	// opportunity_campaigns data
	OpportunityCampaigns []opportunity_campaigns.Base `json:"campaigns,omitempty"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "opportunities"
}
