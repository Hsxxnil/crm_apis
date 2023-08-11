package accounts

import (
	"app.eirc/internal/entity/postgresql/db/account_contacts"
	"app.eirc/internal/entity/postgresql/db/industries"
	"app.eirc/internal/entity/postgresql/db/users"
	model "app.eirc/internal/interactor/models/accounts"
	"app.eirc/internal/interactor/models/sort"
	"app.eirc/internal/interactor/models/special"
	"github.com/lib/pq"
)

// Table struct is accounts database table struct
type Table struct {
	// 帳戶ID
	AccountID string `gorm:"<-:create;column:account_id;type:uuid;not null;primaryKey;" json:"account_id"`
	// 帳戶名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 帳戶電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 行業ID
	IndustryID *string `gorm:"column:industry_id;type:uuid;" json:"industry_id"`
	// industries data
	Industries industries.Table `gorm:"foreignKey:IndustryID;references:IndustryID" json:"industries,omitempty"`
	// 帳戶類型
	Type pq.StringArray `gorm:"column:type;type:text[];not null;" json:"type"`
	// 父系帳戶ID
	ParentAccountID *string `gorm:"column:parent_account_id;type:uuid;" json:"parent_account_id"`
	// 業務員ID
	SalespersonID string `gorm:"column:salesperson_id;type:uuid;not null;" json:"salesperson_id"`
	// salespeople  data
	Salespeople users.Table `gorm:"foreignKey:SalespersonID;references:UserID" json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	// account_contacts data
	AccountContacts []account_contacts.Table `gorm:"foreignKey:AccountID;" json:"contacts,omitempty"`
	special.Table
}

// Base struct is corresponding to accounts table structure file
type Base struct {
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 帳戶名稱
	Name *string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty"`
	// industries data
	Industries industries.Base `json:"industries,omitempty"`
	// 帳戶類型
	Type *[]string `json:"type,omitempty"`
	// 父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty"`
	// salespeople  data
	Salespeople users.Base `json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// account_contacts data
	AccountContacts []account_contacts.Base `json:"contacts,omitempty"`
	special.Base
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "accounts"
}
