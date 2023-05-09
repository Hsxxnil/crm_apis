package orders

import (
	"time"

	"app.eirc/internal/interactor/models/sort"

	"app.eirc/internal/entity/postgresql/db/order_products"
	"app.eirc/internal/entity/postgresql/db/users"
	model "app.eirc/internal/interactor/models/orders"

	"app.eirc/internal/entity/postgresql/db/contracts"

	"app.eirc/internal/entity/postgresql/db/accounts"

	"app.eirc/internal/interactor/models/special"
)

// Table struct is orders database table struct
type Table struct {
	// 訂單ID
	OrderID string `gorm:"<-:create;column:order_id;type:uuid;not null;primaryKey;" json:"order_id"`
	// 訂單狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 訂單開始日期
	StartDate time.Time `gorm:"column:start_date;type:date;not null;" json:"start_date"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 契約ID
	ContractID string `gorm:"column:contract_id;type:uuid;not null;" json:"contract_id"`
	// contracts data
	Contracts contracts.Table `gorm:"foreignKey:ContractID;references:ContractID" json:"contracts,omitempty"`
	// 訂單描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 訂單號碼
	Code string `gorm:"->;column:code;type:text;not null;" json:"code"`
	// 啟用時間
	ActivatedAt *time.Time `gorm:"column:activated_at;type:TIMESTAMP;" json:"activated_at,omitempty"`
	// 啟用者
	ActivatedBy *string `gorm:"column:activated_by;type:uuid;" json:"activated_by,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	// activate_users data
	ActivatedByUsers users.Table `gorm:"foreignKey:ActivatedBy;references:UserID" json:"activated_by_users,omitempty"`
	// order_products data
	OrderProducts []order_products.Table `gorm:"foreignKey:OrderID;" json:"products,omitempty"`
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
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 契約ID
	ContractID *string `json:"contract_id,omitempty"`
	// contracts data
	Contracts contracts.Base `json:"contracts,omitempty"`
	// 訂單描述
	Description *string `json:"description,omitempty"`
	// 訂單號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// activate_users data
	ActivatedByUsers users.Base `json:"activated_by_users,omitempty"`
	// order_products data
	OrderProducts []order_products.Base `json:"products,omitempty"`
	special.UseBase
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "orders"
}
