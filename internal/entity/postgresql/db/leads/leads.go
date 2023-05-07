package leads

import (
	"app.eirc/internal/entity/postgresql/db/accounts"
	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/special"
)

// Table struct is leads database table struct
type Table struct {
	// 線索ID
	LeadID string `gorm:"<-:create;column:lead_id;type:uuid;not null;primaryKey;" json:"lead_id"`
	// 線索狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	// 線索描述
	Description string `gorm:"column:description;type:text;not null;" json:"description"`
	// 線索來源
	Source string `gorm:"column:source;type:text;" json:"source"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 線索分級
	Rating string `gorm:"column:rating;type:text;not null;" json:"rating"`
	// 業務員ID
	SalespersonID string `gorm:"column:salesperson_id;type:uuid;not null;" json:"salesperson_id"`
	// salespeople  data
	Salespeople users.Table `gorm:"foreignKey:SalespersonID;references:UserID" json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.UseTable
}

// Base struct is corresponding to leads table structure file
type Base struct {
	// 線索ID
	LeadID *string `json:"lead_id,omitempty"`
	// 線索狀態
	Status *string `json:"status,omitempty"`
	// 線索描述
	Description *string `json:"description,omitempty"`
	// 線索來源
	Source *string `json:"source,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 線索分級
	Rating *string `json:"rating,omitempty"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty"`
	// salespeople  data
	Salespeople users.Base `json:"salespeople,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "leads"
}
