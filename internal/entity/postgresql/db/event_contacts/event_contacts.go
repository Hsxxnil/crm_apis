package event_contacts

import (
	"crm/internal/entity/postgresql/db/contacts"
	"crm/internal/entity/postgresql/db/users"
	"crm/internal/interactor/models/special"
)

// Table struct is event_contacts database table struct
type Table struct {
	// 事件聯絡人ID
	EventContactID string `gorm:"<-:create;column:event_contact_id;type:uuid;not null;primaryKey;" json:"event_contact_id"`
	// 事件ID
	EventID string `gorm:"column:event_id;type:uuid;not null;" json:"event_id"`
	// 聯絡人ID
	ContactID string `gorm:"column:contact_id;type:uuid;not null;" json:"contact_id"`
	// contacts data
	Contacts contacts.Table `gorm:"foreignKey:ContactID;references:ContactID" json:"contacts,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to event_contacts table structure file
type Base struct {
	// 事件聯絡人ID
	EventContactID *string `json:"event_contact_id,omitempty"`
	// 事件ID
	EventID *string `json:"event_id,omitempty"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty"`
	// contacts data
	Contacts contacts.Base `json:"contacts,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "event_contacts"
}
