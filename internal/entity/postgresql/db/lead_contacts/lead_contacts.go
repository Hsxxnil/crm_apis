package lead_contacts

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is lead_contacts database table struct
type Table struct {
	// 商機線索聯絡人ID
	LeadContactID string `gorm:"<-:create;column:lead_contact_id;type:uuid;not null;primaryKey;" json:"lead_contact_id"`
	// 商機線索聯絡人名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 商機線索聯絡人職稱
	Title string `gorm:"column:title;type:text;" json:"title"`
	// 商機線索聯絡人電話
	PhoneNumber string `gorm:"column:phone_number;type:text;not null;" json:"phone_number"`
	// 商機線索聯絡人行動電話
	CellPhone string `gorm:"column:cell_phone;type:text;" json:"cell_phone"`
	// 商機線索聯絡人電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// 商機線索聯絡人LINE
	Line string `gorm:"column:line;type:text;" json:"line"`
	special.UseTable
}

// Base struct is corresponding to lead_contacts table structure file
type Base struct {
	// 商機線索聯絡人ID
	LeadContactID *string `json:"lead_contact_id,omitempty"`
	// 商機線索聯絡人名稱
	Name *string `json:"name,omitempty"`
	// 商機線索聯絡人職稱
	Title *string `json:"title,omitempty"`
	// 商機線索聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 商機線索聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty"`
	// 商機線索聯絡人電子郵件
	Email *string `json:"email,omitempty"`
	// 商機線索聯絡人LINE
	Line *string `json:"line,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_lead_contacts"
}
