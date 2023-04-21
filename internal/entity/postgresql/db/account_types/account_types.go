package account_types

import "app.eirc/internal/interactor/models/page"

// Table struct is account_types database table struct
type Table struct {
	// 帳戶類型ID
	AccountTypeID string `gorm:"<-:create;column:account_type_id;type:uuid;not null;primaryKey;" json:"account_type_id"`
	// 帳戶類型名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
}

// Base struct is corresponding to account_types table structure file
type Base struct {
	// 帳戶類型ID
	AccountTypeID *string `json:"account_type_id,omitempty"`
	// 帳戶類型名稱
	Name *string `json:"name,omitempty"`
	// 引入page
	page.Pagination
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "account_types"
}
