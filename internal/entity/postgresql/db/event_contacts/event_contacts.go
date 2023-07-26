package event_contacts

import (
	"app.eirc/internal/entity/postgresql/db/contacts"
	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/special"
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
	// 事件聯絡人是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;not null;" json:"is_deleted"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.UseTable
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
	// 事件聯絡人是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "event_contacts"
}
