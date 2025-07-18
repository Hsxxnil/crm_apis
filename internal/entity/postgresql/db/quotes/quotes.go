package quotes

import (
	"time"

	"crm/internal/interactor/models/sort"

	"crm/internal/entity/postgresql/db/opportunities"
	"crm/internal/entity/postgresql/db/quote_products"
	model "crm/internal/interactor/models/quotes"

	"crm/internal/entity/postgresql/db/users"

	"crm/internal/interactor/models/special"
)

// Table struct is quotes database table struct
type Table struct {
	// 報價ID
	QuoteID string `gorm:"<-:create;column:quote_id;type:uuid;not null;primaryKey;" json:"quote_id"`
	// 報價名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 報價狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 報價與商機是否同步化
	IsSyncing bool `gorm:"column:is_syncing;type:bool;not null;" json:"is_syncing"`
	// 報價是否為最終版
	IsFinal bool `gorm:"column:is_final;type:bool;not null;" json:"is_final"`
	// 商機ID
	OpportunityID string `gorm:"column:opportunity_id;type:uuid;not null;" json:"opportunity_id"`
	// opportunities data
	Opportunities opportunities.Table `gorm:"foreignKey:OpportunityID;references:OpportunityID" json:"opportunities,omitempty"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 報價到期日期
	ExpirationDate time.Time `gorm:"column:expiration_date;type:timestamp;" json:"expiration_date"`
	// 報價描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 報價稅額
	Tax float64 `gorm:"column:tax;type:numeric;" json:"tax"`
	// 報價運費及其他費用
	ShippingAndHandling float64 `gorm:"column:shipping_and_handling;type:numeric;" json:"shipping_and_handling"`
	// 報價號碼
	Code string `gorm:"->;column:code;type:text;not null;" json:"code"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	// quote_products data
	QuoteProducts []quote_products.Table `gorm:"foreignKey:QuoteID;" json:"products,omitempty"`
	special.Table
}

// Base struct is corresponding to quotes table structure file
type Base struct {
	// 報價ID
	QuoteID *string `json:"quote_id,omitempty"`
	// 報價名稱
	Name *string `json:"name,omitempty"`
	// 報價狀態
	Status *string `json:"status,omitempty"`
	// 報價與商機是否同步化
	IsSyncing *bool `json:"is_syncing,omitempty"`
	// 報價是否為最終版
	IsFinal *bool `json:"is_final,omitempty"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty"`
	// opportunities data
	Opportunities opportunities.Base `json:"opportunities,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 報價到期日期
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// 報價描述
	Description *string `json:"description,omitempty"`
	// 報價稅額
	Tax *float64 `json:"tax,omitempty"`
	// 報價運費及其他費用
	ShippingAndHandling *float64 `json:"shipping_and_handling,omitempty"`
	// 報價號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// quote_products data
	QuoteProducts []quote_products.Base `json:"products,omitempty"`
	special.Base
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "quotes"
}
