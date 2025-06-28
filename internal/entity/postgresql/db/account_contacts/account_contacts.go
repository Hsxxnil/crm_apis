package account_contacts

import (
	"crm/internal/entity/postgresql/db/users"

	"crm/internal/interactor/models/special"
)

// Table struct is account_contacts database table struct
type Table struct {
	// 帳戶聯絡人ID
	AccountContactID string `gorm:"<-:create;column:account_contact_id;type:uuid;not null;primaryKey;" json:"account_contact_id"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;not null;" json:"account_id"`
	// 聯絡人ID
	ContactID string `gorm:"column:contact_id;type:uuid;not null;" json:"contact_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to account_contacts table structure file
type Base struct {
	// 帳戶聯絡人ID
	AccountContactID *string `json:"account_contact_id,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table account_id for this struct type
func (t *Table) TableName() string {
	return "account_contacts"
}
