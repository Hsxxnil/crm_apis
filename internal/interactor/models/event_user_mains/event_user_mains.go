package event_user_mains

import (
	"app.eirc/internal/interactor/models/page"

	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 事件ID
	EventID string `json:"event_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 主要人員ID
	MainID string `json:"main_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 事件主要人員是否刪除
	IsDeleted bool `json:"is_deleted,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 事件主要人員ID
	EventUserMainID string `json:"event_user_main_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" form:"event_id"`
	// 主要人員ID
	MainID *string `json:"main_id,omitempty" form:"main_id"`
	// 事件主要人員是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" form:"is_deleted"`
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
	EventUserMains []*struct {
		// 事件主要人員ID
		EventUserMainID string `json:"event_user_main_id,omitempty"`
		// 事件ID
		EventID string `json:"event_id,omitempty"`
		// 主要人員ID
		MainID string `json:"main_id,omitempty"`
		// 主要人員名稱
		MainName string `json:"main_name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"event_user_mains"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 事件主要人員ID
	EventUserMainID string `json:"event_user_main_id,omitempty"`
	// 事件ID
	EventID string `json:"event_id,omitempty"`
	// 主要人員ID
	MainID string `json:"main_id,omitempty"`
	// 主要人員名稱
	MainName string `json:"main_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 事件主要人員ID
	EventUserMainID string `json:"event_user_main_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件ID
	EventID *string `json:"event_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 主要人員ID
	MainID *string `json:"main_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 事件主要人員是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
