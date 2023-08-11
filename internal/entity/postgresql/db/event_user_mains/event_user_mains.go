package event_user_mains

import (
	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/special"
)

// Table struct is event_user_main_users database table struct
type Table struct {
	// 事件主要人員ID
	EventUserMainID string `gorm:"<-:create;column:event_user_main_id;type:uuid;not null;primaryKey;" json:"event_user_main_id"`
	// 事件ID
	EventID string `gorm:"column:event_id;type:uuid;not null;" json:"event_id"`
	// 主要人員ID
	MainID string `gorm:"column:main_id;type:uuid;not null;" json:"main_id"`
	// main_users data
	Mains users.Table `gorm:"foreignKey:MainID;references:UserID" json:"mains,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to event_user_mains table structure file
type Base struct {
	// 事件主要人員ID
	EventUserMainID *string `json:"event_user_main_id,omitempty"`
	// 事件ID
	EventID *string `json:"event_id,omitempty"`
	// 主要人員ID
	MainID *string `json:"main_id,omitempty"`
	// main_users data
	Mains users.Base `json:"mains,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "event_user_mains"
}
