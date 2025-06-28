package event_user_attendees

import (
	"crm/internal/entity/postgresql/db/users"
	"crm/internal/interactor/models/special"
)

// Table struct is event_user_attendee_users database table struct
type Table struct {
	// 事件參與人員ID
	EventUserAttendeeID string `gorm:"<-:create;column:event_user_attendee_id;type:uuid;not null;primaryKey;" json:"event_user_attendee_id"`
	// 事件ID
	EventID string `gorm:"column:event_id;type:uuid;not null;" json:"event_id"`
	// 參與人員ID
	AttendeeID string `gorm:"column:attendee_id;type:uuid;not null;" json:"attendee_id"`
	// attendee_users data
	Attendees users.Table `gorm:"foreignKey:AttendeeID;references:UserID" json:"attendees,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to event_user_attendees table structure file
type Base struct {
	// 事件參與人員ID
	EventUserAttendeeID *string `json:"event_user_attendee_id,omitempty"`
	// 事件ID
	EventID *string `json:"event_id,omitempty"`
	// 參與人員ID
	AttendeeID *string `json:"attendee_id,omitempty"`
	// attendee_users data
	Attendees users.Base `json:"attendees,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "event_user_attendees"
}
