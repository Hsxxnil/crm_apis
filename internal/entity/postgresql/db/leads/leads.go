package leads

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is leads database table struct
type Table struct {
	// 線索ID
	LeadID string `gorm:"<-:create;column:lead_id;type:uuid;not null;primaryKey;" json:"lead_id"`
	// 線索狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 線索描述
	Description string `gorm:"column:description;type:text;not null;" json:"description"`
	// 線索來源
	Source string `gorm:"column:source;type:text;" json:"source"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 線索分級
	Rating string `gorm:"column:rating;type:text;not null;" json:"rating"`
	special.UseTable
}

// Base struct is corresponding to leads table structure file
type Base struct {
	// 線索ID
	LeadID *string `json:"lead_id,omitempty"`
	// 線索狀態
	Status *string `json:"status,omitempty"`
	// 線索描述
	Description *string `json:"description,omitempty"`
	// 線索來源
	Source *string `json:"source,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 線索分級
	Rating *string `json:"rating,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "leads"
}
