package leads

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is leads database table struct
type Table struct {
	// 商機線索ID
	LeadID string `gorm:"<-:create;column:lead_id;type:uuid;not null;primaryKey;" json:"lead_id"`
	// 商機線索狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 商機線索客戶名稱
	CompanyName string `gorm:"column:company_name;type:text;not null;" json:"company_name"`
	// 商機線索來源ID
	SourceID string `gorm:"column:source_id;type:uuid;not null;" json:"source_id"`
	// 商機線索客戶行業ID
	IndustryID string `gorm:"column:industry_id;type:uuid;not null;" json:"industry_id"`
	// 商機線索分級
	Rating string `gorm:"column:rating;type:text;not null;" json:"rating"`
	special.UseTable
}

// Base struct is corresponding to leads table structure file
type Base struct {
	// 商機線索ID
	LeadID *string `json:"lead_id,omitempty"`
	// 商機線索狀態
	Status *string `json:"status,omitempty"`
	// 商機線索客戶名稱
	CompanyName *string `json:"company_name,omitempty"`
	// 商機線索來源ID
	SourceID *string `json:"source_id,omitempty"`
	// 商機線索客戶行業ID
	IndustryID *string `json:"industry_id,omitempty"`
	// 商機線索分級
	Rating *string `json:"rating,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_leads"
}
