package contracts

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/opportunities"

	"app.eirc/internal/interactor/models/sort"

	"app.eirc/internal/entity/postgresql/db/accounts"
	"app.eirc/internal/entity/postgresql/db/users"
	model "app.eirc/internal/interactor/models/contracts"

	"app.eirc/internal/interactor/models/special"
)

// Table struct is contracts database table struct
type Table struct {
	// 契約ID
	ContractID string `gorm:"<-:create;column:contract_id;type:uuid;not null;primaryKey;" json:"contract_id"`
	// 契約狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 契約開始日期
	StartDate time.Time `gorm:"column:start_date;type:date;not null;" json:"start_date"`
	// 契約有效期限(月)
	Term int `gorm:"column:term;type:int;not null;" json:"term"`
	// 契約結束日期
	EndDate time.Time `gorm:"column:end_date;type:date;not null;" json:"end_date"`
	// 商機ID
	OpportunityID string `gorm:"column:opportunity_id;type:uuid;not null;" json:"opportunity_id"`
	// opportunities  data
	Opportunities opportunities.Table `gorm:"foreignKey:OpportunityID;references:OpportunityID" json:"opportunities,omitempty"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 契約描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 契約號碼
	Code string `gorm:"->;column:code;type:text;not null;" json:"code"`
	// 業務員ID
	SalespersonID string `gorm:"column:salesperson_id;type:uuid;not null;" json:"salesperson_id"`
	// salespeople  data
	Salespeople users.Table `gorm:"foreignKey:SalespersonID;references:UserID" json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.UseTable
}

// Base struct is corresponding to contracts table structure file
type Base struct {
	// 契約ID
	ContractID *string `json:"contract_id,omitempty"`
	// 契約狀態
	Status *string `json:"status,omitempty"`
	// 契約開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 契約有效期限(月)
	Term *int `json:"term,omitempty"`
	// 契約結束日期
	EndDate *time.Time `json:"end_date,omitempty"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty"`
	// opportunities  data
	Opportunities opportunities.Base `json:"opportunities,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 契約描述
	Description *string `json:"description,omitempty"`
	// 契約號碼
	Code *string `json:"code,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty"`
	// salespeople  data
	Salespeople users.Base `json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "contracts"
}
