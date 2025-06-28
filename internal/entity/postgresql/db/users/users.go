package users

import (
	"crm/internal/interactor/models/special"
)

// Table struct is users database table struct
type Table struct {
	// 使用者ID
	UserID string `gorm:"<-:create;column:user_id;type:uuid;not null;primaryKey;" json:"user_id"`
	// 公司ID
	CompanyID string `gorm:"column:company_id;type:uuid;not null;" json:"company_id"`
	// 使用者名稱
	UserName string `gorm:"column:user_name;type:text;not null;" json:"user_name"`
	// 使用者中文名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 使用者密碼
	Password string `gorm:"column:password;type:text;not null;" json:"password"`
	// 使用者電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 使用者電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// 角色ID
	RoleID string `gorm:"column:role_id;type:uuid;not null;" json:"role_id"`
	special.Table
}

// Base struct is corresponding to users table structure file
type Base struct {
	// 使用者ID
	UserID *string `json:"user_id,omitempty"`
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者密碼
	Password *string `json:"password,omitempty"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"`
	// 角色ID
	RoleID *string `json:"role_id,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "users"
}
