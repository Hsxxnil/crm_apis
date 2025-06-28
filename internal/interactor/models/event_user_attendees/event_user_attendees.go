package event_user_attendees

import (
	"crm/internal/interactor/models/page"

	"crm/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 事件ID
	EventID string `json:"event_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 參與人員ID
	AttendeeID string `json:"attendee_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 事件參與人員ID
	EventUserAttendeeID string `json:"event_user_attendee_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" form:"event_id"`
	// 參與人員ID
	AttendeeID *string `json:"attendee_id,omitempty" form:"attendee_id"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	// 多筆
	EventUserAttendees []*struct {
		// 事件參與人員ID
		EventUserAttendeeID string `json:"event_user_attendee_id,omitempty"`
		// 事件ID
		EventID string `json:"event_id,omitempty"`
		// 參與人員ID
		AttendeeID string `json:"attendee_id,omitempty"`
		// 參與人員名稱
		AttendeeName string `json:"attendee_name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"event_user_attendees"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 事件參與人員ID
	EventUserAttendeeID string `json:"event_user_attendee_id,omitempty"`
	// 事件ID
	EventID string `json:"event_id,omitempty"`
	// 參與人員ID
	AttendeeID string `json:"attendee_id,omitempty"`
	// 參與人員名稱
	AttendeeName string `json:"attendee_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 事件參與人員ID
	EventUserAttendeeID string `json:"event_user_attendee_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 參與人員ID
	AttendeeID *string `json:"attendee_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// EventSingle return structure file for events
type EventSingle struct {
	// 參與人員ID
	AttendeeID string `json:"attendee_id,omitempty"`
	// 參與人員名稱
	AttendeeName string `json:"attendee_name,omitempty"`
}
