package contracts

import (
	"time"

	"app.eirc/internal/interactor/models/special"
)

// Table struct is contracts database table struct
type Table struct {
	// 契約ID
	ContractID string `gorm:"<-:create;column:contract_id;type:uuid;not null;primaryKey;" json:"contract_id"`
	// 契約狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 契約開始日期
	StartDate time.Time `gorm:"column:start_date;type:TIMESTAMP;not null;" json:"start_date"`
	// 契約有效期限(月)
	Term int `gorm:"column:term;type:int;not null;" json:"term"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 契約描述
	Description string `gorm:"column:description;type:text;" json:"description"`
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
	// 契約描述
	Description *string `json:"description,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "contracts"
}
