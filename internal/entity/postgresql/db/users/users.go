package users

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is users database table struct
type Table struct {
	// 使用者ID
	UserID string `gorm:"<-:create;column:user_id;type:uuid;not null;primaryKey;" json:"user_id"`
	// 公司ID
	CompanyID string `gorm:"column:company_id;type:uuid;not null;" json:"company_id"`
	// 使用者名稱
	UserName string `gorm:"column:user_name;type:text;not null;" json:"user_name"`
	// 中文名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 密碼
	Password string `gorm:"column:password;type:text;not null;" json:"password"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;not null;default:false" json:"is_deleted"`
	// 電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	special.UseTable
}

// Base struct is corresponding to users table structure file
type Base struct {
	// 使用者ID
	UserID *string `json:"user_id,omitempty"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 密碼
	Password *string `json:"password,omitempty"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// 電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 電子郵件
	Email *string `json:"email,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "users"
}
