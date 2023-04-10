package accounts

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is accounts database table struct
type Table struct {
	//帳號ID
	AccountID string `gorm:"<-:create;column:account_id;type:uuid;not null;primaryKey;" json:"account_id"`
	//公司ID
	CompanyID string `gorm:"column:company_id;type:uuid;not null;" json:"company_id"`
	//帳號
	Account string `gorm:"<-:create;column:account;type:varchar;not null;" json:"account"`
	//中文名稱
	Name string `gorm:"column:name;type:varchar;not null;" json:"name"`
	//密碼
	Password string `gorm:"column:password;type:varchar;not null;" json:"password"`
	//是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;not null;default:false" json:"is_deleted"`
	//電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	//電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	special.UseTable
}

// Base struct is corresponding to accounts table structure file
type Base struct {
	//帳號ID
	AccountID *string `json:"account_id,omitempty"`
	//公司ID
	CompanyID *string `json:"company_id,omitempty"`
	//帳號
	Account *string `json:"account,omitempty"`
	//中文名稱
	Name *string `json:"name,omitempty"`
	//密碼
	Password *string `json:"password,omitempty"`
	//是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	//電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	//電子郵件
	Email *string `json:"email,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_accounts"
}
