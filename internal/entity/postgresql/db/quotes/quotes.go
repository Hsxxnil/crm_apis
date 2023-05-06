package quotes

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/opportunities"

	"app.eirc/internal/entity/postgresql/db/users"

	"app.eirc/internal/interactor/models/special"
	"github.com/shopspring/decimal"
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
	// 商機ID
	OpportunityID string `gorm:"column:opportunity_id;type:uuid;not null;" json:"opportunity_id"`
	// opportunities data
	Opportunities opportunities.Table `gorm:"foreignKey:OpportunityID;references:OpportunityID" json:"opportunities,omitempty"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 報價到期日期
	ExpirationDate time.Time `gorm:"column:expiration_date;type:date;" json:"expiration_date"`
	// 報價描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 報價稅額
	Tax decimal.Decimal `gorm:"column:tax;type:numeric;" json:"tax"`
	// 報價運輸和處理費
	ShippingAndHandling decimal.Decimal `gorm:"column:shipping_and_handling;type:numeric;" json:"shipping_and_handling"`
	// 報價號碼
	Code uint `gorm:"->;column:code;type:serial;auto_increment" json:"code"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.UseTable
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
	Tax *decimal.Decimal `json:"tax,omitempty"`
	// 報價運輸和處理費
	ShippingAndHandling *decimal.Decimal `json:"shipping_and_handling,omitempty"`
	// 報價號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "quotes"
}
