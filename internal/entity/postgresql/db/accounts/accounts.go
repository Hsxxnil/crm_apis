package accounts

import (
	"app.eirc/internal/entity/postgresql/db/contacts"
	"app.eirc/internal/interactor/models/special"
)

// Table struct is accounts database table struct
type Table struct {
	// 帳戶ID
	AccountID string `gorm:"<-:create;column:account_id;type:uuid;not null;primaryKey;" json:"account_id"`
	// 帳戶名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 帳戶電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 帳戶行業ID
	IndustryID string `gorm:"column:industry_id;type:uuid;not null;" json:"industry_id"`
	// 帳戶類型ID
	AccountTypeID string `gorm:"column:account_type_id;type:uuid;not null;" json:"account_type_id"`
	// 帳戶父系帳戶ID
	ParentAccountID string `gorm:"column:parent_account_id;type:uuid;not null;" json:"parent_account_id"`
	special.UseTable
	// contacts data
	Contacts []contacts.Table `gorm:"foreignKey:AccountID;" json:"contacts"`
}

// Base struct is corresponding to accounts table structure file
type Base struct {
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 帳戶名稱
	Name *string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 帳戶行業ID
	IndustryID *string `json:"industry_id,omitempty"`
	// 帳戶類型ID
	AccountTypeID *string `json:"account_type_id,omitempty"`
	// 帳戶父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty"`
	special.UseBase
	// contacts data
	Contacts []contacts.Base `json:"contacts,omitempty"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "accounts"
}
