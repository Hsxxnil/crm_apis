package lead_contacts

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is lead_contacts database table struct
type Table struct {
	// 線索聯絡人ID
	LeadContactID string `gorm:"<-:create;column:lead_contact_id;type:uuid;not null;primaryKey;" json:"lead_contact_id"`
	// 線索聯絡人名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 線索聯絡人職稱
	Title string `gorm:"column:title;type:text;" json:"title"`
	// 線索聯絡人電話
	PhoneNumber string `gorm:"column:phone_number;type:text;not null;" json:"phone_number"`
	// 線索聯絡人行動電話
	CellPhone string `gorm:"column:cell_phone;type:text;" json:"cell_phone"`
	// 線索聯絡人電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// 線索聯絡人LINE
	Line string `gorm:"column:line;type:text;" json:"line"`
	// 線索ID
	LeadID string `gorm:"column:lead_id;type:uuid;not null;" json:"lead_id"`
	special.UseTable
}

// Base struct is corresponding to lead_contacts table structure file
type Base struct {
	// 線索聯絡人ID
	LeadContactID *string `json:"lead_contact_id,omitempty"`
	// 線索聯絡人名稱
	Name *string `json:"name,omitempty"`
	// 線索聯絡人職稱
	Title *string `json:"title,omitempty"`
	// 線索聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 線索聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty"`
	// 線索聯絡人電子郵件
	Email *string `json:"email,omitempty"`
	// 線索聯絡人LINE
	Line *string `json:"line,omitempty"`
	// 線索ID
	LeadID *string `json:"lead_id,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_lead_contacts"
}
