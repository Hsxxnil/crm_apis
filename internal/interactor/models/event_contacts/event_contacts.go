package event_contacts

import (
	"app.eirc/internal/interactor/models/page"

	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 事件ID
	EventID string `json:"event_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 事件聯絡人ID
	EventContactID string `json:"event_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" form:"event_id"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty" form:"contact_id"`
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
	EventContacts []*struct {
		// 事件聯絡人ID
		EventContactID string `json:"event_contact_id,omitempty"`
		// 事件ID
		EventID string `json:"event_id,omitempty"`
		// 聯絡人ID
		ContactID string `json:"contact_id,omitempty"`
		// 聯絡人名稱
		ContactName string `json:"contact_name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"event_contacts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 事件聯絡人ID
	EventContactID string `json:"event_contact_id,omitempty"`
	// 事件ID
	EventID string `json:"event_id,omitempty"`
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	ContactName string `json:"contact_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 事件聯絡人ID
	EventContactID string `json:"event_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// EventSingle return structure file for events
type EventSingle struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	ContactName string `json:"contact_name,omitempty"`
}
