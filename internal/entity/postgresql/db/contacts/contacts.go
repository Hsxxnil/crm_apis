package contacts

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is contacts database table struct
type Table struct {
	// 聯絡人ID
	ContactID string `gorm:"<-:create;column:contact_id;type:uuid;not null;primaryKey;" json:"contact_id"`
	// 聯絡人名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 聯絡人職稱
	Title string `gorm:"column:title;type:text;" json:"title"`
	// 聯絡人電話
	PhoneNumber string `gorm:"column:phone_number;type:text;not null;" json:"phone_number"`
	// 聯絡人行動電話
	CellPhone string `gorm:"column:cell_phone;type:text;" json:"cell_phone"`
	// 聯絡人電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// 聯絡人稱謂
	Salutation string `gorm:"column:salutation;type:text;" json:"salutation"`
	// 聯絡人部門
	Department string `gorm:"column:department;type:text;" json:"department"`
	// 聯絡人直屬上司ID
	SupervisorID string `gorm:"column:supervisor_id;type:uuid;not null;" json:"supervisor_id"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	special.UseTable
}

// Base struct is corresponding to contacts table structure file
type Base struct {
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	Name *string `json:"name,omitempty"`
	// 聯絡人職稱
	Title *string `json:"title,omitempty"`
	// 聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	Email *string `json:"email,omitempty"`
	// 聯絡人稱謂
	Salutation *string `json:"salutation,omitempty"`
	// 聯絡人部門
	Department *string `json:"department,omitempty"`
	// 聯絡人直屬上司ID
	SupervisorID *string `json:"supervisor_id,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "contacts"
}
