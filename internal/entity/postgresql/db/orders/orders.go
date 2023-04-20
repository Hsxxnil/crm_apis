package orders

import (
	"time"

	"app.eirc/internal/interactor/models/special"
)

// Table struct is orders database table struct
type Table struct {
	// 訂單ID
	OrderID string `gorm:"<-:create;column:order_id;type:uuid;not null;primaryKey;" json:"order_id"`
	// 訂單狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 訂單開始日期
	StartDate time.Time `gorm:"column:start_date;type:TIMESTAMP;not null;" json:"start_date"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 契約ID
	ContractID string `gorm:"column:contract_id;type:uuid;not null;" json:"contract_id"`
	// 訂單描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	special.UseTable
}

// Base struct is corresponding to orders table structure file
type Base struct {
	// 訂單ID
	OrderID *string `json:"order_id,omitempty"`
	// 訂單狀態
	Status *string `json:"status,omitempty"`
	// 訂單開始日期
	StartDate *string `json:"start_date,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 契約ID
	ContractID *string `json:"contract_id,omitempty"`
	// 訂單描述
	Description *string `json:"description,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "orders"
}
