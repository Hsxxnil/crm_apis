package events

import (
	"time"

	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 事件主題
	Subject string `json:"subject,omitempty" binding:"required" validate:"required"`
	// 主要人員IDs
	Main []string `json:"main,omitempty" binding:"required" validate:"required"`
	// 參與人員IDs
	Attendee []string `json:"attendee,omitempty"`
	// 事件是否為全天事件
	IsWhole bool `json:"is_whole,omitempty"`
	// 事件開始日期
	StartDate time.Time `json:"start_date,omitempty" binding:"required" validate:"required"`
	// 事件結束日期
	EndDate time.Time `json:"end_date,omitempty" binding:"required" validate:"required"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 聯絡人IDs
	Contact []string `json:"contact,omitempty"`
	// 事件類型
	Type string `json:"type,omitempty"`
	// 事件地址
	Location string `json:"location,omitempty"`
	// 事件描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 事件ID
	EventID string `json:"event_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件主題
	Subject *string `json:"subject,omitempty" form:"subject"`
	// 事件是否為全天事件
	IsWhole *bool `json:"is_whole,omitempty" form:"is_whole"`
	// 事件開始日期
	StartDate *time.Time `json:"start_date,omitempty" form:"start_date"`
	// 事件結束日期
	EndDate *time.Time `json:"end_date,omitempty" form:"end_date"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 事件類型
	Type *string `json:"type,omitempty" form:"type"`
	// 事件地址
	Location *string `json:"location,omitempty" form:"location"`
	// 事件描述
	Description *string `json:"description,omitempty" form:"description"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 搜尋欄位
	Filter `json:"filter"`
}

// Filter struct is used to store the search field
type Filter struct {
	// 事件主題
	FilterSubject string `json:"subject,omitempty"`
	// 事件主要人員ID
	FilterMainID string `json:"main_id,omitempty"`
	// 事件參與人員ID
	FilterAttendeeID string `json:"attendee_id,omitempty"`
	// 事件類型
	FilterType string `json:"type,omitempty"`
	// 事件開始日期
	FilterStartDate string `json:"start_date,omitempty"`
	// 事件結束日期
	FilterEndDate time.Time `json:"end_date,omitempty" swaggerignore:"true"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Events []*struct {
		// 事件ID
		EventID string `json:"event_id,omitempty"`
		// 事件主題
		Subject string `json:"subject,omitempty"`
		// 事件主要人員
		Main []*Main `json:"main,omitempty"`
		// 事件參與人員
		Attendees []*Attendees `json:"attendees,omitempty"`
		// 事件是否為全天事件
		IsWhole bool `json:"is_whole,omitempty"`
		// 事件開始日期
		StartDate time.Time `json:"start_date,omitempty"`
		// 事件結束日期
		EndDate time.Time `json:"end_date,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		AccountName string `json:"account_name,omitempty"`
		// 聯絡人
		Contacts []*Contacts `json:"contacts,omitempty"`
		// 事件類型
		Type string `json:"type,omitempty"`
		// 事件地址
		Location string `json:"location,omitempty"`
		// 事件描述
		Description string `json:"description,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"events"`
}

// Single return structure file
type Single struct {
	// 事件ID
	EventID string `json:"event_id,omitempty"`
	// 事件主題
	Subject string `json:"subject,omitempty"`
	// 事件主要人員
	Main []*Main `json:"main,omitempty"`
	// 事件參與人員
	Attendees []*Attendees `json:"attendees,omitempty"`
	// 事件是否為全天事件
	IsWhole bool `json:"is_whole,omitempty"`
	// 事件開始日期
	StartDate time.Time `json:"start_date,omitempty"`
	// 事件結束日期
	EndDate time.Time `json:"end_date,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	AccountName string `json:"account_name,omitempty"`
	// 聯絡人
	Contacts []*Contacts `json:"contacts,omitempty"`
	// 事件類型
	Type string `json:"type,omitempty"`
	// 事件地址
	Location string `json:"location,omitempty"`
	// 事件描述
	Description string `json:"description,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 事件ID
	EventID string `json:"event_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 事件主題
	Subject *string `json:"subject,omitempty"`
	// 主要人員IDs
	Main *[]string `json:"main,omitempty"`
	// 參與人員IDs
	Attendee *[]string `json:"attendee,omitempty"`
	// 事件是否為全天事件
	IsWhole *bool `json:"is_whole,omitempty"`
	// 事件開始日期
	StartDate *time.Time `json:"start_date,omitempty"`
	// 事件結束日期
	EndDate *time.Time `json:"end_date,omitempty"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 聯絡人IDs
	Contact *[]string `json:"contact,omitempty"`
	// 事件類型
	Type *string `json:"type,omitempty"`
	// 事件地址
	Location *string `json:"location,omitempty"`
	// 事件描述
	Description *string `json:"description,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Main struct is used to get 事件主要人員
type Main struct {
	// 主要人員ID
	MainID string `json:"main_id,omitempty"`
	// 事件主要人員名稱
	MainName string `json:"main_name,omitempty"`
}

// Attendees struct is used to get 事件參與人員
type Attendees struct {
	// 參與人員ID
	AttendeeID string `json:"attendee_id,omitempty"`
	// 事件參與人員名稱
	AttendeeName string `json:"attendee_name,omitempty"`
}

// Contacts struct is used to get 聯絡人
type Contacts struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	ContactName string `json:"contact_name,omitempty"`
}
