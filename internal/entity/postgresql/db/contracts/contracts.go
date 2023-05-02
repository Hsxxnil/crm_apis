package contracts

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/users"

	"app.eirc/internal/entity/postgresql/db/accounts"

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
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 契約描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 契約號碼
	Code uint `gorm:"->;column:code;type:serial;auto_increment" json:"code"`
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
	StartDate *string `json:"start_date,omitempty"`
	// 契約有效期限(月)
	Term *int `json:"term,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 契約描述
	Description *string `json:"description,omitempty"`
	// 契約號碼
	Code *uint `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "contracts"
}
