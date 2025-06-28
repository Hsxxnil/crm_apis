package events

import (
	"time"

	"crm/internal/entity/postgresql/db/event_contacts"
	"crm/internal/entity/postgresql/db/event_user_attendees"
	"crm/internal/entity/postgresql/db/event_user_mains"

	"crm/internal/entity/postgresql/db/accounts"

	"crm/internal/entity/postgresql/db/users"
	model "crm/internal/interactor/models/events"
	"crm/internal/interactor/models/special"
)

// Table struct is events database table struct
type Table struct {
	// 事件ID
	EventID string `gorm:"<-:create;column:event_id;type:uuid;not null;primaryKey;" json:"event_id"`
	// 事件主題
	Subject string `gorm:"column:subject;type:text;not null;" json:"subject"`
	// 事件是否為全天事件
	IsWhole bool `gorm:"column:is_whole;type:boolean;not null;" json:"is_whole"`
	// 事件開始日期
	StartDate time.Time `gorm:"column:start_date;type:timestamp;not null;" json:"start_date"`
	// 事件結束日期
	EndDate time.Time `gorm:"column:end_date;type:timestamp;not null;" json:"end_date"`
	// 帳戶ID
	AccountID string `gorm:"column:account_id;type:uuid;" json:"account_id"`
	// accounts data
	Accounts accounts.Table `gorm:"foreignKey:AccountID;references:AccountID" json:"accounts,omitempty"`
	// 事件類型
	Type string `gorm:"column:type;type:text;" json:"type"`
	// 事件地址
	Location string `gorm:"column:location;type:text;" json:"location"`
	// 事件描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	// event_user_mains data
	EventUserMains []event_user_mains.Table `gorm:"foreignKey:EventID" json:"mains,omitempty"`
	// event_user_attendees data
	EventUserAttendees []event_user_attendees.Table `gorm:"foreignKey:EventID" json:"attendees,omitempty"`
	// event_contacts data
	EventContacts []event_contacts.Table `gorm:"foreignKey:EventID" json:"contacts,omitempty"`
	special.Table
}

// Base struct is corresponding to events table structure file
type Base struct {
	// 事件ID
	EventID *string `json:"event_id,omitempty"`
	// 事件主題
	Subject *string `json:"subject,omitempty"`
	// 事件是否為全天事件
	IsWhole *bool `json:"is_whole,omitempty"`
	// 事件開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 事件結束日期
	EndDate *time.Time `json:"end_date,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty"`
	// accounts data
	Accounts accounts.Base `json:"accounts,omitempty"`
	// 事件類型
	Type *string `json:"type,omitempty"`
	// 事件地址
	Location *string `json:"location,omitempty"`
	// 事件描述
	Description *string `json:"description,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// event_user_mains data
	EventUserMains []event_user_mains.Base `json:"mains,omitempty"`
	// event_user_attendees data
	EventUserAttendees []event_user_attendees.Base `json:"attendees,omitempty"`
	// event_contacts data
	EventContacts []event_contacts.Base `json:"contacts,omitempty"`
	special.Base
	// 搜尋欄位
	model.Filter `json:"filter"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "events"
}
